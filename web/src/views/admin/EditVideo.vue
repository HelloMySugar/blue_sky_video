<template>
    <div id="editVideo">
        <el-row id="navbar" class="main el-table__fixed-header-wrapper">
            <!--logo-->
            <el-col :span="12" class="right">
                <a href="/home" class="tit">
                    <img src="~assets/img/icon/logo.png" class="logo">
                    <span>Blue Sky Video - Edit Video</span>
                </a>
            </el-col>
        </el-row>
        <!--防止下面的盒子上移动，nav-bar脱离了标准流-->
        <div style="height: 50px;"/>
        <div class="main">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm" style="margin-top: 30px; border: 1px lightgray solid; border-radius: 6px; padding: 20px 20px 0px 0px">
                <el-form-item label="Name" prop="vname">
                    <el-input v-model="ruleForm.vname" placeholder="Please input video name"></el-input>
                </el-form-item>
                <el-form-item label="Category" prop="cid">
                    <el-select v-model="ruleForm.cid" placeholder="Select category">
                        <el-option v-for="(object, index) in categoryList"
                                   :key="index"
                                   :label="object.cname"
                                   :value="object.cid"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Release Time" required>
                    <el-col :span="24">
                        <el-form-item prop="releaseTime">
                            <el-date-picker type="date"
                                            placeholder="Choose date"
                                            v-model="ruleForm.releaseTime"
                                            value-format="yyyy-MM-dd"
                                            @change="getSTime"
                                            style="width: 100%;"></el-date-picker>
                        </el-form-item>
                    </el-col>
                </el-form-item>
                <el-form-item label="Keywords" prop="keyword">
                    <el-input v-model="ruleForm.keyword" placeholder="Separate keywords by ','"></el-input>
                </el-form-item>
                <el-form-item style="text-align: center; margin-left: -50px">
                    <el-button type="primary" @click="submitForm('ruleForm')">Edit</el-button>
                    <el-button @click="$router.go(-1)">Cancel</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
    import {getCategory} from '../../network/video'
    import {editVideo} from '../../network/admin'
    export default {
        name: 'EditVideo',
        data() {
            return {
                categoryList: [],
                ruleForm: {
                    vname: '',
                    cid: '',
                    date: '',
                    keyword: ''
                },
                rules: {
                    vname: [
                        { required: true, message: 'Please input video name', trigger: 'blur' },
                        // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
                    ],
                    cid: [
                        { required: true, message: 'Please choose category', trigger: 'change' }
                    ],
                    releaseTime: [
                        { required: true, message: 'Please choose release time', trigger: 'change' }
                    ],
                    keyword: [
                        { required: true, message: 'Please input keywords', trigger: 'blur' }
                    ]
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.requestEditVideo()
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            async getCategoryList() {
                let a = await getCategory()
                this.categoryList = a
            },
            getSTime (val) {
                this.ruleForm.releaseTime = val;
            },
            async requestEditVideo() {
                let a = await editVideo(this.ruleForm.vid, {
                    vname: this.ruleForm.vname,
                    releaseTime: this.ruleForm.releaseTime,
                    cid: this.ruleForm.cid,
                    keywords: this.ruleForm.keyword
                })
                if(a !== undefined) {
                    this.$message.success('Edit success')
                    this.$router.go(-1)
                } else {
                    this.$message.error('Edit failed')
                }
            }
        },
        mounted() {
            this.getCategoryList()
            this.ruleForm = JSON.parse(localStorage.getItem('editVideo'))
        }
    }
</script>
<style>
    body{
        height: 100%;
        overflow-y: scroll;
    }
    .el-select {
        display: block;
    }
    .el-upload {
        display: block;
    }
    .el-upload-dragger {
        width: 100%;
    }
</style>
<style scoped>
    #navbar {
        height: 50px;
        border-bottom: 1px solid rgba(34, 36, 38, .15);
        box-shadow: 0 1px 2px 0 rgba(34, 36, 38, .15);

        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        background-color: #ffffff;
        z-index: 9;
    }

    .main {
        padding-left: 20%;
        padding-right: 20%;
    }

    .right {
        word-break: keep-all; /* 不换行 */
        white-space: nowrap; /* 不换行 */
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .tit {
        display: block;
        font-weight: 700;
        height: 45px;
        text-decoration: none;
        color: #333;
    }

    .tit span {
        margin-left: 5px;
    }

    .logo {
        width: 30px;
        position: relative;
        top: 5px;
    }
</style>
