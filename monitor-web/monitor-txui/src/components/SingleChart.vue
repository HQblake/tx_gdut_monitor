<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-08 11:20:37
 * @LastEditors: yzq
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
            value-format="yyyy-MM-dd HH:mm:ss"
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

export default {
  name: 'SingleChart',
  
  data () {

    return {
      // list: [],
      ip: '',
      local: '',
      time: '',
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
          value: '5h',
          label: '5h'
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
      method: '',
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
        },
        
      ],
      selector: '30',
      timePickerValue: [],
      // value1: [new Date(), new Date()],
      elId: null,
      name: 1,
      chartTitle: null,
      noDataTip: false,
      config: '',
      myChart: '',
      runtime: [
        [30, 40],
        [40, 50],
        [30, 40],
        [45, 60],
        [50, 40],
        [30, 40],
        [40, 50],
        [30, 40],
        [45, 60],
        [80, 40],
        [30, 40],
        [45, 60],
        [50, 40],
        [30, 40],
        [45, 60],
        [50, 40],
        [30, 40],
        [45, 60],
        [50, 40],
      ],
      thisTime: Date()
    }
  },
  computed: {
    cpu: function () {
      let cpu = []
      for (let i = 0; i < this.runtime.length; i++) {
        cpu.push(this.runtime[i][0])
      }
      return cpu
    },
    mem: function () {
      let mem = []
      for (let i = 0; i < this.runtime.length; i++) {
        mem.push(this.runtime[i][1])
      }
      return mem
    }
  },
  props: ['data1'],
  created () {
    this.ip = this.$route.params.ip
    this.local = this.$route.params.local
  },
  watch: {
    cpu (val) {
      this.drawChart()
    },
    mem (val) {
      this.drawChart()
    },
  },
  destroyed () {
    clearInterval(this.interval)
  },
  methods: {
    parseTime() {

    },
    drawChart () {
      var timeX = []
      const now = new Date()
      // this.cpu.push(20)
      // console.log(timeS);
      // value = value + Math.random() * 21; // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/random
      // console.log(cpu[index])
      for (let i = 0; i < this.cpu.length; i++) {
        let timeS = new Date(now.setMinutes(now.getMinutes() - 1))
        var curr_date = timeS.getDate()
        var curr_month = timeS.getMonth() + 1
        // var curr_year = timeS.getFullYear();
        String(curr_month).length < 2
          ? (curr_month = '0' + curr_month)
          : curr_month
        String(curr_date).length < 2
          ? (curr_date = '0' + curr_date)
          : curr_date
        var yyyyMMdd = curr_month + '-' + curr_date
        var curr_hour = timeS.getHours()
        if (curr_hour < 10) {
          curr_hour = '0' + curr_hour
        }
        var curr_min = timeS.getMinutes()
        if (curr_min < 10) {
          curr_min = '0' + curr_min
        }
        var curr_sec = timeS.getSeconds()
        if (curr_sec < 10) {
          curr_sec = '0' + curr_sec
        }
        var tmp = yyyyMMdd + '\n' + curr_hour + ':' + curr_min + ':' + curr_sec
        
        timeX.unshift(tmp)
      }
      
      const echarts = require('echarts/lib/echarts')
      let c = 'main'
      let myEchart = this.$echarts.init(document.getElementById(c))

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
          data: timeX
        },
        yAxis: [
          {
            type: 'value',
            name: 'Cpu',
            // show: true,
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
      myEchart.setOption(option)
    },
    search () {
      
      console.log(this.time)
      console.log(this.method)
      let start = dayjs().format("YYYY-MM-DD HH:mm:ss")
      console.log('dayjs',start);
      this.timePickerValue[0].toLocaleString('chinese', {hour12:false}).split('/').join('-')
      let end = this.timePickerValue[1].toLocaleString('chinese', {hour12:false}).split('/').join('-')
      console.log(start);
      console.log(end);
      console.log(this.ip);
      console.log(this.local);
      let cpu_metric = GetMetricsWithTime(this.ip, this.local, 'cpu_rate', start, end, this.method, -1)
      let mem_metric = GetMetricsWithTime(this.ip, this.local, 'mem_rate', start, end, this.method, -1)
      if (cpu_metric.length != mem_metric.length) {
        this.runtime = []
      }
      for (let i=0; i<mem_metric.length; i++) {
        this.runtime.push([cpu_metric[0], mem_metric[0]])
      }
      // ip, local, metric, begin, end, method, limit
      // let time = this.value1[0].toString()
      // console.log(time)
    },
    now () {
      var thisTime = new Date()
      let timeS = new Date(thisTime.setMinutes(thisTime.getMinutes() - this.selector))
      this.timePickerValue = [timeS, new Date()]
    },

  },
  mounted () {
    this.drawChart()
    this.now()
    // this.testProps()
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
