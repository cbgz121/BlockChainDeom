<template>
    <div class="user-management">
        <el-row>
            <el-col :span="12">
                <el-input v-model="searchValue" placeholder="请输入搜索关键字" suffix-icon="el-icon-search"
                    @change="search"></el-input>
            </el-col>
            <el-col :span="12" class="text-right">
                <el-button type="primary" icon="el-icon-plus" @click="showAdd()">新增用户</el-button>
            </el-col>
        </el-row>
        <el-table :data="tableData" class="user-table">
            <el-table-column prop="ID" label="ID" sortable></el-table-column>
            <el-table-column prop="userName" label="用户名" sortable></el-table-column>
            <el-table-column prop="passWord" label="密码" sortable></el-table-column>
            <el-table-column prop="email" label="邮箱" sortable></el-table-column>
            <el-table-column prop="phone" label="手机号" sortable></el-table-column>
            <el-table-column label="状态">
                <template slot-scope="scope">
                    <el-switch v-model="scope.row.status" active-value="1" inactive-value="0" active-text="启用"
                        inactive-text="禁用" @change="updateStatus(scope.row)"></el-switch>
                </template>
            </el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button type="text" size="small" @click="editUser(scope.row)">编辑</el-button>
                    <el-button type="text" size="small" class="danger" @click="deleteUser(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination :total="total" :current-page="currentPage" :page-size="pageSize" @current-change="handlePageChange"
            class="pagination"></el-pagination>
        <el-dialog :visible.sync="showEditDialog" title="编辑用户" :append-to-body="true" :fullscreen="true"
            style="width: 700px; top: 20%; left: 30%; height: 500px;" center>
            <el-form ref="editForm" :model="editForm" label-width="120px">
                <el-form-item label="用户名" prop="username" :rules="[{ required: true, message: '请输入用户名', trigger: 'blur' }]">
                    <el-input class="short-input" v-model="editForm.username" placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password"
                    :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }]">
                    <el-input class="short-input" type="password" v-model="editForm.password"
                        placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email"
                    :rules="[{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }]">
                    <el-input class="short-input" v-model="editForm.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="phone"
                    :rules="[{ required: true, message: '请输入手机号', trigger: 'blur' }, { pattern: /^1[34578]\d{9}$/, message: '手机号格式不正确', trigger: 'blur' }]">
                    <el-input class="short-input" v-model="editForm.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>
                <el-form-item label="状态" prop="status" :rules="[{ required: true, message: '请选择状态', trigger: 'blur' }]">
                    <el-radio-group v-model="editForm.status">
                        <el-radio label="1">启用</el-radio>
                        <el-radio label="0">禁用</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <div slot="footer">
                <el-button @click="showEditDialog = false">取消</el-button>
                <el-button type="primary" @click="saveUser">保存</el-button>
            </div>
        </el-dialog>

        <el-dialog :visible.sync="showAddDialog" title="新增用户" :append-to-body="true" :fullscreen="true"
            style="width: 700px; top: 20%; left: 30%; height: 600px;" center>
            <el-form ref="AddForm" :model="AddForm" label-width="120px">
                <el-form-item label="用户名" prop="username" :rules="[{ required: true, message: '请输入用户名', trigger: 'blur' }]">
                    <el-input class="short-input" v-model="AddForm.username" placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password"
                    :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }]">
                    <el-input class="short-input" type="password" v-model="AddForm.password" placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item label="确认密码" prop="password"
                    :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }]">
                    <el-input class="short-input" type="password" v-model="AddForm.confirmPassword"
                        placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email"
                    :rules="[{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }]">
                    <el-input class="short-input" v-model="AddForm.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="phone"
                    :rules="[{ required: true, message: '请输入手机号', trigger: 'blur' }, { pattern: /^1[34578]\d{9}$/, message: '手机号格式不正确', trigger: 'blur' }]">
                    <el-input class="short-input" v-model="AddForm.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>
                <el-form-item label="状态" prop="status" :rules="[{ required: true, message: '请选择状态', trigger: 'blur' }]">
                    <el-radio-group v-model="AddForm.status">
                        <el-radio label="1">启用</el-radio>
                        <el-radio label="0">禁用</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <div slot="footer">
                <el-button @click="showAddDialog = false">取消</el-button>
                <el-button type="primary" @click="saveAddUser('AddForm')">保存</el-button>
            </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogVisible" title="删除用户" style="width: 500px;top: 20%; left: 30%;">
            <el-form :model="form">
                <el-form-item prop="userName">
                    <span>确认删除？</span>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="handleDelete">确认</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script>
import { allusers, keywordsearch, deleteuser, updatedata, register } from '@/api/account'
export default {
    data() {
        return {
            tableData: [],
            userData: [],
            total: 0,
            currentPage: 1,
            pageSize: 10,
            searchValue: '',
            showAddDialog: false,
            dialogVisible: false,
            showEditDialog: false,
            editForm: {
                id: '',
                username: '',
                password: '',
                email: '',
                phone: '',
                status: '1',
            },
            AddForm: {
                username: '',
                password: '',
                email: '',
                phone: '',
                status: '',
            },
            roles: [
                {
                    label: '管理员',
                    value: 'admin',
                },
                {
                    label: '普通用户',
                    value: 'user',
                },
            ],
        };
    },
    created() {
        this.fetchUserData()
            .then(() => {
                this.search()
            })
    },
    methods: {
        fetchUserData() {
            return allusers().then(response => {
                if (response !== null) {
                    this.userData = response
                }
            })
        },
        handlePageChange(newPage) {
            this.currentPage = newPage;
            this.fetchUserData();
        },
        editUser(user) {
            this.showEditDialog = true;
            this.editForm = {
                id: user.ID,
                username: user.userName,
                email: user.email,
                password: user.passWord,
                phone: user.phone,
                status: user.status.toString(),
            };

        },
        deleteUser(user) {
            this.form = Object.assign({}, user);
            this.dialogVisible = true;
            this.deleteuser = user

        },
        handleDelete() {
            deleteuser({ id: this.deleteuser.ID })
            const index = this.tableData.indexOf(this.deleteuser);
            this.tableData.splice(index, 1);
            this.dialogVisible = false;
        },
        showAdd() {
            this.showAddDialog = true
        },
        saveAddUser(formData) {
            if (
                this.AddForm.username &&
                this.AddForm.password &&
                this.AddForm.confirmPassword &&
                this.AddForm.email &&
                this.AddForm.phone &&
                this.AddForm.status
            ) {
                if (this.AddForm.password !== this.AddForm.confirmPassword) {
                    this.$message.error("两次输入的密码不一致");
                    return;
                }
                register({
                    username: this.AddForm.username,
                    password: this.AddForm.password,
                    email: this.AddForm.email,
                    phone: this.AddForm.phone,
                    status: this.AddForm.status,
                }).then(() => {
                    this.$message.success("添加成功")
                    this.showAddDialog = false
                    setTimeout(() => {
                        window.location.reload()
                    }, 0)
                }).catch((error) => {
                    this.$message.error(error.message)
                    this.showAddDialog = false
                });
            } else {
                this.$message("请完整填写信息")
            }
        },
        saveUser() {
            this.$refs.editForm.validate((valid) => {
                if (valid) {
                    updatedata(this.editForm).then(() => {
                        this.$message.success("编辑成功")
                        this.showEditDialog = false;
                        setTimeout(() => {
                            window.location.reload()
                        }, 0)
                    }).catch((error) => {
                        this.$message.error(error.message)
                        this.showEditDialog = false;
                    });
                } else {
                    return false;
                }
            });
        },
        search() {
            if (!this.searchValue) {
                // 如果搜索关键字为空，则显示所有数据
                this.tableData = this.userData;
                //alert(JSON.stringify(this.tableData));
                return;
            } else {
                keywordsearch({ keyword: this.searchValue }).then(response => {
                    this.tableData = response
                    // alert(JSON.stringify(this.tableData));
                })
            }

        },
        updateStatus(row) {
            updatedata({ id: row.ID, status: row.status.toString() })
                .then(response => {
                    console.log(response.data)
                })
                .catch(error => {
                    console.log(error)
                })
        },
    },
};
</script>
</template>
<style scoped>
.user-management {
    padding: 30px;
    background-color: #f5f5f5;
}

.user-table {
    margin-top: 20px;
}

.table-action {
    display: flex;
    justify-content: space-around;
}

.search-container {
    display: flex;
    justify-content: flex-end;
    margin-top: 20px;
}

.search-input {
    margin-right: 10px;
    width: 200px;
}

.user-form {
    width: 400px;
}

.short-input {
    width: 350px;
}

.form-item-inline {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}

.form-item-inline label {
    flex-basis: 120px;
    flex-shrink: 0;
    margin-right: 20px;
}</style>

