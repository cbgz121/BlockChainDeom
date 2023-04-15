<template>
    <div class="login-container">
        <el-form ref="loginForm" class="login-form" auto-complete="on" label-position="left">
            <div class="title-container">
                <h3 class="title">基于区块链的房地产交易系统</h3>
            </div>
            <el-form-item label="用户名">
                <el-input v-model="username" placeholder="请输入用户名"></el-input>
            </el-form-item>

            <el-form-item label="密码">
                <el-input v-model="password" type="password" placeholder="请输入密码"></el-input>
            </el-form-item>

            <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;"
                @click.native.prevent="handleLogin">登录</el-button>
            <div style="text-align: center;">
                还没有账号？
                <a href="javascript:void(0);" @click="goToRegister">立即注册</a>
                <!-- <router-link href="/register">立即注册</router-link> -->
            </div>

        </el-form>
    </div>
</template>
<script>
import { login } from '@/api/account'
export default {
    name: 'Login',
    data() {
        return {
            loading: false,
            username: '',
            password: ''
        }
    },
    watch: {
        $route: {
            handler: function (route) {
                this.redirect = route.query && route.query.redirect
            },
            immediate: true
        }
    },
    methods: {
        handleLogin() {
            if (this.username && this.password) {
                this.loading = true
                this.$store.dispatch('account/login', {username:this.username, password:this.password}).then(() => {
                    this.$router.push({ path: this.redirect || '/' })
                    this.loading = false
                }).catch(() => {
                    this.$message.error('用户名或密码错误')
                    this.loading = false
                })
            } else {
                this.$message('请输入用户名和密码')
            }
        },
        goToRegister() {
            this.$router.push('/register')
        }
    }
}
</script>

<style lang="scss" scoped>
$bg: #4a6476;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #e4e0e0;
}

.login-form {
  width: 400px;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  background-color: #ffffff;
}

.title-container {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: #333333;
  text-align: center;
}

.el-form-item__label {
  font-size: 16px;
  font-weight: bold;
  color: #333333;
}

.el-input__inner {
  border-radius: 5px;
  border: 1px solid #dcdfe6;
  font-size: 16px;
}

.el-button--primary {
  background-color: #409eff;
  border-color: #409eff;
  font-size: 16px;
}

.el-button--primary:hover {
  background-color: #66b1ff;
  border-color: #66b1ff;
}

a {
  color: #409eff;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
</style> 
