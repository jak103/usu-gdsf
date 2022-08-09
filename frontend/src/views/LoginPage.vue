<template>
	<h2 style="text-align: center; color:#10243c; font-weight: normal;">USU Game Dev Store Front</h2>

	<v-container style="max-width: 600px; ">
		<v-card style="background-color: #10243c; padding: 20px;">
			<v-card-title style="text-align: center; padding-bottom: 3px;">
				<h1 style=" color: #FFFFFF; padding-bottom: 36px;">Store Login</h1>
			</v-card-title>

			<v-divider></v-divider>

			<v-list-item style="padding:16px; text-align: center;">
				<input type="text" name="username" v-model="input.username" placeholder="Username" 
				style="background-color: #FFFFFF; padding: 16px;" />
			</v-list-item>
			<v-list-item style="padding:16px; text-align: center;">
				<input type="password" name="password" v-model="input.password" placeholder="Password" 
				style="background-color: #FFFFFF; padding: 16px" />
			</v-list-item>

			<v-list-item style="text-align: end;">
				<button type="button" v-on:click="login()" 
				style=" color: #FFFFFF; padding: 16px; font-weight: bold; border-radius: 0.25rem;
				background-color: #3577c7;">Login</button>
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
				password: "",
				access: false
			}
		}
	},
	methods: {
		login() {
			if (this.input.username == "" || this.input.password == "") {
				console.log("Username/Password cannot be empty")
			} else {
				if (this.input.username == "admin" && this.input.password == "adminpass") {
					this.input.access = true
					return this.$router.push('/')
				}
			}
		}
	}
}
</script>