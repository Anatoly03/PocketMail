import PocketBase from "pocketbase";
// import type { AuthRecord } from "pocketbase";
// import { useLocalStorage } from "@vueuse/core";

// /**
//  *  Save ~~multiple logins~~ **currently one login** in local storage.
//  */
// const pbAuthLocalStore = useLocalStorage<{ token?: string; record?: AuthRecord }>("pocketbase_auth", {});

// PocketBase client instance
const pb = new PocketBase(import.meta.env.VITE_API_ROUTE);

function getTokenFromUrl(): string | null {
    const params = new URLSearchParams(window.location.search);
    const token = params.get("token");
    return token && token.length > 0 ? token : null;
}

(async () => {
    try {
        // allow preview mode by passing the token in the query params
        // TODO: think about security implications
        // TODO: import.meta.env.MODE !== "production"
        const token = getTokenFromUrl();
        if (import.meta.env.VITE_ALLOW_AUTO_LOGIN && token) {
            pb.authStore.save(token, undefined);
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
