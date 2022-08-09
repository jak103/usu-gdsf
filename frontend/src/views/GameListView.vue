<template>
	<div>
		<v-row class="mt-4">
			<GameList></GameList>
		</v-row>
	</div>
</template>

<script>
import GameList from '../components/GameList.vue';
export default {
    name: "GameListView",
components: { GameList }
}
</script>