<template>
	<div>
		<v-row>
			<v-col>
				<h1 class="mt-10 ml-13">Game Management</h1>
			</v-col>
			<v-spacer></v-spacer>
			<v-col class="d-flex align-end flex-column">
				<v-btn 
					class="mt-10 mr-13"
					color="secondary"
					@click="handleCreateGame()">
					Upload a new game
				</v-btn>
			</v-col>
		</v-row>

		<v-table>
			<thead>
				<tr>
					<th> Game Title </th>
					<th> Game Developer </th>
					<th> Game Version </th>
				</tr>
			</thead>
			<tbody>
      			<tr v-for="game in games"
				:key="game.Id"
				:id="game.Id"
				style="cursor: pointer"
				@click="handleClickGame(game)">
        			<td>{{ game.Name }}</td>
    	 			<td>{{ game.Developer }}</td>
					<td>{{ game.Version }}</td>
      			</tr>
			</tbody>
		</v-table>

	<NewGame :showCreate=showCreate @cancel="closeForm()" :game = "selectedGame" @save="handleCreateSave(game)"/>
	<EditGame :showEdit="showEdit" @cancel="closeForm()" :game="selectedGame" @save="handleEditSave(game)"/>

	</div>
	<Footer></Footer>
</template>

<!-- when connected with database remove the selectedGame data and use id to call specified game. -->
<script>
import NewGame from '../components/NewGame.vue'
import EditGame from '../components/EditGame.vue'
import Footer from "../components/Footer.vue"
import axios from 'axios'
import {ref} from "vue"
export default {
	name: 'GameManagementPage',
    components: {
		NewGame,
		EditGame,
		Footer,
	},
	data() {
		
		return {
			games: [],
			showCreate: ref(false),
			showEdit: ref(false),
			selectedGame: ref({
				name: " ",
				developer: " ",
				version: " ",
				description: " ",
				imagePath: " ",
				downloadLink: " "
			}),
			error: ''
		}
	},

	methods: {
		async getGames() {
			await axios.get('http://127.0.0.1:8080/games')
				.then(response => {
					this.games = response.data[0]
				}).catch(error => {
					console.log(error.response.data)
				});
		},
		handleCreateGame() {
			this.showCreate = true	
		},

		handleClickGame(game) {
			this.selectedGame.name = game.Name
			this.selectedGame.developer = game.Developer
			this.selectedGame.version = game.Version
			this.selectedGame.description = game.Description
			this.selectedGame.imagePath = game.ImagePath
			this.selectedGame.downloadLink = game.DownloadLink
		
			this.showEdit = true
		},

		closeForm() {
			this.selectedGame.name = " "
			this.selectedGame.developer = " "
			this.selectedGame.version = " "
			this.selectedGame.description = " "
			this.selectedGame.imagePath = " "
			this.selectedGame.downloadLink = " "
	
			this.showEdit = false
			this.showCreate = false
		},

		handleCreateSave(game) {
			this.closeForm()
			this.selectedGame
			this.selectedGame.name = game.name
			this.selectedGame.developer = game.developer
			this.selectedGame.version = game.version
			this.selectedGame.description = game.description
			this.selectedGame.imagePath = game.imagePath
			this.selectedGame.downloadLink = game.downloadLink
		},

		handleEditSave(game) {
			this.closeForm()
			this.selectedGame.name = game.name
			this.selectedGame.developer = game.developer
			this.selectedGame.version = game.version
			this.selectedGame.description = game.description
			this.selectedGame.imagePath = game.imagePath
			this.selectedGame.downloadLink
		}
	},
	created() {
		this.getGames()
	}
}
</script>