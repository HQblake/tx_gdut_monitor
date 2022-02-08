package judgment

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"

// IJudgment 服务接口化，对外提供Update方法，内部调用Grpc
type IJudgment interface {
	Update(rule *model.AgentRule) error
}
