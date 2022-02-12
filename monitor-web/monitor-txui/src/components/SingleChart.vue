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
        <el-select v-model="selector" placeholder="时长">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
        <div class="block">
          <span class="demonstration"></span>
          <el-date-picker
            v-model="value1"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="yyyy-MM-dd HH:mm:ss"
          >
          </el-date-picker>
        </div>
        <el-row>
          <el-button type="primary" @click="search()">
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
  </div>
</template>

<script>

export default {
  name: 'SingleChart',
  data () {
    return {
      options: [
        {
          value: '30',
          label: '30m'
        },
        {
          value: '60',
          label: '60m'
        },
        {
          value: '300',
          label: '5h'
        },
        {
          value: '720',
          label: '12h'
        },
        {
          value: '1440',
          label: '24h'
        }
      ],
      selector: '30',
      value1: [new Date(2000, 10, 10, 10, 10), new Date(2000, 10, 11, 10, 10)],
      value2: '',
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
        [50, 40]
      ]
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
  created () {},
  watch: {
    cpu (val) {
      this.drawChart()
    },
    mem (val) {
      this.drawChart()
    },
    data1: function (newVal, olVal) {
      this.cData = newVal
    }
  },
  destroyed () {
    clearInterval(this.interval)
  },
  methods: {
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
        // var tmp1 = echarts.time.format("MM-dd\nhh:mm:ss",timeS);
        // console.log(tmp1);
        timeX.unshift(tmp)
      }
      // console.log(timeX);
      // var data = this.cpu
      // console.log(data)
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
    getUlId () {
      console.log('methods', this.elId)
    },
    search () {
      console.log(this.selector)
      console.log(this.value1)
      let time = this.value1[0].toString()
      console.log(time)
    }
    // getCpu(){
    //   for (let i=0; i<this.runtime.length; i++) {
    //     this.cpu.push(this.runtime[i][0])
    //   }
    // },
    // getMem(){
    //   for (let i=0; i<this.runtime.length; i++) {
    //     this.mem.push(this.runtime[i][1])
    //   }
    // }
  },
  mounted () {
    // this.elId = guid();
    // generateUuid().then((elId)=>{
    //   this.elId =  elId;
    //   console.log('create',this.elId);
    // })
    // console.log(this.elId);
    // this.getUlId();
    this.drawChart()
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
  margin: auto;
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
</style>
