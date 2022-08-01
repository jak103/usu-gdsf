<template>
  <v-container>
    <v-row class="ma-2">
      <v-col>
        <GameCardView></GameCardView>
      </v-col>
      <v-col>
        <GameCardView></GameCardView>
      </v-col>
      <v-col>
        <GameCardView></GameCardView>
      </v-col>
      <v-col>
        <GameCardView></GameCardView>
      </v-col>
      <v-col>
        <GameCardView></GameCardView>
      </v-col>
      <v-col>
        <GameCardView></GameCardView>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import GameCardView from '../components/GameCardView.vue'

export default {
  name: 'GameList',
  
  data: () => ({
    something: 'This is where we will implement a dynamic view of different games. We will need to pass in props.'
  }),
  
  methods: {
    
  },

  components: {
    GameCardView,
  }
}
</script>