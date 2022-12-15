<script>

export default {
    props: {
        receivedata: Object,
        images: Array
    },
    data() {
        console.log(this.images)
        return {
            data: this.receivedata.N_Photos
        }
    }
}
</script>

<template>
    <div class="gallery">
        <div v-for="p in images" class="cont">
            <img :src="p">
        </div>
    </div>
</template>

<style>
.gallery {
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
    padding: 10px;
}
.cont img {
    height: 150px;
    object-fit: cover;
    object-position: bottom;
}
</style>