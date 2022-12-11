<script>

export default {
    props: {
        userAuth: Object
    },
    data() {
        var followers = 0, following = 0
        try {
            followers = this.userAuth.followers.length
        }
        catch {}
        try {
            following = this.userAuth.following.length
        }
        catch {}
        return {
            followers,
            following
        }
    },

}
</script>

<template>
    <div id="info">

        <div class="container">

            <div class="profile">

                <div class="profile-user-settings">

                    <h1 class="profile-user-name">{{userAuth.username}}</h1>

                </div>

                <div class="profile-stats">

                    <ul>
                        <li><span class="profile-stat-count">{{userAuth.N_Photos}}</span> posts</li>
                        <li><span class="profile-stat-count">{{followers}}</span> followers</li>
                        <li><span class="profile-stat-count">{{following}}</span> following</li>
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
    margin-right: 4rem;
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
        grid-template-columns: 1fr 2fr;
        grid-template-rows: repeat(3, auto);
        grid-column-gap: 3rem;
        align-items: center;
        position: relative;
        top: 15px;
    }
}

</style>