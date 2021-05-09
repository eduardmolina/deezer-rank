<template>
    <kinesis-container class='kinesis-container'>
        <kinesis-element class='track-container radius shadow' :strength='-20' type='depth'>
            <kinesis-element
                class='track-img radius'
                :strength='10'
                type='depth'
            >
                <img class='track-img radius shadow' :src='imgSrc' />
                <kinesis-element
                    class='play-btn shadow'
                    :strength='25'
                    type='depth'
                    @click='play'
                >
                    <svg v-if='!playing' xmlns='http://www.w3.org/2000/svg' viewBox='0 0 26 26'>
                        <polygon
                            class='play-btn__svg'
                            points='9.33 6.69 9.33 19.39 19.3 13.04 9.33 6.69'
                        />
                    </svg> 
                    <svg v-else
                        xmlns='http://www.w3.org/2000/svg'
                        viewBox='0 0 1200 1200'
                    >
                        <path
                            d='M 600,0 C 268.62914,0 0,268.62914 0,600 c 0,331.37086 268.62914,600 600,600 331.37086,0 600,-268.62914 600,-600 C 1200,268.62914 931.37086,0 600,0 z m -269.16515,289.38 181.71397,0 0,621.24 -181.71397,0 0,-621.24 z m 356.61633,0 181.71399,0 0,621.24 -181.71399,0 0,-621.24 z'
                        />
                    </svg>
                </kinesis-element>
            </kinesis-element>
            <div class='title'> {{ title }} </div>
        </kinesis-element>
    </kinesis-container>
</template>

<script>
export default {
    name: 'Track',
    props: {
        imgSrc: String,
        previewSrc: String,
        title: String
    },
    data: () => ({
        playing: false,
        player: new Audio()
    }),
    methods: {
        play: function () {
            if (!this.playing) {
                this.player.play();
            } else {
                this.player.pause();
            }

            this.playing = !this.playing;
        }
    },
    mounted: function () {
        this.player.src = this.previewSrc;
        this.player.onended = () => this.playing = false;
    }
}
</script>

<style scoped>
    .kineses-container {
        width: 150px;
        height: 150px;
    }

    .radius {
        border-radius: 10%;
    }

    .shadow {
        box-shadow: 2px 5px 5px #111111;
    }

    .title {
        font-size: 16px;
        font-weight: bold;
        max-width: 120px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .track-container {
        width: 150px;
        height: 180px;
        background-color: rgba(255, 255, 255, 0.1);
        display: flex;
        align-items: center;
        justify-content: space-around;
        flex-direction: column;
        cursor: pointer;
        margin: 20px;
    }

    .track-img {
        width: 120px;
        height: 120px;
        position: relative;
        transform: translateY(3px);
    }

    .play-btn {
        width: 50px;
        height: 50px;
        background-color: rgba(255, 255, 255, 0.7);
        border-radius: 50%;
        position: absolute;
        top: calc(50% - 25px);
        left: calc(50% - 25px);
    }
</style>
