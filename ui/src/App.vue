<template>
    <main>
        <div class='appbar'>
            <div class='appbar-header'>
                <span
                    v-if='!mobileSearchBarOpened || screenWidth >= 700'
                    style='animation: 0.5s fadein'
                >
                    Deezer Rank
                </span>
                <span
                    @click='toggleSearchBar'
                    class='material-icons search-icon'
                >
                    search
                </span>
            </div>
            <SearchBar
                v-if='mobileSearchBarOpened || screenWidth >= 700'
                @onsearch='onSearch'
            />
            <span class='slider' />
        </div>
        <div v-if='!searchingTracks.length && mounted' class='not-found'>
            Nada encontrado
        </div>
        <div v-else class='tracks-container'>
            <div v-for='(track, index) in searchingTracks' :key='track.id'>
                <Track :title='(index + 1) + "# " + track.title' :imgSrc='track.album.cover' :previewSrc='track.preview' />
            </div>
        </div>
    </main>
</template>

<script>
import SearchBar from './components/SearchBar.vue';
import Track from './components/Track.vue';

import api from './services/api.js';

export default {
    name: 'App',
    data: () => ({
        tracks: [],
        searchingTracks: [],
        mounted: false,
        mobileSearchBarOpened: false,
        screenWidth: window.innerWidth
    }),
    components: {
        SearchBar,
        Track
    },
    methods: {
        onSearch: function (searchTerm) {
            this.searchingTracks = this.tracks.filter(
                (track) => track.title.toLowerCase()
                           .indexOf(searchTerm.toLowerCase()) !== -1);
        },
        toggleSearchBar: function () {
            this.mobileSearchBarOpened = !this.mobileSearchBarOpened;
        },
        onResize: function () {
            this.screenWidth = window.innerWidth;
        }
    },
    mounted: function () {
        api.get('/api/tracks').then(response => {
            this.tracks = response.data.data;
            this.searchingTracks = this.tracks;
            this.mounted = true;
        });
        window.addEventListener('resize', this.onResize);
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

    @keyframes horizontal-slide {
        from { width: 0; }
        to { width: 100%; }
    }

    @keyframes vertical-slide {
        from { height: 0; }
        to { height: 100vh; }
    }

    .tracks-container {
        width: 100%;
        height: calc(100vh - 80px);
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        padding: 80px;
        animation: 4s fadein;
        overflow: scroll;
    } 

    .appbar {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 60px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        font-size: 36px;
        font-weight: bold;
        box-shadow: 5px 0px 5px #111111;
        animation: 4s fadein;
        color: rgba(255, 255, 255, 0.8);
        z-index: 99;
        background-color: #c31432;
    }

    .not-found {
        width: 100%;
        height: 100vh;
        font: normal normal 600 24px Cairo;
        color: rgba(255, 255, 255, 0.8);
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .search-icon {
        font-size: 30px;
        position: absolute;
        right: 0;
        bottom: calc(50% - 16px);
        margin-right: 16px;
    }

    .slider {
        position: fixed;
        top: 58px;
        width: 100%;
        height: 2px;
        background-color: rgba(255, 255, 255, 0.6);
        animation: 4s horizontal-slide;
    }

    @media (min-width: 700px) {
        .appbar {
            position: relative;
            justify-content: flex-start;
            width: 264px;
            min-width: 264px;
            height: 100vh;
            padding-top: 56px;
            background-color: rgba(0, 0, 0, 0);
        }

        .tracks-container {
            padding: 50px;
        }

        .search-icon {
            display: none;
        }

        .not-found {
            width: 84%;
        }

        .slider {
            width: 2px;
            height: 100vh;
            left: 262px;
            top: 0;
            animation: 4s vertical-slide;
            background-color: rgba(255, 255, 255, 0.2);
        }
    }
</style>
