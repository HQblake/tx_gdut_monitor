package mysql

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"strconv"
	"time"

	//"time"
)


// SaveAgentInfo 调用MySQL的存储过程
// 传入参数：IP、Local、Port、Metric_Name，保存Agent信息
func (c *Client) SaveAgentInfo(metric *model.Metric) error {
	res,err := c.db.Exec("call AddAgentInfo(?,?,?,?)","jhgjh","656yut","jkh","jkhjk")
	if err != nil {
		return err
	}
		fmt.Println(res)
	return nil
}

// SaveAlertInfo 存储告警信息
func (c *Client) SaveAlertInfo(alert *model.HistoryInfo) error {
	var historyAgentid int32
	var historyMetricid int32
	fmt.Println("opoipo")
	rows1, err1 := c.db.Query("select id from agent where ip=? and local=?;",alert.IP,alert.Local)
	if err1 != nil {
		fmt.Println(err1)
		fmt.Println(rows1.Columns())
	}
	fmt.Println(rows1.Columns())
	for rows1.Next(){
		rows1.Scan(&historyAgentid)
	}
	defer rows1.Close()
	rows2, err2 := c.db.Query("select id from metric where name=?;",alert.Metric)
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println(rows2.Columns())
	}
	fmt.Println(rows2.Columns())
	for rows2.Next(){
		rows2.Scan(&historyMetricid)
	}
	defer rows2.Close()
	fmt.Println("historyAgentid,historyMetricid:",historyAgentid,historyMetricid)
	rows, err := c.db.Query("INSERT into history(agentId,metricId,value,threshold,method,level,duration,start) VALUES(?,?,?,?,?,?,?,FROM_UNIXTIME( ?, '%Y-%m-%d %h:%m:%s' )\n);",historyAgentid,historyMetricid,
		alert.Value,alert.Threshold,alert.Method, alert.Level,alert.Duration,alert.Start)

	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// GetAllAgentInfo 提取agent表中的所有agent及其相应的metric列表
func (c *Client) GetAllAgentInfo() []model.AgentInfo {
	// 可以考虑循环调用 GetMetricsByAgentID
	var u1 model.AgentInfo
	var res []model.AgentInfo
	rows, err := c.db.Query("select * from agent;")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.Port, &u1.IsLive)
		res=append(res,u1)
		fmt.Println(u1)
	}
	defer rows.Close()
	return res
}

// GetAgentInfoByIPAndLocal 根据IP与Local获取指定的Agent及其metric列表
func (c *Client) GetAgentInfoByIPAndLocal(ip, local string) model.AgentInfo {
	var u1 model.AgentInfo
	rows, err := c.db.Query("select * from agent where ip=? and local=?;",ip,local)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.Port, &u1.IsLive)

		fmt.Println(u1)
	}
	defer rows.Close()
	return u1
}

// GetMetricsByIPAndLocal 根据IP与Local获取相应Agent的metric列表
func (c *Client) GetMetricsByIPAndLocal(ip, local string) []string {
	var tempName string
	var res []string
	rows, err := c.db.Query("SELECT `name` FROM metric ,agent_metric,agent " +
		"WHERE ip=? and local=? and metric.id=agent_metric.metricId and agent_metric.agentId=agent.id;",ip,local)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&tempName)
		res=append(res,tempName)
		fmt.Println(res)
	}
	defer rows.Close()
	return res

}

// GetAllAlertInfo 获取所有历史告警信息
func (c *Client) GetAllAlertInfo() []model.HistoryInfo {
	var u1 model.HistoryInfo
	var res []model.HistoryInfo
	rows, err := c.db.Query("SELECT history.id,a.ip,a.`local`,m.`name`," +
		"history.`value`,history.threshold,history.method,history.`level`,UNIX_TIMESTAMP(history.`start`),history.duration " +
		"FROM history ,agent as a,metric as m WHERE history.agentId=a.id and history.metricId=m.id")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.Metric, &u1.Value,
			&u1.Threshold, &u1.Method, &u1.Level, &u1.Start, &u1.Duration)
		res=append(res,u1)
		fmt.Println(u1)
	}
	defer rows.Close()
	return res
}

// GetAlertInfo 根据给定的条件获取告警信息
// 若给定的参数值为其类型的零值，则表示该条件未设定
// level 参数在此处的零值为"负数"
func (c *Client) GetAlertInfo(id, level int32, ip, local, metric string, begin, end int64) []model.HistoryInfo {
	var u1 model.HistoryInfo
	var res []model.HistoryInfo
	if id > 0 {
		// 设定了history表的主键id，仅返回一条记录
		sql := "SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, " +
			"h.method, h.level, h.start, h.duration " +
			"FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) " +
			"LEFT JOIN metric AS m ON h.metricId=m.id) " +
			"WHERE h.id=?"
		rows, err := c.db.Query(sql, id)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.Metric, &u1.Value,
				&u1.Threshold, &u1.Method, &u1.Level, &u1.Start, &u1.Duration)
			res=append(res,u1)
		}
		fmt.Println(res)
		return res
	} else {
		// 对设定了的条件进行判定
		sql := "SELECT h.id, a.ip, a.local, m.name, h.value, h.threshold, " +
			"h.method, h.level, UNIX_TIMESTAMP(h.start), h.duration " +
			"FROM ((history AS h LEFT JOIN agent AS a ON h.agentId=a.id) " +
			"LEFT JOIN metric AS m ON h.metricId=m.id) " +
			"WHERE 1=1"
		if ip != "" {
			sql += " AND a.ip=" + "'"+ip+"'"
		}
		if local != "" {
			sql += " AND a.local="+ "'"+local+"'"
		}
		if metric != "" {
			sql += " AND m.name=" + "'"+metric+"'"
		}
		if level >= 0 {
			sql += " AND h.level=" + strconv.Itoa(int(level))
		}
		sql += " AND h.start >=FROM_UNIXTIME(?) AND h.start <= FROM_UNIXTIME(?)"

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
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.Metric, &u1.Value,
				&u1.Threshold, &u1.Method, &u1.Level, &u1.Start, &u1.Duration)
			res=append(res,u1)
		}
		fmt.Println(res)
		return res
	}
	return nil
}

// DelAlterInfo 根据ID删除告警信息
func (c *Client) DelAlterInfo(id int32) error {
	res, err := c.db.Exec("delete from history where id=?;",id)
	if err != nil {
		fmt.Println(res)
		return err
	}
	return nil
}

// GetCheckConfigsByIPAndLocal 根据IP与Local获取check表中的数据
func (c *Client) GetCheckConfigsByIPAndLocal(ip, local string) []model.CheckConfig {
	var u1 model.CheckConfig
	var res []model.CheckConfig
	rows, err := c.db.Query("SELECT `check`.id,agent.ip,agent.`local`, metric.`name`," +
		"`check`.method,`check`.period,`check`.threshold from `check`,agent,metric " +
		"WHERE agent.ip=? and agent.`local`=? and agent.id=`check`.agentId " +
		"and metric.id=`check`.metricId",ip,local)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.Metric, &u1.Method, &u1.Period, &u1.Threshold)
		res=append(res,u1)
		fmt.Println(u1)
	}
	defer rows.Close()
	return res
}

// UpdateCheckConfig 更新check表中除ID、IP、Local外的所有字段
func (c *Client) UpdateCheckConfig(check *model.CheckConfig) error {
	affect, err := c.db.Exec("update `check` set metricId=(SELECT id from metric " +
		"where metric.`name`=?),method=?,period=?,threshold=? WHERE id=?;",
		check.Metric, check.Method, check.Period, check.Threshold, check.ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("出错")
	}
	fmt.Println(affect)
	return nil
}

// DelCheckConfigByID 根据ID删除check表中的记录
func (c *Client) DelCheckConfigByID(id int32) error {
	res, err := c.db.Exec("delete from  `check` where id=?;",id)
	if err != nil {
		fmt.Println("出错",res)
		return err
	}

	return nil
}

// SaveAlertConfig  插入将告警配置保存到alert表中*****
func (c *Client) SaveAlertConfig(alert *model.AlertConfig) error {
	affect, err := c.db.Exec("INSERT into alert(agentId,sendType,`level`,config) " +
		"SELECT a.id,?,?,? FROM agent AS a WHERE a.`ip`=? AND a.`local`=? ;",
		alert.SendType,alert.Level,alert.Config,alert.IP,alert.Local)
	if err != nil {
		fmt.Println(err)
		fmt.Println("出错")
	}
	fmt.Println(affect)
	return nil
}

// UpdateAlertConfig 更新alert表中除ID、IP、Local外的所有字段
func (c *Client) UpdateAlertConfig(alert *model.AlertConfig) error {
	affect, err := c.db.Exec("update alert set sendType=?," +
		"`level`=? ,config=? where id=?",
		alert.SendType,alert.Level,alert.Config,alert.ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("出错")
	}
	fmt.Println(affect)
	return nil
}

// DelAlertConfigByID 根据ID删除alert表中的记录
func (c *Client) DelAlertConfigByID(id int32) error {
	res, err := c.db.Query("delete from alert where id=?;",id)
	if err != nil {
		fmt.Println(res)
		fmt.Println("出错")
		return err
	}
	fmt.Println(res.Columns())
	return nil
}

// GetAlertConfigByID 根据ID获取alert表中的记录
func (c *Client) GetAlertConfigByID(id int32) model.AlertConfig {
	var u1 model.AlertConfig
	rows, err := c.db.Query("select al.id, a.ip, a.`local`," +
		"al.sendType,al.`level`,al.config from alert as al,agent as a where al.id=? and al.agentId=a.id;",id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.SendType, &u1.Level, &u1.Config)
		fmt.Println(u1)
	}
	defer rows.Close()
	return u1
}

// GetAlertConfigByIPAndLocal 根据IP、Local获取所有alert记录
func (c *Client) GetAlertConfigByIPAndLocal(ip, local string) []model.AlertConfig {
	var u1 model.AlertConfig
	var res []model.AlertConfig
	rows, err := c.db.Query("select al.id, a.ip, a.`local`,al.sendType,al.`level`,al.config " +
		"from alert as al,agent as a where a.ip=? and a.`local`=? and al.agentId=a.id;",ip,local)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.SendType, &u1.Config,
			&u1.Level)
		res=append(res,u1)
		fmt.Println(u1)
	}
	defer rows.Close()
	return res

}

// GetAllAlertConfig 获取所有alert记录
func (c *Client) GetAllAlertConfig() []model.AlertConfig {
	var u1 model.AlertConfig
	var res []model.AlertConfig
	rows, err := c.db.Query("SELECT alert.id,agent.ip,agent.`local`," +
		"alert.sendType,alert.`level`,alert.config from alert,agent where alert.agentId=agent.id")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.ID, &u1.IP, &u1.Local, &u1.SendType, &u1.Level, &u1.Config)
		res=append(res,u1)
		fmt.Println(u1)
	}
	defer rows.Close()
	return res
}