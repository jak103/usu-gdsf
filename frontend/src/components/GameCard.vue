<template>
  <v-card
    class="mx-4 my-4"
    max-width="600px"
  >
    <v-img 
      :src="'data:image/png;base64,'+ this.image"
    ></v-img>

    <v-card-title>
      {{title}}
    </v-card-title>

    <v-card-subtitle>
      {{author}}
    </v-card-subtitle>

    <v-card-actions>
      <v-btn
        :disabled="!notDownloading"
        color="secondary"
        text
        @click="downloadGame"
      >Download</v-btn>

      <v-spacer></v-spacer>

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
  import axios from 'axios';

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
      image: "",
      notDownloading: true
    }),

    mounted() {
      axios.get("http://localhost:8080/game/download?downloadType=screenshots&bucket=breakout")
        .then(res => {
          res.data[0].forEach(element => {
            this.image = element.ObjectData
            return;
          })
        })
    },

    methods: {
      sleep: function(ms) {
        return new Promise(r => setTimeout(r, ms))
      },

      downloadGame: async function () {
        if (this.notDownloading) {
          this.notDownloading = false
          await axios.get("http://localhost:8080/game/download?downloadType=game&bucket=breakout", { responseType: "blob" })
            .then(res => {
              const blob = new Blob([res.data])
              const link = document.createElement("a")
              link.href = URL.createObjectURL(blob)
              link.download = "test-the-breakout-game.zip"
              link.click()
              URL.revokeObjectURL(link.href)
            })
            .catch(err => {
              console.error(err);
            })
        }
        this.notDownloading = true;
      },
    }
  }
</script>