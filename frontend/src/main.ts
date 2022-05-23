import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import Axios from "axios"

loadFonts()

Axios.defaults.baseURL = "http://localhost:8890"

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .mount('#app')
