<template>
    <input @input='$emit("onsearch", $event.target.value)' type='text' placeholder='Search for tracks' />
</template>

<script>
export default {
    name: 'SearchBar',
}
</script>

<style scoped>
    input[type='text'] {
        border-radius: 25px;
        border: none;
        height: 30px;
        outline: none;
        padding-left: 16px;
        font-size: 14px;
        color: rgba(60, 60, 60);
        animation: 0.5s fadein;
    }
</style>

