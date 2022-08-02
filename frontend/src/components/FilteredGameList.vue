<template>
  <v-sheet
    class="mx-auto"
    max-width="100%"
  >
    <v-slide-group
      show-arrows
    >
      <v-slide-group-item
        v-for="(game, index) in getTopRated"
        :key="index"
        v-slot="{ isSelected, toggle }"
      >
        <GameCard
          :title="game.title"
          :author="game.author"
          :description="game.description"
          :image="game.image"
        ></GameCard>
      </v-slide-group-item>
    </v-slide-group>
  </v-sheet>
</template>

<script>
import GameCard from './GameCard.vue'

export default {
    name: "FilteredGameList",
    computed: {
        getTopRated() {
          //TODO run some query or something
          return [
            { 
              title: 'Cat Attack', 
              description: 'Craig the cat embarks on an adventure', 
              author: 'Craig',
              image: 'https://cdn.theatlantic.com/media/mt/science/cat_caviar.jpg'
            },
            { 
              title: 'Big Blue Game',
              description: 'This game is big. This game is blue. This game is fun.', 
              author: 'Big Blue',
              image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
            },
                        { 
              title: 'Cat Attack', 
              description: 'Craig the cat embarks on an adventure', 
              author: 'Craig',
              image: 'https://cdn.theatlantic.com/media/mt/science/cat_caviar.jpg'
            },
            { 
              title: 'Big Blue Game',
              description: 'This game is big. This game is blue. This game is fun.', 
              author: 'Big Blue',
              image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
            },
                        { 
              title: 'Cat Attack', 
              description: 'Craig the cat embarks on an adventure', 
              author: 'Craig',
              image: 'https://cdn.theatlantic.com/media/mt/science/cat_caviar.jpg'
            },
            { 
              title: 'Big Blue Game',
              description: 'This game is big. This game is blue. This game is fun.', 
              author: 'Big Blue',
              image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
            }
          ];
        },
    },
    components: { GameCard }
}
</script>