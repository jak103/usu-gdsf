<template>
  <v-container>
    <v-row>
      <!-- Back Button -->
      <BackButton></BackButton>
    </v-row>

    <v-divider></v-divider>
    <div></div>
    <v-row class="ma-3">
      <!-- Carousel -->
      <GameCarousel></GameCarousel>
    </v-row>

    <v-divider></v-divider>

    <!-- Break out Games by: Semester, Type -->
    <v-row class="mt-4">
      <h1>Please view our new arrivals to the store!</h1>
    </v-row>
    <v-row justify="center">
        <v-btn class="ml-3 mr-12 mb-12 mt-12" color="primary" @click="">Last 3 Months
        </v-btn>
        <v-btn class="ma-12" color="secondary" @click="">Last Month
        </v-btn>
        <v-btn class="ma-12" color="info" @click="">Last Week
        </v-btn>
    </v-row>
  </v-container>
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import BackButton from '../components/BackButton.vue';


export default defineComponent({
  name: 'arrivals',

  components: {
    GameCarousel,
    BackButton
},

  computed: {

  },

  methods: {

  }
});
</script>