import PocketBase from "pocketbase";
import { useRoute } from "vue-router";
// import type { AuthRecord } from "pocketbase";
// import { useLocalStorage } from "@vueuse/core";

const route = useRoute();

// /**
//  *  Save ~~multiple logins~~ **currently one login** in local storage.
//  */
// const pbAuthLocalStore = useLocalStorage<{ token?: string; record?: AuthRecord }>("pocketbase_auth", {});

// PocketBase client instance
const pb = new PocketBase(import.meta.env.VITE_API_ROUTE);

(async () => {
    try {
        // allow preview mode by passing the token in the query params
        // TODO: think about security implications
        // TODO: import.meta.env.MODE !== "production"
        if (route.query.token) {
            pb.authStore.save(String(route.query.token), undefined);
        }

        // if (!pbAuthLocalStore.value.token) return;
        // pb.authStore.save(pbAuthLocalStore.value.token, pbAuthLocalStore.value.record);

        if (!pb.authStore.isValid) {
            throw new Error("Invalid auth store");
        }

        await pb.collection("users").authRefresh();
    } catch (err) {
        console.warn("Failed to refresh auth store, clearing it:", err);
        pb.authStore.clear();
    }
})();

export default pb;
// export { pbAuthLocalStore as authStore };
