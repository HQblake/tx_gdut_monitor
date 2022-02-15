package worker

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/client"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"sync"
	"time"
)

type Worker struct {
	workers  *ants.PoolWithFunc
	judgment JudgFunc
}

func NewWorker(s *setting.Setting, judg JudgFunc) *Worker {
	ws := WorkersSetting{}
	err := s.ReadSection("Workers", &ws)
	if err != nil {
		log.Fatalln(err)
	}

	duration, _ := time.ParseDuration(ws.ExpiryDuration)
	workers, err := ants.NewPoolWithFunc(ws.Capacity, handler, ants.WithOptions(ants.Options{
		ExpiryDuration:   duration,
		PreAlloc:         ws.PreAlloc,
		MaxBlockingTasks: ws.MaxBlockingTasks,
		Nonblocking:      ws.Nonblocking,
		PanicHandler:     nil,
		Logger:           ants.Logger(log.New(os.Stderr, "", log.LstdFlags)),
	}))
	if err != nil {
		log.Fatalln(err)
	}
	return &Worker{
		workers:  workers,
		judgment: judg,
	}
}

func (w *Worker) Finish(agent *model.AgentReport, rule *model.AgentRule, client *client.Client) model.AlertInfo {
	var wg sync.WaitGroup
	alert := model.AlertInfoPool.Get().(*model.AlertInfo)
	alert.IP, alert.Local, alert.Metrics = agent.IP, agent.Local, make(map[string]model.MetricInfo)
	for k, _ := range agent.Metrics {
		wg.Add(1)
		_ = w.workers.Invoke(task{k, agent, rule, alert, client, &wg, w.judgment})
	}
	wg.Wait()
	return *alert
}
