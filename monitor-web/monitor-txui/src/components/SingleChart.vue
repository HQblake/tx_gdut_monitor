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
      chart: '',
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
      cpu:[],
      mem:[],
      thisTime: Date(),
      myEchart:null
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
    parseTime() {

    },
    getCpu(){
      for (let i = 0; i < this.runtime.length; i++) {
        this.cpu.push(this.runtime[i][0])
      }
    },
    getMem(){
      for (let i = 0; i < this.runtime.length; i++) {
        this.mem.push(this.runtime[i][1])
      }
    },
    chartInit() {
      let c = 'main'
      let dom = document.getElementById(c)

      console.log('echart',this.$echarts.myEchart);
      this.myEchart = this.$echarts.init(dom)
    },
    drawChart () {
      var timeX = []
      const now = new Date()
      var timeGap
      console.log('time',this.time);
      switch (this.time) {
        case '1s':
          timeGap = 1
          break;
        case '30m':
          timeGap = 1800
          break;
        case '60m':
          timeGap = 3600
          break;
        case '6h':
          timeGap = 21600
          break;
        case '12h':
          timeGap = 43200
          break;
        case '24h':
          timeGap = 86400
          break;
      }
      console.log('timeGap',timeGap);
      for (let i = 0; i < this.cpu.length; i++) {
        console.log();
        let timeS = new Date(now.setSeconds(now.getSeconds() - timeGap))
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
      this.myEchart.setOption(option)
    },
    search () {

      let start1 = this.timePickerValue[0].toLocaleString('chinese', {hour12:false}).split('/').join('-')
      let end = this.timePickerValue[1].toLocaleString('chinese', {hour12:false}).split('/').join('-')
      let startFormat = dayjs(start1).format("YYYY-MM-DD HH:mm:ss")
      let endFormat = dayjs(end).format("YYYY-MM-DD HH:mm:ss")

      console.log('startFormat', startFormat);
      console.log('endFormat', endFormat);
      console.log('method',this.method)
      console.log('time',this.time);
      // console.log(start);
      // console.log(end);
      // console.log(this.ip);
      // console.log(this.local);
      let cpu_metric = GetMetricsWithTime(this.ip, this.local, 'cpu_rate', startFormat, endFormat, this.method, -1)
      let mem_metric = GetMetricsWithTime(this.ip, this.local, 'mem_rate', startFormat, endFormat, this.method, -1)
      // this.runtime = []
      
      // for (let i=0; i<mem_metric.length; i++) {
      //   this.runtime.push([cpu_metric[i], mem_metric[i]])
      // }
      this.cpu = cpu_metric
      this.mem = mem_metric
      this.drawChart()

    },
    now () {
      var thisTime = new Date()
      let timeS = new Date(thisTime.setMinutes(thisTime.getMinutes() - this.selector))
      this.timePickerValue = [timeS, new Date()]
    },

  },
  mounted () {
    this.now()
    this.getCpu()
    this.getMem()
    this.chartInit()
    this.drawChart()
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
