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
          <el-col :span="6"><div class="grid-content bg-purple"></div></el-col>
          <el-col :span="4"
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
          <el-col :span="4"
            ><div class="grid-content bg-purple">
              <el-select v-model="warnContent" placeholder="级别">
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
                value-format="yyyy-MM-dd HH:mm:ss"
                :default-time="['08:00:00', '12:00:00']"
                :default-value="new Date()"
              >
              </el-date-picker>
            </div>
          </el-col>
          <el-col :span="2"
            ><div class="grid-content bg-purple">
              <el-button type="primary">搜索</el-button>
            </div>
          </el-col>
        </el-row>
      </div>

      <el-table
        :data="
          tableData.filter(
            (data) =>
              !search ||
              data.id == search.toLowerCase() ||
              data.level.toLowerCase().includes(search.toLowerCase()) ||
              data.tabName.toLowerCase().includes(search.toLowerCase()) ||
              data.startTime.includes(search)
          )
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
          prop="tabName"
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
        </el-table-column>
        <el-table-column
          min-width="100px"
          label="开始时间"
          prop="startTime"
          align="center"
          header-aligh="center"
        >
        </el-table-column>
        <!-- <el-table-column label="指标" prop="5" > </el-table-column> -->
        <el-table-column
          label="异常值"
          prop="outliers"
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
          prop="during"
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
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >Edit</el-button
            >
            <!-- deleteData(scope.$index, scope.row) -->
            <el-button
              size="mini"
              type="danger"
              @click="openDelete(scope.$index, scope.row)"
              >Delete</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { GetWarnList } from "@/api/show";

export default {
  name: "warning",
  components: {},
  data() {
    return {
      timePickerValue: [],
      level: "",
      levelOptions: [
        {
          label: "严重",
          value: "严重",
        },
        {
          label: "中等",
          value: "中等",
        },
        {
          label: "警告",
          value: "警告",
        },
      ],
      warnContent: "",
      warnOptions: [
        {
          label: "cpu使用率",
          value: "cpu",
        },
        {
          label: "mem使用率",
          value: "mem",
        },
      ],
      Outliers: "",
      threshold: "",
      duration: "",
      dialogVisible: false,
      search: "",
      tableData: [
        {
          id: 1,
          ip: '127.0.0.1',
          local: '北京',
          tabName: "cpu利用率",
          level: "严重",
          startTime: "2022-02-10 18:00:00",
          outliers: 85,
          threshold: 80,
          during: 30,
        },
        {
          id: 2,
           ip: '127.0.0.2',
          local: '北京',
          tabName: "cpu利用率",
          level: "严重",
          startTime: "2022-02-10 19:00:00",
          outliers: 85,
          threshold: 80,
          during: 30,
        },
        {
          id: 3,
          ip: '127.0.0.2',
          local: '上海',
          tabName: "mem利用率",
          level: "中等",
          startTime: "2022-02-10 20:00:00",
          outliers: 85,
          threshold: 80,
          during: 30,
        },
      ],
    };
  },
  methods: {
    now() {
      var thisTime = new Date()
      let timeS = new Date(thisTime.setMinutes(thisTime.getMinutes() - this.selector))
      this.timePickerValue = [timeS, new Date()]
    },
    deleteData(index, row) {
      this.dialogVisible = false;
      console.log(index);
      console.log(row);

      this.tableData.splice(index, 1);
    },
    openDelete(index, row) {
      this.$confirm("此操作将永久删除该文件, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.deleteData(index, row);
          this.$message({
            type: "success",
            message: "删除成功!",
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    handleEdit(index, row) {
      console.log(index, row);
    },
    handleDelete(index, row) {
      console.log(index, row);
    },
  },
  created (){
    
  },
  mounted (){
    this.now()
  }
};
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
</style>
