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
	0: Method{0, "sum", "总和"},
	1: Method{1, "mean", "平均值"},
	2: Method{2, "median", "中位数"},
	3: Method{3, "integral", "积分"},
	4: Method{4, "spread", "极值"},
	5: Method{5, "stddev", "标准差"},
	6: Method{6, "max", "最大值"},
	7: Method{7, "min", "最小值"},
}

const (
	aggregateWitLimit = `from(bucket: "%s")
	|> range(start: %v, stop: %v)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> filter(fn: (r) => r["ip"] == "%s")
	|> filter(fn: (r) => r["local"] == "%s")
	|> filter(fn: (r) => r["_field"] == "value")
	|> aggregateWindow(every: %s, fn: %s)
	|> top(n: %v, columns: ["_time"])
	|> sort(columns: ["_time"], desc: false)`

	aggregateWitoutLimit = `from(bucket: "%s")
	|> range(start: %v, stop: %v)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> filter(fn: (r) => r["ip"] == "%s")
	|> filter(fn: (r) => r["local"] == "%s")
	|> filter(fn: (r) => r["_field"] == "value")
	|> aggregateWindow(every: %s, fn: %s)`

	everyWitLimit = `from(bucket: "%s")
	|> range(start: %v, stop: %v)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> filter(fn: (r) => r["ip"] == "%s")
	|> filter(fn: (r) => r["local"] == "%s")
	|> filter(fn: (r) => r["_field"] == "value")
	|> top(n: %v, columns: ["_time"])
	|> sort(columns: ["_time"], desc: false)`

	everyWitoutLimit = `from(bucket: "%s")
	|> range(start: %v, stop: %v)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> filter(fn: (r) => r["ip"] == "%s")
	|> filter(fn: (r) => r["local"] == "%s")
	|> filter(fn: (r) => r["_field"] == "value")`
)
