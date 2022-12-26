<script>
import PageNotFound from './PageNotFound.vue'
import FootPost from './footerPost.vue'

export default {
    emits: ["postedComment","likeAct"],
    data() {
        return {
            url: null,
            badr: false,
            poster: "",
            comments: [],
            likes: []
        }
    },
    methods: {
        likeAct() {
            this.$emit("likeAct")
        },
        clickOutside(event) {
            if (event.composedPath()['0'].className == 'popupP') {this.exit()}
        },
        exit() {
            this.$router.push("/stream/")
            document.body.style.overflow = "scroll"
            return
        },
        async getImage(path) {
            if (path.slice(-1) != "/") {
                path+="/"
            }
            try {
                let response = await this.$axios.get(path, { responseType: 'blob' });
                this.url = URL.createObjectURL(response.data) 
                return this.url
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
        },
        async getComments() {
            try {
                var response = await this.$axios.get(`/users/${this.poster}/media/${this.$route.params.photo}/comments/`)
                if (response.data) {this.comments = response.data}
                else {this.comments = []}
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
            this.$emit("postedComment")
        },
        checkPosterAuth() {
            return localStorage.username == this.poster
        },
        async deleteCmnt(cmntId) {
            try {
                var response = await this.$axios.delete(`/users/${this.poster}/media/${this.$route.params.photo}/comments/${cmntId}/`)
                var _ = response.data
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
            let cs = await this.getComments()
            if (cs) {this.comments=cs}
        }
    },
    async mounted() {
        this.poster = this.$route.params.username
        this.getImage(this.$route.path)
        await this.getComments()
        document.body.style.overflow = "hidden"
        var clss = document.body.getElementsByClassName("swcmt")
        clss[clss.length-1].style.display = "none"
        clss = document.body.getElementsByClassName("hgy")
        clss[clss.length-1].style.display = "none"

    },
    components: {
        PageNotFound,
        FootPost
    }
}
</script>

<template>
    <div class="popupP" @click="clickOutside">
        <div class="innerP">
            <slot />
            <div id="fle">
                <img id="med" :src="url">
            </div>
            <div class="bdsdf">
                <div class="pkjbk">
                    {{ poster }}
                </div>
                <div class="cmgts">
                    <div v-for="c in comments" class="jbcr">
                        <div class="ds">
                            <div id="propi">
                                {{c.user}}
                            </div>
                            <span id="dvfi" style="display: inline!important;margin: 0!important;">&nbsp;</span>
                            <div id="cmrt">
                                {{ c.text }}
                            </div>
                            <div class="delc" v-if="checkPosterAuth()">
                                <svg @click="deleteCmnt(c.id)" style="color: red" width="16" height="16" class="bi bi-trash" viewBox="0 0 16 16"> <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z" fill="red"></path> <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z" fill="red"></path> </svg>
                            </div>
                        </div>
                    </div>
                </div>
                <FootPost :receivedata="{poster: poster, photo: $route.params.photo}" @postedComment="getComments" @likeAction="likeAct"></FootPost>
            </div>
            <div v-if="badr">
                <PageNotFound></PageNotFound>
            </div>
        </div>
    </div>

</template>

<style>
.delc svg {
    cursor: pointer;
}
.delc {
    position: absolute;
    right: 10px;
}
.jbcr {
    position: relative;
    height: 50px;
}
#cmrt {
    overflow: hidden;
}
.cmgts {
    position: relative;
    overflow: auto;
    height: 88%;
    padding-top: 10px;
}
.piac {
    text-align: left;
}
.bdsdf {
    position: relative;
    width: 346px;
    display: inline-block;
    margin: 2.5% 0 8.5% 0;
    height: 80%;
}
.pkjbk {
    position: relative;
    width: 100%;
    font-size: 20px;
    font-weight: 600;
}
.popup {
    overflow: hidden;
}
#fle {
    margin: 8.5% 0 8.5% 0;
    justify-content: center;
    align-items: center;
    position: relative;
    text-align: center;
    display: flex;
    height: 80%;
    width: 350px;
    float: left;
}
#fle img {
    max-width: 100%;
    max-height: 100%;
}
#med {
    position: relative;
    max-height: 800px;
}
</style>