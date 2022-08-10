<template>
<div v-for = "item in gameInfo">
  <GameCard 
    :title="'title'" 
    :author="'author'" 
    :description="'game description'" 
    image="https://cdn.cloudflare.steamstatic.com/steam/apps/489830/capsule_616x353.jpg?t=1650909796"/>

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

    },

    data: () => ({
      show: false,
      gameInfo: [1,2,3,4,5,6,7,8]
    }),

    components: {
      GameCard,
    },

    methods: {
      getData() {
        let res =  this.$axios.get('/game')
      }
    }
  }
</script>