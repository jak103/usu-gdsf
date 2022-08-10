<template>
  <div>
    <h1>{{ tempData.title }}</h1>
    <h3>{{ tempData.shortDescription }}</h3>
    <v-card>
      <v-card-text>
        {{ tempData.description }}
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'GameDetail',

  props: {
    gameId: { }
  },

  data: () => ({
    tempData: {
      title: 'Game Title',
      description: 'A long Game Description',
      shortDescription: 'short descript',
      rating: 3,
      creationDate: '2022-01-01',
      version: 1.0,
      image: ''
    }
  }),

  methods: {
    async getGameData() {
      await this.$axios.$get('');
    }
  }

}
</script>