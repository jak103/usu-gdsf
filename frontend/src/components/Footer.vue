<template>
  <v-footer color="teal" app>
    <v-container>
        <v-row justify="center">
        <v-btn v-for="link of links" :href="link.link" color="teal"> {{ link.name }}</v-btn>
        <v-col cols="12" class="text-center">
          Created by Summer 2022 CS 5890 Intro to DevOps Course &copy; 2022
        </v-col>
        </v-row>
    </v-container>
  </v-footer>
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