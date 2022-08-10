<template>
  <v-card width="240" height="375" @click.stop="onClick(game.Id, game.Name)">
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
		:isReadOnly="true"
		:isDense="true"
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
			required: true
		}
	},
	components: {
		Rating,
	},
	methods: {
		onClick(id, name) {
			this.$router.push(`/games/info/${name}/${id}`)
		}
	}
};

</script>