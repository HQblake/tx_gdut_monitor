package judgment

import (
	proto "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/client/judgment/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	"google.golang.org/grpc"
	"sync"
)

// IJudgment 服务接口化，对外提供Update方法，内部实际调用Grpc，尽可能解耦
type IJudgment interface {
	Update(rule *model.AgentRule) error
}
type Service struct {
	client *proto.RuleUpdaterClient
	conn   *grpc.ClientConn
	once   sync.Once
}


