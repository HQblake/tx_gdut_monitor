package http

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
)

type Http struct {
	level      output.Level
	formatType string
}

func (h *Http) Level() output.Level {
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
	h.level = level
	h.formatType = conf.FormatType
	return nil
}

func (h *Http) Output(info model.Info) error {
	panic("implement me")
}

func (h *Http) Finish() error {
	panic("implement me")
}

