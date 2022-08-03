<template>
	<div>
		<!-- File uploads: FormData -->
		<h1>This is where you login. Thanks for stopping by.</h1>
	</div>
	<v-container style="max-width: 600px; ">
		<v-card style="background-color: #10243c;">
			<v-card-title>
				<h3 style=" color: #FFFFFF;" >Login</h3>
			</v-card-title>
			<v-divider></v-divider>
			<v-list-item>
				<input type="text" name="username" v-model="input.username" placeholder="Username" style="background-color: #FFFFFF; padding: 16px;" />
			</v-list-item>
			<v-list-item>
				<input type="password" name="password" v-model="input.password" placeholder="Password" style="background-color: #FFFFFF; padding: 16px" />
			</v-list-item>
			<v-list-item>
				<button type="button" onclick="login()" style=" color: #FFFFFF; padding: 16px; background-color: #679ad6;" >Login</button>
			</v-list-item>
		</v-card>
	</v-container>
</template>
<script>
export default {
	name: 'AdminPage',
	data() {
		return {
			input: {
				username: "",
				password: ""
			}
		}
	},
	methods: {
		login() {
			if (this.input.user == "" || this.input.password == "") {
				console.log("Username/Password cannot be empty")
			} else {
				//authenticate user
			}
		}
	}
}
</script>