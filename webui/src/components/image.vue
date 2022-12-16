<script>
import PageNotFound from './PageNotFound.vue'
export default {
    data() {
        return {
            url: null,
            badr: false
        }
    },
    methods: {
        async getImage(path) {
            if (path.slice(-1) != "/") {
                path+="/"
            }
            try {
                let response = await this.$axios.get(path, { responseType: 'blob' });
                this.url = URL.createObjectURL(response.data) 
                return this.url
            }
            catch (e) {
                console.log(e)
                if (e.response.status == 404) {
                    this.badr = true;
                }
            }
        }
    },
    mounted() {
        this.getImage(this.$route.path)
    },
    components: {
        PageNotFound
    }
}
</script>

<template>
    <div id="fle">
        <img id="med" :src="url">
    </div>
    <div v-if="badr">
        <PageNotFound></PageNotFound>
    </div>
</template>

<style>
#fle {
    position: relative;
}
#med {
    position: relative;
    max-height: 800px;
}
</style>