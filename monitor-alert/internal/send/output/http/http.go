package http

import (
	"bytes"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Http struct {
	level      output.Level
	formatType string
	client     *http.Client
	headers    http.Header
	target     string
	method     string
	info       chan []byte
	stopCh     chan bool

	// 避免并发读写
	lock *sync.RWMutex
}

func NewHttp(level output.Level, config *Config) (*Http, error) {
	h := &Http{
		method:     config.Method,
		target:     config.Url,
		level:      level,
		formatType: config.FormatType,
		client:     &http.Client{Timeout: 30 * time.Second},
		headers:    toHeader(config.Headers),
		info:       make(chan []byte, 5),
		stopCh:     make(chan bool),
		lock:       &sync.RWMutex{},
	}
	go h.send()
	return h, nil
}

func (h *Http) Level() output.Level {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.level
}

func (h *Http) Reset(level output.Level, config interface{}) error {
	conf, ok := config.(*Config)
	if !ok {
		return fmt.Errorf("config type is invalid")
	}
	err := conf.doCheck()
	if err != nil {
		return err
	}
	h.lock.Lock()
	defer h.lock.Unlock()
	h.level = level
	h.formatType = conf.FormatType
	h.headers = toHeader(conf.Headers)
	h.method = conf.Method
	h.target = conf.Url
	return nil
}

func (h *Http) Output(infos []model.Info) error {
	msg, err := format.Format(h.formatType, infos)
	if err != nil {
		return err
	}
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.client == nil {
		return nil
	}
	h.info <- msg
	return nil
}

func (h *Http) Finish() error {
	h.lock.Lock()
	defer h.lock.Unlock()
	close(h.info)
	<-h.stopCh
	h.client = nil
	return nil
}

func (h *Http) send() {
	for in := range h.info {
		request, err := http.NewRequest(h.method, h.target, bytes.NewReader(in))
		if err != nil {
			log.Printf("new request error :%s \n", err.Error())
			continue
		}
		request.Header = h.headers
		if h.formatType == "json" {
			request.Header.Set("Content-Type", "application/json")
		}
		if h.formatType == "line" {
			request.Header.Set("Content-Type", "application/plain")
		}
		if h.formatType == "html" {
			request.Header.Set("Content-Type", "application/html")
		}
		log.Println("http 告警发送")
		response, err := h.client.Do(request)
		if err != nil {
			log.Printf("send request error :%s \n", err.Error())
			continue
		}
		if response.StatusCode != 200 {
			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("send to http error :%s \n", err.Error())
			} else {
				log.Printf("http response error :%s \n", string(b))
			}
			continue
		}
		response.Body.Close()
		log.Println("http 告警结束")
	}
	h.stopCh <- true
}
