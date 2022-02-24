package mail

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"net/textproto"
	"sync"
	"time"
)

// amklofqapqejcabf
// arcxevvtlmzjcada
const (
	address = "smtp.qq.com:25"
	user    = "zekew@foxmail.com"
	code    = "arcxevvtlmzjcada"
	host    = "smtp.qq.com"
	subject = "TX_GDUT_ALERT"
)

type Mail struct {
	pool       *sync.Pool
	level      output.Level
	mail       *email.Pool
	target     []string
	formatType string
	infoCh     chan *email.Email
	stopCh     chan bool
	enable     bool

	lock *sync.RWMutex
}

func NewMail(level output.Level, conf *EMailConf) (*Mail, error) {
	pool, err := email.NewPool(address, 5, smtp.PlainAuth("", user, code, host))
	if err != nil {
		return nil, err
	}
	m := &Mail{
		lock:       &sync.RWMutex{},
		enable:     false,
		level:      level,
		target:     conf.Target,
		formatType: conf.FormatType,
		mail:       pool,
		pool: &sync.Pool{
			New: func() interface{} {
				return &email.Email{
					Headers: textproto.MIMEHeader{},
					From:    user,
					Subject: subject,
				}
			},
		},
		infoCh: make(chan *email.Email, 5),
		stopCh: make(chan bool),
	}
	go m.sendMail()
	// 可以写入
	m.enable = true
	return m, nil
}

func (m *Mail) sendMail() {
	for mail := range m.infoCh {
		log.Println("mail 告警开始, mail to", mail.To)
		err := m.mail.Send(mail, 10*time.Second)
		if err != nil {
			log.Println("send mail:", err)
		}
		m.pool.Put(mail)
		log.Println("mail 告警结束")
	}
	m.stopCh <- true
}

func (m *Mail) Level() output.Level {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.level
}

func (m *Mail) Reset(level output.Level, config interface{}) error {
	conf, ok := config.(*Config)
	if !ok {
		return fmt.Errorf("config type is invalid")
	}
	c, err := conf.doCheck()
	if err != nil {
		return err
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.level = level
	m.formatType = c.FormatType
	m.target = c.Target
	return nil
}

func (m *Mail) Output(infos []model.Info) error {
	if !m.enable {
		return nil
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	if m.mail == nil {
		return nil
	}
	em := m.pool.Get().(*email.Email)
	msg, err := format.Format(m.formatType, infos)
	if err != nil {
		return err
	}
	if m.formatType == "html" {
		em.HTML = msg
		em.Headers.Add("Content-Type", "text/html; charset=UTF-8")
	} else {
		em.Text = msg
		em.Headers.Add("Content-Type", "text/plain; charset=UTF-8")
	}
	em.To = m.target
	m.infoCh <- em
	return nil
}

func (m *Mail) Finish() error {
	if !m.enable {
		// 避免重复关
		return nil
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	close(m.infoCh)
	<-m.stopCh
	m.mail.Close()
	m.mail = nil
	m.enable = false
	return nil
}
