<template>
  <v-card width="100%">
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
  </v-card>
</template>

<script>
export default {
  name: 'GameCarousel',

  data: () => ({
    colors: ['primary','secondary','info','warning','white','teal', 'watermelon'],
    cycleTime: 4000,
  }),
}
</script>