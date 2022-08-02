<template>
		<div id="form-container">

			<div id="form-item">
				<label for="gameName"> Game Title </label>
				<input id="gameName" type="text" v-model="gameName">
			</div>


			<div id="form-item">
				<label for="gameAuthor"> Author(s) </label>
				<input id="gameAuthor" type="text" v-model="gameAuthor">
			</div>

			<!-- Game Description -->
			<div id="form-item">
				<label for="gameDesc"> Description </label>
				<textarea type="text" id="gameDesc" v-model="gameDesc"> </textarea>
			</div>

			<div id="form-item">
				<label> Game Files: </label>
				<input type="text" v-model="gameFile">
			</div>

			<div id="image-upload">
			<!-- Image Upload -->
				<form enctype="multipart/form-data">
						<div>
							<input type="file" name="image">
							<p> 
								Drag your file(s) here to begin
								<br> or click to browse 
							</p>

								<v-img width="200" :src="image"></v-img>
					</div>
				</form>
			</div> 
	</div>
</template>

<script>
export default {
	name: 'AdminPage',


	data: () => ({
		gameName: '',
		gameDesc: '',
		gameAuthor: '',
		gameFile: '',
		image:'',
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
		createImage(file) {
			const reader = new FileReader();

			reader.onload = e => {
				this.imageUrl = e.target.result;
			};
			reader.readAsDataURL(file);
  	},
		onFileChange(file) {
			if (!file) {
				return;
			}
			this.createImage(file);
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
		padding: 10px;
		border-radius: 5px;
		border: 1px solid rgb(49, 49, 49);
		margin: 10px;
	}

	#form-item label {
		flex: 1;
	}

	#form-item textarea {
		box-sizing: border-box;
		width: 100%;
		height: 100%;
		padding: 2px;
		border: 1px solid rgb(49, 49, 49);
		flex: 3;
	}

	#form-item input {
		box-sizing: border-box;
		width: 100%;
		height: 100%;
		padding: 2px;
		border: 1px solid rgb(49, 49, 49);
		flex: 3;
	}

	#image-upload {
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 10px;
		border-radius: 5px;
		border: 1px solid rgb(49, 49, 49);
		margin: 10px;
	}

</style> 