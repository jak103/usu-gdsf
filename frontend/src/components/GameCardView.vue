<template>
  <v-card max-width="240">
    <v-img
        height="160"
        src="./src/assets/minecraft.png"
    ></v-img>

    <v-card-title>{{ gameTitle }}</v-card-title>
    <v-card-subtitle>{{ developer }}</v-card-subtitle>
	<Rating
		starColor="#0F2439"
		starBackgroundColor="#8a8d8f"
		isHover="true"
		isHalfIncrements="true"
		>

	</Rating>
    <v-card-text>{{ description }}</v-card-text>
  </v-card>

</template>

<script>
import Rating from './Rating.vue';
export default {
	name: 'GameCardView',

	components: {
		Rating,
	},
	

	data: () => ({
		gameTitle: 'Minecraft',
		developer: 'Mojang Studios',
		rating: 3,
		image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg',
		description: 'Explore your own unique world, survive the night, and create anything you can imagine!'
	}),

	computed: {

	},

	methods: {

	}
};

</script>