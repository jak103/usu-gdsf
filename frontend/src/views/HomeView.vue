<template>
  <v-container>
    <p> Team 1's Frontend</p>
    <v-row class="ma-3">
      <!-- Carousel -->
      <GameCarousel></GameCarousel>
    </v-row>

    <v-divider></v-divider>

    <!-- Break out Games by: Semester, Type -->
    <v-row class="ma-2">
      <v-col cols="6">
        <GameList></GameList>
      </v-col>

      <v-divider vertical></v-divider>

      <v-col cols="6">
        <GameList></GameList>
      </v-col>
    </v-row>

    <v-divider></v-divider>

    <v-row class="ma-2">
      <GameList></GameList>
    </v-row>

    <v-row class="ma-2">
        <GameCardView></GameCardView>
    </v-row>
  </v-container>
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import GameList from '../components/GameList.vue'
import GameCardView from '../components/GameCardView.vue'


export default defineComponent({
  name: 'HomeView',

  components: {
    GameCarousel,
    GameList,
    GameCardView
  },

  computed: {

  },

  methods: {

  }
});
</script>
