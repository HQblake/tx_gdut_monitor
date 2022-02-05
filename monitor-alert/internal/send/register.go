package send

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format/html"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format/line"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output/http"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output/kafka"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output/mail"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output/nsq"
)

func Register()  {
	InitOutputFactory()
	InitFormatPool()
}

func InitOutputFactory()  {
	http.Register()
	kafka.Register()
	nsq.Register()
	mail.Register()
}

func InitFormatPool()  {
	html.Register()
	json.Register()
	line.Register()
}
