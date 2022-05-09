import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { Quasar } from 'quasar'
import lang from 'quasar/lang/ja.js'

import './styles/quasar.sass'
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/material-icons-outlined/material-icons-outlined.css'
import '@quasar/extras/material-icons-round/material-icons-round.css'
import '@quasar/extras/material-icons-sharp/material-icons-sharp.css'
import '@quasar/extras/fontawesome-v5/fontawesome-v5.css'

createApp(App)
  .use(Quasar, {
    plugins: {
    },
    config: {
      brand: {
        primary: '#d2ecf7',
        secondary: '#0d95a0',
        accent: '#58456b',

        dark: '#3a3535',

        positive: '#409140',
        negative: '#bf2e16',
        info: '#9dd6c0',
        warning: '#fcdc59'
      }
    },
    lang: lang
  })
  .use(store)
  .use(router)
  .mount('#app')
