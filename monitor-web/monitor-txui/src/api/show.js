import { fetchGet, fetchPost } from '@/request.js'

export function GetMetricsWithTime (ip, local, metric, begin, end, method, limit) {
  let metricReq = {"ip":ip, "local":local, "metricName":metric, "begin":begin, "end":end, "method":method, "limit":limit}
  console.log(ip, local, metric, begin, limit)
  let url = '/warn/metrics'
  return fetchGet(url, {
    body: metricReq
  })
}

// 获取所有告警信息
export function GetWarnList () {
  let url = '/show/warnList'
  return fetchGet(url, {})
}

// 根据id获取告警信息
export function GetWarnInfoWithId (id) {
  console.log(id)
  let url = '/warn/warnId/' + id
  return fetchGet(url, {})
}

// 根据其他参数搜索获取告警信息
export function GetWarnInfoWithParams (ip, local, metric, level, start, end) {
  console.log('show/api - ',ip, local, metric, level, start, end);
  let url = '/show/warnParams'
  return fetchGet(url, {
    'ip':ip,  
    'local':local,
    'metric':metric,
    'level':level, 
    'start':start,
    'end':end
  })
}

// 删除指定id的告警信息
export function DelWarnInfo (id) {
  let url = 'warn/del/' + id
  return fetchPost(url, {})
}
