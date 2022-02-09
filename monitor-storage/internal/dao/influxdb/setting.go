package influxdb

type InfluxDBSetting struct {
	URL                string
	ORG                string
	Token              string
	Bucket             string
	LogLevel           int
	BatchSize          int
	FlushIntervalMs    int
	RetryInterval      int
	MaxRetryIntervalMs int
	MaxRetries         int
	MaxRetryTime       int
	ExponentialBase    int
	UseGZip            bool
	RetryBufferLimit   int
	Precision          int
	HttpRequestTimeout int
}

type Method struct {
	ID      int32
	English string
	Chinese string
}

var Methods = map[int32]Method{
	0: Method{0, "count", "计数"},
	1: Method{1, "sum", "总和"},
	2: Method{2, "mean", "平均值"},
	3: Method{3, "median", "中位数"},
	4: Method{4, "integral", "积分"},
	5: Method{5, "mode", "众数"},
	6: Method{6, "spread", "极值"},
	7: Method{7, "stddev", "标准差"},
	8: Method{8, "max", "最大值"},
	9: Method{9, "min", "最小值"},
}
