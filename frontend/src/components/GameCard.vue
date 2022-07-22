<template>
  <v-container>
    <v-card class="flex">
        <v-card-title>{{gameTitle}}</v-card-title>
        <v-row class="ma-2">
            <v-col cols="6">
                <v-img width="200" :src="image"></v-img>            
            </v-col>
            <v-col cols="6" class="flex">
                <v-card-subtitle>{{gameOverview}}</v-card-subtitle>
                <v-card-actions>
                    <span class="grey--text text--lighten-2 text-caption mr-2">
                        ({{ rating }})
                    </span>
                    <v-rating
                    value="rating"
                    ></v-rating>
                    <!-- <v-rating
                    :value="rating"
                    ></v-rating> -->
                    <!-- <v-spacer></v-spacer> -->
                    
                </v-card-actions>
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


        <!-- <v-card-title>{{gameTitle}}</v-card-title>
        <v-container>
        <v-row>
            <v-col>
                <v-container>
                    <v-img width="200" :src="image"></v-img>            
                </v-container>
            </v-col>
            <v-col>
                <v-container class="text-left">
                    <v-card-subtitle>{{gameOverview}}</v-card-subtitle>
                    <v-rating
                        :value="rating"
                        color="amber"
                        dense
                        half-increments
                        readonly
                        size="14"
                    ></v-rating>
                    <v-card-actions>
                        <v-btn>
                            Play    
                        </v-btn>
                    </v-card-actions>
                </v-container>
            </v-col>
        </v-row>
        </v-container> -->


<script>
export default {
  name: 'GameCard',
  
  data: () => ({
		gameTitle: 'Big Blue Game',
        gameOverview: 'This is big. This is blue. And this...is a game.',
		rating: 3.0,
		image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
    }),
  
  methods: {
    
  }
}
</script>