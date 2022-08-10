<template>
    <v-container>
      <v-divider class="ma-3"></v-divider>
        <v-row justify="center">
        <v-col class="text-center">
          Summer 2022 CS 5890 Intro to DevOps &copy; 2022
        </v-col>
          <v-col class="text-center" v-for="link of links">
            <a :href="link.link" class="text-primary"> {{ link.name }}</a>
          </v-col>
          <v-col>
             <v-img src="../src/assets/UtahStateLogoBlue.jpeg" width="75" />
          </v-col>
        </v-row>
    </v-container>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Footer',
  components: {},
  data: () => ({
    links: [
      { name: "Utah State University", link: "https://www.usu.edu" },
      { name: "USU Computer Science Department", link: "https://www.usu.edu/cs/" }
    ]
  }),
  computed: {},
})
</script>