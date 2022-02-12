<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-10 18:32:28
 * @LastEditors: yzq
-->

<template>
  <div class="header_container">

    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      router
    >
     <!-- <div class='logo'>
      <img alt="Vue logo" src="../assets/logo.png">
      </div> -->
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/warnConfig">告警配置</el-menu-item>
      <el-menu-item index="/sendConfig">发送配置</el-menu-item>
      <el-menu-item index="/warning">告警信息</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'HeadTop',
  data () {
    return {
      activeIndex: ''
      // baseImgPath,
    }
  },
  created () {},
  computed: {},
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style >
.header_container {
  height: 60px;
  padding-bottom: 20px;
  font-size: 14px;
}
.el-menu-demo .el-menu--horizontal .el-menu{
  width: 100vw;
}
</style>
