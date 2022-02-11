<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
-->
<template>
  <div id="app">
    <head-top></head-top>
    <div>
      <!-- <router-link to='/'>HelloWorld</router-link>
      <router-link to='/about'>About</router-link> -->
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import HeadTop from './components/headTop.vue'
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
    HelloWorld,
    HeadTop
  }
}
</script>

<style lang='less'>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
}
* {
  margin: 0px;
  padding: 0px;
  list-style: none;
}
.headline {
  text-align: left;
  padding-left: 20px;
  padding-bottom: 20px;
}
.headline h3 {
  font-size: 25px;
}
.container {
  width: 100vw;
  height: 100vh;
  padding: 0px;
}

</style>
