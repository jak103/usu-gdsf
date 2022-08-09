// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Vuetify
import { createVuetify } from 'vuetify'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default createVuetify({
  // https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
  theme: {
    themes: {
      light: {
        colors: {
          primary: '#0F2439',
          secondary: '#165A7D',
          info: '#226BAA',
          warning: '#AE6002',
          white: '#FFFFFF',
          gray: '#A2AAAD',
          black: '#000000',
          sky_blue: '#288DC2',
          teal: '#00938F',
          watermelon: '#F16278',
          tangerine: '#F58220',
          sunflower: '#F6BD17',
          grape: '#543035',
        }
      },
      dark: {

      }
    },
  },
  icons: {

  },
  components,
  directives,
})
