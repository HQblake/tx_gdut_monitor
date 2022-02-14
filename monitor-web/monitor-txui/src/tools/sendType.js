export var SendType = {
  0: 'email',
  1: 'kafka',
  2: 'nsq',
  3: 'http'
}
export var SendMap = {
  'email': 0,
  'kafka': 1,
  'nsq': 2,
  'http': 3
}

export function ParseSendtype (t) {
  let v = SendType[t]
  if (v) {
    return v
  }
  return 'unknown'
}

export function CheckType (m) {
  let v = SendType[m]
  if (v) {
    return true
  }
  return false
}
