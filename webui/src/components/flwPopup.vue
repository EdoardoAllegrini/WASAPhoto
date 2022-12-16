<script>
import Profile from './Profile.vue'
export default {
    data() {
        return {
        }
    },
    methods: {
        clickOutside(event) {
            if (event.composedPath()['0'].className == 'popup') {this.exit()}
        },
        exit() {
            this.$parent.$data.flw = 0
            document.body.style.overflow = "scroll"
            return
        },
        handleChange(e) {
            this.image = e.target.files[0];
            this.url = URL.createObjectURL(this.image);
            if (this.url != null) {this.phSel = true}
            else {this.phSel = false}
        },
        wers(w) {
            console.log("Followers",w)
        },
        wing(w) {
            console.log("Following",w)
        },
		async refresh() {
			window.location.reload()
		}
    },
    components: {
        Profile
    }
}
</script>

<template>
    <Profile @followers="wers($event)" @following="wing($event)"></Profile>
    <!-- <div class="popup" @click="clickOutside">
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
            <h2 id="descr">Followers</h2>
            <hr>
            <div v-for="p in rec" class="person">
                <a>{{p}}</a>
            </div>
        </div>
    </div> -->
</template>

<style>
.person {
    position: relative;
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