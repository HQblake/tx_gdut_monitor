import { fetchGet, fetchPost } from '@/request.js'

export function GetWarnList (ip, local, metric, begin, limit) {
  console.log(ip, local, metric, begin, limit);
  let url = "warn/metrics/" + ip + "/" + local + "/" + metric + "/" + begin + "/" + limit
  return fetchGet(url, {})
}

// 获取所有告警信息
export function GetWarnList () {
  let url = "warn/warnList"
  return fetchGet(url, {})
}

// 根据id获取告警信息
export function GetWarnList (id) {
  console.log(id)
  let url = "warn/warnId/" + id
  return fetchGet(url, {})
}

// 根据其他参数搜索获取告警信息
export function GetWarnInfoWithParams (ip, local, metric, level, start, end) {
  console.log(ip, local, metric, level, start, end)
  let url = "warn/warnParams/" + ip + "/" + local + "/" + metric + "/" +level + "/" + start + "/" + end
  return fetchGet(url, {})
}

// 删除指定id的告警信息
export function DelWarnInfo (id) {
  let url = "warn/del/" + id
  return fetchPost(url, {})
}

