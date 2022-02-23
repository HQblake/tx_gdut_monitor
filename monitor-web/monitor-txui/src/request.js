/*
 * @Description:
 * @Autor: zeke
 * @Date: 2022-02-17 00:52:47
 * @LastEditors: yzq
 */
import axios from 'axios'
import Vue from 'vue'
import qs from 'qs'

function getBaseUrl () {
  return Vue.prototype.BASE_URL
}

axios.defaults.timeout = 10000 // 响应时间
axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded'
// axios.defaults.transformRequest = [
//   function (data) {
//     let ret = ''
//     for (let it in data) {
//       ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
//     }
//     return ret
//     // return qs.stringify(data)
//   }
// ]
// axios.defaults.baseURL = getBaseUrl()

// POST传参序列化(添加请求拦截器)
axios.interceptors.request.use((config) => {
  // 在发送请求之前做某件事
  // if (config.method === 'POST') {
  //   config.data = this.qs.stringify(config.data)
  // }
  return config
}, (error) => {
  return Promise.reject(error)
})

// 返回状态判断(添加响应拦截器)
axios.interceptors.response.use((res) => {
  // 对响应数据做些事
  if (res.status == 200) {
    // 配置文件拦截
    if (res.data.URL) {
      return Promise.resolve(res.data)
    }
    if (res.data.code == '000000') {
      return Promise.resolve(res.data)
    }
  }
  return Promise.reject(res.data)
}, (error) => {
  return Promise.reject({ msg: error.message })
})

// 返回一个Promise(发送post请求)
// export function fetchPost (url, params) {
//   return new Promise((resolve, reject) => {
//     axios.post(getBaseUrl() + url, params)
//       .then(response => {
//         resolve(response)
//       }, err => {
//         reject(err)
//       })
//       .catch((error) => {
//         reject(error)
//       })
//   })
// }

export function fetchPost (url, params, json = false) {
  // json格式请求头
  const headerJSON = {
    'Content-Type': 'application/json'
  }
  // FormData格式请求头
  const headerFormData = {
    'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
  }
  return new Promise((resolve, reject) => {
    axios
      .post(getBaseUrl() + url, json ? JSON.stringify(params) : qs.stringify(params), {
        headers: json ? headerJSON : headerFormData
      })
      .then(response => {
        resolve(response)
      }, err => {
        reject(err)
      })
      .catch((error) => {
        reject(error)
      })
  })
}
// 返回一个Promise(发送get请求)
export function fetchGet (url, param) {
  return new Promise((resolve, reject) => {
    console.log('getParams', param)
    axios.get(getBaseUrl() + url, { param: param })
      .then(response => {
        console.log('resolve')
        resolve(response)
      }, err => {
        console.log('reject')
        reject(err)
      })
      .catch((error) => {
        reject(error)
      })
  })
}
export default {
  fetchPost,
  fetchGet
}
