<template>
  <div>
    <div class='headline'>
      <h3 > 告警配置 </h3>
    </div>
    <div class="table-container">
  <el-table
    :data="tableData.filter(data => !search || (data.local.toLowerCase().includes(search.toLowerCase()) || data.ip.toLowerCase().includes(search.toLowerCase())))" fit>
    <el-table-column
      label="Host"
      :formatter="formateAgent">
    </el-table-column>
    <el-table-column
      prop="local"
      label="区域">
    </el-table-column>
    <el-table-column
    :formatter="formateMetric"
      prop="metric"
      label="监控指标">
    </el-table-column>
    <el-table-column
      :formatter="formateLive"
      prop="is_live"
      width="100"
      label="是否存活">
    </el-table-column>
    <el-table-column
      align="center">
      <template slot="header" slot-scope="scope">
        <el-input
          v-model="search"
          size="mini"
          placeholder="输入区域或ip关键字搜索"/>
      </template>
      <template slot-scope="scope">
        <el-button type="success"
          size="mini"
          @click="handleEdit(scope.row)">Edit</el-button>
      </template>
    </el-table-column>
  </el-table>

  </div>
  </div>

</template>

<script>
import { GetAllAgent } from '@/api/agent'
export default {
  name: 'warnConfig',
  components: {
  },
  data () {
    return {
      tableData: [
        //   {
        //   ip: '127.0.0.1',
        //   port: '2016-05-03',
        //   local: '北京',
        //   is_live: true,
        //   metric: ['cpu利用率', 'mem']
        // }, {
        //   ip: '127.0.0.2',
        //   port: '2016-05-03',
        //   local: '上海',
        //   is_live: false,
        //   metric: ['cpu利用率', 'mem']
        // }, {
        //   ip: '127.0.0.1',
        //   port: '2016-05-03',
        //   local: '北京',
        //   is_live: true,
        //   metric: ['cpu利用率', 'mem']
        // }, {
        //   ip: '127.0.0.1',
        //   port: '2016-05-03',
        //   local: '北京',
        //   is_live: true,
        //   metric: ['cpu利用率', 'mem']
        // }, {
        //   ip: '127.0.0.1',
        //   port: '2016-05-03',
        //   local: '北京',
        //   is_live: true,
        //   metric: ['cpu利用率', 'mem']
        // }, {
        //   ip: '127.0.0.1',
        //   port: '2016-05-03',
        //   local: '北京',
        //   is_live: true,
        //   metric: ['cpu利用率', 'mem']
        // }, {
        //   ip: '127.0.0.1',
        //   port: '2016-05-03',
        //   local: '北京',
        //   is_live: true,
        //   metric: ['cpu利用率', 'mem']
        // }
      ],
      search: ''
    }
  },
  created () {
    GetAllAgent().then(data => {
      this.tableData = data.data
    })
      .catch(err => {
        if (err.msg) {
          this.$alert(err.msg)
          return
        }
        this.$alert(err)
      })
  },
  methods: {
    handleEdit (row) {
      // this.$router.push({ path: '/warn/detail', query: { ip: row.ip, local: row.local } })
      this.$router.push('/warn/' + row.ip + '/' + row.local)
    },
    formateMetric (row, column, cellValue) {
      if (cellValue) {
        return cellValue.join(' , ').toString()
      }
      return ''
    },
    formateAgent (row, column, cellValue) {
      if (row.port) {
        return row.ip + ':' + row.port
      }
      return row.ip
    },
    formateLive (row, column, cellValue) {
      if (cellValue) {
        return '是'
      }
      return '否'
    }

  }
}
</script>
<style lang='less'>
.headline {
  padding: 10px 50px;
}
.table-container{
  padding: 0 50px;
  margin: 0 auto;
}
</style>
