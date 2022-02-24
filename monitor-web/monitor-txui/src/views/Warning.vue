<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
-->
<template>
  <div class="warning">
    <div class="content">
      <div class="headline">
        <h3>告警信息</h3>
      </div>
      <div class="search">
        <el-row :gutter="20" type="flex" justify="end">
          <el-col :span="2"><div class="grid-content bg-purple"></div></el-col>
          <el-col :span="3"
            ><div class="grid-content bg-purple">
             <el-input v-model="ip" placeholder="请输入ip" ></el-input></div
          ></el-col>
          <el-col :span="3"
            ><div class="grid-content bg-purple">
              <el-input
                v-model="local"
                placeholder="请输入local"
              ></el-input></div
          ></el-col>
          <el-col :span="3"
            ><div class="grid-content bg-purple">
              <el-select v-model="warnContent" placeholder="告警内容">
                <el-option
                  v-for="item in warnOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                >
                </el-option>
              </el-select></div
          ></el-col>
          <el-col :span="3"
            ><div class="grid-content bg-purple">
              <el-select v-model="level" placeholder="级别">
                <el-option
                  v-for="item in levelOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                >
                </el-option>
              </el-select></div
          ></el-col>
          <el-col :span="8">
            <div class="grid-content bg-purple">
              <span class="demonstration"></span>
              <el-date-picker
                v-model="timePickerValue"
                type="datetimerange"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期"

                value-format="timestamp"

              >
              </el-date-picker>
            </div>
          </el-col>
          <el-col :span="2"
            ><div class="grid-content bg-purple">
              <el-button type="primary" @click="searchWarnInfo">搜索</el-button>
            </div>
          </el-col>
        </el-row>
      </div>
      <!--   .filter(
            (data) =>
              !search ||
              data.ip == search.toLowerCase() ||
              data.local == search.toLowerCase() ||
              data.level == search.toLowerCase() ||
              data.metric.toLowerCase().includes(search.toLowerCase()) ||
              data.value == search.toLowerCase() ||
              data.threshold == search.toLowerCase() ||
              data.duration == search.toLowerCase()
          ) -->
      <el-table
        :data="
          tableData

        "
        :header-cell-style="{ height: '100px' }"
        style="width: 100%"
      >
        <el-table-column
          label="ip"
          prop="ip"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <el-table-column
          label="区域"
          prop="local"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <el-table-column
          label="告警内容"
          prop="metric"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <el-table-column
          label="级别"
          prop="level"
          align="center"
          header-aligh="center"
        >
          <template slot-scope="scope">
            {{scope.row.level | changeLevel}}
          </template>
        </el-table-column>
        <el-table-column
          min-width="100px"
          label="开始时间"
          prop="start"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <!-- <el-table-column label="指标" prop="5" > </el-table-column> -->
        <el-table-column
          label="异常值"
          prop="value"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <el-table-column
          label="阈值"
          prop="threshold"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <el-table-column
          label="持续时间"
          prop="duration"
          align="center"
          header-aligh="center"
        >
        </el-table-column>

        <el-table-column align="center" min-width="100px">
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="输入关键字进行搜索 "
            />
          </template>
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="danger"
              @click="openDelete(scope.$index, scope.row)"
              >Delete</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <!-- 试图添加分页功能 -->
      <!-- <div class="page">
        <el-pagination
          layout="prev, pager, next"
          :total="1000">
        </el-pagination>
      </div> -->
    </div>

  </div>
</template>

<script>
import { GetWarnList, GetWarnInfoWithParams, DelWarnInfo } from '@/api/show'

import dayjs from 'dayjs'

export default {
  name: 'warning',
  components: {},
  filters: {
    changeLevel: (level) => {
      switch (level) {
        case 3:
          return '严重'
        case 2:
          return '中等'
        case 1:
          return '告警'
        case 0:
          return '信息'
      }
    }
  },
  data () {
    return {
      ip: '',
      local: '',
      data1: '',
      timePickerValue: [],
      level: '',
      levelOptions: [
        {
          label: '严重',
          value: '3'
        },
        {
          label: '中等',
          value: '2'
        },
        {
          label: '警告',
          value: '1'
        },
        {
          label: '信息',
          value: '0'
        }
      ],
      warnContent: '',
      warnOptions: [
        {
          label: 'cpu使用率',
          value: 'cpu_rate'
        },
        {
          label: 'mem使用率',
          value: 'mem_rate'
        }
      ],
      Outliers: '',
      threshold: '',
      duration: '',
      dialogVisible: false,
      search: '',
      tableData: [
        // {
        //   id: 1,
        //   ip: '127.0.0.1',
        //   local: '北京',
        //   metric: 'cpu_rate',
        //   level: 1,
        //   start: '2022-02-10 18:00:00',
        //   value: 85,
        //   threshold: 80,
        //   duration: 30
        // },
        // {
        //   id: 2,
        //   ip: '127.0.0.2',
        //   local: '北京',
        //   metric: 'mem_rate',
        //   level: 2,
        //   start: '2022-02-10 19:00:00',
        //   value: 85,
        //   threshold: 80,
        //   duration: 30
        // },
        // {
        //   id: 3,
        //   ip: '127.0.0.2',
        //   local: '上海',
        //   metric: 'mem_rate',
        //   level: 2,
        //   start: '2022-02-10 20:00:00',
        //   value: 85,
        //   threshold: 80,
        //   duration: 30
        // }
      ],
      selector: 30
    }
  },
  methods: {
    searchWarnInfo () {
      console.log('timepicker', this.timePickerValue[0])
      console.log('timepicker', this.timePickerValue[1])
      // let start = this.timePickerValue[0].toLocaleString('chinese', {hour12:false}).split('/').join('-')
      // let end = this.timePickerValue[1].toLocaleString('chinese', {hour12:false}).split('/').join('-')
      let start1 = this.timePickerValue[0]
      let end1 = this.timePickerValue[1]
      let levelInt = Number(this.level)
      // let startFormat = dayjs(start1).format('YYYY-MM-DD HH:mm:ss')
      let startFormat = start1
      // let endFormat = dayjs(end1).format('YYYY-MM-DD HH:mm:ss')
      let endFormat = end1
      console.log('startFormat', startFormat)
      console.log('endFormat', endFormat)
      // console.log(start);
      // console.log(end);
      let start = '2022-02-22 23:26:14'
      let end = '2022-02-22 23:56:14'
      console.log(this.ip, this.local, this.warnContent, levelInt, startFormat, endFormat)
      GetWarnInfoWithParams(this.ip, this.local, this.warnContent, levelInt, startFormat, endFormat).then(data => {
        this.tableData = data.data
      })
        .catch((err) => {
          if (err.msg) {
            this.$alert(err.msg)
            return
          }
          this.$alert(err)
        })
    },
    now () {
      var thisTime = new Date()
      let timeS = new Date(
        thisTime.setMinutes(thisTime.getMinutes() - this.selector)
      )
      this.timePickerValue = [timeS, new Date()]
    },
    deleteData (index, row) {
      this.dialogVisible = false
      console.log('index', index)
      console.log(row.id)

      this.tableData.splice(index, 1)
      DelWarnInfo(row.id)
    },
    openDelete (index, row) {
      this.$confirm('此操作将永久删除该消息, 是否继续?', '提示', {
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
    }
  },
  created () {
    GetWarnList()
      .then((data) => {
        this.tableData = data.data
      })
      .catch((err) => {
        if (err.msg) {
          this.$alert(err.msg)
          return
        }
        this.$alert(err)
      })
  },
  mounted () {
    this.now()
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
.el-select {
  width: 200px;
}
.page{
  float: right;
}
</style>
