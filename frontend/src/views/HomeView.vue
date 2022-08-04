<template>
  <!-- <v-container> -->
    <v-row>
      <v-card height="200" width="2000" color=primary>
        <p class="text-center font-weight-thin" style="color:#FFFFFF;font-size: 100px">
          USU Game Dev Store Front
        </p>
      </v-card>
    </v-row>
    <v-row class="ma-3">
      <v-col>
        <v-card height="300" color=primary>
            <p class="text-center font-weight-thin" style="color:white;font-size: 30px">
              Our Most Popular Game
            </p>
            <p class="text-center font-weight-thin" style="color:white;font-size: 15px">
              Game here
            </p>
        </v-card>
      </v-col>
      <v-col>
        <GameCarousel :screenshotUrls="['https://i.ytimg.com/vi/-Gq2S0AXjNw/mqdefault.jpg', 'https://www.techspot.com/images2/news/bigimage/2016/02/2016-02-11-image.jpg', 'https://www.howtogeek.com/wp-content/uploads/2022/03/MinecraftRAMHeader.jpg']"></GameCarousel>
      </v-col>
      <v-col>
        <v-card height="300" color=primary>
            <p class="text-center font-weight-thin" style="color:#FFFFFF;font-size: 30px">
              Your Favorited Games
            </p>
            <p class="text-center font-weight-thin" style="color:#FFFFFF;font-size: 15px">
              Games here
            </p>
        </v-card>
      </v-col>    
    </v-row>

    <!-- Break out Games by: Semester, Type -->
    <!-- These will be the GameList Components-->
    <v-row>
      <h1>Spring 2022</h1>
    </v-row>
    <v-row class="ma-2">
      <GameList :tag="spring2022"></GameList>
    </v-row>

    <v-row>
      <h1>Puzzles</h1>
    </v-row>
    <v-row class="ma-2">
      <GameList :tag="puzzle"></GameList>
    </v-row>

    <v-row>
      <h1>Most Popular</h1>
    </v-row>
    <v-row class="ma-2">
      <GameList :tag="mostPopular"></GameList>
    </v-row>
  <v-row>
    <Footer></Footer>
  </v-row>
  <!-- </v-container> -->
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import GameList from '../components/GameList.vue';
import GameCardView from '../components/GameCardView.vue';
import Footer from "../components/Footer.vue";


export default defineComponent({
  name: 'HomeView',

  components: {
    GameCarousel,
    GameList,
    GameCardView,
    Footer
  },

  computed: {

  },

  methods: {

  }
});
</script>
