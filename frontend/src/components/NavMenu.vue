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
      <v-tooltip
        v-if="rail"
        activator="parent"
        location="end"
      >Game Dev Store</v-tooltip>
      <template v-slot:append>
        <v-btn
          variant="text"
          icon="mdi-chevron-left"
          @click.stop="rail = !rail"></v-btn>
      </template>
    </v-list-item>

    <v-list-item
      v-if="loggedIn"
      :title="userName"
      :prepend-avatar="avatar"
      to="/account"
    >
      <v-tooltip
        v-if="rail"
        activator="parent"
        location="end"
      >My Account</v-tooltip>
    </v-list-item>

    <v-divider></v-divider>

    <v-list-item
      prepend-icon="mdi-magnify"
      link
    >
      <v-tooltip
        v-if="rail"
        activator="parent"
        location="end"
      >Search</v-tooltip>
      <v-text-field
        @keydown.enter="goSearch()"
        v-model="searchValue"
        label="Search Games"
        single-line
        full-width
        hide-details
        clear-icon="mdi-close-circle"
        clearable
        type="text"
      ></v-text-field>
    </v-list-item>

    <v-list>
    <!-- All users can access these items -->
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
      
    <!-- Only admins can access these. -->
      <div v-if="loggedIn">
        <v-divider></v-divider>
        <v-list-item 
          v-for="item in adminNavItems"
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
      </div>
    </v-list>

    <template v-slot:append>
      <v-list-item
        v-if="!loggedIn"
        title="Admin Login"
        prepend-icon="mdi-login"
        color="white"
        to=/admin/login
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
        to="/"
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
    loggedIn: true, // This will need to be dynamically set once logging in is defined.
    userName: "John Doe", // This will need to be updated once logging in is implemented.
    avatar: "https://randomuser.me/api/portraits/men/85.jpg", // Same here
    navItems: [
      { title: 'Home', icon: 'mdi-home', path: '/' },
      { title: 'All Games', icon: 'mdi-gamepad-square', path: '/games/all' },
      { title: 'About', icon: 'mdi-information', path: '/about' },
    ],
    adminNavItems: [
      { title: 'Manage Users', icon: 'mdi-account-edit', path: '/admin/users' },
      { title: 'Manage Games', icon: 'mdi-circle-edit-outline', path: '/admin/games/manage' },
    ],
    rail: true,
    searchValue: ""
  }),

  methods: {
    selected(path) {
      return this.$router.currentRoute.value.path === path;
    },

    logout() {
      this.loggedIn = !this.loggedIn;

      // Add logic to perform logout.
    },

    goSearch() {
      // Do search stuff
      console.log(this.searchValue);
    }
  },

  components: {

  },
});
</script>
