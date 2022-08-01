// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import { aliases, mdi } from 'vuetify/iconsets/mdi'

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
          black: '#1E252B',
          sky_blue: '#288DC2',
          teal: '#00938F',
          watermelon: '#F16278',
          tangerine: '#F58220',
          sunflower: '#F6BD17',
          grape: '#543035',
        }
      },
      dark: {
        colors: {
          primary: '#E1EAF9',
          secondary: '#4494DA',
          info: '#226BAA',
          warning: '#AE6002',
          white: '#FFFFFF',
          gray: '#A2AAAD',
          black: '#343A40',
          aggie_blue: '#0F2439',
          sky_blue: '#288DC2',
          teal: '#00938F',
          watermelon: '#F16278',
          tangerine: '#F58220',
          sunflower: '#F6BD17',
          grape: '#543035',
        }
      }
    },
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
    }
  },
  components,
  directives,
})
