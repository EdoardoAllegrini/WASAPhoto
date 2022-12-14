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
    methods: {
        async getProfile() {
            var path = this.$route.path
            if (path.slice(-1) != "/") {
                path+="/"
            }
            try {

                let response = await this.$axios.get(path,
                {
                    headers: { Authorization: `Bearer ${localStorage.identifier}` }
                });
                this.sendata = response.data
                this.found = true
            }
            catch (e) {
                if (e.response.status == 404) {
                    this.found = false
                    return
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
