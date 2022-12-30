<script>
import NavBar from './NavBar.vue'
import PageNotFound from './PageNotFound.vue';
import InfoProfile from './infoProfile.vue';
import MediaProfile from './mediaProfile.vue';

export default {
    // emits: ["flw"],
    data() {
        return {
            badr: false,
            unauth: false,
            sendata: {},
        };
    },
    methods: {
        async getProfile() {
            var path = "/users/"+this.$route.params.username+"/"
            try {

                let response = await this.$axios.get(path);
                this.badr = false
                this.sendata = response.data
                // this.$emit('flw', {followers: this.sendata.followers, following: this.sendata.following})
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true
                    return
                } else if (e.response.status == 401) {
                    this.unauth = true
                    return
                }
            }
            document.getElementById("searchIn").value = null
            this.found = true
        }
    },
    // mounted() {
    //     this.getProfile()
    // },
    watch: {
        '$route.params': {
            handler(newValue) {
                document.body.style.overflow = "scroll"
                if (this.$route.path.match(/\/users\/[a-zA-Z]*\/?$/)) {
                    this.getProfile()
                }
            },
            immediate: true,
        }
    },
    components: { NavBar, PageNotFound, InfoProfile, MediaProfile },
}

</script>

<template>
    <NavBar @succPost="getProfile"></NavBar>
    <div id="allPr" v-if="!badr && !unauth">
        <InfoProfile :receivedata="sendata"></InfoProfile>
        <MediaProfile :receivedata="sendata"></MediaProfile>
    </div>
    <div v-else-if="badr">
        <PageNotFound></PageNotFound>
    </div>
    <div v-else-if="unauth">
        <h1 style="text-align: center;position: relative;">401 Unauthorized</h1>
    </div>
    <RouterView :recv="{followers: this.sendata.followers, following: this.sendata.following}" @refr="getProfile"/>
</template>

<style>
body {
    overflow: scroll;
}
#allPr {
    padding: 30px 20px 0;
}
</style>
