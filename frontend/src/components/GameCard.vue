<template>
  <v-container>
    <v-card class="flex">
        <v-card-title>{{gameTitle}}</v-card-title>
        <v-row class="ma-2">
            <v-col class="flex xs=12 md=6">
                <v-img ref="gameImage" width="200" :src="image"></v-img>       
                <v-card-actions class="center">
                    <span class="grey--text text--lighten-2 text-caption">
                        ({{ rating }})
                    </span>
                    <v-rating readonly v-model="rating" half-increments color="amber" size="x-small"></v-rating>
                </v-card-actions>     
            </v-col>
            <v-col class="flex xs=12 md=6">
                <v-card-subtitle v-if=showGameOverview>{{gameOverview}}</v-card-subtitle>
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
    showGameOverview: {
        type: Boolean,
        default: false
    },
    rating: Number,
    image: String
  },
  
  methods: {
    
  }
}
</script>