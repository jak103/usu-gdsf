<template>
	<div>
		<v-table>
			<thead>
				<tr>
					<th width="100px"></th> <!--Image-->
					<th class="text-left" width="15%">
					Name
					</th>
					<th class="text-left" width="15%">
					Developer
					</th>
					<th class="text-left" width="15%">
					Tags
					</th>
					<th class="text-left" width="10%">
					Times Played/Downloaded
					</th>
					<th class="text-left" width="15%">
					Rating
					</th>
					<th></th><!--Tooltip info with description-->
				</tr>
			</thead>
			<tbody>
				<tr	link v-for="item in games" :key="item.id" @click.stop="handleClickGame(item.id)">
					<td>
						<v-img
							height="75"
							:src="item.imagePath"
						></v-img>
					</td>
					<td>{{ item.name }}</td>
					<td>{{ item.developer }}</td>
					<td>
						<v-chip v-for="tag in item.tags">
							{{ tag }}
						</v-chip>
					</td>
					<td>{{ item.timesPlayed }}</td>
					<td>
						<v-tooltip>
							<template v-slot:activator="{ props }">
								<v-rating
									v-bind="props"
									:model-value="item.rating"
									density="compact"
									half-increments
								></v-rating>
							</template>
							<span>{{ item.rating }} Stars</span>
						</v-tooltip>	
					</td>
					<td>
						<v-tooltip>
							<template v-slot:activator="{ props }">
								<v-icon v-bind="props">
									mdi-information
								</v-icon>
							</template>
							<span>{{ item.description }}</span>
						</v-tooltip>
					</td>
				</tr>
			</tbody>
		</v-table>
	</div>
</template>

<script>
import Game from "../models/game.js"
import Rating from '../components/Rating.vue';
export default {
	name: 'AllGamesPage',
	components: {
		Rating
	},
	data: () => ({
		games: [new Game(), new Game(), new Game()]
	}),

	computed: {

	},

	methods: {
		handleClickGame(id) {
			this.$router.push("/games/info/" + id)
		}
	}

}
</script>