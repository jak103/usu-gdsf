<template>
  <v-card max-width="240">
    <v-img
        height="160"
        :src="image"
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
	name: 'AxiosExampleGameCard',
	components: {
		Rating,
	},
  props: {
    gameTitle: {
      required: true,
      type: String
    },
    developer: {
      required: true,
      type: String
    },
    rating: {
      required: false,
      type: String
    },
    image: {
      required: false,
      type: String,
      default: "https://assets-prd.ignimgs.com/2022/06/14/zelda-breath-of-the-wild-1655249167687.jpg"
    },
    description: {
      required: true,
      type: String
    },
  },
};
</script>