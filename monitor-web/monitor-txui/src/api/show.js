/*
 * @Description: 
 * @Autor: yzq
 * @Date: 2022-02-18 16:59:22
 * @LastEditors: yzq
 */
import { fetchGet, fetchPost } from '@/request.js'

export function GetMetricsWithTime (ip, local, metric, begin, limit) {
  console.log(ip, local, metric, begin, limit)
  let url = 'warn/metrics/' + ip + '/' + local + '/' + metric + '/' + begin + '/' + limit
  return fetchGet(url, {})
}

// 获取所有告警信息
export function GetWarnList () {
  let url = '/show/warnList'
  return fetchGet(url, {})
}

// 根据id获取告警信息
export function GetWarnInfoWithId (id) {
  console.log(id)
  let url = 'warn/warnId/' + id
  return fetchGet(url, {})
}

// 根据其他参数搜索获取告警信息
export function GetWarnInfoWithParams (ip, local, metric, level, start, end) {
  warnInfoReq = {"ip":ip, "local":local, "metric":metric, "level":level, "start":start, "end":end}
  console.log(ip, local, metric, level, start, end)
  let url = 'warn/warnParams'
  return fetchGet(url, {
    body:warnInfoReq
  })
}

// 删除指定id的告警信息
export function DelWarnInfo (id) {
  let url = 'warn/del/' + id
  return fetchPost(url, {})
}
