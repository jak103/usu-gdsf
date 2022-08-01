<!-- This component operates as a model to edit users -->
<template>
  <v-row justify="center">
    <v-dialog
      persistent
      v-model="showSelf"
    >
      <v-card>
        <v-card-title class="text-center">
          <span class="text-h5">Edit User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
                sm="4"
                md="6"
              >
                <v-text-field
                  label="First Name*"
                  v-model="selectedUser.firstName"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
                sm="4"
                md="6"
              >
                <v-text-field
                  label="Last Name*"
                  v-model="selectedUser.lastName"
                  persistent-hint
                  required
                ></v-text-field>
              </v-col>
              <v-col 
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Email*"
                  v-model="selectedUser.email"
                  required
                ></v-text-field>
              </v-col>
              <v-col 
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Password*"
                  type="password"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-btn
            color="error"
            text
            @click="$emit('delete', selectedUser)">
            Delete User
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn
            color="secondary"
            text
            @click="$emit('close')">
            Close
          </v-btn>
          <v-btn
            color="secondary"
            text
            @click="$emit('save', selectedUser)">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
  
</template>

<script setup>
  defineProps({
    showSelf: Boolean,
    selectedUser: {
      firstName: String,
      lastName: String,
      email: String,
    }
  });

</script>