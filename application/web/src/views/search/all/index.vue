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
        </el-row>
        <el-row>
          <el-col :span="24" class="center">
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
import { queryRealEstateList } from '@/api/realEstate'
export default {
  name: 'search',
  data() {
    return {
      ruleForm: {
        propertyType: '',
        totalArea: 0
      },
      estateList: [],
      filteredEstateList: [], // 按面积和类型过滤后的数据项
      rules: {
        totalArea: [
          { required: true, message: '请输入房产面积', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          queryRealEstateList().then(response => {
            if (response !== null) {
              this.estateList = response
            }
            this.filterByAreaAndType(this.ruleForm.totalArea, this.ruleForm.totalArea1, this.ruleForm.propertyType); // 传递指定面积范围
          })
        } else {
          return false
        }
      })
    },
    filterByAreaAndType(minArea, maxArea, type) {
      this.filteredEstateList = this.estateList.filter(
        item => item.totalArea >= minArea && item.totalArea <= maxArea && item.estateType === type
      ).map(item => ({ ...item, encumbrance: item.encumbrance.toString() }));
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    }
  }

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

