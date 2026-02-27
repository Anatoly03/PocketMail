<template>
    <div id="app">
        <n-loading-bar-provider :to="loadingBarRef" container-style="position: absolute;">
            <router-view />
        </n-loading-bar-provider>
    </div>
</template>

<script lang="ts" setup>
import pb from "../services/api";
import { ref } from "vue";
import { NLoadingBarProvider } from "naive-ui";
import NotFound from "./NotFound.vue";
import ViewHome from "./mail/ViewHome.vue";

const loadingBarRef = ref<undefined | HTMLElement>(undefined);

defineOptions({
    routes: [
        {
            path: "/",
            component: () => (pb.authStore.isValid ? import("./mail/ViewHome.vue") : import("./auth/ViewAuth.vue")),
            children: ViewHome.routes,
        },
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
