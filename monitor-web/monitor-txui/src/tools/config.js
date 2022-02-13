import { Parse, Stringfy } from '@/tools/json'

export var defaultConfig = {
    0: {
        target: '',
        format_type: ''
    },
    1: {
        topic: '',
        format_type: '',
        address: '',
        version: '',
        partition_type: '',
        partition: '',
        partition_key: ''
    },
    2: {
        topic: '',
        format_type: '',
        address: ''
    },
    3: {
        url: '',
        format_type: '',
        method: '',
        headers: {}
    }
}
export var SendType = {
  0: 'email',
  1: 'kafka',
  2: 'nsq',
  3: 'http'
}

export function ParseObj (s) {
  let levelObj = Parse(s)
  let obj = {}
  Object.keys(levelObj).forEach(function (key) {
    obj[levelObj[key]] = {
      'label': levelObj[key],
      'value': levelObj[key] ? levelObj[key] : ''
    }
  })
  return obj
}

// 直接json序列化
export function StringObj (obj) {
  return Stringfy(obj)
}
