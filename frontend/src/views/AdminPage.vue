<template>
	<v-tabs-item v-model="tab">
    <v-tab-item>
      <GameCard></GameCard>
    </v-tab-item>
  </v-tabs-item>
</template>

<script>
import LoginForm from '../components/LoginForm.vue';
import GameCard from '../components/GameCard.vue';
export default {
    name: "AdminPage",
    components: { LoginForm, GameCard }
}
</script>