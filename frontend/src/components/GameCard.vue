<template>
<<<<<<< HEAD
  <v-card
    class="mx-4 my-4"
    max-width="600px"
  >
    <v-img
      :src="image"
    ></v-img>

    <v-card-title>
      {{title}}
    </v-card-title>

    <v-card-subtitle>
      {{author}}
    </v-card-subtitle>

    <v-card-actions>
      <v-btn
        color="secondary"
        text
      >
        Description
      </v-btn>

      <v-spacer></v-spacer>

      <v-btn
        icon
        @click="show = !show"
      >
        <v-icon>{{ show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <div v-show="show">
        <v-divider></v-divider>

        <v-card-text>
            {{description}}
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>
</template>

<script>
  export default {
    name: 'GameCard',

    props: {
        title: String,
        author: String,
        description: String,
        image: String,
    },

    data: () => ({
      show: false,
    }),
  }
=======
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
>>>>>>> team2-main
</script>