<script>

export default {
    props: {
        receivedata: Object
    },
    data() {
        return {
            n_followers: 0,
            n_following: 0,
            stateFlw: "",
            userAuth:   {
                username: "edoardo",
                identifier: "ID_edoardo"
            }        
        }
    },
    methods: {
        handleFlw() {
            // TO ASK non ha accesso a receivedata???
            console.log(this.receivedata)
            try {
                this.n_followers = this.receivedata.followers.length
            }
            catch {}
            try {
                this.n_following = this.receivedata.following.length
            }
            catch {}
            this.userAuth = {
                username: "edoardo",
                identifier: "ID_edoardo"
            }
            return
        } 
    },
    mounted() {
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
                    <div v-if="userAuth.username==receivedata.username">
                        Edit
                    </div>
                    <div v-else-if="receivedata.followers && receivedata.followers.includes(userAuth.username)">
                        Following
                    </div>
                    <div v-else>
                        Follow
                    </div>
                </div>
                <div class="profile-stats">

                    <ul>
                        <li><span class="profile-stat-count">{{receivedata.N_Photos}}</span> posts</li>
                        <li v-if="receivedata.followers==undefined"><span class="profile-stat-count">0</span> followers</li>
                        <li v-else><span class="profile-stat-count">{{receivedata.followers.length}}</span> followers</li>
                        <li v-if="receivedata.following==undefined"><span class="profile-stat-count">0</span> following</li>
                        <li v-else><span class="profile-stat-count">{{receivedata.following.length}}</span> following</li>
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
    }
}

</style>