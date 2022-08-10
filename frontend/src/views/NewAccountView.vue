<template>
  <h2>Create New User</h2>
	<v-form
    ref="form"
    v-model="valid"
    lazy-validation
  	>
    <v-text-field
      v-model="username"
      label="Enter Username"
      required
    ></v-text-field>

    <v-text-field
      v-model="displayName"
      label="Enter Display Name"
      required
    ></v-text-field>

    <v-text-field
      v-model="userEmail"
      label="Enter E-mail"
      :rules="emailRules"
      required
    ></v-text-field>

    <v-text-field
      v-model="userPassword"
      :type="'password'"
      :rules="passwordRules"
      name="input-10-1"
      label="Enter Password"
      counter
    ></v-text-field>

    <v-select
      v-model="userRole"
      :items="items"
      :rules="[v => !!v || 'Role is required']"
      label="Role"
      solo
    ></v-select>

    <v-btn
      color="success"
      class="mr-4"
      v-on:click="createUser"
    >
      Create User
    </v-btn>
  </v-form>
</template>

<script>
import router from '../router'

export default {
	name: 'NewAccount',

  data: () => ({
    items: ['Admin', 'Publisher', 'Basic'],
    username: '',
    usernameRules: [
      value => !!value || 'Required.',
      value => (value || '').length <= 20 || 'Max 20 characters',
    ],
    userEmail: '',
    emailRules: [
        value => !!value || 'E-mail is required',
        value => {
          const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Invalid e-mail.'
      }
    ],
    displayName: '',
    userPassword: '',
    passwordRules: [
        value => !!value || 'Password is required',
        value => (value || '').length > 7 || 'At least 8 characters required',
    ],
    userRole: 'Basic'
  }),

  methods: {
    createUser() {
      // TODO save this to the model
      console.log("USERNAME ", this.username)
      console.log("DISPLAY NAME ", this.displayName)
      console.log("USER EMAIL ", this.userEmail)
      console.log("USER PASSWORD ", this.userPassword)
      console.log("ROLE ", this.userRole)
      // navigate to the home page
      // TODO push params
      router.push({path: '/'})
    }
  }
}
</script>