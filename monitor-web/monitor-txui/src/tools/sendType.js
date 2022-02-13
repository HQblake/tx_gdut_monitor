var sendType = {
  0: 'email',
  1: 'kafka',
  2: 'nsq',
  3: 'http'
}
export function PaerseSendtype (t) {
  let v = sendType[t]
  if (v) {
    return v
  }
  return 'unknown'
}
