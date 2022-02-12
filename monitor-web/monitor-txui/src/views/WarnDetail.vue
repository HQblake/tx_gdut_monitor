<template>
  <div class="container">
    <div class='headline'>
      <h3 > {{ip}} - {{local}} 判定配置详情 </h3>
    </div>
  </div>
</template>

<script>

export default {
  name: 'warnDetail',
  components: {
  },
  data () {
    return {
      test: 'test',
      ip: '',
      local: ''
    }
  },
  created () {
    this.ip = this.$route.params.ip
    this.local = this.$route.params.local
  },
  methods: {
  }
}
</script>
<style lang='less'>
</style>
