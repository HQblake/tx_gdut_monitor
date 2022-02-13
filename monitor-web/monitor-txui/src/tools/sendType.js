var sendType = {
  0: 'email',
  1: 'kafka',
  2: 'nsq',
  3: 'http'
}
export function ParseSendtype (t) {
  let v = sendType[t]
  if (v) {
    return v
  }
  return 'unknown'
}

export function CheckType (m) {
  let v = sendType[m]
  if (v) {
    return true
  }
  return false
}
