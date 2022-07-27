<template>
  <v-carousel cycle>
    <v-carousel-item v-for="url in screenshotUrls">
      <img :src="url">
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