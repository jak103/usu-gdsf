<template>
  <v-container v-if="!dataLoading">
    <v-row class="ma-2">
      <v-col v-for="game in taggedGames.slice(0,5)">
        <GameCardView :game="game"></GameCardView>
      </v-col>
    </v-row>
  </v-container>
  <Loading data-test="loadbar" v-if="dataLoading" text="Loading Game Data" containerStyle="height: 75vh"/>
</template>

<script>
import GameCardView from '../components/GameCardView.vue'
import axios from "axios"
import Loading from '../components/Loading.vue';
import Game from '../models/game';


export default {
  name: 'GameList',
  
  data: () => ({
    taggedGames: [],
    dataLoading: true,
  }),
  methods: {
    async getTaggedGames() {
      this.dataLoading = true
      await axios.get('http://127.0.0.1:8080/games')
        .then(response => {
          this.taggedGames = 
            response.data[0]
              .filter(game => game['Tags'] && game['Tags'].includes(this.tag))
              .map(rawGame => new Game(
                rawGame['Id'], 
                rawGame['Name'], 
                rawGame['Developer'], 
                rawGame['Rating'],
                rawGame['TimesPlayed'],
                rawGame["ImagePath"],
                rawGame['Description'],
                rawGame['Tags'],
                )
              )
          this.dataLoading = false
        }).catch(error => {
						console.log(error.response.data);
						this.dataLoading = false
					});
    }
  },
  
  props: ["tag"],

  components: {
    GameCardView,
    Loading,
  },

  created() {
    this.getTaggedGames()
  }
}
</script>