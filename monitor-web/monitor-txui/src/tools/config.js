import { Parse, Stringfy } from '@/tools/json'

export var defaultConfig = {
  0: {
    target: {
      label: '邮箱地址',
      value: '',
      tip: '多个邮箱则以,分割'
    },
    format_type: {
      label: '告警格式',
      value: 'line',
      selectList: ['line', 'json', 'html']
    }
  },
  1: {
    topic: {
      label: 'Topic',
      value: ''
    },
    format_type: {
      label: '告警格式',
      value: 'line',
      selectList: ['line', 'json', 'html']
    },
    address: {
      label: '集群地址',
      value: '',
      tip: '多个地址则以,分割'
    },
    version: {
      label: '版本号',
      value: '',
      tip: 'kafka版本,如0.2.0.0'
    },
    partition_type: {
      label: '选区方式',
      value: 'random',
      selectList: ['random', 'robin', 'hash', 'manual']
    },
    partition: {
      label: '分区号',
      value: '',
      tip: '分区号, 选区方式为manual时必填'
    },
    partition_key: {
      label: '分区hash',
      value: '',
      tip: '分区号, 选区方式为hash时必填'
    }
  },
  2: {
    topic: {
      label: 'Topic',
      value: ''
    },
    format_type: {
      label: '告警格式',
      value: 'line',
      selectList: ['line', 'json', 'html']
    },
    address: {
      label: 'nsq地址',
      value: ''
    }
  },
  3: {
    url: {
      label: '目标URL',
      value: ''
    },
    format_type: {
      label: '告警格式',
      value: 'line',
      selectList: ['line', 'json', 'html']
    },
    method: {
      label: '请求方式',
      value: 'POST',
      selectList: ['GET', 'POST', 'HEAD', 'PUT', 'DELETE', 'CONNECT', 'OPTIONS', 'TRACE']
    }
  }
}
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

export function ParseObj (t, s) {
  let levelObj = Parse(s)
  let obj = {}
  let demo = defaultConfig[t]
  Object.keys(demo).forEach(function (key) {
    obj[key] = {
      'label': demo[key].label,
      'value': levelObj[key] ? levelObj[key] : demo[key].value
    }
    if (demo[key].selectList) {
      obj[key].selectList = demo[key].selectList
    }
    if (demo[key].tip) {
      obj[key].tip = demo[key].tip
    }
  })
  return obj
}

// 直接json序列化
export function StringObj (config) {
  let obj = {}
  Object.keys(config).forEach(function (key) {
    obj[key] = config[key].value
  })
  return Stringfy(obj)
}

export function CheckConfig (t, config) {
  if (t == 0) {
    if (!config.target.value) {
      return { check: false, msg: '邮箱地址不能为空' }
    }
    if (config.target.value.indexOf('@') == -1) {
      return { check: false, msg: '邮箱地址格式有误' }
    }
  }
  if (t == 1) {
    if (!config.address.value) {
      return { check: false, msg: '地址不能为空' }
    }
    if (!config.topic.value) {
      return { check: false, msg: 'topic不能为空' }
    }
    if (config.partition_type.value == 'manual' && !config.partition.value) {
      return { check: false, msg: '选区方式为manual时,分区号不能为空' }
    }
    if (config.partition_type.value == 'hash' && !config.partition_key.value) {
      return { check: false, msg: '选区方式为hash时,分区hash值不能为空' }
    }
  }
  if (t == 2) {
    if (!config.topic.value) {
      return { check: false, msg: 'topic不能为空' }
    }
    if (!config.address.value) {
      return { check: false, msg: '地址不能为空' }
    }
  }
  if (t == 3) {
    if (!config.url.value) {
      return { check: false, msg: '请求地址不能为空' }
    }
    if (config.url.value.startsWith('http://') && config.url.value.startsWith('https://')) {
      return { check: false, msg: '请求地址有误' }
    }
  }
  return { check: true, msg: '' }
}
