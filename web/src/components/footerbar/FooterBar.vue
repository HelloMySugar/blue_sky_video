<template>
	<div id="footer-bar">
		<div class="content">
			© 2020 Blue Sky Video
			<br>
			Address: 1800 Wuhou District, Chengdu, Sichuan, China.
			<br>
		</div>

	</div>
</template>

<script>
	export default {
		name: "FooterBar"
	}
</script>

<style scoped>

	#footer-bar {
		position: relative;
		bottom: 0;
		left: 0;
		right: 0;

		padding-left: 6%;
		padding-right: 6%;
		margin-top: 30px;
		height: 80px;
		border-top: 2px solid rgba(34, 36, 38, .15);
		box-shadow: 0 1px 2px 0 rgba(34, 36, 38, .15);
		text-align: center;

	}

	.content {
		padding-top: 20px;
		font-size: 14px;
	}
	.logo {
		width: 30px;
		cursor: pointer;
	}
	.author,.statement {
		padding-top: 5px;
	}

	.weixin {
		width: 30px;
		cursor: pointer;

	}
	.er {
		width: 150px;
	}




</style>
