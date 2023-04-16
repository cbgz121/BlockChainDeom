<template>
    <div class="user-management">
        <el-row>
            <el-col :span="12">
                <el-input v-model="searchValue" placeholder="请输入搜索关键字" suffix-icon="el-icon-search"
                    @change="search"></el-input>
            </el-col>
            <el-col :span="12" class="text-right">
                <el-button type="primary" icon="el-icon-plus">新增用户</el-button>
            </el-col>
        </el-row>
        <el-table :data="tableData" class="user-table">
            <el-table-column prop="ID" label="ID" sortable></el-table-column>
            <el-table-column prop="userName" label="用户名" sortable></el-table-column>
            <el-table-column prop="passWord" label="密码" sortable></el-table-column>
            <el-table-column prop="email" label="邮箱" sortable></el-table-column>
            <el-table-column prop="phone" label="手机号" sortable></el-table-column>
            <!-- <el-table-column prop="role" label="角色" sortable>
                <template slot-scope="scope">
                    <el-tag :type="getRoleTagType(scope.row.role)">{{ scope.row.role }}</el-tag>
                </template>
            </el-table-column> -->
            <el-table-column label="状态">
                <template slot-scope="scope">
                    <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0" active-text="启用"
                        inactive-text="禁用"></el-switch>
                </template>
            </el-table-column>
            <!-- <el-table-column label="注册时间" sortable>
                <template slot-scope="scope">
                    {{ formatDate(scope.row.created_at) }}
                </template>
            </el-table-column> -->
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button type="text" size="small" @click="editUser(scope.row)">编辑</el-button>
                    <el-button type="text" size="small" class="danger" @click="deleteUser(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination :total="total" :current-page="currentPage" :page-size="pageSize" @current-change="handlePageChange"
            class="pagination"></el-pagination>
        <el-dialog :visible.sync="showEditDialog" title="编辑用户" :append-to-body="true" :fullscreen="true">
            <el-form ref="editForm" :model="editForm" label-width="120px">
                <el-form-item label="用户名" prop="username" :rules="[{ required: true, message: '请输入用户名', trigger: 'blur' }]">
                    <el-input v-model="editForm.username" placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password"
                    :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }]">
                    <el-input type="password" v-model="editForm.password" placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email"
                        :rules="[{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }]">
                        <el-input v-model="editForm.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="phone"
                    :rules="[{ required: true, message: '请输入手机号', trigger: 'blur' }, { pattern: /^1[34578]\d{9}$/, message: '手机号格式不正确', trigger: 'blur' }]">
                    <el-input v-model="editForm.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>
                <!-- <el-form-item label="状态" prop="status" :rules="[{ required: true, message: '请选择状态', trigger: 'blur' }]">
                    <el-radio-group v-model="editForm.status">
                        <el-radio :label="1">启用</el-radio>
                        <el-radio :label="0">禁用</el-radio>
                    </el-radio-group>
                </el-form-item> -->
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
        <el-dialog :visible.sync="dialogVisible" title="删除用户"  style="width: 500px;top: 20%; left: 30%;">
                    <el-form :model="form">
                        <el-form-item prop="userName">
                            <!-- <el-input v-model="form.username"></el-input> -->
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
import { allusers, keywordsearch, deleteuser, updatedata } from '@/api/account'
export default {
    data() {
        return {
            tableData: [],
            userData: [],
            total: 0,
            currentPage: 1,
            pageSize: 10,
            searchValue: '',
            dialogVisible: false,
            showEditDialog: false,
            editForm: {
                username: '',
                password: '',
                // confirmPassword: '',
                email: '',
                phone: '',
                // role: '',
                status: '1',
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
            // set editForm data and show dialog
            this.showEditDialog = true;
            this.editForm = {
                username: user.userName,
                email: user.email,
                password: user.passWord,
                // confirmPassword: '',
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
            // Deleted.push(this.deleteuser.userName)
            this.dialogVisible = false;
        },
        saveUser() {
            alert("9999999999999")
            this.$refs.editForm.validate((valid) => {
                if (valid) {
                    updatedata(this.editForm).then(() => {
                        this.$message.success("编辑成功")
                        this.showEditDialog = false;
                    }).catch((error) => {
                        this.$message.error(error.message)
                        this.showEditDialog = false;
                    });
                } else {
                    return false;
                }
            });
        },
        // getRoleTagType(role) {
        //     if (role === 'admin') {
        //         return 'danger';
        //     } else if (role === 'user') {
        //         return 'success';
        //     } else {
        //         return '';
        //     }
        // },
        search() {
            if (!this.searchValue) {
                // 如果搜索关键字为空，则显示所有数据
                this.tableData = this.userData;
                return;
            } else {
                keywordsearch({ keyword: this.searchValue }).then(response => {
                    this.tableData = response
                })
            }

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
}</style>
