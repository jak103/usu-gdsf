<template>
	<v-container>
		<v-row class="mt-4">
			<GameList></GameList>
		</v-row>
	</v-container>
</template>

<script>
import GameList from '../components/GameList.vue';
export default {
    name: "GameListView",
components: { GameList }
}
</script>