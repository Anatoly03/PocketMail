<template>
    <n-icon
        size="20"
        @click.prevent="toggle"
        class="favorite-icon"
        :class="{ active: isFavorite, wiggle: isWiggling }"
        @animationend="isWiggling = false"
    >
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
const isWiggling = ref<boolean>(false);

const emit = defineEmits<{
    "update:isFavorite": [newValue: boolean];
}>();

function toggle() {
    isFavorite.value = !isFavorite.value;
    emit("update:isFavorite", isFavorite.value);
    isWiggling.value = false;
    void document.body.offsetWidth;
    isWiggling.value = true;
}
</script>

<style lang="scss" scoped>
.favorite-icon {
    display: flex;
    cursor: pointer;
    color: black;
    transition: color 0.2s;

    &.active {
        color: gold;
    }

    &.wiggle {
        animation: wiggle-zoom 0.3s;
    }
}

@keyframes wiggle-zoom {
    0% {
        transform: scale(1);
    }
    20% {
        transform: scale(1.2);
    }
    50% {
        transform: scale(0.9);
    }
    80% {
        transform: scale(1.1);
    }
    100% {
        transform: scale(1);
    }
}
</style>
