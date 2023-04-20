<template>
  <div class="container">
    <div v-if="realEstateList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in realEstateList" :key="index" :span="6" :offset="1">
        <el-card class="realEstate-card">
          <div slot="header" class="clearfix">
            担保状态:
            <span style="color: rgb(255, 0, 0);">{{ val.encumbrance }}</span>
          </div>

          <div class="item">
            <el-tag>房产ID: </el-tag>
            <span>{{ val.realEstateId }}</span>
          </div>
          <div class="item">
            <el-tag type="success">房产类型: </el-tag>
            <span>{{ val.estateType }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">总空间: </el-tag>
            <span>{{ val.totalArea }} ㎡</span>
          </div>
          <div class="item">
            <el-tag type="danger">居住空间: </el-tag>
            <span>{{ val.livingSpace }} ㎡</span>
          </div>

          <div v-if="!val.encumbrance&&roles[0] !== 'admin'">
            <el-button type="text" @click="openDialog(val)">出售</el-button>
            <el-divider direction="vertical" />
            <el-button type="text" @click="openDonatingDialog(val)">详细信息</el-button>
          </div>
          <el-rate v-if="roles[0] === 'admin'" />
        </el-card>
      </el-col>
    </el-row>
    <el-dialog v-loading="loadingDialog" :visible.sync="dialogCreateSelling" :close-on-click-modal="false" @close="resetForm('realForm')">
      <el-form ref="realForm" :model="realForm" :rules="rules" label-width="100px">
        <el-form-item label="价格 (元)" prop="price">
          <el-input-number v-model="realForm.price" :precision="2" :step="10000" :min="0" />
        </el-form-item>
        <el-form-item label="有效期 (天)" prop="salePeriod">
          <el-input-number v-model="realForm.salePeriod" :min="1" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="createSelling('realForm')">立即出售</el-button>
        <el-button @click="dialogCreateSelling = false">取 消</el-button>
      </div>
    </el-dialog>
    <el-dialog v-loading="loadingDialog" :visible.sync="dialogCreateDonating" :close-on-click-modal="false" @close="resetForm('DonatingForm')">
      <el-form ref="DonatingForm" :model="DonatingForm" :rules="rulesDonating" label-width="100px">
        <el-form style="width: 100%" v-for="(account, index) in accountList" :key="index">
        <el-form-item :label="'在售状态: ' + account.encumbrance"></el-form-item>
        <el-form-item :label="'房产地址: ' + account.estateAddress"></el-form-item>
        <el-form-item :label="'建造年份: ' + account.buildYear"></el-form-item>
        <el-form-item :label="'房产类型: ' + account.estateType"></el-form-item>
        <el-form-item :label="'房产状态描述: ' + account.estateStatus"></el-form-item>
        <el-form-item :label="'总空间 ㎡: ' + account.totalArea"></el-form-item>
        <el-form-item :label="'居住空间 ㎡: ' + account.livingSpace"></el-form-item>
      </el-form>
        
      </el-form>
      <!-- <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="createDonating('DonatingForm')">立即出租</el-button>
        <el-button @click="dialogCreateDonating = false">取 消</el-button>
      </div> -->
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { queryRealEstateList } from '@/api/realEstate'
import { createSelling } from '@/api/selling'
import { createDonating } from '@/api/donating'

export default {
  name: 'RealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      loading: true,
      loadingDialog: false,
      realEstateList: [],
      dialogCreateSelling: false,
      dialogCreateDonating: false,
      realForm: {
        price: 0,
        salePeriod: 0
      },
      rules: {
        price: [
          { validator: checkArea, trigger: 'blur' }
        ],
        salePeriod: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      DonatingForm: {
        proprietor: ''
      },
      rulesDonating: {
        proprietor: [
          { required: true, message: '请选择业主', trigger: 'change' }
        ]
      },
      accountList: [],
      valItem: {}
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ])
  },
  created() {
    if (this.roles[0] === 'admin') {
      queryRealEstateList().then(response => {
        if (response !== null) {
          this.realEstateList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryRealEstateList({ proprietor: this.accountId }).then(response => {
        if (response !== null) {
          this.realEstateList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
    openDialog(item) {
      this.dialogCreateSelling = true
      this.valItem = item
    },
    openDonatingDialog(item) {
      this.dialogCreateDonating = true
      this.valItem = item
      queryRealEstateList({ proprietor: this.accountId }).then(response => {
        if (response !== null) {
          this.accountList = response.filter(item =>
            item.realEstateId === this.valItem.realEstateId,
          )
        }
      })
    },
    createSelling(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即出售?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createSelling({
              objectOfSale: this.valItem.realEstateId,
              seller: this.valItem.proprietor,
              price: this.realForm.price,
              salePeriod: this.realForm.salePeriod,
              estateType: this.valItem.estateType
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '出售成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '出售失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateSelling = false
            this.$message({
              type: 'info',
              message: '已取消出售'
            })
          })
        } else {
          return false
        }
      })
    },
    createDonating(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即出租?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createDonating({
              objectOfDonating: this.valItem.realEstateId,
              donor: this.valItem.proprietor,
              grantee: this.DonatingForm.proprietor
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '出租成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '出租失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateDonating = false
            this.$message({
              type: 'info',
              message: '已取消出租'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.DonatingForm.proprietor = accountId
    }
  }
}

</script>

<style>
  /* .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
  .tag {
    float: left;
  }*/

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  /*.clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .realEstate-card {
    width: 280px;
    height: 340px;
    margin: 18px;
  } */

.container {
  text-align: center;
  margin: 30px auto;
  max-width: 1200px;
}

.el-alert {
  margin-bottom: 30px;
  border-radius: 4px;
  font-size: 16px;
}

.el-alert p {
  margin-bottom: 10px;
}

.el-row {
  margin-bottom: 30px;
}

.el-card {
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
}

.el-card .item {
  margin-top: 20px;
}

.el-card .item el-tag {
  margin-right: 10px;
  font-size: 1px;
}

.el-card .encumbrance {
  font-size: 18px;
  font-weight: bold;
}

.el-card .realEstate-card .el-button {
  color: #333;
  font-size: 14px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
}

.el-card .realEstate-card .el-button:hover {
  background-color: #f5f7fa;
}

.el-card .realEstate-card .el-divider {
  margin: 0 10px;
}

.el-card .realEstate-card .el-rate {
  margin-top: 10px;
  font-size: 20px;
}

.el-dialog {
  text-align: center;
}

.el-dialog__header {
  font-size: 18px;
}

.el-dialog__body {
  padding: 20px;
}

.el-dialog__footer {
  padding: 20px;
}

.el-button--primary {
  background-color: #1890ff;
  border-color: #1890ff;
}

.el-button--primary:hover {
  background-color: #40a9ff;
  border-color: #40a9ff;
}

.el-select {
  width: 100%;
  font-size: 14px;
}

.el-select__caret {
  color: #c0c4cc;
}

.el-option {
  font-size: 14px;
}

.el-option__label {
  display: flex;
  justify-content: space-between;
}

.el-option__label span:first-child {
  margin-right: 10px;
}

.el-input-number {
  width: 100%;
  font-size: 14px;
}
</style>
