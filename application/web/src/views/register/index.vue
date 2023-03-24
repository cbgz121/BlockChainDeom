<template>
    <div class="login-container">
        <el-form ref="registerForm" class="register-form" auto-complete="on" label-position="left">
            <div class="title-container">
                <h3 class="title">基于区块链的房地产交易系统</h3>
            </div>
            <el-form-item label="用户名">
                <el-input v-model="username" placeholder="请输入用户名" :disabled="registering"></el-input>
            </el-form-item>

            <el-form-item label="密码">
                <el-input v-model="password" type="password" placeholder="请输入密码" :disabled="registering"></el-input>
            </el-form-item>

            <el-form-item label="确认密码">
                <el-input v-model="confirmPassword" type="password" placeholder="请再次输入密码"
                    :disabled="registering"></el-input>
            </el-form-item>

            <el-form-item label="邮箱">
                <el-input v-model="email" placeholder="请输入邮箱" :disabled="registering"></el-input>
            </el-form-item>

            <el-form-item label="手机号">
                <el-input v-model="phone" placeholder="请输入手机号" :disabled="registering"></el-input>
            </el-form-item>

            <el-button :loading="registering" type="primary" style="width:100%;margin-bottom:30px;"
                @click.native.prevent="handleRegister">
                注册
            </el-button>

        </el-form>
    </div>
</template>
<script>
import { register } from "@/api/account";

export default {
    name: "Register",
    data() {
        return {
            registering: false,
            username: "",
            password: "",
            confirmPassword: "",
            email: "",
            phone: "",
        };
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
        handleRegister() {
            if (
                this.username &&
                this.password &&
                this.confirmPassword &&
                this.email &&
                this.phone
            ) {
                if (this.password !== this.confirmPassword) {
                    this.$message.error("两次输入的密码不一致");
                    return;
                }
                this.registering = true;
		// this.$store.dispatch('account/register', this.username,this.password,this.email,this.phone).then(() => {
		// this.$message.success("注册成功")
        //   	this.$router.push({ path: this.redirect || '/' })
                register({
                   username: this.username,
                   password: this.password,
                   email: this.email,
                   phone: this.phone,
                }).then(() => {
                   this.$message.success("注册成功")
                   this.$router.push('/')
                // window.location.href('/')
                    // this.registering = false
                }).catch((error) => {
                    this.$message.error(error.message)
                    this.registering = false
                });
            } else {
                this.$message("请完整填写注册信息")
            }
        },
    },
};
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

    .register-form {
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
