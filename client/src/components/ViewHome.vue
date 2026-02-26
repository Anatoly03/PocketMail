<template>
    <div class="ui-thoughts">
        <n-tabs key="home-tabs" type="card" animated placement="left">
            <n-tab-pane name="oasis" tab="All"> All Mails </n-tab-pane>
            <n-tab-pane name="the beatles" tab="Junk"> Junk Mails </n-tab-pane>
            <n-tab-pane name="jay chou" tab="Sent"> Sent Mails </n-tab-pane>
        </n-tabs>
    </div>
</template>

<script lang="ts" setup>
import { NTabs, NTabPane, NSpace } from "naive-ui";
</script>

<style lang="scss" scoped>
.ui-thoughts {
    padding: 20px;
}
</style>
