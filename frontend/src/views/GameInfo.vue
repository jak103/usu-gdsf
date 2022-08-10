<template>
	<v-container>
		<h2>{{ gameTitle }} {{ $route.params.id }}</h2>

		<v-card>
			{{ rating }}
			<v-img width="200" :src="image"></v-img>
			<router-link :to="{ name: 'info', params: { id: '2' }}">To another game</router-link>

		</v-card>
	</v-container>
</template>

<script>
export default {
	name: 'GameInfo',

	data: () => ({
		gameTitle: 'Big Blue Game',
		rating: 3,
		image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
	}),

	computed: {

	},

	methods: {

	}
};

</script>