<template>
  <v-container>
    <h1>Edit User</h1>
    <v-form
      ref="form"
      v-model="valid"
      lazy-validation
    >
      <v-text-field
        v-model="name"
        :counter="10"
        :rules="nameRules"
        label="Name"
        required
      ></v-text-field>

      <v-text-field
        v-model="username"
        :rules="usernameRules"
        label="Username"
        required
      ></v-text-field>

      <v-text-field
        v-model="email"
        :rules="emailRules"
        label="E-mail"
        required
      ></v-text-field>

      <v-btn
        :disabled="!valid"
        color="success"
        class="mr-4"
        @click="validate"
      >
        Save Changes
      </v-btn>
    </v-form>
  </v-container>
</template>

<script>
  export default {
    name: 'EditUser',
    data: () => ({
      valid: true,
      name: '',
      nameRules: [
        v => !!v || 'Name is required',
      ],
      email: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      usernameRules: [
        v => !!v || 'Username is required',
        v => !/\s/g.test(v) || 'Username cannot contain spaces',
      ],
      select: null,
      checkbox: false,
      name: '',
      username: '',
      email: ''
    }),

    methods: {
      validate () {
        this.$refs.form.validate()
        //Post user changes to database here
      },
    },

    watch: {
      '$route.params': {
        handler: function(params) {
          this.username = params.username
          this.email = params.email
          this.name = params.name
        },
        deep: true,
        immediate: true
      }
    }
  }
</script>