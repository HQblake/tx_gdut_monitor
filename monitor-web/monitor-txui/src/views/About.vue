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
