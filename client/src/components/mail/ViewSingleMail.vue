<template>
    <div class="mail-view">
        <div class="mail-subject">{{ mailSubject }}</div>
        <div class="mail-body" v-html="mailBody"></div>
    </div>
</template>

<script lang="ts" setup>
import pb from '../../services/api';
import { onBeforeMount, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useLoadingBar } from 'naive-ui';

const route = useRoute();
const loadingBar = useLoadingBar();
const mailBody = ref<string>("");
const mailSubject = ref<string>("");

onBeforeMount(async () => {
    const result = await pb.collection('mails').getOne(route.params.id as string);

    mailBody.value = result.body;
    mailSubject.value = result.subject;
})

onMounted(() => {
    loadingBar.finish();
});
</script>

<style lang="scss" scoped>
.mail-view {
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 16px;
    height: 100%;
    box-sizing: border-box;

    .mail-subject {
        font-size: 1.5em;
        font-weight: bold;
    }

    .mail-body {
        flex-grow: 1;
        overflow-y: auto;
        background-color: #fff;
        padding: 16px;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
}
</style>
