<template>
	<div v-if="!dataLoading && allGames.length > 0">
		<v-table 
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
					<th class="text-left" width="10%">
					Created
					</th>
					<th class="text-left" width="10%">
					Tags
					</th>
					<th class="text-left" width="10%">
					Downloads
					</th>
					<th class="text-left" width="15%">
					Rating
					</th>
					<th></th><!--Tooltip info with description-->
				</tr>
			</thead>
			<tbody>
				<tr	link v-for="item in allGames.slice(((page - 1) * perPage), ((page - 1) * perPage) + perPage)" :key="item.id" @click.stop="handleClickGame(item.id)">
					<td>
						<v-img
							height="50"
							:src="item.imagePath"
						></v-img>
					</td>
					<td>{{ item.Name }}</td>
					<td>{{ item.Author }}</td>
					<td>v{{ item.Version }}</td>
					<td>{{ getDateString(new Date(item.CreationDate)) }}</td>
					<td>
						<v-chip v-for="tag in item.Tags">
							{{ tag }}
						</v-chip>
					</td>
					<td>{{ item.downloads }}</td>
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
							<span>{{ item.rating ? item.rating : 0 }} Stars</span>
						</v-tooltip>	
					</td>
					<td>
						<v-tooltip>
							<template v-slot:activator="{ props }">
								<v-icon v-bind="props">
									mdi-information
								</v-icon>
							</template>
							<span>{{ item.description ? item.description : "No description available" }}</span>
						</v-tooltip>
					</td>
				</tr>
			</tbody>
		</v-table>
		<v-pagination
			v-model="page"
			:length="Math.ceil(allGames.length / perPage)"
		></v-pagination>
	</div>
	<Loading v-if="dataLoading" text="Loading Game Data" containerStyle="height: 75vh"/>
</template>

<script>
	import Rating from '../components/Rating.vue';
	import Loading from '../components/Loading.vue';
	import TableFilter from '../components/TableFilter.vue';
	import axios from "axios";
	export default {
		name: 'AllGamesPage',
		components: {
			Rating,
			Loading,
			TableFilter
		},
		data() {
			return {
				allGames: {},
				dataLoading: false,
				page: 1,
				perPage: 13
			};
		},

		computed: {

		},

		methods: {
			handleClickGame(id) {
				this.$router.push("/games/info/" + id)
			},
			getGames() {
				this.dataLoading = true;
				// we may want to configure a base-url for this, because it won't work on production
				axios.get('http://127.0.0.1:8080/games')
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