<template>
    <v-carousel v-if="screenshotUrls.length > 0" cycle height="auto" hide-delimiters>
      <v-carousel-item v-for="url in screenshotUrls.slice(0,100)">
        <v-img :src="url" contain max-height="300"></v-img>
      </v-carousel-item>
    </v-carousel>
    <div v-else>
      Hey backend team, tag some games with "HomePage" and make sure they have image URLs and this carousel will have images.
    </div>
</template>

<script>
export default {
  name: 'GameCarousel',

  data: () => ({
    screenshotUrls: [],
  }),

  methods: {
    async getHomeViewImages() {
      // TODO: This very much should be moved up and out of this component. 
      // Components in general should only ever display data, it should
      // never be the one to extract data from the backend. Also, use the
      // general gamesServices.js rather than calling this out with axios here.
      await axios.get('http://127.0.0.1:8080/games')
        .then(response => {
          this.screenshotUrls = 
            response.data
              .filter(game => game['Tags'] && game['Tags'].includes("HomePage"))
              .map(rawGame => rawGame["ImagePath"])
        }).catch(error => {
						console.log(error.response.data);
					});
    }
  },

  

}
</script>