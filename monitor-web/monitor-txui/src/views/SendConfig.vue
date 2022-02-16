<template>
  <div>
    <div class='headline'>
      <h3 > 发送配置 </h3>
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
        :formatter="formateSend"
          prop="send"
          label="配置类型">
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
import { GetAllSendAgent } from '@/api/agent'
import { ParseSendtype } from '@/tools/sendType'
export default {
  name: 'sendConfig',
  components: {
  },
  data () {
    return {
      tableData: [{
        ip: '127.0.0.1',
        port: '2016-05-03',
        local: '北京',
        is_live: true,
        send: [1, 2]
      }, {
        ip: '127.0.0.2',
        port: '2016-05-03',
        local: '上海',
        is_live: false,
        send: [1, 2]
      }, {
        ip: '127.0.0.1',
        port: '2016-05-03',
        local: '北京',
        is_live: true,
        send: [1, 2]
      }, {
        ip: '127.0.0.1',
        port: '2016-05-03',
        local: '北京',
        is_live: true,
        send: [1, 2]
      }, {
        ip: '127.0.0.1',
        port: '2016-05-03',
        local: '北京',
        is_live: true,
        send: [1, 2]
      }, {
        ip: '127.0.0.1',
        port: '2016-05-03',
        local: '北京',
        is_live: true,
        send: [1, 2]
      }, {
        ip: '127.0.0.1',
        port: '2016-05-03',
        local: '北京',
        is_live: true,
        send: [1, 2]
      }],
      search: ''
    }
  },
  created () {
    GetAllSendAgent().then(data => {
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
      this.$router.push({ path:'/send/detail', query: {ip: row.ip, local: row.local} });
    },
    formateSend (row, column, cellValue) {
      var arr = []
      for (var j = 0; j < cellValue.length; j++) {
        arr.push(ParseSendtype(cellValue[j]))
      }
      return arr.join(' , ').toString()
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
