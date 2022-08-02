<template>
	<v-tabs-item v-model="tab">
    <v-tab-item>
      <LoginForm></LoginForm>
    </v-tab-item>
  </v-tabs-item>
</template>

<script>
import LoginForm from '../components/LoginForm.vue';
export default {
    name: "AdminPage",
    components: { LoginForm }
}
</script>