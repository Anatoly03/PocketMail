<template>
    <div class="view-home">
        <n-split v-model:size="splitLeft">
            <template #1>
                <!-- 
        <n-tabs key="home-tabs" type="card" animated placement="left">
            <n-tab-pane name="oasis" tab="All"> All Mails </n-tab-pane>
            <n-tab-pane name="the beatles" tab="Junk"> Junk Mails </n-tab-pane>
            <n-tab-pane name="jay chou" tab="Sent"> Sent Mails </n-tab-pane>
        </n-tabs> -->
            </template>
            <template #2>
                <n-split v-model:size="splitRight">
                    <template #1>
                        <view-mails />
                    </template>
                    <template #2>
                        Right pane
                    </template>
                </n-split>
            </template>
        </n-split>
    </div>
</template>

<script lang="ts" setup>
import pb from "../services/api";
import { onMounted, reactive, ref } from "vue";
import { NSplit, NDataTable } from "naive-ui";
import type { DataTableColumns, DataTableRowKey, PaginationInfo } from "naive-ui";
import ViewMails from "./ViewMails.vue";

const splitLeft = ref<number>(0.2);
const splitRight = ref<number>(0.95);
</script>

<style lang="scss" scoped>
.view-home {
    width: 100%;
}
</style>
