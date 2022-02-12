package mysql

//func TestSaveAgentInfo(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	m:=model.Metric{}
//
//	err:=client.SaveAgentInfo(&m)
//	if err != nil {
//		println(err)
//	}
//}
//func TestSaveAlertInfo(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	h:=model.HistoryInfo{
//		IP:        "i",
//		Local:     "0",
//		Metric:    "ho",
//		Value:     23,
//		Threshold: 14,
//		Level:     1,
//		Start:     1521507600,
//		Duration:  "m",
//	}
//
//	err:=client.SaveAlertInfo(&h)
//	if err != nil {
//		println(err)
//	}
//}
//func TestGetAllAgentInfo(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	res:=client.GetAllAgentInfo()
//	if res != nil {
//		fmt.Println(res)
//
//	}
//}
//func TestGetAllAlertInfo(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	res:=client.GetAllAlertInfo()
//	if res != nil {
//		fmt.Println(res)
//
//	}
//}
//
//func TestGetAlertInfo(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	//测试id
//	//res:=client.GetAlertInfo(1,-1,"","","",1231,13144)
//	//if res != nil {
//	//	fmt.Println(res)
//	//}
//	//测试level
//	//res:=client.GetAlertInfo(0,0,"","","",0,0)
//	//if res != nil {
//	//	fmt.Println(res)
//	//}
//	//测试ip,local
//	//res:=client.GetAlertInfo(0,-1,"i","0","",0,0)
//	//if res != nil {
//	//	fmt.Println(res)
//	//}
//	//测试metric
//	//res:=client.GetAlertInfo(0,-1,"","","ho",0,0)
//	//if res != nil {
//	//	fmt.Println(res)
//	//}
//	//测试metric
//	res:=client.GetAlertInfo(0,-1,"","","",1521507600,1521507600)
//	if res != nil {
//		fmt.Println(res)
//	}
//}

//func TestGetAgentInfoByIPAndLocal(t *testing.T) {
//	s := MySQLSetting{}
//	client := NewClient(&s)
//	res := client.GetAgentInfoByIPAndLocal("i", "0")
//	fmt.Println(res)
//}
//func TestGetMetricsByIPAndLocal(t *testing.T) {
//	s := MySQLSetting{}
//	client := NewClient(&s)
//	res := client.GetMetricsByIPAndLocal("i","0")
//	fmt.Println(res)
//}
//func TestDelAlterInfo(t *testing.T) {
//	s := MySQLSetting{}
//	client := NewClient(&s)
//	res := client.DelAlterInfo(6)
//	fmt.Println(res)
//}

//func TestGetCheckConfigsByIPAndLocal(t *testing.T) {
//	s := MySQLSetting{}
//	client := NewClient(&s)
//	res := client.GetCheckConfigsByIPAndLocal("iou","090")
//	fmt.Println(res)
//}

//func TestUpdateCheckConfig(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	c:=model.CheckConfig{
//		ID:        1,
//		IP:        "iou",
//		Local:     "090",
//		Metric:    "ho",
//		Method:    2,
//		Period:    "2m",
//		Threshold: "th2",
//	}
//	err:=client.UpdateCheckConfig(&c)
//	if err != nil {
//		println(err)
//	}
//}

//func TestDelCheckConfigByID(t *testing.T) {
//	s := MySQLSetting{}
//	client := NewClient(&s)
//	res := client.DelCheckConfigByID(2)
//	fmt.Println(res)
//}
//func TestSaveAlertConfig(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	al:=model.AlertConfig{
//		IP:       "i",
//		Local:    "0",
//		SendType: 3,
//		Config:   "config3",
//		Level:    3,
//	}
//
//	err:=client.SaveAlertConfig(&al)
//	if err != nil {
//		println(err)
//	}
//}
//func TestUpdateAlertConfig(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	al:=model.AlertConfig{
//		ID: 1,
//		IP:       "i",
//		Local:    "0",
//		SendType: 4,
//		Config:   "config3",
//		Level:    3,
//	}
//
//	err:=client.UpdateAlertConfig(&al)
//	if err != nil {
//		println(err)
//	}
//}

//func TestDelAlertConfigByID(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	res:=client.DelAlertConfigByID(4)
//	if res != nil {
//		fmt.Println(res)
//	}
//}
//func TestGetAlertConfigByID(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	res:=client.GetAlertConfigByID(5)
//	fmt.Println(res)
//}

//func TestGetAlertConfigByIPAndLocal(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	res:=client.GetAlertConfigByIPAndLocal("iou","090")
//	fmt.Println(res)
//}
//func TestGetAllAlertConfig(t *testing.T) {
//	s:=MySQLSetting{}
//	client:=NewClient(&s)
//	res:=client.GetAllAlertConfig()
//	fmt.Println(res)
//	fmt.Println(time.Now())
//}



