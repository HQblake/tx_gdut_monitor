package mysql

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"time"
)

// SaveAgentInfo 调用MySQL的存储过程
// 传入参数：IP、Local、Port、Metric_Name，保存Agent信息
func (c *Client) SaveAgentInfo(metric *model.Metric) error {
	return nil
}

// SaveAlertInfo 存储告警信息
func (c *Client) SaveAlertInfo(alert *model.HistoryInfo) error {
	return nil
}

// GetAllAgentInfo 提取agent表中的所有agent及其相应的metric列表
func (c *Client) GetAllAgentInfo() []model.AgentInfo {
	// 可以考虑循环调用 GetMetricsByAgentID
	return nil
}

// GetAgentInfoByIPAndLocal 根据IP与Local获取指定的Agent及其metric列表
func (c *Client) GetAgentInfoByIPAndLocal(ip, local string) model.AgentInfo {
	return model.AgentInfo{}
}

// GetMetricsByIPAndLocal 根据IP与Local获取相应Agent的metric列表
func (c *Client) GetMetricsByIPAndLocal(ip, local string) []string {
	return nil
}

// GetAllAlertInfo 获取所有历史告警信息
func (c *Client) GetAllAlertInfo() []model.HistoryInfo {
	return nil
}

// GetAlertInfo 根据给定的条件获取告警信息
// 若给定的参数值为其类型的零值，则表示该条件未设定
// level 参数在此处的零值为"负数"
func (c *Client) GetAlertInfo(id, level int32, ip, local, metric string, begin, end int64) []model.HistoryInfo {
	if id <= 0 {
		// 设定了history表的主键id，仅返回一条记录
		sql := "SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, " +
			"h.method, h.level, h.start, h.duration " +
			"FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) " +
			"LEFT JOIN metric AS m ON h.metricId=m.id) " +
			"WHERE h.id=?"
		row, err := c.db.Query(sql, id)
		if err != nil {

		}
		for row.Next() {

		}

		return nil
	} else {
		// 对设定了的条件进行判定
		sql := "SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, " +
			"h.method, h.level, h.start, h.duration " +
			"FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) " +
			"LEFT JOIN metric AS m ON h.metricId=m.id) " +
			"WHERE 1=1"
		if ip != "" {
			sql += " AND a.ip=" + ip
		}
		if local != "" {
			sql += " AND a.local=" + local
		}
		if metric != "" {
			sql += " AND m.name=" + metric
		}
		if level >= 0 {
			sql += " AND h.level=" + string(level)
		}
		sql += " AND h.start >= ? AND h.start <= ?"

		if begin <= 0 {
			begin = 0
		}
		if end <= 0 {
			end = time.Now().Unix()
		}
		if begin > end {
			begin, end = end, begin
		}

		row, err := c.db.Query(sql, begin, end)
		if err != nil {

		}
		for row.Next() {

		}
		return nil
	}
}

// DelAlterInfo 根据ID删除告警信息
func (c *Client) DelAlterInfo(id int32) error {
	return nil
}

// GetCheckConfigsByIPAndLocal 根据IP与Local获取check表中的数据
func (c *Client) GetCheckConfigsByIPAndLocal(ip, local string) []model.CheckConfig {
	return nil
}

// UpdateCheckConfig 更新check表中除ID、IP、Local外的所有字段
func (c *Client) UpdateCheckConfig(check *model.CheckConfig) error {
	return nil
}

// DelCheckConfigByID 根据ID删除check表中的记录
func (c *Client) DelCheckConfigByID(id int32) error {
	return nil
}

// SaveAlertConfig 将告警配置保存到alert表中
func (c *Client) SaveAlertConfig(alert *model.AlertConfig) (int32, error) {
	result, err := c.db.Exec("INSERT INTO ......")
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	return int32(id), err
}

// UpdateAlertConfig 更新alert表中除ID、IP、Local外的所有字段
func (c *Client) UpdateAlertConfig(alert *model.AlertConfig) error {
	return nil
}

// DelAlertConfigByID 根据ID删除alert表中的记录
func (c *Client) DelAlertConfigByID(id int32) error {
	return nil
}

// GetAlertConfigByID 根据ID获取alert表中的记录
func (c *Client) GetAlertConfigByID(id int32) model.AlertConfig {
	return model.AlertConfig{}
}

// GetAlertConfigByIPAndLocal 根据IP、Local获取所有alert记录
func (c *Client) GetAlertConfigByIPAndLocal(ip, local string) []model.AlertConfig {
	return nil
}

// GetAllAlertConfig 获取所有alert记录
func (c *Client) GetAllAlertConfig() []model.AlertConfig {
	return nil
}
