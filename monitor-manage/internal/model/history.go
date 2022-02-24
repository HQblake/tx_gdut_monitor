/*
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-11 10:35:39
 * @LastEditors: yzq
 */
package model

type HistoryInfo struct {
	Id        int32   `json:"id"`
	Ip        string  `json:"ip"`
	Local     string  `json:"local"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Method    int32   `json:"method"`
	Level     int32   `json:"level"`
	Start     int64   `json:"start"`
	End       int64   `json:"end"`
	Duration  string  `json:"duration"`
}

type HistoryResponse struct {
	Id        int32   `json:"id"`
	Ip        string  `json:"ip"`
	Local     string  `json:"local"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Method    int32   `json:"method"`
	Level     int32   `json:"level"`
	Start     string  `json:"start"`
	End       string  `json:"end"`
	Duration  string  `json:"duration"`
}
