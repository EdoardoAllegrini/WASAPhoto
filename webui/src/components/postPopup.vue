<script>
export default {
    emits: ["succPost"],
    data() {
        return {
            url: null,
            phSel: false,
            caption: "",
            image: null,
        }
    },
    methods: {
        clickOutside(event) {
            if (event.composedPath()['0'].className == 'popupP') {this.exit()}
        },
        exit() {
            this.$parent.$data.popupPost = false
            document.body.style.overflow = "scroll"
            return
        },
        handleChange(e) {
            this.image = e.target.files[0];
            this.url = URL.createObjectURL(this.image);
            if (this.url != null) {this.phSel = true}
            else {this.phSel = false}
        },
        async postPhoto() {
            try {
                var path = `/users/${localStorage.username}/media/`
                var formData = new FormData()
                formData.append("photoCaption", this.caption)
                formData.append("photoFile", this.image)
				let response = await this.$axios.post(path, formData);
				this.some_data = response.data;
                this.exit()
                this.succPost()
			} catch (e) {
                console.log(e)
				this.errormsg = e.toString();
			}
        },
		async refresh() {
			window.location.reload()
		},
        succPost() {
            document.getElementById("sscp").style.display = "block"
            setTimeout(function () {            
                document.getElementById("sscp").style.display = "none"
                return 1},
                3000);
            this.$emit("succPost")
        }
    }
}
</script>

<template>
    <div class="popupP" @click="clickOutside">
        <div class="innerP">
            <slot />
            <div class="tit">
                <div class="close" @click="exit">
                    <svg class="xB" viewPort="0 0 12 12">
                        <line x1="1" y1="11" 
                            x2="11" y2="1" 
                            stroke="black" 
                            stroke-width="2"/>
                        <line x1="1" y1="1" 
                            x2="11" y2="11" 
                            stroke="black" 
                            stroke-width="2"/>
                    </svg>
                </div>
                <h2 id="descr">Create a new post</h2>
            </div>
            <hr>
            <div id="preview">
                <img v-if="true" :src="url" />
            </div>
            <div v-if="phSel" id="capt">
                <textarea v-model="caption" placeholder="Write a caption..." autocomplete="off" autocorrect="off" maxlength="250"></textarea>
            </div>
            <div class="lopi">
                <label class="containerPost" id="file" for="grio">
                    Select File
                </label>
                <input type="file" @change="handleChange" accept="image/png, image/jpeg, image/jpg" id="grio">
                <button v-if="phSel" class="containerPost" id="send" @click="postPhoto">post</button>
            </div>

        </div>
    </div>
</template>

<style>

#capt {
    position: relative;
    height: 86px;
    padding: 10px;
    overflow: hidden;
}
#capt textarea {
    width: 100%;
    height: 100%;
    resize: none;
    border: none;
    background-color: transparent;
    overflow: auto;
    outline: none;
}
#preview {
    margin-top: 20px;
    justify-content: center;
    align-items: center;
    position: relative;
    text-align: center;
    display: flex;
    height: 380px;
}

#preview img {
    max-width: 100%;
    max-height: 100%;
}
input[type="file"] {
    display: none;
}
.containerPost {
    border: 1px solid #ccc;
    cursor: pointer;
    background-color: rgb(0, 149, 246);
    border: 1px solid transparent;
    color: rgb(255,255,255);
    border-radius: 4px;
    text-align: center;
    position: relative;
}
#send {
    height: 23px;
    left: 320px;
    width: 36px;
    background-color: white;
    border: 1px solid black;
    color: black;
}
.lopi {
    position: relative;
    width: inherit;
    display: flex;
    justify-content: center;
}
#send:hover {
    background-color: #ccc;
}
#file {
    height: 23px;
    width: 100px;
    position: absolute;
}
#descr {
    top: 10px;
    text-align: center;
    height: 40px;
    position: relative;
    bottom: 0px;
    margin: 0;
    width: fit-content;
}
.inner hr {
    position: relative;
    bottom: 0px;
    margin: 0;
}
.close {
    width: 20px;
    height: 20px;
    cursor: pointer;
    position: relative;
    float: right;
    bottom: 0px;
    left: 460px;
}
.close:hover {
    transform: scale(1.2);
}
.xB {
    vertical-align: text-top;
    height: 20px;
    width: 20px;
}
.tit {
    display: flex;
    height: 8%;
    justify-content: center;
}
.popupP {
    overflow: auto;
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    bottom: 0;
    z-index: 99;
    background-color: rgba(0,0,0,0.2);
    display: flex;
    align-items: center;
    justify-content: center;
}
.innerP {
    text-align: center;
    border-radius: 10px;
    background: #FFF;
    padding: 0px;
    width: 700px;
    height: 600px;
}
.popup {
    overflow: auto;
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    bottom: 0;
    z-index: 99;
    background-color: rgba(0,0,0,0.2);
    display: flex;
    align-items: center;
    justify-content: center;
}
.inner {
    text-align: center;
    border-radius: 10px;
    background: #FFF;
    padding: 0px;
    width: 700px;
    height: 600px;
}
</style>