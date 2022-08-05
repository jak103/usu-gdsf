<template>
	<div
		data-test="table"
		v-if="!dataLoading && allGames.length > 0"
	>
		<h1 data-test="title" class="ml-10 pb-5">All Games</h1>
		<v-table
			data-test="data-table"
			height="80vh"
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
					<th class="text-left" width="10%">
					Tags
					</th>
					<th class="text-left" width="10%">
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
					:key="item.Id"
					:id="item.Id"
					@click.stop="handleClickGame(item.Id)">
					<td>
						<v-img v-if="item.ImagePath"
							height="50"
							:src="item.ImagePath"
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
						<div v-if="item.downloads != null">
							Downloads: {{ item.downloads }}
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
	<Loading data-test="loadbar" v-if="dataLoading" text="Loading Game Data" containerStyle="height: 75vh"/>
</template>

<script>
	import Rating from '../components/Rating.vue';
	import Loading from '../components/Loading.vue';
	import axios from "axios";
	export default {
		name: 'AllGamesPage',
		components: {
			Rating,
			Loading
		},
		data() {
			return {
				allGames: [],
				dataLoading: false,
				page: 1,
				perPage: 13
			};
		},
		methods: {
			handleClickGame(id) {
				this.$router.push("/games/info/" + id)
			},
			async getGames() {
				this.dataLoading = true;
				// we may want to configure a base-url for this, because it won't work on production
				await axios.get('http://127.0.0.1:8080/games')
					.then(response => {
						this.allGames = response.data[0];
						this.dataLoading = false
					}).catch(error => {
						console.log(error.response.data);
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