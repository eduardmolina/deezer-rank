<template>
    <main>
        <div class='page-title'>
            Deezer Rank
        </div>
        <div class='tracks-container'>
            <div v-for='(track, index) in tracks' :key='track.id'>
                <Track :title='(index + 1) + "# " + track.title' :imgSrc='track.album.cover' :previewSrc='track.preview' />
            </div>
        </div>
    </main>
</template>

<script>
import Track from './components/Track.vue'
import api from './services/api.js';

export default {
    name: 'App',
    data: () => ({
        tracks: []
    }),
    components: {
        Track
    },
    mounted: function () {
        api.get('/chart/0/tracks?limit=50').then(response => {
            this.tracks = response.data.data;
        });
    }
}
</script>

<style>
    body {
        font-family: Cairo;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        background: #c31432;
        background: -webkit-linear-gradient(to right, #240b36, #c31432);
        background: linear-gradient(to right, #240b36, #c31432);
        margin: 0;
        user-select: none;
        overflow: hidden;
    }

    main {
        width: 100%;
        height: 100%;
        display: flex;
    }

    @keyframes fadein {
        from { opacity: 0; }
        to   { opacity: 1; }
    }

    .tracks-container {
        width: 80%;
        height: calc(100vh - 80px);
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        padding: 80px;
        animation: 4s fadein;
        overflow: scroll;
    } 

    .page-title {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 60px;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 36px;
        font-weight: bold;
        box-shadow: 5px 5px 5px #111111;
        animation: 4s fadein;
        color: rgba(255, 255, 255, 0.8);
        z-index: 99;
        background-color: #c31432;
    }

    @media (min-width: 700px) {
        .page-title {
            position: relative;
            width: 20%;
            min-width: 264px;
            height: 100vh;
            padding-top: 56px;
            background-color: rgba(0, 0, 0, 0);
            align-items: flex-start;
        }

        .tracks-container {
            padding: 50px;
        }
    }
</style>
