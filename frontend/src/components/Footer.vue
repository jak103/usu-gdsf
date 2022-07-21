<template>
  <v-footer color="teal" app>
    <v-container>
      <v-row>
        <v-col> About </v-col>
      </v-row>
      <v-row>
        <v-col> <a href="https://usu.edu"> Utah State University</a> </v-col>
      </v-row>
      <v-row>
        <v-col> <a href="https://www.usu.edu/cs/"> USU CS Department</a></v-col>
      </v-row>
      <v-row>
        <v-col>
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
  data: () => ({}),
  computed: {},
})
</script>