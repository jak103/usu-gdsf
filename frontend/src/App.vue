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

  <Footer/>

  </v-app>
</template>

<script>
import HomeView from './views/HomeView.vue'
import NavMenu from './components/NavMenu.vue'
import Footer from './components/Footer.vue'

export default {
  name: 'App',

  components: {
    HomeView,
    NavMenu,
    Footer
},

  data: () => ({

  }),

  computed: {

  }
}
</script>
