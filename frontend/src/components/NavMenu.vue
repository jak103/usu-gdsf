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
      <v-list-item @click="gameSelected = !gameSelected"> Games </v-list-item>
      <v-divider v-if="gameSelected"></v-divider> 
      <v-list-item v-if="gameSelected" :to="`/info/${defaultGame}`">Game</v-list-item>
      <v-divider v-if="gameSelected"></v-divider> 
      <v-list-item v-if="gameSelected" :to="`/info/2`">Game 2</v-list-item>
    </v-list>
    <v-btn to="/admin" color="secondary">Login</v-btn>
  </div>
</template>

<script>
import { defineComponent } from 'vue';


export default defineComponent({
  name: 'NavMenu',

  data: () => ({
    defaultGame: 1,
    gameSelected: false,
  }),
  
});
</script>
