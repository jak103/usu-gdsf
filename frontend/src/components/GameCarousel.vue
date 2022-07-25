<template>
  <v-card width="100%">
    <v-card-body>
      <v-carousel cycle :interval="cycleTime">
        <v-carousel-item v-for="color in colors" :key="color">
          <v-sheet :color="color" height="100%" tile>
            <v-row class="fill-height" align="center" justify="center">
              <div class="text-h2">
                {{ color }}
              </div>
            </v-row>
          </v-sheet>
        </v-carousel-item>
      </v-carousel>
    </v-card-body>
    <!-- <v-card-text>
      This is where the vuetify carousel will be implemented.
    </v-card-text> -->
  </v-card>
</template>

<script>
export default {
  name: 'GameCarousel',

  data: () => ({
    colors: ['red','orange','yellow','green','blue','indigo', '#7f00ff'],
    cycleTime: 4000,
  }),
}
</script>