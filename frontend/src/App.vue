<template>
  <v-app>
    <NavMenu app/>

    <!-- Sizes your content based upon application components -->
    <v-main>

      <!-- Provides the application the proper gutter -->
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>

    <v-footer color="teal" app>
      Bottom navigation info and copyright stuff. &copy; 2022
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

  }),

  computed: {

  }
}
</script>
