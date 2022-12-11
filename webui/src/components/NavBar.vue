<script>
export default {
	data() {
		return {
			profile: ""
		}
	},
	methods: {
		async search() {
			this.$router.push("/users/"+this.profile)
		}
	}
}
</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
		<input type="search" placeholder="Search" id="searchIn" v-model="profile">
		<button type="submit" id="btn" @click="search">
			<svg id="search-icon" class="search-icon" viewBox="0 0 24 24">
				<path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
				<path d="M0 0h24v24H0z" fill="none"/>
			</svg>
		</button>
	</header>
</template>

<style>
	/* Removes the browser default x from search input */
	input[type="search"]::-webkit-search-cancel-button,
	input[type="search"]::-webkit-search-results-decoration { display: none; }
	#bar {
		position: relative;
		height: 60px;
		text-align: center;
	}

	svg {
	color: #fff;
	fill: currentColor;
	width: 10%;
	padding-top: 2px;
	}

	#btn {
		width: 44px;
  		height: 44px;
		all: unset;
	  	cursor: pointer;
		background-color: rgba(var(--bs-dark-rgb), var(--bs-bg-opacity)) !important;
		border: none;
		left: -77px;
		position: relative;
	}

	#search-container {
		position: relative;
		top: 10px;
	}

	#searchIn {
		border-radius: 5px;
   		border: 2px solid black;
		text-align: center;
	}

	#magn {
		width: 20px;
	}
</style>
