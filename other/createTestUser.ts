import PocketBase from "pocketbase";

const pb = new PocketBase("http://127.0.0.1:8090");

(async () => {
    try {
        await pb.collection("users").create({
            email: "test@example.com",
            emailVisibility: true,
            mailName: "test",
            name: "Test",
            password: "password",
            passwordConfirm: "password",
        });
        // console.log("Test user created");

        const authData = await pb.collection("users").authWithPassword("test@example.com", "password");
        // console.log("Auth data:", authData);

        console.log(authData.token);
    } catch (e) {
        if (e?.data?.email?.code === "validation_unique") {
            console.error("Test user already exists");
        } else {
            console.error(e);
            process.exit(1);
        }
    }
})();
