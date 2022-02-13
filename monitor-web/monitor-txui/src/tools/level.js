import { Parse, Stringfy } from '@/tools/json'
export var LevelType = {
  0: 'info',
  1: 'warn',
  2: 'error',
  3: 'panic'
}
export var LevelMap = {
  'info': 0,
  'warn': 1,
  'error': 2,
  'panic': 3
}

export function ParseObj (threshold) {
    let levelObj = Parse(threshold)
    let obj = {}
    Object.keys(LevelType).forEach(function (key) {
        obj[LevelType[key]] = {
            "label": LevelType[key],
            "value": levelObj[key] ? levelObj[key] : ''
        }
    })
    return obj
  }

  export function StringObj (threshold) {
    let obj = {}
    Object.keys(threshold).forEach(function (key) {
        if(threshold[key].value) {
            obj[LevelMap[key]] = threshold[key].value
        }
      
    })
    return Stringfy(obj)
  }

  export function CheckThreshold (threshold) {
      let flag = false
    Object.keys(LevelType).forEach(function (key) {
        if(threshold[LevelType[key]].value){
            flag = true
        }
    })
    return flag
  }

