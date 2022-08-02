<template>
	<div>
		<!-- File uploads: FormData -->
		<h1>This is where you login. Thanks for stopping by.</h1>
	</div>
	<v-container>
		<v-card>
			<v-card-title>
				<h3>Login</h3>
			</v-card-title>
			<v-divider></v-divider>
			<v-list-item>
				<input type="text" name="username" v-model="input.username" placeholder="Username" />
				<input type="password" name="password" v-model="input.password" placeholder="Password" />
			</v-list-item>
			<v-list-item>
				<button type="button" onclick="login()" >Login</button>
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