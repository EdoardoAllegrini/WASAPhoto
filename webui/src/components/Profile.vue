<script>
import NavBar from './NavBar.vue'
import PageNotFound from './PageNotFound.vue';
import InfoProfile from './infoProfile.vue';
import MediaProfile from './mediaProfile.vue';

export default {
    data() {
        return {
            badr: false,
            unauth: false,
            sendata: {},
            listIm: []
        };
    },
    methods: {
        async getProfile() {
            var path = this.$route.path
            if (path.slice(-1) != "/") {
                path+="/"
            }
            try {

                let response = await this.$axios.get(path);
                this.sendata = response.data
            }
            catch (e) {
                if (e.response.status == 404) {
                    this.badr = true
                    return
                } else if (e.response.status == 401) {
                    this.unauth = true
                    return
                }
            }
            document.getElementById("searchIn").value = null
            for (let i=0;i<this.sendata.photos.length;i++) {
                this.getImage(this.sendata.photos[i].URL)
            }
            this.found = true
        },
        async getImage(path) {
            try {
                let response = await this.$axios.get(path, { responseType: 'blob' });
                var url = URL.createObjectURL(response.data) 
                this.listIm.push(url)
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
            console.log(path)
        }
    },
    mounted() {
        this.getProfile()
    },
    components: { NavBar, PageNotFound, InfoProfile, MediaProfile },
}

</script>

<template>
    <NavBar></NavBar>
    <div v-if="!badr && !unauth">
        <InfoProfile :receivedata="sendata"></InfoProfile>
        <MediaProfile :receivedata="sendata" :images="listIm"></MediaProfile>
    </div>
    <div v-else-if="badr">
        <PageNotFound></PageNotFound>
    </div>
    <div v-else-if="unauth">
        <h1 style="text-align: center;position: relative;">401 Unauthorized</h1>
    </div>
</template>

<style>
</style>
