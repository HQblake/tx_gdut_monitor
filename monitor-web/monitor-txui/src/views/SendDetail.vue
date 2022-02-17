<template>
  <div class="container">
    <div class='headline'>
      <h3 > {{ip}} - {{local}} 发送配置详情 </h3>
    </div>
    <div class="table-container">
      <el-row>
          <el-col :span="1">
            <el-button type="primary"
              size="mini"
              @click="showAdd()">新增配置</el-button>
          </el-col>
          <el-col :span="17"><div>&nbsp;</div></el-col>
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
        :data="tableData.filter(data => !search || (data.ip.toLowerCase().includes(search.toLowerCase()) || data.local.toLowerCase().includes(search.toLowerCase()) || data.config.toLowerCase().includes(search.toLowerCase())))" fit>
        <el-table-column
        :formatter="formateLevel"
          label="告警等级"
          prop="level">
        </el-table-column>
        <el-table-column
          :formatter="formateSendType"
          prop="send_type"
          label="告警类型">
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
      <el-col :span="3">告警等级:</el-col>
      <el-col :span="9">
        <el-select v-model="level" placeholder="请选择告警等级">
          <el-option v-for="(item, key) in levelList" :key="key" :label="item" :value="key"></el-option>
        </el-select>
      </el-col>
      <el-col :span="3">告警类型:</el-col>
      <el-col :span="9">
          <el-select v-model="sendType" @change="TypeChange" placeholder="请选择告警类型">
            <el-option v-for="(item, key) in sendList" :key="key" :label="item" :value="key"></el-option>
          </el-select>
      </el-col>
    </el-form-item>
    <el-form-item v-for="(item, key) in config" :key="key">
      <el-col :span="3">{{item.label}}:</el-col>
      <el-col :span="21">

        <el-select v-if="item.selectList" v-model="item.value" placeholder="请选择">
            <el-option v-for="it in item.selectList" :key="it" :label="it" :value="it"></el-option>
        </el-select>
        <el-input v-else v-model="item.value"></el-input>
      </el-col>
      <p v-if="item.tip" class="tip">{{item.tip}}</p>
    </el-form-item>

  </el-form>
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false">取 消</el-button>
    <el-button type="primary" @click="handleEdit()">确 定</el-button>
  </div>
</el-dialog>
  </div>

</template>

<script>

import { GetSendConfigs, AddSendConfig, UpdateSendConfig, DelSendConfig } from '@/api/send'
import { SendMap, SendType, ParseSendtype, CheckType } from '@/tools/sendType'
import { ParseObj, defaultConfig, StringObj, CheckConfig } from '@/tools/config'
import { LevelMap, LevelType, ParseLevel, CheckLevel } from '@/tools/level'
export default {
  name: 'sendDetail',
  components: {
  },
  data () {
    return {
      id: -1,
      ip: '',
      local: '',
      sendType: '',
      level: '',
      config: {},
      sendList: SendType,
      levelList: LevelType,
      search: '',
      dialogFormVisible: false,
      tableData: [
        {
          id: 1,
          ip: '127.0.0.1',
          local: '北京',
          send_type: 0,
          level: 0,
          config: `{"target": "526756656@qq.com,123456789@163.com", "format_type": "html"}`
        },
        {
          id: 11,
          ip: '127.0.0.1',
          local: '广州',
          send_type: 1,
          level: 1,
          config: `{"topic": "test", "format_type": "line", "address": "127.0.0.1:8554,127.0.0.2:7894", "version": "0.2.0.1", "partition_type": "hash","partition":0, "partition_key": "test"}`
        },
        {
          id: 12,
          ip: '127.0.0.1',
          local: '深圳',
          send_type: 2,
          level: 2,
          config: `{"topic": "test_nsq", "format_type": "json", "address": "127.0.0.1:1234"}`
        },
        {
          id: 13,
          ip: '127.0.0.1',
          local: '上海',
          send_type: 0,
          level: 3,
          config: `{"target": "526756656@qq.com", "format_type": "line"}`
        },
        {
          id: 14,
          ip: '127.0.0.1',
          local: '北京',
          send_type: 3,
          level: 0,
          config: `{"url": "http://127.0.0.1/test", "format_type": "json", "method": "GET"}`
        }
      ]
    }
  },

  created () {
    this.ip = this.$route.query.ip
    this.local = this.$route.query.local
    if(!this.ip || !this.local){
      this.$router.push({ path:'/'});
    }
    this.id = -1
    this.sendType = ParseSendtype(1)
    this.level = ParseLevel(1)
    this.config = {}
    GetSendConfigs(this.ip, this.local)
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
      if (!this.ip || !this.local || !this.config) {
        this.$alert('操作有误, 请重试~')
        return false
      }
      // 检查告警类型
      if (!CheckType(this.sendType)) {
        this.$alert('告警类型有误, 请重试~')
        return false
      }
      // 检查告警等级
      if (!CheckLevel(this.level)) {
        this.$alert('告警类型有误, 请重试~')
        return false
      }
      let checkRes = CheckConfig(this.sendType, this.config)
      if (!checkRes.check) {
        this.$alert(checkRes.msg)
        return false
      }
      return true
    },
    formateSendType (row, column, cellValue) {
      return ParseSendtype(cellValue)
    },
    formateLevel (row, column, cellValue) {
      return ParseLevel(cellValue)
    },
    TypeChange (data) {
      this.config = defaultConfig[data]
    },
    showAdd () {
      this.level = ParseLevel(0)
      this.sendType = ParseSendtype(0)
      this.id = -1
      this.config = defaultConfig[0]
      this.dialogFormVisible = true
    },
    showEdit (row) {
      this.level = ParseLevel(row.level)
      this.sendType = ParseSendtype(row.send_type)
      this.id = row.id
      // 清空告警类型，方便判断整合
      this.config = ParseObj(row.send_type, row.config)
      // 页面显示
      this.dialogFormVisible = true
    },
    handleEdit () {
      if (!CheckType(this.sendType)) {
        this.sendType = SendMap[this.sendType]
      }
      if (!CheckLevel(this.level)) {
        this.level = LevelMap[this.level]
      }
      if (!this.check()) {
        return
      }
      if (this.id == -1) {
        // 新增
        AddSendConfig(this.ip, this.local, this.sendType, this.level, StringObj(this.config))
          .then(data => {
            this.$router.go(0)
          })
          .catch(err => {
            if (err.msg) {
              this.$alert(err.msg)
            } else {
              this.$alert(err)
            }
            this.$router.go(0)
          })
      } else {
        UpdateSendConfig(this.ip, this.local, this.id, this.sendType, this.level, StringObj(this.config))
          .then(data => {
            this.$router.go(0)
          })
          .catch(err => {
            if (err.msg) {
              this.$alert(err.msg)
            } else {
              this.$alert(err)
            }
            this.$router.go(0)
          })
      }
    },
    handleDel (row) {
      DelSendConfig(this.ip, this.local, row.id)
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
          this.$router.go(0)
        })
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
    position: absolute;
    bottom: -30px;
    right: 0px;
}
</style>
