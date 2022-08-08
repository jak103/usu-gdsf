<template>
<div v-for = "item in [0,1,2,3,4,5,6,7]">
  <GameCard :title="'Game Title'" :author="'Game Author'" :description="'Game Description'" image="https://cdn.cloudflare.steamstatic.com/steam/apps/489830/capsule_616x353.jpg?t=1650909796" />
  </div>
</template>

<script>
  import GameCard from './GameCard.vue';
  export default {
    name: 'GameList',

    props: {
        title: String,
        author: String,
        description: String,
        image: String,
    },

    data: () => ({
      show: false,
    }),

    components: {
      GameCard,
    },
  }
</script>