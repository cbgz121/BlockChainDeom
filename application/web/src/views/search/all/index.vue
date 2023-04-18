<template>
  <div class="container">
    <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="120px">
      <el-row>
        <el-col :span="8">
          <el-form-item label="房产类型" prop="propertyType">
            <el-input v-model="ruleForm.propertyType" placeholder="请输入房产类型"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="总空间 ㎡" prop="totalArea">
            <el-row>
              <el-col :span="10">
                <el-input-number v-model="ruleForm.totalArea" :precision="2" :step="0.1" :min="0" :max="10000" />
              </el-col>
              <el-col :span="4" class="center">
                <span>~</span>
              </el-col>
              <el-col :span="10">
                <el-input-number v-model="ruleForm.totalArea1" :precision="2" :step="0.1" :min="0" :max="10000" />
              </el-col>
            </el-row>
          </el-form-item>
        </el-col>
        <el-col :span="8" class="center">
          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">搜索</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>

    <el-table :data="filteredEstateList" style="width: 100%" stripe>
      <el-table-column prop="encumbrance" label="在售状态"></el-table-column>
      <el-table-column prop="estateAddress" label="房产地址"></el-table-column>
      <el-table-column prop="buildYear" label="建造年份"></el-table-column>
      <el-table-column prop="estateType" label="房产类型"></el-table-column>
      <el-table-column prop="estateStatus" label="房产状态描述"></el-table-column>
      <el-table-column prop="totalArea" label="总空间 ㎡"></el-table-column>
      <el-table-column prop="livingSpace" label="居住空间 ㎡"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      ruleForm: {
        propertyType: '',
        totalArea: '',
        totalArea1: '',
      },
      rules: {
        propertyType: [{ required: true, message: '请输入房产类型', trigger: 'blur' }],
      },
    };
  },
  computed: {
    filteredEstateList() {
      // 计算过滤后的列表数据
      // ...
    },
  },
};
</script>

<style scoped>
.container {
  padding: 20px;
}

.center {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
