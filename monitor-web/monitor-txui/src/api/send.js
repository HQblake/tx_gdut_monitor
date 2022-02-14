import { fetchGet, fetchPost } from '@/request.js'

// 获取指定agent的判定规则信息
export function GetSendConfigs (ip, local) {
  let url = '/send/info/' + ip + '/' + local
  return fetchGet(url, {})
}

export function AddSendConfig (ip, local, sendType, level, config) {
  console.log(ip, local, sendType, level, config)
  let url = '/send/info/' + ip + '/' + local
  return fetchPost(url, {
    'sendType': sendType,
    'level': level,
    'config': config
  })
}

// 更新agent的指定id判定规则信息
export function UpdateSendConfig (ip, local, id, sendType, level, config) {
  console.log(ip, local, id, sendType, level, config)
  let url = '/send/update/' + ip + '/' + local + '/' + id
  return fetchPost(url, {
    'sendType': sendType,
    'level': level,
    'config': config
  })
}

// 删除agent的指定id判定规则信息
export function DelSendConfig (ip, local, id) {
  let url = '/send/del/' + ip + '/' + local + '/' + id
  return fetchPost(url, {})
}
