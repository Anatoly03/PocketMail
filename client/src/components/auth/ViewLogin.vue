<template>
    <div class="view-login">
        <n-form @submit.prevent="submit" label-placement="left" :label-width="80" style="gap: 0px">
            <n-form-item label="Username" path="inputValue" :validation-status="validationStatus">
                <n-input v-model:value="username" placeholder="User Login" autofocus :input-props="{ autocomplete: 'username email webauthn' }" />
            </n-form-item>
            <n-form-item label="Password" path="inputValue" :validation-status="validationStatus" :feedback="inputFeedback">
                <n-input v-model:value="password" type="password" placeholder="****" :input-props="{ autocomplete: 'current-password' }" />
            </n-form-item>
        </n-form>
        <n-button type="success" @click="submit()"> Success </n-button>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import pb from "../../services/api";
import { NButton, NForm, NFormItem, NInput, useLoadingBar } from "naive-ui";

const username = ref("");
const password = ref("");
const validationStatus = ref<"success" | "warning" | "error" | undefined>();
const inputFeedback = ref("");
const loadingBar = useLoadingBar();

async function submit() {
    loadingBar.start();

    try {
        await pb.collection("users").authWithPassword(username.value, password.value);
        validationStatus.value = "success";
        inputFeedback.value = "Login successful";
        loadingBar.finish();
    } catch (err) {
        loadingBar.error();
        validationStatus.value = "error";
        inputFeedback.value = "Invalid username or password";
        return;
    }

    window.location.reload();
}
</script>

<style lang="scss" scoped>
.view-login {
    display: flex;
    flex-direction: column;
    padding: 10px;
    gap: 10px;
}
</style>
