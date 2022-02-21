package model

type MetricsInfo struct {
	Timestamp int64   `json:"timestamp" yaml:"timestamp"`
	Metric    string  `json:"metric" yaml:"metric"`
	Value     float64 `json:"value" yaml:"value"`
}

type MetricsReq struct {
	IP     string `json:"ip"`
	Local  string `json:"local"`
	Metric string `json:"metricName"`
	Period string `json:"period"`
	Begin  int64  `json:"begin"`
	End    int64  `json:"end"`
	Method int32  `json:"method"`
	Limit  int32  `json:"limit"`
}

// type MetricRequest struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	IP     string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"` // ip和local组成agent的唯一标识
// 	Local  string `protobuf:"bytes,2,opt,name=Local,proto3" json:"Local,omitempty"`
// 	Metric string `protobuf:"bytes,3,opt,name=Metric,proto3" json:"Metric,omitempty"`
// 	Begin  int64  `protobuf:"varint,4,opt,name=Begin,proto3" json:"Begin,omitempty"`
// 	End    int64  `protobuf:"varint,5,opt,name=End,proto3" json:"End,omitempty"`
// 	Period string `protobuf:"bytes,6,opt,name=Period,proto3" json:"Period,omitempty"`
// 	Method int32  `protobuf:"varint,7,opt,name=Method,proto3" json:"Method,omitempty"`
// 	Limit  int32  `protobuf:"varint,8,opt,name=Limit,proto3" json:"Limit,omitempty"` // 返回条数限制
// }
