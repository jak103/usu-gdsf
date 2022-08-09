<!-- 
  This component operates as an admin tool to edit users and create admin accounts.
  It could easily be used to update accounts from the /accounts page, just make sure isAdmin
  is set to false.
 -->
<template>
  <v-form>
    <v-row justify="center">
      <v-dialog
        persistent
        v-model="showSelf"
        max-width="500px"
        @afterLeave="$emit('leftDialog')"
      >
        <v-card>
          <v-card-title class="text-center">
            <span v-if="isAdminCreation" class="text-h5">Create Admin Account</span>
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
                    id="firstNameField"
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
                    id="lastNameField"
                    label="Last Name*"
                    v-model="selectedUser.lastName"
                    persistent-hint 
                    :rules="[rules.counter, rules.required]"
                    required
                  ></v-text-field>
                </v-col>

                <v-col 
                  style="padding-top: 0; padding-bottom: 0; margin-bottom: 0;"
                  cols="12"
                >
                  <v-text-field
                    id="birthDateField"
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
                    id="emailField"
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
                    v-model="newPassword"
                    :rules="[rules.password]"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <div style="color: red; margin-left: 40px;">{{errorMsg}}</div>
          <v-card-actions>
            <v-btn
              id="deleteUserBtn"
              v-if="!isAdminCreation"
              color="error"
              @click="handleDelete">
              Delete User
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              id="closeBtn"
              color="secondary"
              @click="handleClose">
              Close
            </v-btn>
            <v-btn
              id="saveEditBtn"
              type="submit"
              v-if="!isAdminCreation"
              color="secondary"
              @click="submitUpdate">
              Save
            </v-btn>

              <v-btn
              id="adminCreateBtn"
              v-if="isAdminCreation"
              color="secondary"
              @click="submitCreation">
              Create
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </v-form>
  
</template>

<script>
  export default {
    props: {
        showSelf: Boolean,
        isAdminCreation: Boolean,
        selectedUser: {
          firstName: String,
          lastName: String,
          email: String,
          dob: String,
        },
        newPassword: String
      },

      methods: {

        validate() {
          this.errorMsg = ''
          if (!this.selectedUser.firstName || !this.selectedUser.lastName) {
            this.errorMsg = "Name fields cannot be blank"
            return false
          }
          if (!this.selectedUser.email) {
            this.errorMsg = "Email field cannot be blank"
            return false
          }

          if (!this.selectedUser.dob) {
            this.errorMsg = "Birth Date field cannot be blank"
            return false
          }

          if (this.selectedUser.firstName.length > 30 || this.selectedUser.lastName.length > 30) {
            return false
          }

          if (!this.emailPattern.test(this.selectedUser.email)) {
            return false
          }

          if (this.newPassword != '' && this.newPassword.length < 12) {
            return false
          }

          this.errorMsg = ""
          return true
        },

        submitUpdate() {
          if (this.validate()) {
            this.$emit('save', this.selectedUser)
          }
        },

        submitCreation() {
          if (this.validate()) {
            if (!this.newPassword) {
              this.errorMsg = "Password field cannot be blank"
            } else {
              this.$emit('createAdmin', this.selectedUser, this.newPassword)
            }
          }

        },

        handleClose() {
          this.errorMsg = ''
          this.$emit('close')
        },

        handleDelete() {
          this.errorMsg = ''
          this.$emit('delete', this.selectedUser)
        }
      },

      data() {
        return {
          emailPattern : /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
          // These rules are only for visual aspect in the form right now (:rules), actual validation happens in methods
          rules : {
            required: value => !!value || '',
            counter: value => value.length <= 20 || 'Max 30 characters',
            email: value => {
              return this.emailPattern.test(value) || 'Invalid e-mail.'
            },
            password: value => {
              if (this.isAdminCreation) {
               return value.length >= 12 || 'Password must be at least 12 characters'
              } else {
                return !!value
              }
            }
          },
          errorMsg : ''
        }
      }
  }

</script>