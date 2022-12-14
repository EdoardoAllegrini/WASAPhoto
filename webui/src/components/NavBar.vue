<script>
import PostPopup from './postPopup.vue'

export default {
	data() {
		return {
			profile: "",
			menu: false,
			popupPost: false
		}
	},
	methods: {
		postImage() {
			return
		},
		search() {
			if (this.profile) {
				this.$router.push("/users/"+this.profile)
				document.getElementById("searchIn").value = null
			}
			return
		},
		showMenu() {
			this.menu = !this.menu
			setTimeout(() => this.menu = false, 2000)
		},
		goStream() {
			this.$router.push(`/stream/${localStorage.username}`)
		}
	},
    components: { PostPopup },
}
</script>

<template>
	<header id="hea" class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
		<div class="sr">
			<input type="search" placeholder="Search" id="searchIn" spellcheck="false" v-model="profile">
			<button id="search-button" @click="search">
				<svg id="search-icon" class="search-icon" viewBox="0 0 24 24">
					<path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
					<path d="M0 0h24v24H0z" fill="none"/>
				</svg>
			</button>
		</div>
		<div class="dh" @click="showMenu">
			<svg class="svgL"><line x1="3" x2="21" y1="4" y2="4"></line><line x1="3" x2="21" y1="12" y2="12"></line><line x1="3" x2="21" y1="20" y2="20"></line></svg>
		</div>
		<div v-if="menu" class="menu">
			<button id="sout" type="submit" @click="$router.push('/')">Sign out</button>
		</div>
		<div class="dh" @click="goStream">
			<svg class="svgL"><path d="M9.005 16.545a2.997 2.997 0 0 1 2.997-2.997A2.997 2.997 0 0 1 15 16.545V22h7V11.543L12 2 2 11.543V22h7.005Z"></path></svg>		
		</div>
		<div class="dh" @click="popupPost=!popupPost">
			<svg class="svgL"><path d="M2 12v3.45c0 2.849.698 4.005 1.606 4.944.94.909 2.098 1.608 4.946 1.608h6.896c2.848 0 4.006-.7 4.946-1.608C21.302 19.455 22 18.3 22 15.45V8.552c0-2.849-.698-4.006-1.606-4.945C19.454 2.7 18.296 2 15.448 2H8.552c-2.848 0-4.006.699-4.946 1.607C2.698 4.547 2 5.703 2 8.552Z"></path><line x1="6.545" x2="17.455" y1="12.001" y2="12.001"></line><line x1="12.003" x2="12.003" y1="6.545" y2="17.455"></line></svg>		
		</div>
		<PostPopup v-if="popupPost">

		</PostPopup>
	</header>
</template>

<style>
	.dh {
		box-sizing: border-box;
		height: 24px;
		width: 24px;
		cursor: pointer;
		float: right;
		padding-top: 12px;
		padding-right: 50px;
	}

	.dh:hover {
		transform: scale(1.1);
	}

	.dh line, .dh path {
		fill: none;
		stroke: currentcolor;
		stroke-linecap: round;
		stroke-linejoin: round;
		stroke-width: 2;
	}
	.svgL {
		height: 24px;
    	width: 24px;
		color: rgb(255, 255, 255);
    	fill: rgb(38, 38, 38);
	}

	#hea {
		display: inline-block;
		width: 100%;
	}
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
	width: 25px;
	padding-top: 2px;
	}

	.menu {
		position: absolute;
		right: 1px;
		background-color: rgba(var(--bs-dark-rgb), var(--bs-bg-opacity)) !important;
	}

	#sout {
		margin: 5px;
		background-color: rgba(var(--bs-dark-rgb), var(--bs-bg-opacity)) !important;
		font-size: 17px;
		border: none;
		cursor: pointer;
		color: white;
	}

	#sout:hover {
		color: gray;
	}

	.sr {
		display: inline-block;
		position: relative;
		left: 25%;
	}

	#search-button {
		display: inline-block;
	  	cursor: pointer;
		background-color: rgba(var(--bs-dark-rgb), var(--bs-bg-opacity)) !important;
		border: none;
		position: relative;
	}

	#search-container {
		position: relative;
		top: 10px;
		display: inline-block;
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
