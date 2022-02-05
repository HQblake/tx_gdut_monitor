package line

import (
	"bytes"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
)

type Line struct {

}

func (h *Line) Format(info model.Info) ([]byte, error) {
	b := &bytes.Buffer{}
	b.Reset()
	b.WriteString(fmt.Sprintf("|%s|%s|%.2f|%.2f|%s|%s|%s|%s|", info.Agent, info.Metric, info.Value, info.Threshold, info.Level, info.Method, info.Duration, info.Start))
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func NewLine() *Line {
	return &Line{}
}

