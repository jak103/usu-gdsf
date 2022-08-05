<template>
<v-container>
<v-card>
<v-card-title>
  <p class="text-h3 text--primary">About our Project</p>
</v-card-title>
<v-card-text>
  This is Team 2's CS 5890 project. This project is to help the display and allow consumers access to games built by students
  in the Game Development Class at Utah State University. It is a collaboration of many different students with various tasks 
  both on the front and back end. The project is built using Vue and Go lead by Gordon Fjeldsted and Bridget Staley.
  </v-card-text>
  </v-card>

  <v-card
  elevation="9"
  loading
  shaped
  tile
  class="d-flex mt-3 mb-3">
  <v-card-text> This project was made possible by Jacob Christensen's CS 5890 class and by contributions to Utah State University 
  from gamers like you. Thank You.
</v-card-text>
</v-card>

<v-card>
  <v-card-text>
    <p class="text-h4 text--primary">
      Contributors
    </p>
    <div class="text--primary">
      Who are the amazing people that helped create this site? Click below to find out!
    </div>
  </v-card-text>
  <v-card-actions>
      <v-btn
        color="watermelon"
        text
      >
        View All Contributors
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
        <p>Marshal Taylor <br/>
        Rohit Mishri <br/>
        Jacob White <br/>
        Gordon Fjeldsted <br/>
        Bridget Staley <br/>
        Jared Horton <br/>
        Spencer Robinson <br/>
        Elise Francis <br/>
        Brendon Earl <br/>
        Kailey Bales <br/>
        Richard Snider <br/>
        Bryant Saunders <br/>
        Matthew Shaw <br/>
        Logan Ballard <br/>
        Abbey Allen <br/>
        Rebecca Macias <br/>
        Adam Thomas <br/>
        Jenna Cioffi <br/>
        Brittany Graham <br/>
        Marshal Taylor <br/>
        Ian McLachlan</p>
        <p class="font-italic">**Some contributors missing</p>
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>

</v-container>
</template>

<script>

export default {
name: 'AboutView',

  data: () => ({
    show: false,
    moreWords: 'This project was made in part by Team 2. And gamers like you. Thank you'
  })

}
</script>
