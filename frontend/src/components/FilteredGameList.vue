<template>
  <h3>Filtered Games</h3>
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
            },
            { 
              title: 'Big Blue Game',
              description: 'This game is big. This game is blue. This game is fun.', 
              author: 'Big Blue',
            },
                        { 
              title: 'Cat Attack', 
              description: 'Craig the cat embarks on an adventure', 
              author: 'Craig',
            },
            { 
              title: 'Big Blue Game',
              description: 'This game is big. This game is blue. This game is fun.', 
              author: 'Big Blue',
            },
                        { 
              title: 'Cat Attack', 
              description: 'Craig the cat embarks on an adventure', 
              author: 'Craig',
            },
            { 
              title: 'Big Blue Game',
              description: 'This game is big. This game is blue. This game is fun.', 
              author: 'Big Blue',
            }
          ];
        },
    },
    components: { GameCard }
}
</script>