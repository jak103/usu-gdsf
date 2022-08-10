<template>
<div v-for = "item in gameInfo" :key="item.id"
>
  <GameCard 
  :title="item.title" 
  :author="Game Author" 
  :description="item.description" />
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
      gameInfo: []
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