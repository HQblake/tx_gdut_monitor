/*
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
 */
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import * as echarts from 'echarts'
import './plugins/element.js'
import less from 'less'
import axios from 'axios'
import QS from 'qs'

Vue.prototype.$axios = axios
Vue.prototype.qs = QS
Vue.use(less)
// import ElementUI from 'element-ui';
// import 'element-ui/lib/theme-chalk/index.css';

Vue.config.productionTip = false
Vue.prototype.$echarts = echarts
// Vue.use(ElementUI)

axios.get('/static/config.json').then((res) => {
  // 基础地址
  Vue.prototype.BASE_URL = res.URL
  new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount('#app')
}).catch((err)=>{
  console.log('err:',err)
})

