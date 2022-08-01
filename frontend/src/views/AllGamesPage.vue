<template>
	<div>
		<v-table>
			<thead>
				<tr>
					<th></th> <!--Image-->
					<th class="text-left">
					Name
					</th>
					<th class="text-left">
					Developer
					</th>
					<th class="text-left">
					Tags
					</th>
					<th class="text-left">
					Times Played/Downloaded
					</th>
					<th class="text-left">
					Rating
					</th>
					<th></th><!--Tooltip info with description-->
				</tr>
			</thead>
			<tbody>
				<tr	v-for="item in games" :key="item.id">
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
export default {
	name: 'AllGamesPage',

	data: () => ({
		games: [new Game(), new Game(), new Game()]
	}),

	computed: {

	},

	methods: {

	}

}
</script>