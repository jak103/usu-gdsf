<template>
  <v-navigation-drawer
    permanent
    color="primary"
    :rail="rail"
    @click="rail = false"
  >
    <!-- TODO: This icon is only a place holder until we can get the appropriate URL for the official Utah State Icon, or we decide on some other icon. -->
    <v-list-item
      title="Game Dev Store"
      prepend-avatar="https://gmedia.playstation.com/is/image/SIEPDC/game-library-white-icon-01-en-09nov21?$native--t$"
    >
      <template v-slot:append>
        <v-btn
          variant="text"
          icon="mdi-chevron-left"
          @click.stop="rail = !rail"
        ></v-btn>
      </template>
    </v-list-item>

    <!-- TODO: Add another entry here somehow for the logged in user, when logged in. -->

    <v-divider></v-divider>

    <v-list>
      <v-list-item 
        v-for="item in navItems"
        :value="item.title"
        :title="item.title"
        :to="item.path"
        :prepend-icon="item.icon"
        color="secondary"
      >
        <v-tooltip
          v-if="rail"
          activator="parent"
          location="end"
        >{{ item.title }}</v-tooltip>
      </v-list-item>
    </v-list>

    <template v-slot:append>
      <v-list-item
        v-if="!loggedIn"
        title="Admin Login"
        prepend-icon="mdi-login"
        color="white"
        to=/admin
      >
        <v-tooltip
          v-if="rail"
          activator="parent"
          location="end"
        >Admin Login</v-tooltip>
      </v-list-item>
      
      <v-list-item
        v-if="loggedIn"
        title="Logout"
        prepend-icon="mdi-logout"
        color="white"
        @click.stop="logout()"
      >
        <v-tooltip
          v-if="rail"
          activator="parent"
          location="end"
        >Logout</v-tooltip>
      </v-list-item>
    </template>
  </v-navigation-drawer>

</template>

<script>
import { defineComponent } from 'vue';


export default defineComponent({
  name: 'NavMenu',

  data: () => ({
    defaultGame: 1,
    drawer: true,
    navItems: [
      { title: 'Home', icon: 'mdi-home', path: '/' },
      { title: 'Games', icon: 'mdi-gamepad-square', path: '/games' },
      { title: 'About', icon: 'mdi-information', path: '/about' },
    ],
    rail: true,
    loggedIn: false // This will need to be dynamically set once logging in is defined.
  }),

  computed: {
    selected(path) {
      return this.$router.currentRoute.value.path === path;
    },

    logout() {
      this.loggedIn = !this.loggedIn

      // Add logic to perform logout.
    }
  },

  components: {

  },
});
</script>
