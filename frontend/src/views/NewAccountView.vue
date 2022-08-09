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
      required
    ></v-text-field>

    <v-text-field
      v-model="userPassword"
      :type="'password'"
      name="input-10-1"
      label="Enter Password"
      hint="At least 8 characters"
      counter
    ></v-text-field>

    <v-select
      :items="items"
      label="Role"
      solo
    ></v-select>

    <v-btn
      :disabled="!valid"
      color="success"
      class="mr-4"
      @click="validate"
    >
      Create User
    </v-btn>
  </v-form>
</template>

<script>
export default {
	name: 'NewAccount',
  data: () => ({
    items: ['Admin', 'Publisher', 'Basic'],
  })
}
</script>