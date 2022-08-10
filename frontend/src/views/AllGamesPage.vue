<template>
	<div
		data-test="table"
		v-if="!dataLoading && allGames && allGames.length > 0"
	>
		<div data-test="title" class="ml-10 pb-3 text-h4">
			All Games
		</div>
		<v-table
			data-test="data-table"
			height="70vh"
		>
			<thead>
				<tr>
					<th width="100px" ></th> <!--Image-->
					<th class="text-left" width="10%">
					Name
					</th>
					<th class="text-left" width="12%">
					Developer
					</th>
					<th class="text-left" width="5%">
					Version
					</th>
					<th class="text-left" width="7%">
					Created
					</th>
					<th class="text-left" width="20%">
					Tags
					</th>
					<th class="text-left" width="15%">
					Times Played/Downloads
					</th>
					<th class="text-left" width="10%">
					Rating
					</th>
					<th></th><!--Tooltip info with description-->
				</tr>
			</thead>
			<tbody>
				<tr	
					v-for="item in allGames.slice(((page - 1) * perPage), ((page - 1) * perPage) + perPage)"
					class="pt-2"
					:key="item.Id"
					:id="item.Id"
					style="cursor: pointer"	
					@click.stop="handleClickGame(item.Id, item.Name)">
					<td>
						<v-img v-if="item.ImagePath"
							height="50"
							:src="item.ImagePath"
							@error="item.ImagePath = null"
						></v-img>
						<v-icon v-else color="gray">
							mdi-gamepad-square
						</v-icon>
					</td>
					<td>{{ item.Name }}</td>
					<td>{{ item.Developer }}</td>
					<td>v{{ item.Version }}</td>
					<td>{{ getDateString(new Date(item.CreationDate)) }}</td>
					<td>
						<div v-if="!item.Tags">-</div>
						<v-chip v-if="item.Tags" v-for="tag in item.Tags">
							{{ tag }}
						</v-chip>
					</td>
					<td>
						<div v-if="item.TimesPlayed != null">
							Times Played: {{ item.TimesPlayed }}
						</div>
						<div v-if="item.Downloads != null">
							Downloads: {{ item.Downloads }}
						</div>
						<div v-if="item.TimesPlayed == null && item.downloads == null">
							-
						</div>
					</td>
					<td>
						<v-tooltip>
							<template v-slot:activator="{ props }">
								<v-rating
									v-bind="props"
									:model-value="item.Rating"
									density="compact"
									half-increments
								></v-rating>
							</template>
							<span>{{ item.Rating ? item.Rating : 0 }} Stars</span>
						</v-tooltip>	
					</td>
					<td>
						<v-tooltip>
							<template v-slot:activator="{ props }">
								<v-icon v-bind="props">
									mdi-information
								</v-icon>
							</template>
							<span>{{ item.Description ? item.Description : "No description available" }}</span>
						</v-tooltip>
					</td>
				</tr>
			</tbody>
		</v-table>
		<v-pagination
			data-test="pagination" 
			v-model="page"
			:length="Math.ceil(allGames.length / perPage)"
		></v-pagination>
	</div>

	<Loading data-test="loadbar" v-if="dataLoading" text="Game Data" containerStyle="height: 75vh"/>
	<NoData data-test="noData" v-if="!dataLoading && (!allGames || allGames?.length === 0)" text="All Games" containerStyle="height: 75vh"/>
	<Footer />
</template>

<script>
	import Rating from '../components/Rating.vue';
	import Loading from '../components/Loading.vue';
	import Footer from '../components/Footer.vue';
	import NoData from '../components/NoData.vue';

	import * as GamesServices from '../services/gamesServices';
	import Game from '../models/game';
	export default {
		name: 'AllGamesPage',
		components: {
			Rating,
			Loading,
			Footer,
			NoData
		},
		data() {
			return {
				allGames: [Game],
				dataLoading: false,
				page: 1,
				perPage: 13
			};
		},
		methods: {
			handleClickGame(id, name) {
				this.$router.push(`/games/info/${name}/${id}`)
			},
			async getGames() {
				this.dataLoading = true;
				await GamesServices.getAllGames()
					.then(response => {
						this.allGames = response.data
						this.dataLoading = false
					}).catch(error => {
						console.log(error);
						this.allGames = [];
						this.dataLoading = false
					});
			},
			getDateString(date){
				let year = date.getFullYear()
				let month = date.getMonth()
				let day = date.getDate()

				return `${month}/${day}/${year}`
			}
		},
		created() {
			this.getGames();
		}
	}
</script>
