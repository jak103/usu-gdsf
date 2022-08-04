<template>
  <v-container>
    <v-row class="ma-3">
      <!-- Carousel -->
      <GameCarousel></GameCarousel>
    </v-row>

    <v-divider></v-divider>

    <!-- Break out Games by: Semester, Type -->
    <v-row class="mt-4">
      <GameList></GameList>
    </v-row>
  </v-container>
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import GameList from '../components/GameList.vue'


export default defineComponent({
  name: 'HomeView',

  components: {
    GameCarousel,
    GameList
  },

  computed: {

  },

  methods: {

  }
});
</script>
