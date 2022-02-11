/*
 * @Description: 
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
 */
import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import About from './views/About.vue'
import WarnConfig from './views/WarnConfig'
import Warning from './views/Warning'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component: About
    },
    {
      path: '/warnConfig',
      name: 'warnConfig',
      component: WarnConfig
    },
    {
      path: '/warning',
      name: 'warning',
      component: Warning
    },
  ]
})
