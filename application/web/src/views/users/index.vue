<template>
    <div>
        <el-table :data="tableData">
            <el-table-column label="用户名" prop="userName"></el-table-column>
                <el-table-column label="密码" prop="passWord"></el-table-column>
            <el-table-column label="邮箱" prop="email"></el-table-column>
            <el-table-column label="手机号" prop="phone"></el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <!-- <el-button type="primary" size="small" @click="handleEdit(scope.row)">编辑</el-button> -->
                    <el-button type="danger" size="small" @click="handleEdit(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

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
import { queryAccountList, deleteid } from '@/api/account'
let Deleted = ['admin'];
export default {
    
    data() {
        return {
            tableData: [],
            deleteuser: [],
            dialogVisible: false,
            form: {
                username: '',
                email: '',
                phone: '',
            },
        };
    },
    created() {
        queryAccountList().then(response => {
            if (response !== null) {
                // 过滤掉管理员
                this.tableData = response.filter(item =>
                    !Deleted.includes(item.userName)
                )
            }
        })
    },
    methods: {
        handleEdit(row) {
            this.form = Object.assign({}, row);
            this.dialogVisible = true;
            this.deleteuser = row
        },
        handleDelete() {
            deleteid({username:this.deleteuser.userName, password: this.deleteuser.passWord})
            const index = this.tableData.indexOf(this.deleteuser);
            this.tableData.splice(index, 1);
            Deleted.push(this.deleteuser.userName)
            this.dialogVisible = false;
        },
        handleSave() {
            // 保存用户信息的逻辑
            this.dialogVisible = false;
        },
    },
};
</script>

<style scoped>
</style>