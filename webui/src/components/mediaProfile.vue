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
        },
        async deletePh(phId) {
            try {
                var response = await this.$axios.delete(`/users/${localStorage.username}/media/${phId}/`);
                this.$parent.getProfile()
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
            <a :href="'/#'+p.URL" class="img">
                <img :src="images[p.ID]" id="casu">
            </a>
            <svg @click="deletePh(p.ID)" style="color: red" width="16" height="16" class="bi bi-trash" viewBox="0 0 16 16"> <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z" fill="red"></path> <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z" fill="red"></path> </svg>
        </div>
    </div>
</template>

<style>
.cont svg {
    margin-top: 10px;
    cursor: pointer;
}
.gallery{
    position: relative;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-columns: 33% 33% 33%;
    grid-gap: 20px;
    padding: 10px;
    margin-top: 100px;
    justify-content: center;
}
.cont{
    position: relative;
    text-align: center;
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