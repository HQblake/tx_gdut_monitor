package mysql

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"log"
	"strconv"
	"time"
)

// SaveAgentInfo 调用MySQL的存储过程
// 传入参数：IP、Local、Port、Metric_Name，保存Agent信息
func (c *Client) SaveAgentInfo(metric *model.Metric) error {
	_, err := c.db.Exec("CALL AddAgentInfo(?,?,?,?)", metric.IP, metric.Local, metric.Port, metric.Name)
	if err != nil {
		return err
	}
	return nil
}

// SaveAlertInfo 存储告警信息
func (c *Client) SaveAlertInfo(history *model.HistoryInfo) error {
	_, err := c.db.Exec("INSERT INTO history(agentId, metricId, value, threshold, method, "+
		"level, duration, start) VALUES((SELECT a.id FROM agent AS a WHERE a.ip=? AND a.local=?), "+
		"(SELECT m.id FROM metric AS m WHERE m.name=?), ?, ?, ?, ?, ?, ?)",
		history.IP, history.Local, history.Metric, history.Value, history.Threshold, history.Method,
		history.Level, history.Duration, history.Start)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetAllAgentInfo 提取agent表中的所有agent及其相应的metric列表
func (c *Client) GetAllAgentInfo() []model.AgentInfo {
	// 可以考虑循环调用 GetMetricsByAgentID
	var agent model.AgentInfo
	var res []model.AgentInfo
	rows, err := c.db.Query("SELECT * FROM agent")
	if err != nil {
		log.Println(err)
		return res
	}
	for rows.Next() {
		rows.Scan(&agent.ID, &agent.IP, &agent.Local, &agent.Port, &agent.IsLive)
		res = append(res, agent)
		log.Println(agent)
	}
	defer rows.Close()
	return res
}

// GetAgentInfoByIPAndLocal 根据IP与Local获取指定的Agent及其metric列表
func (c *Client) GetAgentInfoByIPAndLocal(ip, local string) model.AgentInfo {
	var agent model.AgentInfo
	rows, err := c.db.Query("SELECT * FROM agent WHERE ip=? AND local=?", ip, local)
	if err != nil {
		log.Println(err)

	} else if rows.Next() {
		rows.Scan(&agent.ID, &agent.IP, &agent.Local, &agent.Port, &agent.IsLive)
	}
	defer rows.Close()
	return agent
}

// GetMetricsByIPAndLocal 根据IP与Local获取相应Agent的metric列表
func (c *Client) GetMetricsByIPAndLocal(ip, local string) []string {
	var tempName string
	var res []string
	rows, err := c.db.Query("SELECT m.name "+
		"FROM ((agent_metric AS am LEFT JOIN agent AS a ON am.agentId=a.id) "+
		"LEFT JOIN metric AS m ON am.metricId=m.id) "+
		"WHERE a.ip=? AND a.local=?", ip, local)
	if err != nil {
		log.Println(err)
		return res
	}
	for rows.Next() {
		rows.Scan(&tempName)
		res = append(res, tempName)
	}
	defer rows.Close()
	return res

}

// GetAllAlertInfo 获取所有历史告警信息
func (c *Client) GetAllAlertInfo() []model.HistoryInfo {
	var history model.HistoryInfo
	var res []model.HistoryInfo
	rows, err := c.db.Query("SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, h.method, h.level, " +
		"h.start, h.duration FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) " +
		"LEFT JOIN metric AS m ON h.metricId=m.id)")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		rows.Scan(&history.ID, &history.IP, &history.Local, &history.Metric, &history.Value,
			&history.Threshold, &history.Method, &history.Level, &history.Start, &history.Duration)
		res = append(res, history)
	}
	defer rows.Close()
	return res
}

// GetAlertInfo 根据给定的条件获取告警信息
// 若给定的参数值为其类型的零值，则表示该条件未设定
// level 参数在此处的零值为"负数"
func (c *Client) GetAlertInfo(id, level int32, ip, local, metric string, begin, end int64) []model.HistoryInfo {
	var alert model.HistoryInfo
	var res []model.HistoryInfo
	if id > 0 {
		// 设定了history表的主键id，仅返回一条记录
		rows, err := c.db.Query("SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, "+
			"h.method, h.level, h.start, h.duration FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) "+
			"LEFT JOIN metric AS m ON h.metricId=m.id) WHERE h.id=?", id)
		if err != nil {
			log.Println(err)
			return res
		}
		for rows.Next() {
			rows.Scan(&alert.ID, &alert.IP, &alert.Local, &alert.Metric, &alert.Value,
				&alert.Threshold, &alert.Method, &alert.Level, &alert.Start, &alert.Duration)
			res = append(res, alert)
		}
	} else {
		// 对设定了的条件进行判定
		sql := "SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, h.method, h.level, h.start, h.duration " +
			"FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) LEFT JOIN metric AS m ON h.metricId=m.id) " +
			"WHERE 1=1"
		if ip != "" {
			sql += " AND a.ip=" + "'" + ip + "'"
		}
		if local != "" {
			sql += " AND a.local=" + "'" + local + "'"
		}
		if metric != "" {
			sql += " AND m.name=" + "'" + metric + "'"
		}
		if level >= 0 {
			sql += " AND h.level=" + strconv.Itoa(int(level))
		}
		sql += " AND h.start>=? AND h.start<=?"

		if begin <= 0 {
			begin = 0
		}
		if end <= 0 {
			end = time.Now().Unix()
		}
		if begin > end {
			begin, end = end, begin
		}
		//rows, err := c.db.Query(sql)
		rows, err := c.db.Query(sql, begin, end)
		if err != nil {
			log.Println(err)
			return res
		}
		for rows.Next() {
			rows.Scan(&alert.ID, &alert.IP, &alert.Local, &alert.Metric, &alert.Value,
				&alert.Threshold, &alert.Method, &alert.Level, &alert.Start, &alert.Duration)
			res = append(res, alert)
		}
	}
	return res
}

// DelAlterInfo 根据ID删除告警信息
func (c *Client) DelAlterInfo(id int32) error {
	res, err := c.db.Exec("DELETE FROM history WHERE id=?", id)
	if err != nil {
		log.Println(res)
		return err
	}
	return nil
}

// GetCheckConfigsByIPAndLocal 根据IP与Local获取check表中的数据
func (c *Client) GetCheckConfigsByIPAndLocal(ip, local string) []model.CheckConfig {
	var check model.CheckConfig
	var res []model.CheckConfig
	rows, err := c.db.Query("SELECT c.id, a.ip, a.local, m.name, c.method, c.period, c.threshold"+
		"FROM ((`check` AS c LEFT JOIN agent AS a ON c.agentId=a.id) LEFT JOIN metric AS m ON c.metricId=m.id)"+
		"WHERE a.ip=? AND a.local=?", ip, local)
	if err != nil {
		log.Println(err)
		return res
	}
	for rows.Next() {
		rows.Scan(&check.ID, &check.IP, &check.Local, &check.Metric, &check.Method, &check.Period, &check.Threshold)
		res = append(res, check)
	}
	defer rows.Close()
	return res
}

// UpdateCheckConfig 更新check表中除ID、IP、Local、Metric外的所有字段
func (c *Client) UpdateCheckConfig(check *model.CheckConfig) error {
	_, err := c.db.Exec("UPDATE `check` SET method=?, period=?, threshold=? WHERE id=?",
		check.Method, check.Period, check.Threshold, check.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DelCheckConfigByID 根据ID删除check表中的记录
func (c *Client) DelCheckConfigByID(id int32) error {
	_, err := c.db.Exec("DELETE FROM `check` WHERE id=?", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SaveAlertConfig  插入将告警配置保存到alert表中*****
func (c *Client) SaveAlertConfig(alert *model.AlertConfig) (int32, error) {
	affect, err := c.db.Exec("INSERT INTO alert(agentId, sendType, level, config) "+
		"VALUES((SELECT a.id FROM agent AS a WHERE a.ip=? AND a.local=?), ?, ?, ?)",
		alert.IP, alert.Local, alert.SendType, alert.Level, alert.Config)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := affect.LastInsertId()
	return int32(id), err
}

// UpdateAlertConfig 更新alert表中除ID、IP、Local外的所有字段
func (c *Client) UpdateAlertConfig(alert *model.AlertConfig) error {
	_, err := c.db.Exec("UPDATE alert SET sendType=?, level=?, config=? WHERE id=?",
		alert.SendType, alert.Level, alert.Config, alert.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DelAlertConfigByID 根据ID删除alert表中的记录
func (c *Client) DelAlertConfigByID(id int32) error {
	_, err := c.db.Exec("DELETE FROM alert WHERE id=?", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetAlertConfigByID 根据ID获取alert表中的记录
func (c *Client) GetAlertConfigByID(id int32) model.AlertConfig {
	var alert model.AlertConfig
	rows, err := c.db.Query("SELECT al.id, a.ip, a.local, al.sendType, al.level, al.config "+
		"FROM alert AS al LEFT JOIN agent AS a ON al.agentId=a.id WHERE al.id=?", id)
	if err != nil {
		log.Println(err)
		return alert
	} else if rows.Next() {
		rows.Scan(&alert.ID, &alert.IP, &alert.Local, &alert.SendType, &alert.Level, &alert.Config)
	}
	defer rows.Close()
	return alert
}

// GetAlertConfigByIPAndLocal 根据IP、Local获取所有alert记录
func (c *Client) GetAlertConfigByIPAndLocal(ip, local string) []model.AlertConfig {
	var alert model.AlertConfig
	var res []model.AlertConfig
	rows, err := c.db.Query("SELECT al.id, a.ip, a.local, al.sendType, al.level, al.config "+
		"FROM alert AS al LEFT JOIN agent AS a ON al.agentId=a.id WHERE a.ip=? AND a.local=?", ip, local)
	if err != nil {
		log.Println(err)
		return res
	}
	for rows.Next() {
		rows.Scan(&alert.ID, &alert.IP, &alert.Local, &alert.SendType, &alert.Config,
			&alert.Level)
		res = append(res, alert)
	}
	defer rows.Close()
	return res

}

// GetAllAlertConfig 获取所有alert记录
func (c *Client) GetAllAlertConfig() []model.AlertConfig {
	var alert model.AlertConfig
	var res []model.AlertConfig
	rows, err := c.db.Query("SELECT al.id, a.ip, a.local, al.sendType, al.level, al.config " +
		"FROM alert AS al LEFT JOIN agent AS a ON al.agentId=a.id")
	if err != nil {
		log.Println(err)
		return res
	}
	for rows.Next() {
		rows.Scan(&alert.ID, &alert.IP, &alert.Local, &alert.SendType, &alert.Level, &alert.Config)
		res = append(res, alert)
	}
	defer rows.Close()
	return res
}
