import 'vuetify/styles'
import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from '@/router'
import { loadFonts } from './plugins/webfontloader'
import axios from 'axios'

loadFonts()

const app = createApp(App)
  .use(vuetify)
  .use(router)

  app.config.globalProperties.$axios = axios.create({
    baseURL: 'http://localhost:8080'
  });

  app.mount('#app');
