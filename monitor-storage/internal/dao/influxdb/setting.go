package influxdb

type InfluxDBSetting struct {
	URL       string
	ORG       string
	Token     string
	Bucket    string
	BatchSize int
}

type Method struct {
	ID      int32
	English string
	Chinese string
}

var Methods = map[int32]Method{
	0: Method{0, "sum", "求和"},
	1: Method{1, "mean", "平均"},
	2: Method{2, "", ""},
	3: Method{3, "", ""},
	4: Method{4, "", ""},
	5: Method{5, "", ""},
}
