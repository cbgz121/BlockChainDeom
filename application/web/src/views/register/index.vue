<template>
    <div class="login-container">
        <el-form ref="registerForm" :rules="rules" class="register-form" auto-complete="on" label-position="left">
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
            rules:{
                    username:[{ required: true, message: '请输入用户名', trigger: 'blur' }, { min: 4, max: 20, message: '用户名长度应为4到20个字符', trigger: 'blur' }], 
                    email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }],
                    password:[{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }],
                    confirmPassword:[{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }],
                    phone:[{ required: true, message: '请输入手机号', trigger: 'blur' }, { pattern: /^1[34578]\d{9}$/, message: '手机号格式不正确', trigger: 'blur' }],
            },
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
                register({username:this.username, password:this.password, email: this.email,phone:this.phone, status: '1'}).then(() => {
                   this.$message.success("注册成功")
                   this.$router.push('/')
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
$bg: #4a6476;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
 // background-color: #e4e0e0;
}

.register-form {
  width: 400px;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
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
  color: #409eff;
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
</style>
