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
      <el-table
        :data="
          tableData.filter(
            (data) =>
              !search || data.level.toLowerCase().includes(search.toLowerCase()) || data.level.toLowerCase().includes(search.toLowerCase())
          )
        "
        :header-cell-style="{ height: '100px' }"
        style="width: 100%"
      >
        <el-table-column
          label="对象"
          prop="id"
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

        <el-table-column align="center" min-width='100px'>
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="输入级别关键字进行搜索 "
            />
          </template>
          <template slot-scope="scope" >
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
export default {
  name: 'warning',
  components: {},
  data () {
    return {
      dialogVisible: false,
      search: '',
      tableData: [
        {
          id: 1,
          tabName: 'cpu使用',
          level: '严重',
          startTime: '2022-02-10 18:00:00',
          outliers: 85,
          threshold: 80,
          during: 30
        },
        {
          id: 2,
          tabName: 'cpu使用',
          level: '严重',
          startTime: '2022-02-10 18:00:00',
          outliers: 85,
          threshold: 80,
          during: 30
        },
        {
          id: 3,
          tabName: 'cpu使用',
          level: '中等',
          startTime: '2022-02-10 18:00:00',
          outliers: 85,
          threshold: 80,
          during: 30
        }
      ]
    }
  },
  methods: {
    deleteData (index, row) {
      this.dialogVisible = false
      console.log(index)
      console.log(row)

      this.tableData.splice(index, 1)
    },
    openDelete (index, row) {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
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
    },
    handleEdit (index, row) {
      console.log(index, row)
    },
    handleDelete (index, row) {
      console.log(index, row)
    }
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
</style>
