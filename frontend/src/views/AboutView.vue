<template>
  <v-container>
    <v-card>
      <v-card-title>
        About our Project
      </v-card-title>
      <v-card-text>
        This is how our project works and what it is doing.
      </v-card-text>
    </v-card>

    <v-card>
      <v-card-title>
        Just text
      </v-card-title>
      <v-card-text>
        {{ moreWords }}
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: 'AboutView',

  data: () => ({
    moreWords: 'This project was made in part by Team 2. And gamers like you. Thank you'
  })
}
</script>