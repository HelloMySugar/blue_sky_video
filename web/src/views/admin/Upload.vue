<template>
    <div id="upload">
        <div class="uploadBg">
            <el-row id="navbar" class="main el-table__fixed-header-wrapper">
                <!--logo-->
                <el-col :span="12" class="right">
                    <a href="/home" class="tit">
                        <img src="~assets/img/icon/logo.png" class="logo">
                        <span>Blue Sky Video - Upload Video</span>
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
                    <el-form-item label="Video" prop="video" required>
                        <el-upload
                            ref="upload"
                            class="upload-demo"
                            drag
                            :auto-upload="false"
                            :on-change="handleChange"
                            :on-remove="handleRemove"
                            :file-list="uploadList"
                            :http-request="requestUploadVideo"
                            :limit="1"
                            :on-progress="uploadVideoProcess"
                            action="">
                            <i class="el-icon-upload"></i>
                            <div class="el-upload__text">Drag and drop the file here, or <em>Click to upload</em></div>
                        </el-upload>
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
                    <el-form-item label="Keywords" prop="keywords">
                        <el-input v-model="ruleForm.keywords" placeholder="Separate keywords by ','"></el-input>
                    </el-form-item>
                    <el-form-item style="text-align: center; margin-left: -50px">
                        <el-button type="primary" @click="submitForm('ruleForm')">Upload</el-button>
                        <el-button @click="resetForm('ruleForm')" plain type="warning">Reset</el-button>
                        <el-button @click="$router.go(-1)" >Cancel</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
        <div class="progressBg" v-if="videoFlag">
            <el-progress type="circle" :percentage="videoUploadPercent"></el-progress>
        </div>
    </div>
</template>

<script>
    import {getCategory} from '../../network/video'
    import axios from 'axios'
    export default {
        name: 'Upload',
        data() {
            return {
                uploadList: [],
                fileZNames: '.mp4, .mpeg4, .ogg, .webm',
                categoryList: [],
                ruleForm: {
                    vname: '',
                    cid: '',
                    date: '',
                    keywords: '',
                    video: ''
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
                    keywords: [
                        { required: true, message: 'Please input keywords', trigger: 'blur' }
                    ],
                    video: [
                        { required: true, message: 'Please upload video', trigger: 'change' }
                    ],
                },
                videoFlag: false,
                videoUploadPercent: 0
            };
        },
        methods: {
            handleRemove(file, fileList) {
                this.uploadList = []
                this.ruleForm.video = ''
            },
            handleChange(file) {
                //console.log(file)
                this.uploadList = []
                this.ruleForm.video = ''
                let tmpFileName = file.name.toLowerCase()
                let lastIndex = tmpFileName.lastIndexOf('.')
                let ext = tmpFileName.substr(lastIndex)
                if (this.fileZNames.indexOf(ext) < 0) {
                    this.$message.error('Please choose mp4/ogg/webm type file')
                    return
                }
                this.uploadList = [file]
                this.ruleForm.video = file
            },
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.$refs.upload.submit()
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
                this.uploadList = [];
            },
            async getCategoryList() {
                let a = await getCategory()
                this.categoryList = a
            },
            async requestUploadVideo(file) {
                let that = this
                const formData = new FormData();
                formData.append('video', file.file);
                formData.append('vname', this.ruleForm.vname);
                formData.append('cid', this.ruleForm.cid);
                formData.append('releaseTime', this.ruleForm.releaseTime);
                formData.append('keywords', this.ruleForm.keywords);
                axios.post('api/v0/videos', formData, {
                    onUploadProgress(progressEvent) {
                        if (progressEvent.lengthComputable) {
                            that.videoFlag = true;
                            let val = (progressEvent.loaded / progressEvent.total * 100).toFixed(0);
                            that.videoUploadPercent = Number(val);
                        }
                    }
                }).then((response) => {
                    // console.log(response)
                    this.videoFlag = false;
                    this.videoUploadPercent = 0;
                    this.$message.success('Upload success')
                    this.$router.go(-1)
                }).catch((error) => {
                    this.videoFlag = false;
                    this.videoUploadPercent = 0;
                    this.$message.error(error.message)
                    this.handleRemove()
                });
            },
            // 请求OK，但是无法计算进度
            // async requestUploadVideo(file) {
            //     const formData = new FormData();
            //     formData.append('video', file.file);
            //     formData.append('vname', this.ruleForm.vname);
            //     formData.append('cid', this.ruleForm.cid);
            //     formData.append('releaseTime', this.ruleForm.releaseTime);
            //     formData.append('keywords', this.ruleForm.keywords);
            //     const xhr = new XMLHttpRequest();
            //     xhr.open('post', 'api/v0/videos', true);
            //     xhr.onload = (res) => {
            //         this.videoFlag = false;
            //         this.videoUploadPercent = 0;
            //         if(res.currentTarget.status !== 200) {
            //             this.$message.error(res.currentTarget.responseText)
            //             this.handleRemove()
            //             return;
            //         }
            //         this.$message.success('Upload success')
            //         this.$router.go(-1)
            //     };
            //     xhr.onerror = () => {
            //         this.videoFlag = false;
            //         this.videoUploadPercent = 0;
            //         this.$message.error('Upload failed, please retry')
            //     };
            //     xhr.ontimeout = function timeout() {
            //         this.videoFlag = false;
            //         this.videoUploadPercent = 0;
            //         this.$message.error('Upload timeout, please retry')
            //     };
            //     xhr.send(formData);
            // },
            getSTime (val) {
                this.ruleForm.releaseTime = val;
            },
            uploadVideoProcess(event, file, fileList){
                this.videoFlag = true
                this.videoUploadPercent = file.percentage.toFixed(0);
                console.log(this.videoUploadPercent)
            },
        },
        mounted() {
            this.getCategoryList()
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
    .el-progress__text {
        color: white;
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

    .progressBg {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        z-index: 999;
        background-color: black;
        display: flex;
        justify-content: center;
        align-items: center;
        opacity: 0.75;
    }
</style>
