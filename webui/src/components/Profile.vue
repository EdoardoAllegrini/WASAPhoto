<script>
import NavBar from './NavBar.vue'
import PageNotFound from './PageNotFound.vue';
import InfoProfile from './infoProfile.vue';

export default {
    data: () => {
        return {
            user: {},
            found: true,
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
                this.data = response.data;
                this.user = this.data
                this.found = true
            }
            catch (e) {
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
        <InfoProfile :userAuth="user"></InfoProfile>
    </div>
    <div v-else>
        <PageNotFound></PageNotFound>
    </div>
</template>

<style>
</style>
