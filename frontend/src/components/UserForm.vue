<!-- This component operates as a model to edit users -->
<template>
  <v-row justify="center">
    <v-dialog
      persistent
      v-model="showSelf"
    >
      <v-card>
        <v-card-title class="text-center">
          <span v-if="isAdmin" class="text-h5">Create Admin Account</span>
          <span v-else class="text-h5">Edit User Profile</span>
          <span></span>
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
                  label="Birth Date*"
                  type="date"
                  v-model="selectedUser.dob"
                  required
                ></v-text-field>
              </v-col>
              <v-col 
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
              >
                <v-text-field
                  label="Email*"
                  type="email"
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
            v-if="!isAdmin"
            color="error"
            flat
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
            v-if="!isAdmin"
            color="secondary"
            flat
            @click="$emit('save', selectedUser)">
            Save
          </v-btn>

            <v-btn
            v-if="isAdmin"
            color="secondary"
            text
            @click="$emit('createAdmin', selectedUser)">
            Create
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
  
</template>

<script setup>
  defineProps({
    showSelf: Boolean,
    isAdmin: Boolean,
    selectedUser: {
      firstName: String,
      lastName: String,
      email: String,
    }
  });

</script>