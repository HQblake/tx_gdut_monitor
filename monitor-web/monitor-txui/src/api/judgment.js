import { fetchGet, fetchPost } from '@/request.js'

// 获取指定agent的判定规则信息
export function GetAgentRule (ip, local) {
  let url = '/judgment/info/' + ip + '/' + local
  return fetchGet(url, {})
}

// 更新agent的指定id判定规则信息
export function UpdateRule (ip, local, id, method, metric, period, threshold) {
  let url = '/judgment/update/' + ip + '/' + local + '/' + id
  console.log(ip, local, id, method, metric, period, threshold)
  return fetchPost(url, {
    'method': method,
    'metric': metric,
    'period': period,
    'threshold': threshold
  })
}

// 删除agent的指定id判定规则信息
export function DelRule (ip, local, id) {
  let url = '/judgment/del/' + ip + '/' + local + '/' + id
  return fetchPost(url, {})
}
