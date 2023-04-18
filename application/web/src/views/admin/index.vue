<template>
    <div class="login-container">
        <el-form ref="loginForm" class="login-form" auto-complete="on" label-position="left">
            <div class="title-container">
                <h3 class="title">管理员登陆</h3>
            </div>
            <el-form-item label="用户名">
                <el-input v-model="username" placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码">
                <el-input v-model="password" type="password" placeholder="请输入密码"></el-input>
            </el-form-item>

         <!-- <el-form-item label="验证码" class="captcha">
                <el-input v-model="captcha" placeholder="请输入验证码"></el-input>
                <el-image class="captcha-img" :src="`/captcha/${captchaId}.png`" @click="getCaptcha" alt="验证码"></el-image>
            </el-form-item>  -->
              <!--  <div>
                    <el-form :model="form" ref="form">
                    <el-form-item label="验证码">
                        <img :src="captchaUrl" @click="refreshCaptcha" style="cursor: pointer">
                    </el-form-item>
                    <el-form-item label="验证码答案">
                        <el-input v-model="form.captchaSolution"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submitForm">提交</el-button>
                    </el-form-item>
                    </el-form>
                </div> -->

            <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;"
                @click.native.prevent="handleLogin">登录</el-button>

            <div class="extra-links">
                <!-- <a href="javascript:void(0);" @click="goToRegister">立即注册</a> -->
                <a href="javascript:void(0);" @click="goToUserLogin">用户登录</a>
            </div>
        </el-form>
    </div>
</template>
<script>
import { captcha } from '@/api/account'

export default {
    name: 'Login2',
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
    created() {
        this.getCaptcha();
    },

    methods: {
        getCaptcha() {
            captcha().then(response => {
               // alert(response.captchaId)
                this.captchaId = response.captchaId;
                this.captcha = '';
            }).catch(err => {
                console.log(err);
            });
        },
        handleLogin() {
            if (this.username && this.password) {
                this.loading = true
                this.$store.dispatch('account/adminlogin', { username: this.username, password: this.password }).then(() => {
                    //alert(this.username)
                    this.$router.push({ path: this.redirect || '/' })
                    this.loading = false
                }).catch(() => {
                    this.$message.error('用户名或密码错误')
                    this.loading = false
                    this.getCaptcha();
                })
            } else {
                this.$message('请输入用户名和密码')
            }
        },

//        goToRegister() {
//            this.$router.push({ path: '/register' });
//        },
        goToUserLogin() {
            this.$router.push({ path: '/login' });
        },
    },
};
</script>
<style scoped>
.login-container {
    max-width: 400px;
    margin: 0 auto;
    padding-top: 30px;
}

.title-container {
    text-align: center;
    margin-bottom: 30px;
}

.title {
    font-size: 28px;
    font-weight: bold;
    color: #409EFF;
}

.login-form {
    background-color: #fff;
    padding: 40px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
    border-radius: 5px;
}

.captcha {
    display: flex;
    align-items: center;
}

.captcha-img {
    margin-left: 10px;
    cursor: pointer;
}

.extra-links {
    display: flex;
    justify-content: space-between;
    margin-top: 20px;
}

.extra-links a {
    margin-right: 20px;
    color: #666;
}

.extra-links a:hover {
    color: #409EFF;
}
</style>

