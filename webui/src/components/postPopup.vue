<script>
export default {
    data() {
        return {
            url: null,
            phSel: false,
            caption: "",
            image: null
        }
    },
    methods: {
        exit() {
            this.$parent.$data.popupPost = false
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
			} catch (e) {
                console.log(e)
				this.errormsg = e.toString();
			}
        }
    }
}
</script>

<template>
    <div class="popup">
        <div class="inner">
            <slot />
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
            <hr>
            <div id="preview">
                <img v-if="true" :src="url" />
            </div>
            <div v-if="phSel" id="capt">
                <textarea v-model="caption" placeholder="Write a caption..." autocomplete="off" autocorrect="off"></textarea>
            </div>
            <label class="containerPost" id="file">
                <input type="file" @change="handleChange"/>
                Select File
            </label>
            <button v-if="phSel" class="containerPost" id="send" @click="postPhoto">post</button>
        </div>
    </div>
</template>

<style>
#capt {
    position: relative;
    bottom: 38px;
    height: 66px;
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
  justify-content: center;
  align-items: center;
  position: relative;
  bottom: 50px;
  height: 384px;
  text-align: center;
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
    left: 550px;
    bottom: 10px;
    width: 36px;
    background-color: white;
    border: 1px solid black;
    color: black;
}

#file {

    height: 23px;
    bottom: 10px;
    left: 280px;
}
#descr {
    text-align: center;
    height: 40px;
    position: relative;
    bottom: 48px;
}
hr {
    position: relative;
    bottom: 60px;
}
.close {
    width: 20px;
    height: 20px;
    cursor: pointer;
    position: relative;
    bottom: 30px;
    left: 650px;
}
.close:hover {
    transform: scale(1.2);
}
.xB {
    vertical-align: text-top;
    height: 20px;
    width: 20px;
}
.popup {
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
    border-radius: 10px;
    background: #FFF;
    padding: 32px;
    width: 700px;
    height: 600px;
}
</style>