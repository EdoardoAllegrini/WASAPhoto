<script>

export default {
    props: {
        receivedata: Object,
    },
    data() {
        return {
            images: {}
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
            for (let p in this.receivedata.photos) {
                let curr = this.receivedata.photos[p]
                let url = await this.getImage(curr.URL)
                this.images[`${curr.ID}`] = url
            } 
        }
    }
}
</script>

<template>
    <div class="gallery">
        <div v-for="p in receivedata.photos" class="cont">
            <a @click="$router.push(p.URL)" class="img">
                <img :src="images[p.ID]" id="casu">
            </a>
        </div>
    </div>
</template>

<style>
.gallery{
    top: 100px;
    position: relative;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-columns: 33% 33% 33%;
    grid-gap: 20px;
    padding: 10px;
}
.cont{
    position: relative;
}
.cont a {
    cursor: pointer;
}
#casu {
    object-fit: contain;
    height: 216px;
    width: 100%;
}
/* .gallery {
    top: 100px;
    position: relative;
    display: grid;
    grid-template-columns: 33% 33% 33%;
    max-width: 1200px;
    width: 80%;
    margin-left: 10%;
    flex-wrap: wrap;
    justify-content: center;
}
.cont {
    text-align: center;
    margin: 10px;
}
.img {
    cursor: pointer;
    object-fit: contain;
}
.img img {
    height: 264px;
} */
#casu:hover {
    opacity: 0.5;
}
</style>