import { fetchGet } from '@/request.js'

// 获取所有agent列表
export function GetAllAgent () {
  return fetchGet('/agent/list', {})
}

export function GetAllSendAgent () {
  return fetchGet('/agent/sendList', {})
}

// 获取指定agent的信息
export function GetAgentInfo (ip, local) {
  let url = '/agent/info/' + ip + '/' + local
  return fetchGet(url, {})
}
