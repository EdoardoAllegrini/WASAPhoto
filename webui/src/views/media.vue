<script>
import NavBar from '../components/NavBar.vue'
import Imagecomp from '../components/image.vue'
export default {
    components: {
        NavBar,
        Imagecomp
    },
    async beforeRouteUpdate(to, from) {
        // react to route changes...
        window.location.reload()
    }
}
</script>
<template>
    <NavBar></NavBar>
    <Imagecomp :userPoster="$route.params.username" :photoId="Number($route.params.photo)"></Imagecomp>

</template>