export var MethodType = {
  0: '总和',
  1: '平均值',
  2: '中位数',
  3: '积分',
  4: '极值',
  5: '标准差',
  6: '最大值',
  7: '最小值'
}
export var MethodMap = {
  '总和': 0,
  '平均值': 1,
  '中位数': 2,
  '积分': 3,
  '极值': 4,
  '标准差': 5,
  '最大值': 6,
  '最小值': 7
}

export function ParseMethod (m) {
  let v = MethodType[m]
  if (v) {
    return v
  }
  return 'unknown'
}

export function CheckMethod (m) {
  let v = MethodType[m]
  if (v) {
    return true
  }
  return false
}
