<template>
  <div class="app-container">
    <el-form
      ref="ruleForm"
      :model="ruleForm"
      :rules="rules"
      label-width="100px"
      style="max-width: 800px; margin: 0 auto;"
    >
      <el-form-item label="业主ID" prop="proprietor">
        <el-input v-model="ruleForm.proprietor" placeholder="请输入业主ID" />
      </el-form-item>
      <el-form-item label="房产证号" prop="propertyCertificate">
        <el-input v-model="ruleForm.propertyCertificate" placeholder="请输入房产证号" />
      </el-form-item>
      <el-form-item label="房产地址" prop="propertyAddress">
        <el-input v-model="ruleForm.propertyAddress" placeholder="请输入房产地址" />
      </el-form-item>
      <el-form-item label="建造年份" prop="constructionYear">
        <el-date-picker
          v-model="ruleForm.constructionYear"
          type="date"
          placeholder="选择日期"
          style="width: 100%;"
        />
      </el-form-item>
      <el-form-item label="房产类型" prop="propertyType">
        <el-input v-model="ruleForm.propertyType" placeholder="请输入房产类型" />
      </el-form-item>
      <el-form-item label="房产状态描述" prop="propertyStatus">
        <el-input v-model="ruleForm.propertyStatus" placeholder="请输入房产状态描述" />
      </el-form-item>
      <el-form-item label="总空间 ㎡" prop="totalArea">
        <el-input-number v-model="ruleForm.totalArea" :precision="2" :step="0.1" :min="0" />
      </el-form-item>
      <el-form-item label="居住空间 ㎡" prop="livingSpace">
        <el-input-number v-model="ruleForm.livingSpace" :precision="2" :step="0.1" :min="0" />
      </el-form-item>
      <el-form-item label="上传图片">
        <el-upload
          class="upload-demo"
          :on-success="handleUploadSuccess"
          :before-upload="beforeUpload"
          :auto-upload="false"
        >
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="upload">上传文件</el-button>
          <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过2MB</div>
        </el-upload>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      ruleForm: {
        proprietor: '',
        propertyCertificate: '',
        propertyAddress: '',
        constructionYear: '',
        propertyType: '',
        propertyStatus: '',
        totalArea: '',
        livingSpace: '',
        imageUrl: '' // 新增图片URL属性
      },
      rules: {
        // 省略原有的表单校验规则
      }
    };
  },
  methods: {
    handleUploadSuccess(response, file) {
      this.ruleForm.imageUrl = URL.createObjectURL(file.raw); // 保存图片URL
      this.$message({
        message: '图片上传成功',
        type: 'success'
      });
    },
    beforeUpload(file) {
      const isJPG = file.type === 'image/jpeg' || file.type === 'image/png';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG/PNG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          // 发送添加请求
          this.$http.post('/api/properties', this.ruleForm)
            .then(response => {
              this.$message({
                message: '创建成功',
                type: 'success'
              });
              // 清空表单数据和图片URL
              this.$refs[formName].resetFields();
              this.ruleForm.imageUrl = '';
            })
            .catch(error => {
              console.log(error);
              this.$message.error('创建失败');
            });
        } else {
          console.log('校验未通过');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.ruleForm.imageUrl = ''; // 重置时清空图片URL
    }
  }
};
</script>

<style>
.app-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}
.el-form-item__label {
  font-size: 16px;
  font-weight: bold;
}
.el-form-item__content {
  margin-left: 120px;
}
.el-input-number__decrease,
.el-input-number__increase {
  color: #409eff;
}
.el-upload__tip {
  color: #909399;
}
</style>
