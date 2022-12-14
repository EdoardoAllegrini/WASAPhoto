<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>
	<div>
		<main class="routerView">
			<RouterView :key="$route.fullPath"/>
		</main>
	</div>
</template>

<style>
</style>
