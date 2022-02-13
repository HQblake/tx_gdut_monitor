package dao

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao/influxdb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao/mysql"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type StorageDao struct {
	influxdbClient *influxdb.Client
	mysqlClient    *mysql.Client
}

var dao *StorageDao = nil

func NewStorageDao(s *setting.Setting) *StorageDao {
	if dao == nil {
		influxdbSetting := &influxdb.InfluxDBSetting{}
		mysqlSetting := &mysql.MySQLSetting{}
		s.ReadSection("InfluxDB", influxdbSetting)
		s.ReadSection("MySQL", mysqlSetting)
		dao = &StorageDao{influxdb.NewClient(influxdbSetting), mysql.NewClient(mysqlSetting)}
	}
	return dao
}

// GetAggregatedData 操作InfluxDB数据库
func (dao *StorageDao) GetAggregatedData(period string, method int32, metric *model.Metric) (float64, error) {
	// 1. 将上报数据保存到InfluxDB中
	err := dao.influxdbClient.SaveMatricData(metric)
	if err != nil {
		return 0, err
	}

	// 2. 获取聚合数据
	return dao.influxdbClient.GetAggregatedData(metric, period, method, metric.Timestamp)
}

func (dao *StorageDao) SaveAgentInfo(metric *model.Metric) error {
	return dao.mysqlClient.SaveAgentInfo(metric)
}

func (dao *StorageDao) SaveAlertInfo(alert *model.HistoryInfo) error {
	return dao.mysqlClient.SaveAlertInfo(alert)
}

func (dao *StorageDao) GetAllAgentInfo() []model.AgentInfo {
	return dao.mysqlClient.GetAllAgentInfo()
}

func (dao *StorageDao) GetAgentInfoByIPAndLocal(ip, local string) model.AgentInfo {
	return dao.mysqlClient.GetAgentInfoByIPAndLocal(ip, local)
}

func (dao *StorageDao) GetMetricsByIPAndLocal(ip, local string) []string {
	return dao.mysqlClient.GetMetricsByIPAndLocal(ip, local)
}

// GetMetricData 操作InfluxDB数据库
func (dao *StorageDao) GetMetricData(ip, local, metricName, period string, begin, end int64,
	method, limit int32) ([]model.Metric, error) {
	return dao.influxdbClient.GetMetricData(ip, local, metricName, period, begin, end, method, limit)
}

func (dao *StorageDao) GetAllAlertInfo() []model.HistoryInfo {
	return dao.mysqlClient.GetAllAlertInfo()
}

func (dao *StorageDao) GetAlertInfo(id, level int32, ip, local, metric string, begin, end int64) []model.HistoryInfo {
	return dao.mysqlClient.GetAlertInfo(id, level, ip, local, metric, begin, end)
}

func (dao *StorageDao) DelAlterInfo(id int32) error {
	return dao.mysqlClient.DelAlterInfo(id)
}

func (dao *StorageDao) GetCheckConfigsByIPAndLocal(ip, local string) []model.CheckConfig {
	return dao.mysqlClient.GetCheckConfigsByIPAndLocal(ip, local)
}

func (dao *StorageDao) UpdateCheckConfig(check *model.CheckConfig) (int32, error) {
	return dao.mysqlClient.UpdateCheckConfig(check)
}

func (dao *StorageDao) DelCheckConfigByID(id int32) error {
	return dao.mysqlClient.DelCheckConfigByID(id)
}

func (dao *StorageDao) SaveAlertConfig(alert *model.AlertConfig) (int32, error) {
	return dao.mysqlClient.SaveAlertConfig(alert)
}

func (dao *StorageDao) UpdateAlertConfig(alert *model.AlertConfig) error {
	return dao.mysqlClient.UpdateAlertConfig(alert)
}

func (dao *StorageDao) DelAlertConfigByID(id int32) error {
	return dao.mysqlClient.DelAlertConfigByID(id)
}

func (dao *StorageDao) GetAlertConfigByID(id int32) model.AlertConfig {
	return dao.mysqlClient.GetAlertConfigByID(id)
}

func (dao *StorageDao) GetAlertConfigByIPAndLocal(ip, local string) []model.AlertConfig {
	return dao.mysqlClient.GetAlertConfigByIPAndLocal(ip, local)
}

func (dao *StorageDao) GetAllAlertConfig() []model.AlertConfig {
	return dao.mysqlClient.GetAllAlertConfig()
}
