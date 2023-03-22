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
                login(this.username, this.password).then(() => {
                    this.$router.push('/')
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
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
    min-height: 100%;
    width: 100%;
    background-color: $bg;
    overflow: hidden;

    .login-form {
        position: relative;
        width: 520px;
        max-width: 100%;
        padding: 160px 35px 0;
        margin: 0 auto;
        overflow: hidden;
    }

    .svg-container {
        padding: 6px 5px 6px 15px;
        color: $dark_gray;
        vertical-align: middle;
        width: 30px;
        display: inline-block;
    }

    .title-container {
        position: relative;

        .title {
            font-size: 26px;
            color: $light_gray;
            margin: 0px auto 40px auto;
            text-align: center;
            font-weight: bold;
        }
    }

}
</style> 