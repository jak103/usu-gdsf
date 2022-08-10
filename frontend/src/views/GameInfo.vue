<template>
	<v-container data-test="game-data" v-if="!dataLoading && gameData !== null">
		<v-row>
			<div data-test="title" class="ml-10 pb-3 text-h4">
				{{ gameData.Name }}
			</div>
		</v-row>
		<v-row>
			<div data-test="developer" class="ml-10 pb-3 text-subtitle-1">
				{{ gameData.Developer }}
			</div>
		</v-row>
		<v-row height=90vh>
			<v-col data-test="image-col" cols="2">
				<v-row
					class="fill-height"
					align-content="center"
					justify="center"
				>
					<div v-if="gameData.ImagePath !== null && gameData.ImagePath !== ''">
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
			<v-col data-test="data-col" cols="6">
				<v-row>
					<v-sheet class="pa-2 ma-2">
						<div><b>Version:</b> v{{ gameData.Version }}</div>
					</v-sheet>
				</v-row>
				<v-row>
					<v-sheet class="pa-2 ma-2">
						<div><b>Description:</b> {{ gameData.Description }}</div>
					</v-sheet>
				</v-row>
				<v-row v-if="gameData.Tags && gameData.Tags.length > 0">
					<v-sheet class="pa-2 ma-2">
						<span class="pr-2"><b>Tags:</b></span>
						<v-chip v-for="tag in gameData.Tags">
							{{ tag }}
						</v-chip>
					</v-sheet>
				</v-row>
			</v-col>
			<v-col cols="2">
				<v-row>
					<v-sheet class="pa-2 ma-2">
						<Rating
							:rating="gameData.Rating"
							:isDense=true
							:isReadOnly=true
							:isHalfIncrements="true"/>
					</v-sheet>
				</v-row>
				<v-row>
					<v-sheet class="pa-2 ma-2">
						<b>Created on:</b> {{ getDateString(new Date(gameData.CreationDate)) }}
					</v-sheet>
				</v-row>
				<v-row v-if="gameData.Downloads">
					<v-sheet class="pa-2 ma-2">
						<b>Downloads:</b> {{ gameData.Downloads }}
					</v-sheet>
				</v-row>
				<v-row v-if="gameData.TimesPlayed">
					<v-sheet class="pa-2 ma-2">
						<b>Times Played:</b> {{ gameData.TimesPlayed }}
					</v-sheet>
				</v-row>
			</v-col>
		</v-row>
		<v-row justify="end">
			<v-sheet class="pa-2 ma-2 pt-10">
				<v-btn 
					color="primary" 
					@click.stop="downloadGame(gameData.Id, gameData.DownloadLink)"
					prepend-icon="mdi-download"
				>
					Download
				</v-btn>
				<v-btn
					class="ml-2"
					color="primary" 
					@click.stop="playGame(gameData.Id)"
					prepend-icon="mdi-play"
				>
				<!-- TODO: This button should only appear if the game is playable in-browser.-->
					Play
				</v-btn>
			</v-sheet>
		</v-row>
		<v-row v-if="similarGames && similarGames.length > 0">
			<div data-test="title" class="ml-10 pb-3 text-h5">More Games Like {{ gameData.Name }}...</div>
			<GameList :games="similarGames"/>
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
	import NoData from '../components/NoData.vue';
	import GameList from '../components/GameList.vue';

	import * as GamesServices from '../services/gamesServices';
	import Game from '../models/game';
	export default {
		name: 'GameInfo',
		data: () => ({
			dataLoading: true,
			gameData: Game,
			similarGames: []
		}),
		components:{
			Rating,
			Loading,
			Footer,
			NoData,
			GameList
		},
		methods: {
			async getGameInfo(id) {
				this.dataLoading = true;
				await GamesServices.getGameInfo(id)
					.then(response => {
						this.gameData = Game.populateFromObject(response.data);
						if (this.gameData.Tags){
							this.getGamesLikeThis(this.gameData.Tags);
						}
						else {
							this.dataLoading = false
						}
					}).catch(error => {
						console.log(error);
						this.dataLoading = false
						this.gameData = null;
					});
			},
			async getGamesLikeThis(tags) {
				await GamesServices.getGamesWithTags(tags)
					.then(response => {
						this.similarGames = response?.data.filter(g => g.Id !== this.gameData.Id).map(g => Game.populateFromObject(g));
						this.dataLoading = false
					}).catch(error => {
						console.log(error);
						this.dataLoading = false
					});
			},
			downloadGame(id, link){
				// Download this game with id
				console.log(`DOWNLOAD: ${id} LINK: ${link}`)
			},
			playGame(id){
				// Play this game with id
				console.log(`PLAY: ${id}`)
			},
			getDateString(date){
				if (date){
					let year = date.getFullYear()
					let month = date.getMonth()
					let day = date.getDate()

					return `${month}/${day}/${year}`
				}

				return '';
			}
		},
		created() {
			this.getGameInfo(this.$route.params.id);
		}
	};
</script>