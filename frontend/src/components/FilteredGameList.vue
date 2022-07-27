<template>
  <v-container fluid>
    <!--This is a drop down if we want it idk-->
      <v-select
          v-on:input="testFunction"
          v-model="filter"
          label="Select a Filter"
          :items="options"
          solo
        ></v-select>
    <!-- <v-virtual-scroll> -->
    <v-list v-for="game in this.filteredList">
     <!-- gameOverview="game.gameOverview" rating="game.rating" image="game.image" showGameOverview -->
        <GameCard :gameTitle=game.gameTitle :gameOverview=game.gameOverview :rating=game.rating :image=game.image showGameOverview></GameCard>
    </v-list>
    <!-- </v-virtual-scroll> -->
  </v-container>
</template>

<script>
import GameCard from './GameCard.vue'

export const ratings = {
  MOST_POPULAR: 'Most Popular',
  TOP_RATED: 'Top Rated'
}
export default {
    name: "FilteredGameList",
    data: () => ({
        filteredList: [],
        options: [
          ratings.MOST_POPULAR,
          ratings.TOP_RATED
        ], 
        filter: ratings.MOST_POPULAR
    }),
    methods: {
        getFilteredList: function (filter) {
            if (filter == ratings.MOST_POPULAR) {
              this.filteredList = this.getMostPopular()
            }
            else if (filter == ratings.TOP_RATED) {
              this.filteredList = this.getTopRated()
            }
        },

        getTopRated: function() {
          //TODO run some query or something
          return [
            { 
              gameTitle: 'Cat Attack', 
              gameOverview: 'Craig the cat embarks on an adventure', 
              rating: 4,
              image: 'https://cdn.theatlantic.com/media/mt/science/cat_caviar.jpg'
            },
            { 
              gameTitle: 'Big Blue Game',
              rating: 3,
              gameOverview: 'This game is big. This game is blue. This game is fun.', 
              image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
            }
          ];
        },

        getMostPopular: function() {
          //TODO run some query or something
          return [
            { 
              gameTitle: 'Big Blue Game',
              rating: 3,
              gameOverview: 'This game is big. This game is blue. This game is fun.', 
              image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
            },
            { 
              gameTitle: 'Cat Attack', 
              gameOverview: 'Craig the cat embarks on an adventure', 
              rating: 4,
              image: 'https://cdn.theatlantic.com/media/mt/science/cat_caviar.jpg'
            }
          ];
        },
    },
    watch: {
      filter : function(filter){
        this.getFilteredList(filter)
      }
    },
    components: { GameCard }
}
</script>