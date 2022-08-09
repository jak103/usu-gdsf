<template>
  <div style="background-color:#10243c; margin:0%">
    <v-card color="#10243c" style="color:#FFFFFF;padding: 2em;">
      <v-card-title>
        <div style="font-size:48px;text-align:center;line-height:48px;">
          About Our Project
        </div>
      </v-card-title>
      <v-card-text style="font-size:24px;margin:0%;padding:1em;text-align:center;display:flex;line-height:36px;">
        {{ words }}
      </v-card-text>
    </v-card>

    <v-card color="#10243c" style="color:#FFFFFF;padding: 2em;">
      <v-card-title>
        <div style="font-size:48px;text-align:center;line-height:48px;">
          Contributors
        </div>      
      </v-card-title>
      <v-card-text style="font-size:24px;margin:0%;padding:1em;text-align:center;display:flex;line-height:36px;">
          {{ moreWords }}
      </v-card-text>
      <v-card-text style="display:flex;justify-content: center;">
        <div>
          <div style="float:left;padding-right: 1em;">
             <li v-for="name in namesList.firstHalf" :key="name.id">
              <label>
                {{ name.firstName }} {{ name.lastName }}
              </label>
            </li>
          </div>
          <div style="float:right;padding-left: 1em;">
             <li v-for="name in namesList.secondHalf" :key="name.id">
              <label>
                {{ name.firstName }} {{ name.lastName }}
              </label>
            </li>
          </div>
        </div>
      </v-card-text>
    </v-card>
  </div>
  <Footer></Footer>
</template>

<script>
import names from './names.js';
import stringsToObjects from './nameFunctions.js';
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
      let list = stringsToObjects(names.names);
      let index = Math.ceil(list.length / 2);
      let first = list.splice(0, index);
      let second = list.splice(-index);
      return {
        firstHalf: first,
        secondHalf: second
      }
    }
  }
}
</script>