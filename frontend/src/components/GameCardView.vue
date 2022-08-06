<template>
  <v-card max-width="240" @click.stop="OnClick(game.Id, game.Name)">
    <v-img
        height="160"
        :src="game.ImagePath"
    ></v-img>

    <v-card-title>{{ game.Name }}</v-card-title>
    <v-card-subtitle>{{ game.Developer }}</v-card-subtitle>
	<Rating
		starColor="#0F2439"
		starBackgroundColor="#8a8d8f"
		:rating="game.Rating"
		:isHover="true"
		:isHalfIncrements="true"
		>
	</Rating>
    <v-card-text>{{ game.Description }}</v-card-text>
  </v-card>

</template>

<script>
import Rating from './Rating.vue';
import Game from '../models/game';
export default {
	name: 'GameCardView',
	props: {
		game: {
			type: Game,
			required: true,
			default() {
				return new Game() // This will need to be removed at some point once we've only created one of these Game Cards if there is data.
			}
		}
	},
	components: {
		Rating,
	},
	methods: {
		OnClick(id, name) {
			this.$router.push(`/games/info/${name}/${id}`)
		}
	}
};

</script>