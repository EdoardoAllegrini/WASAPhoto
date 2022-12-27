<script>
// import {identifier} from './LoginPage.vue'
import NavBar from '../components/NavBar.vue'
import Feed from '../components/feed.vue'

export default {
    data() {
        return {
            sendata: {}
        };
    },
    methods: {
        async getStream() {
            try {
                let response = await this.$axios.get("/stream/");
                this.sendata = response.data
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 401) {
                    this.unauth = true
                    return
                }
            }
        }
    },
    mounted() {
        document.body.style.overflow = "scroll"
        this.getStream();
    },
	components: {
        NavBar,
        Feed
    }
}
</script>

<template>
	<NavBar></NavBar>
    <Feed :receivedata="sendata" @likeAct="getStream"></Feed>
</template>

<style>

</style>
