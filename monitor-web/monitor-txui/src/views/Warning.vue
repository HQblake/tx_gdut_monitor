<!--
 * @Description: 
 * @Autor: yzq
 * @Date: 2022-02-10 16:05:05
 * @LastEditors: yzq
-->
<template>
  <div class="warning">
    <el-container class="container">
      <el-main>
        <el-table
          :data="
            tableData.filter(
              (data) =>
                !search ||
                data.name.toLowerCase().includes(search.toLowerCase())
            )
          "
          
          
          style="width: 100%; "
        > 

          <el-table-column label="对象" prop="id" align="center" header-aligh="center"> </el-table-column>
          <el-table-column label="告警内容" prop="tabName" align="center" header-aligh="center"> </el-table-column>
          <el-table-column label="级别" prop="level" align="center" header-aligh="center"> </el-table-column>
          <el-table-column label="开始时间" prop="startTime" align="center" header-aligh="center"> </el-table-column>
          <!-- <el-table-column label="指标" prop="5" > </el-table-column> -->
          <el-table-column label="异常值" prop="outliers" align="center" header-aligh="center"> </el-table-column>
          <el-table-column label="阈值" prop="threshold" align="center" header-aligh="center"> </el-table-column>
          <el-table-column label="持续时间" prop="during" align="center" header-aligh="center"> </el-table-column>

          
          <el-table-column align="center">
            <template slot="header" slot-scope="scope">
              <el-input
                v-model="search"
                size="mini"
                placeholder="输入关键字搜索"
              />
            </template>
            <template slot-scope="scope">
              <el-button
                size="mini"
                @click="handleEdit(scope.$index, scope.row)"
                >Edit</el-button
              >
              <el-button
                size="mini"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
                >Delete</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "warning",
  components: {},
  data() {
    return {
      search:'',
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
          id: 1,
          tabName: 'cpu使用',
          level: '严重',
          startTime: '2022-02-10 18:00:00',
          outliers: 85,
          threshold: 80,
          during: 30
        },
        
      ],
    };
  },
  methods: {
    handleEdit(index, row) {
      console.log(index, row);
    },
    handleDelete(index, row) {
      console.log(index, row);
    },
  },
};
</script>
