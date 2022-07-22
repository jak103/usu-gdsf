<template>
  <v-container>
    <v-card class="flex">
        <v-card-title>{{gameTitle}}</v-card-title>
        <v-row class="ma-2">
            <v-col cols="6">
                <v-img width="200" :src="image"></v-img>       
                <v-card-actions class="center">
                    <span class="grey--text text--lighten-2 text-caption">
                        ({{ rating }})
                    </span>
                    <v-rating readonly v-model="rating" half-increments color="amber"></v-rating>
                </v-card-actions>     
            </v-col>
            <v-col cols="6" class="flex">
                <v-card-subtitle>{{gameOverview}}</v-card-subtitle>
                <v-card-actions>
                    <v-btn>
                        Play    
                    </v-btn>
                </v-card-actions>
            </v-col>
        </v-row>
	</v-card>
  </v-container>
</template>

<script>
export default {
  name: 'GameCard',

  props: {
    gameTitle: String,
    gameOverview: String,
    rating: Float32Array,
    image: String
  },
  
  methods: {
    
  }
}
</script>