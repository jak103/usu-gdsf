<template>
  <v-container>
    <v-row>
      <v-card height="200" width="2000" color="#10243c">
        <p class="text-center font-weight-thin" style="color:#FFFFFF;font-size: 100px">
          USU Game Dev Store Front
        </p>
      </v-card>
    </v-row>
    <v-row class="ma-3">
      <v-col>
        <v-card height="300">
          Something else here
        </v-card>
      </v-col>
      <!-- Carousel -->
      <v-col>
        <GameCarousel :screenshotUrls="['https://i.ytimg.com/vi/-Gq2S0AXjNw/mqdefault.jpg', 'https://www.techspot.com/images2/news/bigimage/2016/02/2016-02-11-image.jpg']"></GameCarousel>
      </v-col>
      <v-col>
        <v-card height="300">
          Something else here
        </v-card>
      </v-col>    
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
