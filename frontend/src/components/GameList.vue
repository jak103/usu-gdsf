<template>
<div v-for = "item in [0,1,2,3,4,5,6,7]">
  <GameCard :title="'Game Title'" :author="'Game Author'" :description="'Game Description'" />
  </div>
</template>

<script>
  import GameCard from './GameCard.vue';
  export default {
    name: 'GameList',

    data: () => ({
      show: false,
    }),

    components: {
      GameCard,
    },
  }
</script>