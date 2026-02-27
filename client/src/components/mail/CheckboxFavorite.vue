<template>
    <n-icon size="20" @click.prevent="toggle" class="favorite-icon" :class="{ active: isFavorite }">
        <component :is="isFavorite ? Star16Filled : Star16Regular" />
    </n-icon>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { NIcon } from "naive-ui";
import { Star16Filled, Star16Regular } from "@vicons/fluent";

const props = defineProps<{
    isFavorite: boolean;
}>();

const isFavorite = ref<boolean>(props.isFavorite);

const emit = defineEmits<{
    "update:isFavorite": [newValue: boolean];
}>();

function toggle() {
    isFavorite.value = !isFavorite.value;
    emit("update:isFavorite", isFavorite.value);
}
</script>

<style lang="scss" scoped>
.favorite-icon {
    display: flex;
    cursor: pointer;
    color: black;

    &.active {
        color: gold;
    }
}
</style>
