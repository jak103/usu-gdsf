<template>
  <div>
    <v-toolbar color="tangerine">
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
      <!-- Game List: add route that goes to all games page-->
      <v-list-item to="/games"> Game List </v-list-item>
    </v-list>
    <v-btn to="/admin" color="secondary">Login</v-btn>
    <!-- To Do: add buttons for roles (anything that a default user wouldn't see, like admin/author navigation buttons)-->
  </div>
</template>

<script>
import { defineComponent } from 'vue';


export default defineComponent({
  name: 'NavMenu',

//just keeping these values here until we know we don't need them :)
  data: () => ({
    defaultGame: 1,
    gameSelected: false,
    
  }),
  
});
</script>
