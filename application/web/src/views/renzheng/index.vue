<template>
    <div class="personal-info-authentication">
        <el-card class="personal-info-card">
            <div slot="header" class="personal-info-header">
                <h3>个人信息认证</h3>
            </div>
            <el-form :model="form" :rules="rules" ref="form" label-width="120px" class="personal-info-form">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="form.username"></el-input>
                </el-form-item>
                <el-form-item  label="密码" prop="password">
                    <el-input type="password" v-model="form.password"></el-input>
                </el-form-item>
                <el-form-item label="真实姓名" prop="name">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="身份证号码" prop="idCard">
                    <el-input v-model="form.idCard"></el-input>
                </el-form-item>
                <el-form-item label="手机号码" prop="phone">
                    <el-input v-model="form.phone"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="Authen(val)">认证并获取ID</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-dialog v-loading="loadingDialog" :visible.sync="dialogCreateDonating" :close-on-click-modal="false" @close="resetForm('DonatingForm')" style="width: 700px;top: 20%; left: 30%;">
            <el-form ref="DonatingForm" :model="DonatingForm" :rules="rulesDonating" label-width="100px">
              <el-form style="width: 100%">
                <div ref="formItem" >
                    <el-form-item label="ID:">
                        <span v-text="account" @click="copyToClipboard"></span>
                    </el-form-item>
                </div>
              </el-form>
            </el-form>
            <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="confirm('realForm')">确定</el-button>
          </div>
          </el-dialog>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
import {authentication} from '@/api/account'
export default {
    data() {
        return {
            form: {
                username: '',
                password: '',
                name: '',
                idCard: '',
                phone: '',
                photo: '',
            },
            rules: {
                username: [{ required: true, message: '请输入用户名', trigger: 'blur' }, { min: 4, max: 20, message: '用户名长度应为4到20个字符', trigger: 'blur' }],
                password: [{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }],
                name: [{ required: true, message: '请输入真实姓名', trigger: 'blur' }, { pattern: /^[\u4E00-\u9FA5A-Za-z]+$/, message: '姓名只能输入汉字或字母', trigger: 'blur' }],
                idCard: [{ required: true, message: '请输入身份证号码', trigger: 'blur' }, { pattern: /^\d{18}$/, message: '身份证号码必须为18位数字', trigger: 'blur' }],
                phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }, { pattern: /^1[34578]\d{9}$/, message: '手机号格式不正确', trigger: 'blur' }],
            },
            account: '',
            dialogCreateDonating: false,
            formItemContent: ''
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
    methods: {
        submitForm() {
            this.$refs.form.validate(valid => {
                if (valid) {
                    this.$message.success('个人信息认证成功!');
                } else {
                    this.$message.error('请填写正确的信息!');
                }
            });
        },
        Authen() {
            // 获取表单组件的引用
            const form = this.$refs.form;

            // 获取表单数据
            const formData = form.model;
            this.$refs.form.validate(valid => {
                    if (valid) {
                       // this.dialogCreateDonating = true
                        this.$store.dispatch('account/authentication',
                            {
                                username: formData.username,
                                password: formData.password,
                                name: formData.name,
                                idcard: formData.idCard,
                                phone: formData.phone,
                            }
                        ).then(accountId => {
                            this.$message.success("认证成功")
                            this.dialogCreateDonating = true
                            this.account = accountId
                        }).catch((error) => {
                            alert(error.message)
                            this.$message.error(error.message)
                            this.dialogCreateDonating = false
                        });
                        
                    } else {
                        this.$message.error('请填写正确的信息!');
                    }
                
            });
        },
        confirm() {
            setTimeout(() => {
                window.location.reload()
            }, 10)
        }
        
    },
};
</script>

<style>
.personal-info-authentication {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.personal-info-card {
    width: 500px;
    margin-top: 50px;
}

.personal-info-header {
    background-color: #eef1f6;
    padding: 15px;
    text-align: center;
}

.personal-info-form {
    margin-top: 20px;
}

.upload-demo {
    margin-top: 10px;
}

.upload-demo i {
    font-size: 28px;
    color: #999;
}

.upload-demo .el-upload-list {
    display: flex;
}

.upload-demo .el-upload-list__item {
    margin-right: 10px;
}
</style>
