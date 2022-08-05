<template>
  <v-app>
    <v-navigation-drawer
      app
      color="primary"
    >
      <NavMenu/>
    </v-navigation-drawer>

  <!-- Sizes your content based upon application components -->
  <v-main>

    <!-- Provides the application the proper gutter -->
    <v-container fluid>
      <router-view></router-view>
    </v-container>
  </v-main>

  <v-footer color="primary" app
  paddless
  >
  <v-row
  justify="center"
  no-gutters
  >
  <v-btn
        v-for="link in links"
        :key="link"
        color="secondary"
        text
        rounded
        class="my-2"
        :href="link.url"

      >
        {{ link.text }}
      </v-btn>
      <v-col
        class="primary lighten-2 text-center black--text"
        cols="12"
      >
        {{ new Date().getFullYear() }} — <strong>Vuetify</strong>
      </v-col>
  </v-row>
  </v-footer>
  </v-app>
</template>

<script>
import HomeView from './views/HomeView.vue'
import NavMenu from './components/NavMenu.vue'

export default {
  name: 'App',

  components: {
    HomeView,
    NavMenu
},

  data: () => ({
    links: [
      {
        text: "USU CS Department",
        url: "www.usu.edu/cs"
      },
      {
        text: "About Us",
        url: "/about"
      },

    ],

  }),

  computed: {

  }
}
</script>
