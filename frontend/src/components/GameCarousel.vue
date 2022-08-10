<template>
  <v-card width="100%">
    <v-carousel cycle :interval="cycleTime">
      <v-carousel-item v-for="game in gameData" :key="game">
        <v-sheet height="100%" tile>
          <v-row class="fill-height" align="center" justify="center">
            <v-col align="center" justify="center">
              <v-row class="text-h2 fill-height" align="center" justify="center">
                Title: {{ game.title}}
              </v-row>
              <v-row align="center" justify="center">
                Description: {{ game.description }}
              </v-row>
            </v-col>
          </v-row>
        </v-sheet>
      </v-carousel-item>
    </v-carousel>
  </v-card>
</template>

<script>

export default {
  name: 'GameCarousel',

  props: {
    tag: { default: '' }
  },

  data: () => ({
    colors: ['primary', 'secondary', 'info', 'warning', 'white', 'teal', 'watermelon'],
    cycleTime: 4000,
    gameData: []
  }),

  async mounted() {
    try{
      let query = '';
      if(this.tag) {
        query = `?tag=${this.tag}`
      }
      const response = await this.$axios.get(`/game${query}`);
      this.gameData = response.data[0];
    } catch {
      console.log('Error')
    }
  }
}
</script>