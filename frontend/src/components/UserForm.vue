<!-- This component operates as a model to edit users -->
<template>
  <v-row justify="center">
    <v-dialog
      persistent
      v-model="showSelf"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                cols="12"
                sm="4"
                md="6"
              >
                <v-text-field
                  label="First Name*"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                cols="12"
                sm="4"
                md="6"
              >
                <v-text-field
                  label="Last Name*"
                  persistent-hint
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="Email*"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="Password*"
                  type="password"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue-darken-1"
            text
            @click="$emit('close')">
            Close
          </v-btn>
          <v-btn
            color="blue-darken-1"
            text
            @click="$emit('save')">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script setup>
  defineProps({
    showSelf: Boolean
  });
</script>