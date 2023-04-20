<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ accountId }}</p>
      <p>用户名: {{ userName }}</p>
      <p>余额: ￥{{ balance }} 元</p>
    </el-alert>
    <div v-if="donatingList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in donatingList" :key="index" :span="6" :offset="1">
        <el-card class="d-buy-card">
          <div slot="header" class="clearfix">
            <span>{{ val.donating.donatingStatus }}</span>
            <el-button v-if="val.donating.donatingStatus==='出租中'" style="float: right; padding: 3px 0" type="text" @click="updateDonating(val,'done')">确认接收</el-button>
            <el-button v-if="val.donating.donatingStatus==='出租中'" style="float: right; padding: 3px 6px" type="text" @click="updateDonating(val,'cancelled')">取消</el-button>
          </div>
          <div class="item">
            <el-tag>房产ID: </el-tag>
            <span>{{ val.donating.objectOfDonating }}</span>
          </div>
          <div class="item">
            <el-tag type="success">出租者ID: </el-tag>
            <span>{{ val.donating.donor }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">受赠人ID: </el-tag>
            <span>{{ val.donating.grantee }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">创建时间: </el-tag>
            <span>{{ val.donating.createTime }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryDonatingListByGrantee, updateDonating } from '@/api/donating'

export default {
  name: 'DonatingGrantee',
  data() {
    return {
      loading: true,
      donatingList: []
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'userName',
      'balance'
    ])
  },
  created() {
    queryDonatingListByGrantee({ grantee: this.accountId }).then(response => {
      if (response !== null) {
        this.donatingList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
  },
  methods: {
    updateDonating(item, type) {
      let tip = ''
      if (type === 'done') {
        tip = '确认租赁'
      } else {
        tip = '取消出租操作'
      }
      this.$confirm('是否要' + tip + '?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).then(() => {
        this.loading = true
        updateDonating({
          donor: item.donating.donor,
          grantee: item.donating.grantee,
          objectOfDonating: item.donating.objectOfDonating,
          status: type
        }).then(response => {
          this.loading = false
          if (response !== null) {
            this.$message({
              type: 'success',
              message: tip + '操作成功!'
            })
          } else {
            this.$message({
              type: 'error',
              message: tip + '操作失败!'
            })
          }
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        }).catch(_ => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
        this.$message({
          type: 'info',
          message: '已取消' + tip
        })
      })
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
  } */

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  /* .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .d-buy-card {
    width: 280px;
    height: 300px;
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
