var methodType = {
  0: '总和',
  1: '平均值',
  2: '中位数',
  3: '积分',
  4: '极值',
  5: '标准差',
  6: '最大值',
  7: '最小值'
}

export function ParseMethod (m) {
  let v = methodType[m]
  if (v) {
    return v
  }
  return 'unknown'
}

export function CheckMethod (m) {
  let v = methodType[m]
  if (v) {
    return true
  }
  return false
}
