<template>
  <v-container>
    <v-card>
      <v-card-title>
        About our Project
      </v-card-title>
      <v-card-text>
        {{ words }}
      </v-card-text>
    </v-card>

    <v-card>
      <v-card-title>
        Just text
      </v-card-title>
      <v-card-text>
        {{ moreWords }}
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: 'AboutView',

  data: () => ({
    words: 'This application was created as a project for the Intro to DevOps and Cloud Based Services course at USU. Our Mission is to deliver and showcase student projects created for the CS 5410 - Game Development course at USU. Our Project works by using Docker and Google Cloud to deploy the application. Firestore is used for our database.',
    moreWords: 'The Game Development Store Front is a continuous application developed by students and maintained by professors. Many have contributed and many more will contribute to this project. Below is a list of all of the students who have helped make this project what it is today.',
    listOfNames: [{firstName: 'firstName1',  lastName: 'lastName1'}, {firstName: 'firstName2',  lastName: 'lastName2'}]
  })
}
</script>