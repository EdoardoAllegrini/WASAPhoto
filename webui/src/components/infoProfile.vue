<script>

export default {
    props: {
        receivedata: Object
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
        };
    },
    mounted() {
    },
    methods: {
        async follow() {
            try {
                var path = `/users/${this.username}/following/${this.receivedata.username}/`;
                let response = await this.$axios.put(path, {}, {
                    headers: { Authorization: `Bearer ${localStorage.identifier}` }
                });
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
                var path = `/users/${this.username}/following/${this.receivedata.username}/`;
                let response = await this.$axios.delete(path, {
                    headers: { Authorization: `Bearer ${localStorage.identifier}` }
                });
                this.res = response.data;
            }
            catch (e) {
                console.log(e);
                if (e.response.status == 404) {
                    this.found = false;
                }
            }
            this.$parent.getProfile();
        }
    }
}
</script>

<template>
    <div id="info">

        <div class="container">

            <div class="profile">

                <div class="profile-user-settings">

                    <h1 class="profile-user-name">{{receivedata.username}}</h1>

                </div>
                <div class="flw">
                    <div v-if="username==receivedata.username">
                    </div>
                    <button v-else-if="receivedata.followers && receivedata.followers.includes(username)" @click="unfollow" class="btnFlw" id="following">
                        Following
                    </button>
                    <button v-else @click="follow" class="btnFlw" id="follow">
                        Follow
                    </button>
                </div>
                <div class="profile-stats">

                    <ul>
                        <li><span class="profile-stat-count">{{getFollowers}}</span> followers</li>
                        <li><span class="profile-stat-count">{{getFollowing}}</span> following</li>
                        <li>Post: <span class="profile-stat-count">{{receivedata.N_Photos}}</span></li>
                    </ul>

                </div>
            </div>
        </div>
    </div>
</template>

<style>

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
    width: 90%;
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
        grid-template-columns: 1fr 2fr 3fr;
        grid-template-rows: repeat(3, auto);
        grid-column-gap: 1rem;
        align-items: center;
        position: relative;
        top: 15px;
        text-align: center;
    }
}

</style>