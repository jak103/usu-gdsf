<template>
  <div>
    <v-toolbar color=teal>
      Game Dev Store
      <v-btn icon :to="`/`">
        <v-icon>
          mdi-home
        </v-icon>
      </v-btn>
    </v-toolbar>
    <v-list nav>
      <v-list-item to="/">Home</v-list-item>
      <v-list-item to="/about">About</v-list-item>
      <v-list-item v-if="appBarDisplay" :to="`/info/${defaultGame}`">Game</v-list-item>
    </v-list>

      <v-btn to="/admin" color="secondary">Login</v-btn>
  </div>
</template>

<script>
import { defineComponent } from 'vue';


export default defineComponent({
  name: 'NavMenu',

  data: () => ({
    defaultGame: 1
  }),

  computed: {
    appBarDisplay() {
      return this.$router.currentRoute.value.path !== '/';
    }
  },

  components: {

  },
});
</script>
