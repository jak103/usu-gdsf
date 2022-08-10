<template>
  <v-container v-if="games.length > 0">
    <v-row>
      <GameCardView class="ma-2" v-for="game in games" :game="game"></GameCardView>
    </v-row>
  </v-container>
  <div v-else>
    Hey backend team, if you want game cards to show up here, tag games and make sure to add the image URLs
  </div>
</template>

<script>
import GameCardView from '../components/GameCardView.vue'
import Game from '../models/game';

export default {
  name: 'GameList',
  
  props: { games: [Game] },

  components: {
    GameCardView,
  },
}
</script>