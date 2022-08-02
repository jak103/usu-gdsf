<template>
  <v-card>
    <div color="primary">
      <v-card-title>Login Form</v-card-title>
    </div>

    <v-card-text>
      <v-form>
        <v-text-field 
            label="username" 
            v-model="username" 
            required 
            :rules="[v => !!v || 'Username is required']">
        </v-text-field>
        <v-text-field 
            label="password" 
            v-model="password" 
            type="password" 
            required 
            :rules="[v => !!v || 'Password is required']">
        </v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn color="teal" elevation="7"> Login </v-btn>
    </v-card-actions>
  </v-card>
</template>
<script>
export default {
  name: "LoginForm",
  props: {
    source: String,
  },
};
</script>

<style></style>
