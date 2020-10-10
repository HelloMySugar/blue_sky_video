<template>
    <div id="admin">
        <el-row id="navbar" class="main el-table__fixed-header-wrapper">
            <!--logo-->
            <el-col :span="12" class="right">
                <a href="/home" class="tit">
                    <img src="~assets/img/icon/logo.png" class="logo">
                    <span>Blue Sky Video - Admin System</span>
                </a>
            </el-col>
            <el-col :span="12" style="text-align: right; margin-top: 12px">
                <el-button @click="$router.push('/upload')" style="border-color: #3a8eff; color: #3a8eff; font-size: 15px; padding: 5px 10px" size="mini">Upload Video</el-button>
            </el-col>
        </el-row>
        <!--防止下面的盒子上移动，nav-bar脱离了标准流-->
        <div style="height: 50px;"/>
        <div class="main">
            <el-table
                :data="tableData"
                :header-cell-style="getRowClass"
                border
                style="width: 100%; margin-top: 20px">
                <el-table-column
                    type="index" :index="indexMethod"
                    label="No."
                    align="center"
                    width="60">
                </el-table-column>
                <el-table-column
                    prop="vname"
                    label="Name"
                    align="center">
                </el-table-column>
                <el-table-column
                    prop="keyword"
                    align="center"
                    label="Keyword">
                </el-table-column>
                <el-table-column
                    prop="releaseTime"
                    align="center"
                    label="Release Time"
                    width="120">
                </el-table-column>
                <el-table-column
                    align="center"
                    label="More"
                    width="240">
                    <template slot-scope="scope">
                        <div style="display: block">
                            <div>
                                <el-button @click="gotoEdit(scope.row)" style="border-color: #3a8eff; color: #3a8eff; font-size: 15px; padding: 4px 7px" size="mini">Edit</el-button>
                                <span style="color: black">｜</span>
                                <el-button @click="toPlayer(scope.row)" style="border-color: #3a8eff; color: #3a8eff; font-size: 15px; padding: 4px 7px" size="mini">Detail</el-button>
                                <span style="color: black">｜</span>
                                <el-button @click="deleteDataBy(scope.row)" style="border-color: #3a8eff; color: #3a8eff; font-size: 15px; padding: 4px 7px" size="mini">Delete</el-button>
                            </div>
                        </div>
                    </template>
                </el-table-column>
            </el-table>

            <el-col :span="24" class="pagination">
                <el-pagination
                    background
                    :small="screenWidth <= 768"
                    hide-on-single-page
                    layout="total, prev, pager, next"
                    @current-change="handleCurrentChange"
                    :current-page="currentPage"
                    :page-size="10"
                    :total="totalCount">
                </el-pagination>
            </el-col>
        </div>
    </div>
</template>

<script>
    import {toPlayerMixin} from "../../assets/js/mixin";
    import {getPageBean} from '../../network/video'
    import {deleteVideo} from '../../network/admin'
    export default {
        name: 'Admin',
        mixins: [toPlayerMixin],
        data() {
            return {
                currentPage: 1,
                totalCount: 18,
                screenWidth: document.body.clientWidth,     // 屏幕宽,为了控制分页条的大小
                tableData: []
            }
        },
        methods: {
            getRowClass ({ row, column, rowIndex, columnIndex }) {
                if (rowIndex === 0) {
                    return 'background:#f5f5f5;color:#333;text-align:center;font-size:15px;font-weight:500;'
                } else {
                    return ''
                }
            },
            indexMethod (index) {
                let number = index + 1
                number = (Array(2).join('0') + number).slice(-2)
                return number
            },
            gotoEdit(row) {
                localStorage.setItem('editVideo', JSON.stringify(row))
                this.$router.push('/editVideo')
            },
            async getVideoList() {
                let a = await getPageBean(0, this.currentPage, 10)
                this.totalCount = a.totalCount
                this.tableData = a.list
            },
            handleCurrentChange(currentPage) {
                this.currentPage = currentPage
                this.getVideoList()
                // 回到顶部
                scrollTo(0,0)
            },
            deleteDataBy (row) {
                this.$confirm('Are you sure to delete it?', 'Tips', {
                    confirmButtonText: 'OK',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(async () => {
                    let a = await deleteVideo(row.vid)
                    if(a !== undefined) {
                        this.currentPage = 1
                        this.$message.success('Delete success')
                        this.getVideoList()
                    } else {
                        this.$message.error('Delete failed')
                    }
                })
            },
        },
        mounted() {
            this.getVideoList()
        }
    }
</script>

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
        padding-left: 10%;
        padding-right: 10%;
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

    .pagination {
        text-align: center;
        padding: 10px;
    }

</style>
