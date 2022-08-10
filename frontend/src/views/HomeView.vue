<template>
  <v-container style="overflow:auto">
    <h2>Featured Games</h2>
    <v-row class="ma-3">
      <!-- Carousel -->
      <GameCarousel tag="Featured"></GameCarousel>
    </v-row>

    <v-divider></v-divider>

    <!-- Break out Games by: Semester, Type -->
    <v-row class="mt-4">
      <FilteredGameList></FilteredGameList>
    </v-row>
  </v-container>
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import GameList from '../components/GameList.vue'
import FilteredGameList from '../components/FilteredGameList.vue';


export default defineComponent({
  name: 'HomeView',

  components: {
    GameCarousel,
    GameList,
    FilteredGameList
},

  computed: {

  },

  methods: {

  }
});
</script>
