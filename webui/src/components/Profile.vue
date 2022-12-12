<script>
import NavBar from './NavBar.vue'
import PageNotFound from './PageNotFound.vue';
import InfoProfile from './infoProfile.vue';

export default {
    data() {
        return {
            found: true,
            sendata: {}
        };
    },
    watch: {
        $route() {
            this.getProfile()
        }
    },
    methods: {
        async getProfile() {
            try {
                var path = this.$route.path
                if (path.slice(-1) != "/") {
                    path+="/"
                }
                let response = await this.$axios.get(path, 
                // { 
                // 	headers: { Authorization: `Bearer ${token}` }
                // }
                {
                    headers: { Authorization: `Bearer ID_edoardo` }
                });
                this.sendata = response.data
                this.found = true
                // InfoProfile.methods.handleFlw()
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.found = false
                }
            }
            document.getElementById("searchIn").value = null
        }
    },
    mounted() {
        this.getProfile()
    },
    components: { NavBar, PageNotFound, InfoProfile },
}

</script>

<template>
    <NavBar></NavBar>
    <div v-if="found">
        <InfoProfile :receivedata="sendata"></InfoProfile>
    </div>
    <div v-else>
        <PageNotFound></PageNotFound>
    </div>
</template>

<style>
</style>
