<template>
  <!-- <v-container> -->
    <v-row>
      <v-card height="165" width="2000" color=primary>
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
            <p v-if="!mostPopularGame" class="text-center font-weight-thin" style="color:#FFFFFF;font-size: 15px">
              Unable to get the most popular game at this time
            </p>
            <div v-else class="" @click.stop="handleClickGame(mostPopularGame.Id)">
             <v-img v-if="mostPopularGame.ImagePath"
							height="150"
							:src=mostPopularGame.ImagePath
						  ></v-img>
              <h2 style="text-align: center">{{mostPopularGame["Name"]}}</h2>
              <p style="text-align: center"> Made by: {{mostPopularGame["Developer"]}}  |  Times Played: {{mostPopularGame["TimesPlayed"]}}</p>
            </div>
        </v-card>
      </v-col>
      <v-col>
        <GameCarousel></GameCarousel>
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
      <GameList tag="spring2022"></GameList>
    </v-row>

    <v-row>
      <h1>Puzzles</h1>
    </v-row>
    <v-row class="ma-2">
      <GameList tag="puzzle"></GameList>
    </v-row>

    <v-row>
      <h1>Shooters</h1>
    </v-row>
    <v-row class="ma-2">
      <GameList tag="shooter"></GameList>
    </v-row>
  <v-row>
    <Footer></Footer>
  </v-row>
  <!-- </v-container> -->
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import Game from '../models/game.js'
import GameList from '../components/GameList.vue';
import GameCardView from '../components/GameCardView.vue';
import Footer from "../components/Footer.vue";
import axios from "axios";

export default defineComponent({
  name: 'HomeView',

  components: {
    GameCarousel,
    GameList,
    GameCardView,
    Footer
  },

  data() {
    return {
      exampleGame: new Game(),
      mostPopularGame: {},
    }
  },

  computed: {
  },

  methods: {
    async getMostPupularGame() {
				this.dataLoading = true;
				// we may want to configure a base-url for this, because it won't work on production
				await axios.get('http://127.0.0.1:8080/most_popular')
					.then(response => {
						this.mostPopularGame = response.data;
            console.log(this.mostPopularGame, response.data);
						this.dataLoading = false
					}).catch(error => {
						console.log(error.response.data);
						this.dataLoading = false
					});
			},
      handleClickGame(id) {
				this.$router.push("/info/" + id)
			},
  },
  created() {
			this.getMostPupularGame();
		}
});
</script>
