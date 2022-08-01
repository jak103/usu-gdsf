<template>
  <v-card max-width="240">
    <v-img
        height="160"
        :src="game.imagePath"
    ></v-img>

    <v-card-title>{{ game.name }}</v-card-title>
    <v-card-subtitle>{{ game.developer }}</v-card-subtitle>
    <v-card-text>{{ game.description }}</v-card-text>
  </v-card>

</template>

<script>
export default {
	name: 'GameCardView',
	props: ["game"],
	data: () => ({

	}),

	computed: {

	},

	methods: {

	}
};

</script>