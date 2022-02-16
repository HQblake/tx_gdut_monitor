<!--
 * @Description: 
 * @Autor: yzq
 * @Date: 2022-02-16 14:28:12
 * @LastEditors: yzq
-->
<template>
  <div class="about">
    <div class='headline'>
      <h3 > {{ip}}-{{local}} </h3>
    </div>
    <SingleChart></SingleChart>
  </div>
</template>

<script>
// @ is an alias to /src
import SingleChart from '@/components/SingleChart.vue'

export default {
  name: 'show',
  created() {
     this.ip = this.$route.params.ip
     this.local = this.$route.params.local
  },
  components: {
    SingleChart
  }
}
</script>

<style lang="less">

  
</style>