<template>
  <v-container>
    <v-card color="#10243c" style="color:#FFFFFF;margin-bottom: 1em;">
      <v-card-title style="font-size:2em">
        About our Project
      </v-card-title>
      <v-card-text style="font-size:1em">
        {{ words }}
      </v-card-text>
    </v-card>

    <v-card color="#10243c" style="color:#FFFFFF;margin-top: 1em;">
      <v-card-title style="font-size:2em">
        Contributors
      </v-card-title>
      <v-card-text style="font-size:1em">
        {{ moreWords }}
      </v-card-text>
      <v-card-text>
        <li v-for="name in namesList" :key="name.id">
          <label>
            {{ name.firstName }} {{ name.lastName }}
          </label>
        </li>
      </v-card-text>
    </v-card>
  </v-container>
  <Footer></Footer>
</template>

<script>
import names from './names.js'
import stringsToObjects from './nameFunctions.js'
import Footer from "../components/Footer.vue";


export default {
  name: 'AboutView',
  components: {
    Footer,
  },

  data: () => ({
    words: 'This application was created as a project for the Intro to DevOps and Cloud Based Services course at USU. Our Mission is to deliver and showcase student projects created for the CS 5410 - Game Development course at USU. Our Project works by using Docker and Google Cloud to deploy the application. Firestore is used for our database.',
    moreWords: 'The Game Development Store Front is a continuous application developed by students and maintained by professors. Many have contributed and many more will contribute to this project. Below is a list of all of the students who have helped make this project what it is today.',
    listOfNames: [{id: 1, firstName: 'firstName1',  lastName: 'lastName1'}, {id: 2, firstName: 'firstName2',  lastName: 'lastName2'}],
  }),

  computed: {
    namesList() {
      return stringsToObjects(names.names);
    }
  }
}
</script>