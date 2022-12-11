<script>
// import {identifier} from './LoginPage.vue'
import NavBar from '../components/NavBar.vue'

export default {
    data() {
        return {};
    },
    methods: {
        async getStream() {
            this.username = this.$route.params.username;
        }
    },
    mounted() {
        this.getStream();
    },
	components: {
		NavBar
	}
}
</script>

<template>
	<NavBar></NavBar>
</template>

<style>

</style>
