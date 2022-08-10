<template>
  <v-row>
    <div v-for = "item in gameInfo" :key="item.id">
    <GameCard
      :title="item.title"
      :author="item.title"
      :description="item.description"
      image="https://cdn.cloudflare.steamstatic.com/steam/apps/489830/capsule_616x353.jpg?t=1650909796"/>
    </div>
  </v-row>
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
      gameInfo: [1,2,3,4,5,6,7,8]
    }),

    components: {
      GameCard,
    },

    mounted() {
      this.getData()
    },

    methods: {
      async getData() {
        let res = await this.$axios.get('/game')
        this.gameInfo = res.data[0]
      }
    }
  }
</script>