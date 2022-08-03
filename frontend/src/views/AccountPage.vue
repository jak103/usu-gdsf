<template>
	<v-row>
		<v-card width="2000" height="100" color="primary">
			<v-card-title class="text-h3 text-center " color="primary" height="1000"> User Profile </v-card-title>
		</v-card>
	</v-row>
	<v-row justify-center>
		<v-card color="primary" width="2000">
			<v-row>
				<v-col justify="space-between" style="align-self: center">
					<v-img src="https://randomuser.me/api/portraits/men/85.jpg" height="100" > </v-img>
				</v-col>
				<v-col>
					<v-row class="flex-column ">
						<v-col> {{ user.firstName }} {{user.lastName}}</v-col>
						<v-col> {{ user.email }}</v-col>
						<v-col> Birthday: {{ user.dateOfBirth }}</v-col>
					</v-row>
				</v-col>
			</v-row>
		</v-card>
	</v-row>
</template>

<script>
export default {
	name: 'AccountPage',
	data: () => ({
		user: {
			Id: "123",
			email: "example@usu.edu",
			firstName: "John",
			lastName: "Doe",
			dateOfBirth: "Jan, 1, 2022",
			joinedTimestamp: 123456789,
			image:"https://randomuser.me/api/portraits/men/85.jpg" 
		}
	
	})
}
</script>