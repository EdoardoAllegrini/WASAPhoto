<script>
export default {
    props: {
        receivedata: Object,
    },
    computed: {
        getFollowers() {
            if (this.receivedata.followers == undefined) {
                return 0;
            }
            else
                return this.receivedata.followers.length;
        },
        getFollowing() {
            if (this.receivedata.following == undefined) {
                return 0;
            }
            else
                return this.receivedata.following.length;
        }
    },
    data() {
        return {
            username: localStorage.username,
            flw: 0,
            ban: {
                show: false
            }
        };
    },
    mounted() {
    },
    methods: {
        async follow() {
            try {
                var path = `/users/${this.username}/following/${this.receivedata.User}/`;
                let response = await this.$axios.put(path, {});
                this.res = response.data;
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
            this.$parent.getProfile();
        },
        async unfollow() {
            try {
                var path = `/users/${this.username}/following/${this.receivedata.User}/`;
                let response = await this.$axios.delete(path);
                this.res = response.data;
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
            this.$parent.getProfile();
        },
        popFollowers() {
            this.$router.push(`/users/${this.receivedata.User}/followers`)
        },
        popFollowing() {
            this.$router.push(`/users/${this.receivedata.User}/following`)
        },
        async banPop() {
            this.ban.show = !this.ban.show
            var response = await this.$axios.get(`/users/${localStorage.username}/ban/`)
            if (response.data.includes(this.receivedata.User)) {
                console.log("bannato")
            } else {
                console.log("non bannato")
            }
        }
    }
}
</script>

<template>
    <div id="info">

        <div class="container">

            <div class="profile">

                <div class="profile-user-settings">

                    <h1 class="profile-user-name">{{receivedata.User}}</h1>

                </div>
                <div class="flw">
                    <div v-if="!receivedata.User || username==receivedata.User">
                    </div>
                    <button v-else-if="receivedata.followers && receivedata.followers.includes(username)" @click="unfollow" class="btnFlw" id="following">
                        Following
                    </button>
                    <button v-else @click="follow" class="btnFlw" id="follow">
                        Follow
                    </button>
                </div>
                <div class="profile-stats">

                    <ul style="margin: 0;">
                        <li @click="popFollowers"><span class="profile-stat-count">{{getFollowers}}</span> followers</li>
                        <li @click="popFollowing"><span class="profile-stat-count">{{getFollowing}}</span> following</li>
                        <li style="cursor: text;">Post: <span class="profile-stat-count">{{receivedata.N_Photos}}</span></li>
                    </ul>

                </div>
                <div class="cadaf">
                    <div class="hfd" @click="banPop">
                        <svg color="#262626" fill="#262626" height="32" role="img" viewBox="0 0 24 24" width="32"><circle cx="12" cy="12" r="1.5"></circle><circle cx="6" cy="12" r="1.5"></circle><circle cx="18" cy="12" r="1.5"></circle></svg>
                    </div>
                    <div v-if="ban.show" class="menu">
                        <button v-if="ban.Banned==true" id="sout" type="submit" @click="ban">Ban</button>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<style>
#info {
    display: flex;
}
.hfd svg {
    color: rgb(38, 38, 38);
    fill: rgb(38, 38, 38);
}
.hfd {
    display: flex;
    justify-content: center;
    cursor: pointer;
}
body {
    font-family: "Open Sans", Arial, sans-serif;
    min-height: 100vh;
    background-color: #fafafa;
    color: #262626;
    padding-bottom: 3rem;
}

.container {
    max-width: 93.5rem;
    margin: 0 auto;
    padding: 0 2rem;
}
.flw {
    display: inline-block;
}
.btn {
    display: inline-block;
    font: inherit;
    background: none;
    border: none;
    color: inherit;
    padding: 0;
    cursor: pointer;
}

.btnFlw {
    border-radius: 4px;
    text-align: center;
    background: none;
    box-sizing: border-box;
    cursor: pointer;
    display: block;
    font-size: 14px;
    font-weight: var(--font-weight-system-semibold);
    padding: 5px 9px !important;
    pointer-events: auto;
    text-overflow: ellipsis;
    text-transform: inherit;
    width: auto;
    background-color: transparent;
    color: rgb(var(--ig-secondary-button));
    font-size: 1rem;
    font-weight: 300;
}

#following {
    color: black;
    background-color: transparent;
    border: 1px solid black;
}

#follow {
    background-color: rgb(0, 149, 246);
    border: 1px solid transparent;
    color: rgb(255,255,255);
}

.profile-stats {
    float: left;
}

.profile-user-name {
    display: inline-block;
    font-size: 2.2rem;
    font-weight: 300;
}

.profile-stats li {
    display: inline-block;
    font-size: 1.3rem;
    line-height: 1.5;
    margin-right: 1rem;
    cursor: pointer;
}

.profile-stats li:last-of-type {
    margin-right: 0;
}

.profile-stat-count {
    font-weight: 600;
}

@supports (display: grid) {
    .profile {
        display: grid;
        grid-template-columns: 1fr 1fr 3fr 1fr;
        grid-template-rows: repeat(4, auto);
        grid-column-gap: 1rem;
        align-items: center;
        position: relative;
        top: 15px;
        text-align: center;
        justify-items: end;
    }
}

</style>