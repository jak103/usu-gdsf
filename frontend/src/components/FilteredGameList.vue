<template>
  <v-container fluid>
    <v-data-iterator v-for="game in this.getTopRated()" >
        <GameCard :gameTitle=game.gameTitle :gameOverview=game.gameOverview :rating=game.rating :image=game.image showGameOverview></GameCard>
    </v-data-iterator>
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
    methods: {
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
    },
    components: { GameCard }
}
</script>