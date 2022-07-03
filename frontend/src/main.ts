import { createApp } from "vue";
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { loadFonts } from "./plugins/webfontloader";
import Axios from "axios";
import { firebaseConfig } from "../firebase_config";

loadFonts();

export const FirebaseApp = initializeApp(firebaseConfig);
export const FirebaseAuth = getAuth(FirebaseApp);

Axios.defaults.baseURL = "http://localhost:8890";

createApp(App).use(createPinia()).use(router).use(vuetify).mount("#app");
