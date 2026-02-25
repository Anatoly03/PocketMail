<template>
    <router-view />
</template>

<script lang="ts" setup>
import NotFound from "./NotFound.vue";

defineOptions({
    routes: [
        { path: "/:pathMatch(.*)*", component: NotFound },
    ],
});
</script>

<style lang="scss">
body {
    margin: 0;
    padding: 0;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;

    #app {
        width: 100vw;
        height: 100vh;
        overflow-x: hidden;
        display: block;
    }
}
</style>