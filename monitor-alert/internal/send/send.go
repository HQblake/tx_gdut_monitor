package send

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"

type ISend interface {
	Send(alert model.AlertInfo) error
}

//// SendService 告警系统——判定服务模块
//type SendService struct{}
//
//// 为告警系统的其他服务提供调用 SendService 服务相应功能的接口
//
//// Send 方法用于对判定服务发来的告警信息进行处理，并通过邮箱进行通知
//func (ss *SendService) Send(alert *model.AlertInfo) error {
//	return nil
//}
