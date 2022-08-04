<template>
	<div>
		<!-- File uploads: FormData -->
		<h1>This is where you view all games. Thanks for stopping by.</h1>
    <!-- example of data that comes from a backend api ajax call using axios (see getGames method below) -->
    <v-container>
      <v-container v-for="game of games">
        <AxiosExampleGameCard
          :game-title="game.Name"
          :developer="game.Author"
          description="Description would be passed here from a game object"
        >
        </AxiosExampleGameCard>
      </v-container>
    </v-container>
	</div>
</template>

<script>
import axios from "axios";
import AxiosExampleGameCard from "../components/AxiosExampleGameCard.vue";

export default {
	name: 'AllGamesPage',
  components: {
    AxiosExampleGameCard
  },
  data() {
    return {
      games: {}
    };
  },
  methods: {
    getGames() {
      // we may want to configure a base-url for this, because it won't work on production
      axios.get('http://127.0.0.1:8080/games')
          .then(response => {
            this.games = response.data[0];
          }).catch(error => {
        console.log(error.response.data);
      });
    }
  },
  created() {
    this.getGames();
  }
}
</script>