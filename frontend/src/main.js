import 'vuetify/styles'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createVuetify } from 'vuetify'
import { loadFonts } from './plugins/webfontloader'

import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

const app = createApp(App)
const vuetify = createVuetify(components,directives)
loadFonts()

app.use(router)
  .use(vuetify)
  .mount('#app')
