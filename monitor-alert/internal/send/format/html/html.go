package html

import (
	"bytes"
	"embed"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"html/template"
)

//go:embed tmp.html
var fs embed.FS


type Html struct {
	tmp *template.Template
}

func (h *Html) Format(info []model.Info) ([]byte, error) {
	b := &bytes.Buffer{}
	b.Reset()
	err := h.tmp.Execute(b, info)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func NewHtml() *Html {
	tmp, err := template.ParseFS(fs, "tmp.html")
	if err != nil {
		tmp = template.Must(template.New("html").Parse(`<div class="container"><div style="text-align: center"><h3>监控指标异常告警，请尽快修复！</h3></div><table><tr><th>Agent</th><th>指标类型</th><th>指标值</th><th>阈值</th><th>聚合方式</th><th>告警等级</th><th>持续时间</th><th>开始时间</th></tr><tr><td>{{ .Agent }}</td><td>{{ .Metric }}</td><td>{{ .Value }}</td><td>{{ .Threshold }}</td><td>{{ .Method }}</td><td>{{ .Level }}</td><td>{{ .Duration }}</td><td>{{ .Start }}</td></tr></table></div>`))
	}
	return &Html{
		tmp: tmp,
	}
}

