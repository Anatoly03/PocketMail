<template>
    <div class="view-auth">
        <n-card>
            <n-loading-bar-provider :to="loadingBarRef" container-style="position: absolute;">
                <div ref="loadingBarRef" class="loading-bar"></div>
                <n-tabs type="line" size="large">
                    <n-tab-pane name="Login">
                        <view-login />
                    </n-tab-pane>
                    <n-tab-pane name="Sign up" disabled> Sign up is currently unavailable. Please contact support if you want to use PocketMail. </n-tab-pane>
                </n-tabs>
            </n-loading-bar-provider>
        </n-card>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { NCard, NTabs, NTabPane, NLoadingBarProvider } from "naive-ui";
import ViewLogin from "./ViewLogin.vue";

const loadingBarRef = ref<undefined | HTMLElement>(undefined);
</script>

<style lang="scss" scoped>
.view-auth {
    width: 480px;
    margin: auto;
}

.loading-bar {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    border-radius: var(--n-border-radius);
    overflow: hidden;
    pointer-events: none;
}
</style>
