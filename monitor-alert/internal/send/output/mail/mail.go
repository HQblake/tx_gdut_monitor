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

type Mail struct {
	pool       *sync.Pool
	level      output.Level
	mail       *email.Pool
	target     []string
	formatType string
	infoCh     chan *email.Email
	stopCh     chan bool
	enable     bool
}

func NewMail(level output.Level, conf *Config) (*Mail, error) {
	pool, err := email.NewPool("smtp.qq.com:25",  3, smtp.PlainAuth("", "zekew@foxmail.com", "arcxevvtlmzjcada", "smtp.qq.com"))
	if err != nil {
		return nil, err
	}
	m := &Mail{
		enable:     false,
		level:      level,
		target:     conf.Target,
		formatType: conf.FormatType,
		mail:       pool,
		pool: &sync.Pool{
			New: func() interface{} {
				return &email.Email{
					Headers: textproto.MIMEHeader{},
					From:    "zekew@foxmail.com",
					To:      conf.Target,
					Subject: "TX_GDUT_ALERT",
				}
			},
		},
		infoCh: make(chan *email.Email, 5),
		stopCh: make(chan bool),
	}
	go m.sendMail()
	m.enable = true
	return m, nil
}

func (m *Mail) sendMail()  {
	for mail := range m.infoCh {
		err :=  m.mail.Send(mail, 5 * time.Second)
		if err != nil {
			log.Println(err)
		}
		m.pool.Put(mail)
	}
	m.stopCh<-true
}

func (m *Mail) Level() output.Level {
	return m.level
}

func (m *Mail) Reset(level output.Level, config interface{}) error {
	conf, ok := config.(*Config)
	if !ok {
		return fmt.Errorf("config type is invalid")
	}
	err := conf.doCheck()
	if err != nil {
		return err
	}
	m.level = level
	m.formatType = conf.FormatType
	m.pool.New = func() interface{} {
		return &email.Email{
			Headers: textproto.MIMEHeader{},
			From:    "zekew@foxmail.com",
			To:      conf.Target,
			Subject: "TX_GDUT_ALERT",
		}
	}
	return nil
}

func (m *Mail) Output(info model.Info) error {
	if !m.enable {
		return nil
	}
	em := m.pool.Get().(*email.Email)
	msg, err := format.Format(m.formatType, info)
	if err != nil {
		return err
	}
	if m.formatType == "html" {
		em.HTML = msg
		em.Headers.Add("Content-Type", "text/html; charset=UTF-8")
	}else {
		em.Text = msg
		em.Headers.Add("Content-Type", "text/plain; charset=UTF-8")
	}
	m.infoCh<-em
	return nil
}

func (m *Mail) Finish() error {
	if !m.enable {
		return nil
	}
	m.enable = false
	close(m.infoCh)
	<-m.stopCh
	m.mail.Close()
	return nil
}

