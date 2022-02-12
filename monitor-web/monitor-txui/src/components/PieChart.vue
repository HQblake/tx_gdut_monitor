<!--
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-08 11:20:37
 * @LastEditors: yzq
-->
<template>
  <div>
    <div class="single-chart">
      <div id="main" class="echart">pieChart</div>
    </div>
  </div>
</template>

<script>
export default {
<<<<<<< HEAD
  name: "PieChart",
  data() {
    return {};
=======
  name: 'PieChart',
  data () {
    return {

    }
  },
  computed: {
>>>>>>> 00599733bc39c12b7927691154ac6e5948b210f4
  },
  computed: {},
  props: [],
<<<<<<< HEAD
  created() {},
  watch: {},
  destroyed() {},
  methods: {
    run(_rawData) {
      echarts.registerTransform(window.ecSimpleTransform.aggregate);
      option = {
        dataset: [
          {
            id: "raw",
            source: _rawData,
          },
          {
            id: "since_year",
            fromDatasetId: "raw",
            transform: [
              {
                type: "filter",
                config: {
                  dimension: "Year",
                  gte: 1950,
                },
              },
            ],
          },
          {
            id: "income_aggregate",
            fromDatasetId: "since_year",
            transform: [
              {
                type: "ecSimpleTransform:aggregate",
                config: {
                  resultDimensions: [
                    { name: "min", from: "Income", method: "min" },
                    { name: "Q1", from: "Income", method: "Q1" },
                    { name: "median", from: "Income", method: "median" },
                    { name: "Q3", from: "Income", method: "Q3" },
                    { name: "max", from: "Income", method: "max" },
                    { name: "Country", from: "Country" },
                  ],
                  groupBy: "Country",
                },
              },
              {
                type: "sort",
                config: {
                  dimension: "Q3",
                  order: "asc",
                },
              },
            ],
          },
        ],
        title: {
          text: "Income since 1950",
        },
        tooltip: {
          trigger: "axis",
          confine: true,
        },
        xAxis: {
          name: "Income",
          nameLocation: "middle",
          nameGap: 30,
          scale: true,
        },
        yAxis: {
          type: "category",
        },
        grid: {
          bottom: 100,
        },
        legend: {
          // selected: { detail: false },
        },
        dataZoom: [
          {
            type: "inside",
          },
          {
            type: "slider",
            height: 20,
          },
        ],
        series: [
          {
            name: "boxplot",
            type: "boxplot",
            datasetId: "income_aggregate",
            itemStyle: {
              color: "#b8c5f2",
            },
            encode: {
              x: ["min", "Q1", "median", "Q3", "max"],
              y: "Country",
              itemName: ["Country"],
              tooltip: ["min", "Q1", "median", "Q3", "max"],
            },
          },
          {
            name: "detail",
            type: "scatter",
            datasetId: "since_year",
            symbolSize: 6,
            tooltip: {
              trigger: "item",
            },
            label: {
              show: true,
              position: "top",
              align: "left",
              verticalAlign: "middle",
              rotate: 90,
              fontSize: 12,
            },
            itemStyle: {
              color: "#d00000",
            },
            encode: {
              x: "Income",
              y: "Country",
              label: "Year",
              itemName: "Year",
              tooltip: ["Country", "Year", "Income"],
            },
          },
        ],
      };
      myChart.setOption(option);
    },
  },
  mounted() {
=======
  created () {},
  watch: {

  },
  destroyed () {

  },
  methods: {
    run (_rawData) {
      echarts.registerTransform(window.ecSimpleTransform.aggregate)
      option = {
        dataset: [
          {
            id: 'raw',
            source: _rawData
          },
          {
            id: 'since_year',
            fromDatasetId: 'raw',
            transform: [
              {
                type: 'filter',
                config: {
                  dimension: 'Year',
                  gte: 1950
                }
              }
            ]
          },
          {
            id: 'income_aggregate',
            fromDatasetId: 'since_year',
            transform: [
              {
                type: 'ecSimpleTransform:aggregate',
                config: {
                  resultDimensions: [
                    { name: 'min', from: 'Income', method: 'min' },
                    { name: 'Q1', from: 'Income', method: 'Q1' },
                    { name: 'median', from: 'Income', method: 'median' },
                    { name: 'Q3', from: 'Income', method: 'Q3' },
                    { name: 'max', from: 'Income', method: 'max' },
                    { name: 'Country', from: 'Country' }
                  ],
                  groupBy: 'Country'
                }
              },
              {
                type: 'sort',
                config: {
                  dimension: 'Q3',
                  order: 'asc'
                }
              }
            ]
          }
        ],
        title: {
          text: 'Income since 1950'
        },
        tooltip: {
          trigger: 'axis',
          confine: true
        },
        xAxis: {
          name: 'Income',
          nameLocation: 'middle',
          nameGap: 30,
          scale: true
        },
        yAxis: {
          type: 'category'
        },
        grid: {
          bottom: 100
        },
        legend: {
          selected: { detail: false }
        },
        dataZoom: [
          {
            type: 'inside'
          },
          {
            type: 'slider',
            height: 20
          }
        ],
        series: [
          {
            name: 'boxplot',
            type: 'boxplot',
            datasetId: 'income_aggregate',
            itemStyle: {
              color: '#b8c5f2'
            },
            encode: {
              x: ['min', 'Q1', 'median', 'Q3', 'max'],
              y: 'Country',
              itemName: ['Country'],
              tooltip: ['min', 'Q1', 'median', 'Q3', 'max']
            }
          },
          {
            name: 'detail',
            type: 'scatter',
            datasetId: 'since_year',
            symbolSize: 6,
            tooltip: {
              trigger: 'item'
            },
            label: {
              show: true,
              position: 'top',
              align: 'left',
              verticalAlign: 'middle',
              rotate: 90,
              fontSize: 12
            },
            itemStyle: {
              color: '#d00000'
            },
            encode: {
              x: 'Income',
              y: 'Country',
              label: 'Year',
              itemName: 'Year',
              tooltip: ['Country', 'Year', 'Income']
            }
          }
        ]
      }
      myChart.setOption(option)
    }

  },
  mounted () {

>>>>>>> 00599733bc39c12b7927691154ac6e5948b210f4
    // this.run();
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
.block {
  margin-left: 10px;
  margin-right: 10px;
}
</style>
