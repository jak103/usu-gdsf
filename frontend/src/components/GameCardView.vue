<template>
  <v-card max-width="240">
    <v-img
        height="160"
        :src="game.imagePath"
    ></v-img>

    <v-card-title>{{ game.name }}</v-card-title>
    <v-card-subtitle>{{ game.developer }}</v-card-subtitle>
	<Rating
		starColor="#0F2439"
		starBackgroundColor="#8a8d8f"
		:rating="game.rating"
		:isHover="true"
		:isHalfIncrements="true"
		>
	</Rating>
    <v-card-text>{{ game.description }}</v-card-text>
  </v-card>

</template>

<script>
import Rating from './Rating.vue';
export default {
	name: 'GameCardView',
	props: ["game"],

	components: {
		Rating,
	},
	

	data: () => ({

	}),

	computed: {

	},

	methods: {

	}
};

</script>