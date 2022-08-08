<template>
  <v-row justify="center">
    <v-dialog
      persistent
      v-model="showCreate"
    >
      <v-card>
        <v-card-title class="text-center">
          <span class="text-h5">Create new game</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Game Title*"
                  v-model="game.name"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Developer*"
                  v-model="game.developer"
                  required
                ></v-text-field>
              </v-col>

              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Version*"
                  type="gameVersion"
                  v-model="game.version"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Description*"
                  type="description"
                  v-model="game.description"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="ImagePath*"
                  type="imagePath"
                  v-model="game.imagePath"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="DownloadLink*"
                  type="downloadLink"
                  v-model="game.downloadLink"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-btn
            color="error"
            flat
            @click="$emit('cancel')">
            Cancel
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn
            color="secondary"
            text
            @click="$emit('save', game)">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
  
</template>

<script setup>
  defineProps({
    showCreate: Boolean,
    game: {
      name: String,
      developer: String,
      version: String,
      description: String,
      imagePath: String,
	  downloadLink: String
    }
  })
</script>

<style>

</style>