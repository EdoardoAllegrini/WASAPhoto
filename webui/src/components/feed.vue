<script>
import { toRaw } from 'vue'
import FooterPost from './footerPost.vue'

export default {
    props: {
        receivedata: Object,
    },
    data() {
        return {
            images: [],
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
        }
    },
    watch: {
        async receivedata() {
            var tmp = toRaw(this.receivedata)
            for (let a in tmp) {
                let curr = tmp[a]
                let url = await this.getImage(curr.Ph.URL)
                this.images[`${curr.Ph.URL}`] = url
            } 
        }
    },
    components: {
        FooterPost
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
                    <hr class="hrFeed">
                    <div class="ph">
                        <a :href="'/#'+a.Ph.URL" class="sdsv">
                            <img :src="images[a.Ph.URL]" id="wrte">
                        </a>
                    </div>
                <FooterPost :receivedata="{poster: a.Ph.User, photo: a.Ph.ID, capt: a.Ph.Caption}"></FooterPost>
                </div>
            </article>

        </section>
    </div>
</template>

<style>
.hrFeed {
    margin: 5px;
}
#cmrt {
    position: relative;
    max-height: 62px;
    overflow: auto;
}
#propi {
    display: inline;
    font-size: 14px;
    color: rgb(38, 38, 38);
    font-weight: 600;
}
.ds {
    position: relative;
    display: flex;
    text-align: left;
    padding-left: 10px;
}
#hgtu {
    border: none;
    color: #b3dbff;
    /* color: rgb(0, 149, 246); */
    background-color: white;
    text-align: center;
    font-weight: bold;
    pointer-events: none;
}
.lakd {
    float: right;
    text-align: center;
    position: relative;

}
.cmnty {
    width: 100%;
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
    border-radius: 5px;
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
    height: 25%;
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
    overflow: hidden;
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
    height: 700px;
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