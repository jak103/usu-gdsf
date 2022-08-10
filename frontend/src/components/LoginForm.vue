<template>
  <v-card>
    <div>
      <v-card-title class="bg-primary">Login Form</v-card-title>
    </div>

    <v-card-text>
      <v-form v-model="isValid">
        <v-text-field
            label="username"
            v-model="username"
            required >
        </v-text-field>
        <v-text-field
            label="password"
            v-model="password"
            type="password"
            required >
        </v-text-field>
      </v-form>
    </v-card-text>
    <router-link :to="{ name: 'newAccount', params: { type: 'admin' }}">
      Need an account? Create one here
    </router-link>
    <v-card-actions>
      <v-btn to="/admin" class="bg-teal" elevation="7"> Login </v-btn>
    </v-card-actions>
  </v-card>
</template>
<script>
export default {
  name: "LoginForm",
  props: {
    source: String,
  },
  data: () => ({
    username: null,
    password: null,
  }),

  methods: {
    async save() {
      await this.$axios.post({
        url: '/register',
        data: {
          username: this.username,
          password: this.password
        }
      })
    }

  }
};
</script>

<style></style>
