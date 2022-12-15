<script>
export default {
	data() {
		return {
			username: null,
			badUsername: false,
			identifier: ""
		}
	},
	methods: {
		async doLogin() {
			this.badUsername = false 
			try {
				let response = await this.$axios.post("/session/", {
					username: this.username
				});
				this.identifier = response.data;
				localStorage.setItem("identifier",  this.identifier)
				localStorage.setItem("username",  this.username)
				this.$router.push({ path: `/stream/${this.username}` })
			} catch (e) {
				if (e.code == "ERR_BAD_REQUEST") {
					this.badUsername = true
				}
				return
			}
		}
	}
}
</script>

<template>
	<div id="wrapper">
		<div class="main-content">
			<div class="header">
				<h2>WASAPhoto</h2>
			</div>
			<div class="l-part">
				<input type="text" v-model="username" placeholder="Username" class="input-1" required/>
				<input type="button" @click="doLogin" value="Log in" class="mybtn" />
			</div>
			<div v-if="badUsername" class="badReq">
				Username not valid
			</div>
		</div>

	</div>

</template>

<style>
* {
  margin: 0px;
  padding: 0px;
}

body {
  background-color: #eee;
}

#wrapper {
  width: 500px;
  height: 700px;
  overflow: hidden;
  border: 0px solid #000;
  margin: 50px auto;
  padding: 10px;
}

.main-content {
  width: 300px;
  height: 40%;
  margin: 10px auto;
  background-color: #fff;
  border: 2px solid #e6e6e6;
  padding: 40px 50px;
}

.header {
	margin: 0 auto;
	text-align: center;
	position: relative;
}

.l-part {
	position: relative;
	top: 30px;
}

.input-1 {
  width: 100%;
  margin-bottom: 5px;
  padding: 8px 12px;
  border: 1px solid #dbdbdb;
  box-sizing: border-box;
  border-radius: 3px;
}

.mybtn {
  width: 100%;
  background-color: rgba(0 0 0 / 25%);
  border: 1px solid black;
  padding: 5px 12px;
  color: #fff;
  font-weight: bold;
  border-radius: 3px;
  transition-duration: 10ms;
}

.badReq {
	position: relative;
	top: 50px;
	text-align: center;
	color: red;
}


</style>
