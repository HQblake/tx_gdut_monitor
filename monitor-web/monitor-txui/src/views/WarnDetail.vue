<template>
  <div class="container">
    <div class='headline'>
      <h3 > {{ip}} - {{local}} 判定配置详情 </h3>
    </div>
    <div class="table-container">
      <el-row>
          <el-col :span="18"><div>&nbsp;</div></el-col>
          <el-col :span="6">
            <div>
              <el-input
              v-model="search"
              size="mini"
              placeholder="输入关键字搜索"/>
            </div>
          </el-col>
      </el-row>

      <el-table
        :data="tableData.filter(data => !search || (data.metric.toLowerCase().includes(search.toLowerCase()) || data.period.toLowerCase().includes(search.toLowerCase()) || data.threshold.toLowerCase().includes(search.toLowerCase())))" fit>
        <el-table-column
          label="指标类型"
          prop="metric">
        </el-table-column>
        <el-table-column
          :formatter="formateMethod"
          prop="method"
          label="聚合方式">
        </el-table-column>
        <el-table-column
          prop="period"
          label="聚合周期">
        </el-table-column>
        <el-table-column
          align="center"
          label="管理">
          <template slot-scope="scope">
            <el-button type="primary"
              size="mini"
              @click="handleEdit(scope.row)">编辑</el-button>
              <template v-if="scope.row.id != -1">
              <el-button type="warning"
                size="mini"
                @click="handleDel(scope.row)">删除</el-button>
              </template>
          </template>
        </el-table-column>
      </el-table>

    </div>
  </div>

</template>

<script>
import { ParseMethod, CheckMethod } from '@/tools/method'
import { GetAgentRule, UpdateRule, DelRule } from '@/api/judgement'
export default {
  name: 'warnDetail',
  components: {
  },
  data () {
    return {
      ip: '',
      local: '',
      metric: '',
      method: 1,
      period: '5m',
      warn: undefined,
      info: undefined,
      error: undefined,
      panic: undefined,
      tableData: [
        {
          id: -1,
          ip: '127.0.0.1',
          local: '北京',
          metric: 'cpu利用率',
          method: 1,
          period: '5m',
          threshold: "{'warn':0.01, 'panic': 0.02}"
        },
        {
          id: 11,
          ip: '127.0.0.1',
          local: '广州',
          metric: 'mem',
          method: 2,
          period: '5m',
          threshold: "{'warn':0.01, 'panic': 0.02}"
        },
        {
          id: 12,
          ip: '127.0.0.1',
          local: '深圳',
          metric: 'rate',
          method: 3,
          period: '5m',
          threshold: "{'warn':0.01, 'panic': 0.02}"
        },
        {
          id: 13,
          ip: '127.0.0.1',
          local: '上海',
          metric: 'cpu利用率',
          method: 0,
          period: '5m',
          threshold: "{'warn':0.01, 'panic': 0.02}"
        },
        {
          id: 14,
          ip: '127.0.0.1',
          local: '北京',
          metric: 'cpu利用率',
          method: 1,
          period: '5m',
          threshold: "{'warn':0.01, 'panic': 0.02}"
        }
      ],
      search: ''
    }
  },

  created () {
    this.ip = this.$route.params.ip
    this.local = this.$route.params.local
    this.metric = ''
    this.id = -1
    this.method = 1
    this.period = '5m'
    this.warn = undefined
    this.info = undefined
    this.error = undefined
    this.panic = undefined
    GetAgentRule(this.ip, this.local)
      .then(data => {
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
    check () {
      // 检查必要的信息
      if (this.ip == '' || this.local == '' || this.metric == '' || this.period == '') {
        this.$alert('操作有误, 请重试~')
        return false
      }
      // 检查聚合方式
      if (!CheckMethod(this.method)) {
        this.$alert('聚合方式有误, 请重试~')
        return false
      }
      // 检查告警类型
      if(!this.warn && !this.panic && !this.info && !this.error) {
        this.$alert('至少需要设置一个告警阈值, 请重试~')
        return false
      }
      return true
    },
    showEdit (row) {
      this.metric = row.metric
      this.method = row.method
      this.period = row.period
      this.id = row.id
      // 清空告警类型，方便判断整合
      this.warn = undefined
      this.info = undefined
      this.error = undefined
      this.panic = undefined
      // 显示可编辑表单
    },
    handleEdit (row) {
      if (!this.check()) {
        return
      }
      UpdateRule(this.ip, this.local, this.id)
        .then(data => {
          this.$alert('删除成功')
        })
        .catch(err => {
          if (err.msg) {
            this.$alert(err.msg)
          } else {
            this.$alert(err)
          }
        })
      this.$router.go(0)
    },
    handleDel (row) {
      DelRule(this.ip, this.local, row.id)
        .then(data => {
          this.$alert('删除成功')
        })
        .catch(err => {
          if (err.msg) {
            this.$alert(err.msg)
          } else {
            this.$alert(err)
          }
        })
      this.$router.go(0)
    },
    formateMethod (row, column, cellValue) {
      return ParseMethod(cellValue)
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
