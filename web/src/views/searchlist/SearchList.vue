<template>
	<div id="search-list">
		<!--搜索结果面包屑-->
		<el-breadcrumb class="bread"
						separator-class="el-icon-arrow-right"
						v-if="this.$route.path.indexOf('searchlist') > -1">
			<el-breadcrumb-item :to="{ path: '' }"><a href="/home">Back To Home</a></el-breadcrumb-item>
			<el-breadcrumb-item>Search results：total <span class="result">({{$store.state.pageBean.totalCount}})</span></el-breadcrumb-item>
		</el-breadcrumb>

		<el-row class="movie-list" v-if="$store.state.pageBean.totalCount !== 0 && $store.state.pageBean.totalCount !== undefined ">
			<el-col :md="6" :sm="8" :xs="12" v-for="(video,index) in $store.state.pageBean.list" :key="index">
				<movie-item :video="video"/>
			</el-col>

			<!--
					分页按钮：
						page-size：每页显示条数
						total：总条数
						hide-on-single-page： 页数为一时隐藏
				-->
			<el-col :span="24" class="pagination">
				<el-pagination
								background
								hide-on-single-page
								layout="prev, pager, next"
								@current-change="handleCurrentChange"
								:current-page="$store.state.activePage"
								:page-size="8"
								:total="$store.state.pageBean.totalCount">
				</el-pagination>
			</el-col>
		</el-row>
		<el-row v-else class="not-result">
			<el-col :span="12" :offset="6">
				<img src="~assets/img/icon/not-result.png">
				<div>No results found</div>
			</el-col>
		</el-row>

	</div>

</template>

<script>
	import MovieItem from "components/movieitem/MovieItem";
	import SearchBox from "components/searchbox/SearchBox";
	export default {
		name: "SearchList",
		data() {
			return {
				currentPage: 1
			}
		},
		components: {SearchBox, MovieItem},
		methods: {
			handleCurrentChange(currentPage) {
				this.currentPage = currentPage
				this.$store.commit("updatePage", currentPage)
				this.$store.dispatch("getPageBeanByValue")
			}
		},
		created() {
			// 当页面创建后，页码变为1
			this.$store.commit('updatePage',1)
		}
	}
</script>

<style scoped>
	#search-list {
		padding-left: 6%;
		padding-right: 6%;
		padding-top: 30px;
	}

	.bread {
		font-size: 15px;
	}

	.movie-list {
		padding-top: 15px;
	}

	.pagination {
		text-align: center;
		padding: 10px;
	}

	.result {
		color: red;
	}

	.not-result {
		padding-top: 100px;
		padding-bottom: 100px;
		text-align: center;
	}
</style>
