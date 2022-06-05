// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Vuetify
import { createVuetify  } from 'vuetify'

export default createVuetify({
  theme: {
    defaultTheme: "mainLightTheme",
    themes: {
      mainLightTheme: {
        dark: false,
        colors: {
          primary: '#d2ecf7',
          secondary: '#0d95a0',
          accent: '#58456b',
          // dark: '#3a3535',
          positive: '#409140',
          negative: '#bf2e16',
          info: '#9dd6c0',
          warning: '#fcdc59'
        },
      },
      mainDarkTheme: {
        dark: true,
        colors: {},
      }
    }
  },
})
