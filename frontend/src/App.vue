<template>
   <v-app>
  <v-navigation-drawer app>
    <!-- -->
    <NavMenu/>
  </v-navigation-drawer>



  <!-- Sizes your content based upon application components -->
  <v-main>

    <!-- Provides the application the proper gutter -->
    <v-container fluid>

      <!-- If using vue-router -->
      <router-view></router-view>
      <!-- <HomeView/> -->
    </v-container>
  </v-main>

  <v-footer app>
    <!-- -->
    <p>Footer</p>
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
    //
  }),
}
</script>
