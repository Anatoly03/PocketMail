<template>
    <div id="app">
        <router-view />
    </div>
</template>

<script lang="ts" setup>
import pb from "../services/api";
import NotFound from "./NotFound.vue";
import ViewAuth from "./auth/ViewAuth.vue";
import ViewHome from "./ViewHome.vue";

defineOptions({
    routes: [
        { path: "/", component: () => (pb.authStore.isValid ? ViewHome : ViewAuth) },
        { path: "/", component: ViewHome },
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
        display: flex;
        justify-content: center;
        width: 100vw;
        height: 100vh;
        overflow-x: hidden;
    }
}
</style>