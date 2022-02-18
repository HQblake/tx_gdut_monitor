<template>
  <div class="home">
    <div class='headline'>
      <h3 > 主机列表 </h3>
    </div>
    <div class="table-container">
      <el-row>
          <el-col :span="18"><div>&nbsp;</div></el-col>
          <el-col :span="6">
            <div>
              <el-input
              v-model="search"
              size="mini"
              placeholder="输入区域或ip关键字搜索"/>
            </div>
          </el-col>
      </el-row>

      <el-table
        :data="tableData.filter(data => !search
              || data.ip == search.toLowerCase()
              || (data.ip.toString() + data.port.toString()) == search.toLowerCase()
              || data.local.toLowerCase().includes(search.toLowerCase())
              || (data.is_live ? '是' : '否') == search.toLowerCase())
      ">

        <el-table-column
          label="Host"
          align="center"
          header-aligh="center"
          :formatter="formateAgent">
        </el-table-column>
        <el-table-column
          label="区域"
          align="center"
          header-aligh="center"
          prop="local">
        </el-table-column>
        <el-table-column
          :formatter="formateLive"
          prop="is_live"
          width="100"
          align="center"
          header-aligh="center"
          label="是否存活">
        </el-table-column>
        <el-table-column
          align="center"
          header-aligh="center"
          min-width = '200px'
          label="agent管理">
          <template slot-scope="scope">
            <el-button type="primary"
              size="mini"
              @click="handleWarn(scope.row)">告警规则管理</el-button>
            <el-button type="warning"
              size="mini"
              @click="handleSend(scope.row)">发送配置管理</el-button>
            <el-button type="danger"
              size="mini"
              @click="handleHistory(scope.row)">告警历史详情</el-button>
            <el-button type="info"
              size="mini"
              @click="handleShow(scope.row)">数据详情</el-button>
          </template>
        </el-table-column>
      </el-table>

    </div>
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  </div>
</template>

<script>
// // @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import { GetAllAgent } from '@/api/agent'
export default {
  name: 'home',
  methods: {
    deleteData (index, row) {
      this.dialogVisible = false
      this.tableData.splice(index, 1)
    },
    openDelete (index, row) {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.deleteData(index, row)
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    handleEdit (index, row) {
      console.log(index, row)
    },
    handleDelete (index, row) {
      console.log(index, row)
    }
  },
  components: {
    // HelloWorld
  },
  data () {
    return {
      dialogVisible: false,
      search: '',
      tableData: [{
        ip: '127.0.0.1',
        port: '80',
        local: '北京',
        is_live: false
      }, {
        ip: '127.0.0.2',
        port: '80',
        local: '上海',
        is_live: false
      }, {
        ip: '127.0.0.1',
        port: '80',
        local: '北京',
        is_live: true
      }, {
        ip: '127.0.0.1',
        port: '81',
        local: '北京',
        is_live: true
      }, {
        ip: '127.0.0.1',
        port: '80',
        local: '北京',
        is_live: true
      }, {
        ip: '127.0.0.1',
        port: '80',
        local: '北京',
        is_live: true
      }, {
        ip: '127.0.0.1',
        port: '80',
        local: '北京',
        is_live: true
      }],
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
    handleWarn (row) {
      this.$router.push({ path: '/warn/detail', query: { ip: row.ip, local: row.local } })
    },
    handleSend (row) {
      this.$router.push({ path: '/send/detail', query: { ip: row.ip, local: row.local } })
    },
    handleHistory (row) {
      // 待定
      this.$router.push('/send/' + row.ip + '/' + row.local)
    },
    handleShow (row) {
      this.$router.push('/show/' + row.ip + '/' + row.local)
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

<style>
.el-table__header tr,
.el-table__header th {
  padding: 0;
  height: 30px;
  line-height: 30px;
}
.el-table__body tr,
.el-table__body td {
  padding: 0;
  height: 30px;
  line-height: 30px;
}
.el-table {
  padding: 0px;
}
</style>
