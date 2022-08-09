<template>
		<v-card color="primary" id="form-container">
			<v-card-title>
				<span>Add Game</span>
			</v-card-title>

			<v-card id="form-item">
				<label for="gameName"> Game Title </label>
				<input id="gameName" type="text" v-model="gameName">
			</v-card>


			<v-card id="form-item">
				<label for="gameAuthor"> Author(s) </label>
				<input id="gameAuthor" type="text" v-model="gameAuthor">
			</v-card>

			<!-- Game Description -->
			<v-card id="form-item">
				<label for="gameDesc"> Description </label>
				<textarea type="text" id="gameDesc" v-model="gameDesc"> </textarea>
			</v-card>

			<v-card id="form-item">
				<label> Game Files: </label>
				<input type="text" v-model="gameFile">
			</v-card>

			<v-card id="form-item">
			<!-- Image Upload -->
					<div id="image-upload">
          <label for="gameImage"> Thumbnail: </label>
					<input id="gameImage" type="file" @change="onFileChange" />
					 </div>
          <div id="imageBox">
    				<img v-if="url" :src="url" />
 					</div>
			</v-card> 
			<!-- End Image Upload -->

			<!-- Submission Button-->
			<div id="form-item">
				<v-btn color="secondary" @click="submitGame">Submit</v-btn>
			</div>
	</v-card>
</template>

<script>
export default {
	name: 'AdminPage',


	data: () => ({
		gameName: '',
		gameDesc: '',
		gameAuthor: '',
		gameFile: '',
		gameImage: '',
		image:'',
		url: null,
	}),

	//submit
	methods: {
		submit() {
			//send form data to server
		},
	  addAuthor() {
			this.authors.push({
				name: this.authorName,
				email: this.authorEmail,
			});
			this.authorName = '';
			this.authorEmail = '';
		},
		onFileChange(e) {
      const file = e.target.files[0];
      this.url = URL.createObjectURL(file);
    },
	}

}
</script>

<style>

	#form-container {
		width: 90%;
		height: 90%;
		margin: 0 auto;
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		border: 1px solid rgb(49, 49, 49);
		padding: 10px;
	}

	#form-item {
		width: 100%;
		height: 100%;
		display: flex;
		justify-content: flex-end;
		flex-wrap: wrap;
		padding: 10px;
		border-radius: 5px;
		margin: 10px;
	}

	#form-item label {
		flex: 1;
    font-family: Arial, Helvetica, sans-serif;
	}

	#form-item textarea {
		box-sizing: border-box;
		width: 100%;
		height: 100%;
		padding: 2px;
		border: 1px solid rgb(49, 49, 49);
		flex: 4;
	}

	#form-item input {
		box-sizing: border-box;
		width: 100%;
		height: 100%;
		padding: 4px;
		border: 1px solid rgb(49, 49, 49);
		flex: 4;
	}

	#image-upload {
		width: 100%;
		height: 100%;
		display: flex;
    flex-direction: row;
		justify-content: end;
		flex-wrap: wrap;
	}

  #imageBox {
    width: 100%;
    height: 100%;
    flex: 1;
		padding: 10px;
		display: flex;
		justify-content: center;
  }

	#imageBox img {
		max-width: 100%;
		max-height: 300px;
	}

</style> 