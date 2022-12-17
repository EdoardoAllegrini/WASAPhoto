<script>
import { toRaw } from 'vue'
import Profile from './Profile.vue'
export default {
    data() {
        return {
            descr: null,
            list: []
        }
    },
    methods: {
        clickOutside(event) {
            if (event.composedPath()['0'].className == 'popup') {this.exit()}
        },
        exit() {
            this.$router.push(`/users/${this.$route.params["username"]}/`)
            document.body.style.overflow = "scroll"
            return
        },
        handleChange(e) {
            this.image = e.target.files[0];
            this.url = URL.createObjectURL(this.image);
            if (this.url != null) {this.phSel = true}
            else {this.phSel = false}
        },
        handleF(w) {
            if (this.$route.path.slice(-1) == 's') {
                this.list = toRaw(w.followers)
                this.descr = "Followers"
            } else {
                this.list = toRaw(w.following)
                this.descr = "Following"
            }
        },
		async refresh() {
			window.location.reload()
		},
        async changeFlw(event) {
            var user = toRaw(event.path[1]).querySelector("#name").text
            if (event.target.id == "following") {
                try {
                var path = `/users/${localStorage.username}/following/${user}/`;
                let response = await this.$axios.delete(path);
                this.res = response.data;
                }
                catch (e) {
                    console.log(e);
                    if (e.response.status == 404) {
                        this.found = false;
                    }
                }
                event.path[0].id = "follow"
                event.path[0].innerHTML = "Follow"
            } else if (event.target.id == "follow") {
                try {
                    var path = `/users/${localStorage.username}/following/${user}/`;
                    let response = await this.$axios.put(path, {});
                    this.res = response.data;
                }
                catch (e) {
                    console.log(e);
                    if (e.response.status == 404) {
                        this.found = false;
                    }
                }
                event.path[0].id = "following"
                event.path[0].innerHTML = "Following"
            }
        }
    },
    components: {
        Profile
    },
    mounted() {
        document.body.style.overflow = "hidden"
    }
}
</script>

<template>
    <Profile @flw="handleF($event)"></Profile>
    <div class="popup" @click="clickOutside">
        <div class="inner">
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
                <h2 id="descr">{{descr}}</h2>
            </div>
            <hr>

            <div class="people">
                <div v-for="p in list" class="person">
                    <a id="name" :href="'/#/users/'+p+'/'">{{p}}</a>
                    <button v-if="descr=='Following'" class="acsda" :id="descr.toLowerCase()" @click="changeFlw($event)">{{descr}}</button>
                </div>
            </div>

        </div>
    </div>
</template>

<style>
.acsda {
    float: right;
    top: 15px;
    position: relative;
    border-radius: 4px;
    text-align: center;
    background: none;
    box-sizing: border-box;
    display: block;
    pointer-events: auto;
    text-overflow: ellipsis;
    text-transform: inherit;
    width: auto;
    font-size: 15px;
}
.people {
    position: relative;
    padding: 0 50px 0 50px;
    overflow: auto;
    height: 90%;
}
#name {
    position: relative;
    padding-left: 20px;
    font-weight: bold;
    cursor: pointer;
    top: 15px;
    text-decoration: none;
    color: black;
}
.person {
    position: relative;
    height: 50px;
}
#descr {
    text-align: center;
    height: 40px;
    position: relative;
    bottom: 0px;
    margin: 0;
    width: fit-content;
}
hr {
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
    left: 306px;
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
    height: 10%;
    justify-content: center;
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
    border-radius: 10px;
    background: #FFF;
    padding: 0px;
    width: 500px;
    height: 400px;
}
</style>