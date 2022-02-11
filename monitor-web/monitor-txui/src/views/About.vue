<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
-->
<template>
  <div class="about">
    <SingleChart></SingleChart>
  </div>
</template>

<script>
// @ is an alias to /src
import SingleChart from '@/components/SingleChart.vue'

export default {
  name: 'about',
  components: {
    SingleChart
  }
}
</script>
