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
import SendConfig from './views/SendConfig'
import Warning from './views/Warning'
import warnDetail from './views/WarnDetail'
import sendDetail from './views/SendDetail'
import showDetail from './views/Show'

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
      path: '/warn',
      name: 'warnConfig',
      component: WarnConfig
    },
    {
      path: '/send',
      name: 'sendConfig',
      component: SendConfig
    },
    {
      path: '/warning',
      name: 'warning',
      component: Warning
    },
    {
      path: '/warn/:ip/:local',
      name: 'warnDetail',
      component: warnDetail
    },
    {
      path: '/send/:ip/:local',
      name: 'sendDetail',
      component: sendDetail
    },
    {
      path: '/show/:ip/:local',
      name: 'showDetail',
      component: showDetail
    }
  ]
})
