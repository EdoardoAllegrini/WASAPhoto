<script>
import { toRaw } from 'vue'

export default {
    props: {
        receivedata: Object,
    },
    data() {
        return {
            images: [],
            likers: {},
            comment: ""
        }
    },
    methods: {
        async getImage(path) {
            try {
                var response = await this.$axios.get(path, { responseType: 'blob' });
                return URL.createObjectURL(response.data)
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
                return
            }
        },
        async handleLike(event, art) {
            if (this.likers == undefined) {return}
            if (this.likers[art.Ph.URL] == null) {this.likers[art.Ph.URL] = []}
            try {
                if (this.likers[art.Ph.URL] == undefined || this.likers[art.Ph.URL] == null || !this.likers[art.Ph.URL].includes(localStorage.username)) {
                    var response = await this.$axios.put(`/users/${art.Ph.User}/media/${art.Ph.ID}/likes/${localStorage.username}/`, {})
                    var _ = response.data
                    this.likers[art.Ph.URL].push(localStorage.username)
                    art.Likes++
                } else {
                    var response = await this.$axios.delete(`/users/${art.Ph.User}/media/${art.Ph.ID}/likes/${localStorage.username}/`)
                    var _ = response.data  
                    this.likers[art.Ph.URL].splice(this.likers[art.Ph.URL].indexOf(localStorage.username));
                    art.Likes--
                }
                return 
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
            console.log(this.likers)
        },
        async getLikers(username, photoId) {
            try {
                var response = await this.$axios.get(`/users/${username}/media/${photoId}/likes/`)
                return response.data
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
        },
        statusImg(pic) {
            if (toRaw(this.likers[pic.Ph.URL]) != undefined) {
                if (this.likers[pic.Ph.URL].includes(localStorage.username)) {
                    return "fill: red;"
                }
            }
            return ""
        },
        async postComment(username, photoId) {
            var b = {
                text: this.comment
            }
            try {
                var response = await this.$axios.get(`/users/${username}/media/${photoId}/comments/`, b)
                var _ = response.data
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
            var tmp = toRaw(this.receivedata)
            for (let a in tmp) {
                let curr = tmp[a]
                let url = await this.getImage(curr.Ph.URL)
                this.images[`${curr.Ph.URL}`] = url
                this.likers[curr.Ph.URL] = []
                this.likers[curr.Ph.URL] = await this.getLikers(curr.Ph.User, curr.Ph.ID)
            } 
        }
    }
}
</script>

<template>
    <div style="display: flex; justify-content: center;">
        <section>
            <article v-for="a in receivedata">
                <div class="article">
                    <div class="head">
                        <p class="poster">{{a.Ph.User}}</p>
                    </div>
                    <div class="ph">
                        <a :href="'/#'+a.Ph.URL" class="sdsv">
                            <img :src="images[a.Ph.URL]" id="wrte">
                        </a>
                    </div>
                    <div class="descap">
                        <div class="sts">
                            <div class="hgy" @click="handleLike($event, a)">
                                <svg class="svgL kiop"><path :style="statusImg(a)" d="M12,21.35L10.55,20.03C5.4,15.36 2,12.27 2,8.5C2,5.41 4.42,3 7.5,3C9.24,3 10.91,3.81 12,5.08C13.09,3.81 14.76,3 16.5,3C19.58,3 22,5.41 22,8.5C22,12.27 18.6,15.36 13.45,20.03L12,21.35Z"></path></svg>                            
                            </div>
                            <div class="hgy" @click="setComment">
                                <svg class="svgL kiop" viewBox="0 0 24 24"><path d="M20.656 17.008a9.993 9.993 0 1 0-3.59 3.615L22 22Z"></path></svg>                            
                            </div>
                        </div>
                        <div v-if="a.Likes>0" class="piac">
                            {{a.Likes}} likes
                        </div>
                        <form class="cmnty">
                            <textarea v-model="comment" name="" id="friw" placeholder="Add a comment..." maxlength="250" autocomplete="off" autocapitalize="off" autocorrect="off"></textarea>
                            <div class="lakd">
                                <button id="hgtu"  @click="postComment(a.Ph.User, a.Ph.ID)">Publish</button>
                            </div>
                        </form>
                    </div>
                </div>
            </article>

        </section>
    </div>
</template>

<style>
#hgtu {
    border: none;
    color: rgb(0, 149, 246);
    background-color: white;
    text-align: center;
    font-weight: bold;
}
.lakd {
    float: right;
    text-align: center;
    position: relative;

}
.cmnty {
    padding-top: 10px;
    position: relative;
    width: 100%;
    height: 98%;
    align-items: center;
    border: 0;
    display: flex;
    flex-direction: row;
    flex-grow: 1;
    flex-shrink: 1;
    margin: 0;
    padding: 0;
    position: relative;
    vertical-align: baseline;
}
#friw {
    overflow: auto;
    outline: none;
    border: none;
    padding: 2px 2px 2px 10px;
    width: 90%;
    height: 100%;
    position: relative;
    resize: none;
}
.piac {
    padding-left: 10px;
}
.liked {
    fill: red;
}
.hgy {
    width: 100%;
    padding-left: 10px;
}
.kiop {
    color: black;
}
.hgy {
    box-sizing: border-box;
    width: 24px;
    padding-right: 30px;
    float: left;
}
.hgy path {
		fill: none;
		stroke: currentcolor;
		stroke-linecap: round;
		stroke-linejoin: round;
		stroke-width: 2;
	}
.sts {
    height: 30px;
}
.descap {
    height: 88px;
}
.sdsv {
    position: relative;
    display: flex;
    justify-content: center;
}
#wrte {
    width: 100%;
    object-fit: contain;
    max-height: 100%;
}
.ph {
    position: relative;
    display: flex;
    width: 100%;
    height: 470px;
    text-align: center;
    justify-content: center;
}
section {
    width: 560px;
    overflow: auto;
    position: relative;
    padding-top: 20px;
    top: 0;
    right: 0;
    left: 0;
    bottom: 0;
    z-index: 99;
    display: block;
    align-items: center;
    justify-content: center;
}
section article {
    display: block;
    width: 100%;
    position: relative;
}
.head {
    margin: 10px 15px 0 15px;
    height: 30px;
}
.poster {
    margin: 0;
    font-size: 20px;
}
.article {
    height: 650px;
    background: white;
    width: 100%;
    position: relative;
    margin: auto;
    margin-top: 60px;
    margin-bottom: 60px;
    border: 1px solid black;
    border-radius: 5px;
}
</style>