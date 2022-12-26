<script>

export default {
    emits: ["postedComment", "likeAction"],
    props: {
        receivedata: Object,
    },
    data() {
        return {
            likers: [],
            comments: [],
            stat: ""
        }
    },
    methods: {
        showComments() {
            this.$router.push('/users/'+this.receivedata.poster+'/media/'+this.receivedata.photo+'/')
        },
        async handleLike() {
            try {
                if (this.likers == undefined || this.likers == null || !this.likers.includes(localStorage.username)) {
                    var response = await this.$axios.put(`/users/${this.receivedata.poster}/media/${this.receivedata.photo}/likes/${localStorage.username}/`, {})
                    var _ = response.data
                } else {
                    var response = await this.$axios.delete(`/users/${this.receivedata.poster}/media/${this.receivedata.photo}/likes/${localStorage.username}/`)
                    var _ = response.data  
                }
                let l = await this.getLikers() 
                if (l) {this.likers = l}
                else {this.likers = []}
                this.statusImgLik()
                this.$emit("likeAction")
                return 
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
        },
        statusImgLik() {
            if (this.likers != undefined) {
                if (this.likers.includes(localStorage.username)) {
                    this.stat =  "fill: red;"
                    return
                }
            }
            this.stat = ""
            return
        },
        async postComment(event) {
            var cp = event.composedPath()
            var txar = event.composedPath()[2][0]
            var b = {
                text: txar.value
            }
            try {
                var response = await this.$axios.post(`/users/${this.receivedata.poster}/media/${this.receivedata.photo}/comments/`, b)
                var _ = response.data
                txar.value = ""
                cp[0].style.color = "rgb(179, 219, 255)"
                cp[0].style.pointerEvents = "none"
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
            this.$emit("postedComment")
        },
        handleInputCmnt(ev) {
            if (ev.target.value.length < 1 && ev.composedPath()[1][1].style.color=="rgb(0, 149, 246)") {
                ev.composedPath()[1][1].style.color = "rgb(179, 219, 255)"
                ev.composedPath()[1][1].style.pointerEvents = "none"
            } else if (ev.target.value.length >= 1 && ev.composedPath()[1][1].style.color=="rgb(179, 219, 255)") {
                ev.composedPath()[1][1].style.color = "rgb(0, 149, 246)"
                ev.composedPath()[1][1].style.pointerEvents = "visible"
            }
        },
        async getLikers() {
            try {
                var response = await this.$axios.get(`/users/${this.receivedata.poster}/media/${this.receivedata.photo}/likes/`)
                return response.data
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
        }
    },
    watch: {
        async receivedata() {
            let l = await this.getLikers()
            if (l) {this.likers = l}
            else {this.likers = []}
            this.statusImgLik()
        }
    }
}
</script>

<template>
    <div class="descap">
        <div class="sts">
            <div class="hgy" @click="handleLike()">
                <svg class="svgL kiop"><path :style="stat" ref="lkbtn" d="M12,21.35L10.55,20.03C5.4,15.36 2,12.27 2,8.5C2,5.41 4.42,3 7.5,3C9.24,3 10.91,3.81 12,5.08C13.09,3.81 14.76,3 16.5,3C19.58,3 22,5.41 22,8.5C22,12.27 18.6,15.36 13.45,20.03L12,21.35Z"></path></svg>                            
            </div>
            <div class="hgy" @click="showComments">
                <svg class="svgL kiop" viewBox="0 0 24 24"><path d="M20.656 17.008a9.993 9.993 0 1 0-3.59 3.615L22 22Z"></path></svg>                            
            </div>
        </div>
        <div v-if="likers && likers.length>0" class="piac">
            {{likers.length}} likes
        </div>
        <div class="ds">
            <div id="propi">
                {{receivedata.poster}}
            </div>
            <span id="dvfi" style="display: inline!important;margin: 0!important;">&nbsp;</span>
            <div id="cmrt">
                {{receivedata.capt}}
            </div>
        </div>
        <hr class="hrFeed">
        <form class="cmnty">
            <textarea @input="handleInputCmnt($event)" id="friw" placeholder="Add a comment..." maxlength="250" minlength="1" autocomplete="off" autocapitalize="off" autocorrect="off"></textarea>
            <div class="lakd">
                <button style="border: none;color: #b3dbff;background-color: white;text-align: center;font-weight: bold;" id="hgtu" @click="postComment($event)">Publish</button>
            </div>
        </form>
        <a class="swcmt" :href="'/#/users/'+receivedata.poster+'/media/'+receivedata.photo+'/'">Show comments</a>
    </div>
</template>

<style>
.swcmt {
    color: rgb(168,168,168);
    text-decoration: none;
    padding-left: 10px;
    pointer-events: auto;
    cursor: pointer;
}
.swcmt:hover {
    color: initial;
}
</style>