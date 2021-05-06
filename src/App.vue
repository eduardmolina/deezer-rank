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
        height: calc(100vh - 50px);
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        padding: 50px;
        animation: 4s fadein;
        overflow: scroll;
    } 

    .page-title {
        width: 20%;
        min-width: 264px;
        display: flex;
        justify-content: center;
        font-size: 36px;
        font-weight: bold;
        box-shadow: 5px 5px 5px #111111;
        animation: 4s fadein;
        padding-top: 56px;
        color: rgba(255, 255, 255, 0.7);
    }
</style>
