<template>
    <v-carousel cycle height="auto" hide-delimiters>
      <v-carousel-item v-for="url in screenshotUrls">
        <v-img :src="url" contain max-height="300"></v-img>
      </v-carousel-item>
    </v-carousel>
</template>

<script>
export default {
  name: 'GameCarousel',

  data: () => ({
      
  }),

  props: ["screenshotUrls"]

}
</script>