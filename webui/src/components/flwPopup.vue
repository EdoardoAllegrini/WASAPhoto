<script>
import { toRaw } from 'vue'
export default {
    emits: ["refr"],
    props: {
        recv: Object
    },
    data() {
        return {
            descr: null,
            list: [],
            following: []
        }
    },
    methods: {
        clickOutside(event) {
            if (event.composedPath()['0'].className == 'popupF') {this.exit()}
        },
        exit() {
            this.$router.push(`/users/${this.$route.params["username"]}/`)
            document.body.style.overflow = "scroll"
            return
        },
		async refresh() {
			window.location.reload()
		},
        async follow(person) {
            try {
                var path = `/users/${localStorage.username}/following/${person}/`;
                let response = await this.$axios.put(path, {});
                this.res = response.data;
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
            this.following = await this.getFollowing(localStorage.username)
            this.$emit("refr")
        },
        async unfollow(person) {
            try {
                var path = `/users/${localStorage.username}/following/${person}/`;
                let response = await this.$axios.delete(path);
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
            this.following = await this.getFollowing(localStorage.username)
            this.$emit("refr")
        },
        async getFollowing(us) {
            try {
                var response = await this.$axios.get(`/users/${us}/following/`)
                if (response.data) {return response.data}
                else {return []}
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
        },
        async getFollowers(us) {
            try {
                var response = await this.$axios.get(`/users/${us}/followers/`)
                if (response.data) {return response.data}
                else {return []}
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
        },
        statePerson(per) {
            if (localStorage.username==per) {return ""}
            else if (this.following.includes(per)) {return "Following"}
            return "Followers"
        }
    },
    async mounted() {
        if (this.$route.path.slice(-1) == 's') {
            this.list = this.recv.followers
            this.descr = "Followers"
        } else {
            this.list = this.recv.following
            this.descr = "Following"
        }
        document.body.style.overflow = "hidden"
        this.following = await this.getFollowing(localStorage.username)
    }
}
</script>

<template>
    <div class="popupF" @click="clickOutside">
        <div class="innerF">
            <slot />
            <div class="tit">
                <div class="closeF" @click="exit">
                    <svg class="xBF" viewPort="0 0 12 12">
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
                <h2 id="descrF">{{descr}}</h2>
            </div>
            <hr>

            <div class="people">
                <div v-for="p in list" class="person" :key="p">
                    <a id="name" :href="'/#/users/'+p">{{p}}</a>
                    <!-- <button v-if="descr=='Following'" class="acsda" :id="descr.toLowerCase()" @click="changeFlw($event)">{{descr}}</button> -->
                    <button v-if="statePerson(p)=='Following'" class="acsda" id="following" @click="unfollow(p)">Following</button>
                    <button v-else-if="statePerson(p)=='Followers'" class="acsda" id="follow" @click="follow(p)">Follow</button>
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
#descrF {
    text-align: center;
    height: 40px;
    position: relative;
    bottom: 0px;
    margin: 0;
    width: fit-content;
}
.innerF hr {
    position: relative;
    bottom: 0px;
    margin: 0;
}
.closeF {
    width: 20px;
    height: 20px;
    cursor: pointer;
    position: relative;
    float: right;
    bottom: 0px;
    left: 306px;
}
.closeF:hover {
    transform: scale(1.2);
}
.xBF {
    vertical-align: text-top;
    height: 20px;
    width: 20px;
}
.tit {
    display: flex;
    height: 10%;
    justify-content: center;
}
.popupF {
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
.innerF {
    border-radius: 10px;
    background: #FFF;
    padding: 0px;
    width: 500px;
    height: 400px;
}
</style>