<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-11 16:57:09
 * @LastEditors: yzq
-->
<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
-->
<template>
  <div class="warnListChart">
    <PieChart></PieChart>
  </div>
</template>

<script>
// @ is an alias to /src
import PieChart from '@/components/PieChart.vue'

export default {
  name: 'warnListChart',
  components: {
    PieChart
  }
}
</script>
