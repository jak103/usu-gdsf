<template>
  <v-container v-if="!dataLoading">
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
            <div v-else class="" @click.stop="clickGameCard(mostPopularGame.Name, mostPopularGame.Id)">
             <v-img v-if="mostPopularGame.ImagePath"
							height="150"
							:src=mostPopularGame.ImagePath
						  ></v-img>
              <h2 style="text-align: center">{{mostPopularGame["Name"]}}</h2>
              <p style="text-align: center">Developed by: {{mostPopularGame["Developer"]}}  |  Times Played: {{mostPopularGame["TimesPlayed"]}}</p>
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

    <v-row v-for="genre in genres" class="mt-4">
      <v-sheet v-if="sortedGames[genre]?.length > 0">
        <v-row>
          <div :data-test="`${genre}-title`" class="ml-10 pb-3 text-h5">{{ genre }}</div>
        </v-row>
        <v-row class="ma-2">
          <GameList :games='sortedGames[genre]'></GameList>
        </v-row>
      </v-sheet>
    </v-row>
    <v-row>
      <Footer></Footer>
    </v-row>
  </v-container>
  
	<Loading data-test="loadbar" v-if="dataLoading" text="Game Data" containerStyle="height: 75vh"/>
</template>

<script>
import { defineComponent } from 'vue';
import GameCarousel from '../components/GameCarousel.vue';
import GameList from '../components/GameList.vue';
import GameCardView from '../components/GameCardView.vue';
import Footer from "../components/Footer.vue";
import Loading from '../components/Loading.vue';

import Game from '../models/game.js'
import * as GamesServices from '../services/gamesServices.js';

export default defineComponent({
  name: 'HomeView',

  components: {
    GameCarousel,
    GameList,
    GameCardView,
    Footer,
    Loading
  },

  data() {
    return {
      genres: [],
      sortedGames: {},
      dataLoading: false,
      mostPopularGame: {},
    }
  },

  methods: {
    uniqueVals(a){
      var prims = {"boolean":{}, "number":{}, "string":{}}, objs = [];

      return a.filter(function(item) {
        var type = typeof item;
        if(type in prims)
            return prims[type].hasOwnProperty(item) ? false : (prims[type][item] = true);
        else
            return objs.indexOf(item) >= 0 ? false : objs.push(item);
      });
    },
    clickGameCard(name, id){
			this.$router.push(`/games/info/${name}/${id}`)
    },
    async getGenres(){
      this.dataLoading = true;
      await GamesServices.getAllTags()
        .then(response => {
          this.genres = this.uniqueVals(response?.data);
          this.genres.sort().forEach(genre => this.getGamesWithGenre(genre))
        }).catch(error => {
          console.log(error);
          this.dataLoading = false
        })
    },
    async getGamesWithGenre(genre){
      await GamesServices.getGamesWithTags([ genre ])
        .then(response => {
          this.sortedGames[genre] = response?.data.map(g => Game.populateFromObject(g));
          this.dataLoading = false;
        }).catch(error => {
          console.log(error);
          this.dataLoading = false
        })
    },
    async getMostPupularGame() {
				this.dataLoading = true;
				await GamesServices.getMostPopularGame()
					.then(response => {
						this.mostPopularGame = response.data;
						this.dataLoading = false
					}).catch(error => {
						console.log(error.response.data);
						this.dataLoading = false
					});
			},
  },

  created() {
    this.getGenres();
    this.getMostPupularGame();
  }
});
</script>
