<template>
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
</script>