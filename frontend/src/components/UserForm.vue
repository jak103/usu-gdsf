<!-- 
  This component operates as an admin tool to edit users and create admin accounts.
  It could easily be used to update accounts from the /accounts page, just make sure isAdmin
  is set to false.
 -->
<template>
  <v-row justify="center">
    <v-dialog
      persistent
      v-model="showSelf"
      @afterLeave="$emit('leftDialog')"
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
                sm="6"
                md="6"
              >
                <v-text-field
                  label="First Name*"
                  v-model="selectedUser.firstName"
                  :rules="[rules.counter, rules.required]"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
                sm="6"
                md="6"
              >
                <v-text-field
                  label="Last Name*"
                  v-model="selectedUser.lastName"
                  persistent-hint 
                  :ruels="[rules.counter, rules.required]"
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
                  :rules="[rules.required]"
                  required
                ></v-text-field>
              </v-col>
              <v-col 
                style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                cols="12"
                class="overflow-hidden"
              >
                <v-text-field
                  label="Email*"
                  type="email"
                  :rules="[rules.email, rules.required]"
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
                  :rules="[rules.password]"
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
            @click="$emit('delete', selectedUser)">
            Delete User
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn
            color="secondary"
            @click="$emit('close')">
            Close
          </v-btn>
          <v-btn
            v-if="!isAdmin"
            color="secondary"
            @click="$emit('save', selectedUser)">
            Save
          </v-btn>

            <v-btn
            v-if="isAdmin"
            color="secondary"
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

  let rules = {
    // Password only required if admin creation
    // DOB rules?
    required: value => !!value || '',
    counter: value => value.length <= 20 || 'Max 20 characters',
    email: value => {
      const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
      return pattern.test(value) || 'Invalid e-mail.'
    },
    password: value => {
      if (isAdmin) {
        return value.length <= 20 && value.length >= 12
      }
    }
  }
</script>