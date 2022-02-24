<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-08 11:20:37
 * @LastEditors: zeke
-->
<template>
  <div>
    <div class="nav">
      <div class="subNav">

        <el-select v-model="time" placeholder="时长" class='select'>
          <el-option
            v-for="item in timeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
         <el-select v-model="method" placeholder="聚合方式" class='select'>
          <el-option
            v-for="item in methodOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
        <div class="block">
          <span class="demonstration"></span>
          <el-date-picker
            v-model="timePickerValue"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
          >
          </el-date-picker>
        </div>
        <el-row>
          <el-button type="primary" @click="search">
            <div class="icon">
              搜索
            </div>
          </el-button>
        </el-row>
      </div>
    </div>

    <div class="single-chart">
      <div :name="name" id="main" class="echart"></div>
    </div>
    <!-- {{ip}} -- {{local}} -->
  </div>
</template>

<script>

import { GetMetricsWithTime } from '@/api/show'
import dayjs from 'dayjs'
// import rateChart from '../global'
// import { markRaw } from 'vue'

export default {
  name: 'SingleChart',

  data () {
    return {
      // list: [],
      ip: '',
      local: '',
      time: '1s',
      timeOptions: [
        {
          value: '1s',
          label: '1s'
        },
        {
          value: '30m',
          label: '30m'
        },
        {
          value: '60m',
          label: '60m'
        },
        {
          value: '6h',
          label: '6h'
        },
        {
          value: '12h',
          label: '12h'
        },
        {
          value: '24h',
          label: '24h'
        }
      ],
      method: -1,
      methodOptions: [
        {
          label: '不聚合',
          value: -1
        },
        {
          label: '总和',
          value: 0
        },
        {
          label: '平均值',
          value: 1
        },
        {
          label: '中位数',
          value: 2
        },
        {
          label: '积分',
          value: 3
        },
        {
          label: '极值',
          value: 4
        },
        {
          label: '标准差',
          value: 5
        },
        {
          label: '最大值',
          value: 6
        },
        {
          label: '最小值',
          value: 7
        }

      ],
      selector: '30',
      timePickerValue: [],
      // value1: [new Date(), new Date()],
      elId: null,
      name: 1,
      chartTitle: null,
      noDataTip: false,
      config: '',
      chart: '',
      cpu: [],
      mem: [],
      timeX: [],
      thisTime: Date(),
      myEchart: null
    }
  },
  computed: {

  },
  props: ['data1'],
  created () {
    this.ip = this.$route.params.ip
    this.local = this.$route.params.local
  },
  watch: {
  },
  destroyed () {
    clearInterval(this.interval)
  },
  methods: {
    parseTime () {

    },
    getCpu () {
      for (let i = 0; i < this.runtime.length; i++) {
        this.cpu.push(this.runtime[i][0])
      }
    },
    getMem () {
      for (let i = 0; i < this.runtime.length; i++) {
        this.mem.push(this.runtime[i][1])
      }
    },
    chartInit () {
      let c = 'main'
      let dom = document.getElementById(c)

      console.log('echart', this.$echarts.myEchart)
      this.myEchart = this.$echarts.init(dom)
    },
    drawChart () {
      var timeX = []
      const now = new Date()
      var timeGap
      switch (this.time) {
        case '1s':
          timeGap = 1
          break
        case '30m':
          timeGap = 1800
          break
        case '60m':
          timeGap = 3600
          break
        case '6h':
          timeGap = 21600
          break
        case '12h':
          timeGap = 43200
          break
        case '24h':
          timeGap = 86400
          break
      }

      let option = {
        title: {
          text: 'runtime'
        },
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          data: ['Cpu', 'Mem']
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        toolbox: {
          feature: {
            saveAsImage: {},
            restore: {}
          }
        },
        xAxis: {
          type: 'category',

          boundaryGap: false,
          axisLine: {
            lineStyle: {
              color: '#a1a1a2'
            }
          },
          splitLine: {
            show: true,
            lineStyle: {
              color: ['#a1a1a2'],
              width: 1,
              type: 'solid'
            }
          },
          data: this.timeX
        },
        yAxis: [
          {
            type: 'value',
            name: 'Cpu',
            axisLine: {
              lineStyle: {
                color: '#a1a1a2'
              }
            },
            splitLine: {
              show: true,
              lineStyle: {
                color: ['#a1a1a2'],
                width: 1,
                type: 'solid'
              }
            }
          },
          {
            type: 'value',
            name: 'Mem',
            axisLabel: {
              formatter: '{value} %'
            }
          }
        ],
        series: [
          {
            type: 'line',
            name: 'Cpu',
            showSymbol: false,
            data: this.cpu
          },
          {
            type: 'line',
            name: 'Mem',
            showSymbol: false,
            data: this.mem
          }
        ]
      }
      this.myEchart.setOption(option)
      console.log('drawChart')
    },
    search () {
      if (!this.timePickerValue) {
        this.$alert('请选择合适的起始时间~')
        return
      }
      if (this.timePickerValue.length < 2) {
        this.$alert('请选择合适的起始时间~')
        return
      }

      // let startFormat = dayjs(start1).format('YYYY-MM-DD HH:mm:ss')
      let startFormat = Math.round(this.timePickerValue[0] / 1000)
      // let endFormat = dayjs(end).format('YYYY-MM-DD HH:mm:ss')
      let endFormat = Math.round(this.timePickerValue[1] / 1000)

      var timeGap
      switch (this.time) {
        case '1s':
          timeGap = '1s'
          break
        case '30m':
          timeGap = '1800s'
          break
        case '60m':
          timeGap = '3600s'
          break
        case '6h':
          timeGap = '21600s'
          break
        case '12h':
          timeGap = '43200s'
          break
        case '24h':
          timeGap = '86400s'
          break
      }

      let cpu_metric = GetMetricsWithTime(this.ip, this.local, 'cpu_rate', startFormat, endFormat, timeGap, this.method, -1)
      let mem_metric = GetMetricsWithTime(this.ip, this.local, 'mem_rate', startFormat, endFormat, timeGap, this.method, -1)

      cpu_metric.then((data) => {
        // 处理cpu数据
        this.cpu = []
        this.timeX = []
        for (let i = 0; i < data.data.length; i++) {
          this.cpu.push(data.data[i].value)
          let timeStamp = dayjs.unix(data.data[i].timestamp).format('MM-DD hh:mm:ss')
          this.timeX.push(timeStamp)
        }

        mem_metric.then((data) => {
          this.mem = []
          console.log('object', data.data)
          for (let i = 0; i < data.data.length; i++) {
            this.mem.push(data.data[i].value)
          }
          this.drawChart()
        })
          .catch((err) => {
            if (err.msg) {
              this.$alert(err.msg)
              return
            }
            this.$alert(err)
          })
      })
        .catch((err) => {
          if (err.msg) {
            this.$alert(err.msg)
            return
          }
          this.$alert(err)
        })

      console.log('drawChart')
    },
    now () {
      var thisTime = new Date()
      let timeS = new Date(thisTime.setMinutes(thisTime.getMinutes() - this.selector))
      this.timePickerValue = [timeS, new Date()]
    }

  },
  mounted () {
    this.now()
    this.chartInit()
  },
  components: {}
}
</script>

<style scoped>
#main {
  width: 700px;
  height: 500px;
  margin: auto;
}
.nav {
  width: 100vw;
  height: 100px;
}
.subNav {
  display: flex;
  float: left;
  width: 100%;
  justify-content: center;
}
.block{
  margin-left: 10px;
  margin-right: 10px;
}
.select{
  width: 200px;
  padding-left: 10px;
}
</style>
