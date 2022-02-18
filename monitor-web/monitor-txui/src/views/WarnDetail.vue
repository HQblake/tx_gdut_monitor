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
              @click="showEdit(scope.row)">编辑配置</el-button>
              <template v-if="scope.row.id != -1">
              <el-button type="danger"
                size="mini"
                @click="handleDel(scope.row)">删除配置</el-button>
              </template>
              <template v-else>
              <el-button type="info" disabled
                size="mini"
                >默认配置</el-button>
              </template>
          </template>
        </el-table-column>
      </el-table>

    </div>
    <el-dialog title="配置编辑" :visible.sync="dialogFormVisible">
  <el-form>
    <el-form-item required>
      <el-col :span="3">主机:</el-col>
      <el-col :span="9"><el-input v-model="ip" disabled></el-input></el-col>
      <el-col :span="3">区域:</el-col>
      <el-col :span="9"><el-input v-model="local" disabled></el-input></el-col>
    </el-form-item>
    <el-form-item required>
      <el-col :span="3">指标类型:</el-col><el-col :span="9"><el-input v-model="metric" disabled></el-input></el-col>
      <el-col :span="3">聚合周期:</el-col><el-col :span="9"><el-input v-model="period"></el-input></el-col>
    </el-form-item>
    <el-form-item required>
       <el-col :span="3">聚合方式:</el-col>
       <el-col :span="21">
         <el-select v-model="method" placeholder="请选择聚合方式">
          <el-option v-for="(item, key) in methodList" :key="key" :label="item" :value="key"></el-option>
        </el-select>
       </el-col>
    </el-form-item>
    <el-form-item v-for="(item, key) in threshold" :key="key">
      <el-col :span="3">{{item.label}}阈值:</el-col><el-col :span="21"><el-input v-model="item.value"></el-input></el-col>
    </el-form-item>
    <p class="tip">设置告警阈值, 留空则不监控</p>
  </el-form>
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false">取 消</el-button>
    <el-button type="primary" @click="handleEdit()">确 定</el-button>
  </div>
</el-dialog>
  </div>

</template>

<script>
import { ParseMethod, CheckMethod, MethodType, MethodMap } from '@/tools/method'
import { ParseObj, CheckThreshold, StringObj } from '@/tools/level'
import { GetAgentRule, UpdateRule, DelRule } from '@/api/judgment'
export default {
  name: 'warnDetail',
  components: {
  },
  data () {
    return {
      id: -1,
      ip: '',
      local: '',
      metric: '',
      method: 1,
      period: '5m',
      threshold: {},
      dialogFormVisible: false,
      tableData: [
        {
          id: -1,
          ip: '127.0.0.1',
          local: '北京',
          metric: 'cpu利用率',
          method: 1,
          period: '5m',
          threshold: `{"0":0.01, "1": 0.02}`
        },
        {
          id: 11,
          ip: '127.0.0.1',
          local: '广州',
          metric: 'mem',
          method: 2,
          period: '5m',
          threshold: `{}`
        },
        {
          id: 12,
          ip: '127.0.0.1',
          local: '深圳',
          metric: 'rate',
          method: 3,
          period: '5m',
          threshold: `{}`
        },
        {
          id: 13,
          ip: '127.0.0.1',
          local: '上海',
          metric: 'cpu利用率',
          method: 0,
          period: '5m',
          threshold: `{"1":0.01, "2": 0.02}`
        },
        {
          id: 14,
          ip: '127.0.0.1',
          local: '北京',
          metric: 'cpu利用率',
          method: 1,
          period: '5m',
          threshold: `{"3":0.01}`
        }
      ],
      search: '',
      methodList: MethodType
    }
  },

  created () {
    this.ip = this.$route.query.ip
    this.local = this.$route.query.local
    if (!this.ip || !this.local) {
      this.$router.push({ path: '/' })
    }

    this.metric = ''
    this.id = -1
    this.method = 1
    this.period = '5m'
    this.threshold = {}
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
      if (!this.ip || !this.local || !this.metric || !this.period) {
        this.$alert('操作有误, 请重试~')
        return false
      }
      // 检查聚合方式
      if (!CheckMethod(this.method)) {
        this.$alert('聚合方式有误, 请重试~')
        return false
      }
      // 检查告警类型
      if (!CheckThreshold(this.threshold)) {
        this.$alert('至少需要设置一个告警阈值, 请重试~')
        return false
      }
      return true
    },
    showEdit (row) {
      this.metric = row.metric
      this.method = ParseMethod(row.method)
      this.period = row.period
      this.id = row.id
      // 清空告警类型，方便判断整合
      this.threshold = ParseObj(row.threshold)
      // 显示可编辑表单
      this.dialogFormVisible = true
    },
    handleEdit () {
      if (!CheckMethod(this.method)) {
        this.method = MethodMap[this.method]
      }
      if (!this.check()) {
        return
      }
      UpdateRule(this.ip, this.local, this.id, this.method, this.metric, this.period, StringObj(this.threshold))
        .then(data => {
          this.$alert('更新成功')
          this.$router.go(0)
        })
        .catch(err => {
          if (err.msg) {
            this.$alert(err.msg)
          } else {
            this.$alert(err)
          }
        })
    },
    handleDel (row) {
      DelRule(this.ip, this.local, row.id)
        .then(data => {
          this.$alert('删除成功')
          this.$router.go(0)
        })
        .catch(err => {
          if (err.msg) {
            this.$alert(err.msg)
          } else {
            this.$alert(err)
          }
        })
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
.el-select {
  width: 100%;
}
.tip {
  font-size: 8px;
  text-align: right;
}
</style>
