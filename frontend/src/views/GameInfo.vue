<template>
	<v-container data-test="game-data" v-if="!dataLoading && gameData !== null">
		<h1 data-test="title" class="ml-10 pb-5">{{ gameData.Name }}</h1>
		<v-row height=90vh>
			<v-col data-test="image-col" cols="2">
				<v-row
					class="fill-height"
					align-content="center"
					justify="center"
				>
					<div v-if="gameData.ImagePath !== null || gameData.ImagePath !== ''">
						<v-img					
							width=100%
							max-height="500"
							:src="gameData.ImagePath"
						></v-img>
					</div>
					<div v-else>
						<v-tooltip>
							<template v-slot:activator="{ props }">
								<v-icon x-large color="gray" v-bind="props">
									mdi-image-remove
								</v-icon>
							</template>
							<span>No image(s) available</span>
						</v-tooltip>	
					</div>
				</v-row>
			</v-col>
			<v-col data-test="data-col" cols="10">
				<v-row>
					<Rating
						:rating="gameData.Rating"
						:isReadOnly=true />
				</v-row>
				<v-row>
					<div>Description: {{ gameData.Description }}</div>
				</v-row>
			</v-col>
		</v-row>
		<v-row>
			<h3 data-test="title" class="ml-10 pb-5 pt-10">More Games Like {{ gameData.Name }}...</h3>
		</v-row>
	</v-container>
	<Loading data-test="loadbar" v-if="dataLoading" :text="$route.params.name" containerStyle="height: 75vh"/>
	<NoData v-if="!dataLoading && gameData === null" :text="$route.params.name" containerStyle="height: 75vh"></NoData>
	<Footer />
</template>

<script>
	import Rating from '../components/Rating.vue';
	import Loading from '../components/Loading.vue';
	import Footer from '../components/Footer.vue';
	import axios from "axios";
import NoData from '../components/NoData.vue';
	export default {
		name: 'GameInfo',

		data: () => ({
			dataLoading: true,
			gameData: null
		}),
		components:{
    Rating,
    Loading,
    Footer,
    NoData
},
		computed: {

		},

		methods: {
			async getGameInfo(id) {
				this.dataLoading = true;
				await axios.get(`http://127.0.0.1:8080/info/${id}`)
					.then(response => {
						this.gameData = response.data;
						this.gameData.ImagePath = "https://starwarsblog.starwars.com/wp-content/uploads/2017/05/1-star-wars-poster.jpg"
						this.dataLoading = false
					}).catch(error => {
						console.log(error.response.data);
						this.dataLoading = false
						this.gameData = null
					});
			},
			downloadGame(id){
				// Download this game with id
			}
		},
		created() {
			this.getGameInfo(this.$route.params.id);
		}
	};
</script>