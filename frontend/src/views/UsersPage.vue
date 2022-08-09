<template>
<div>
	<v-row>
		<v-col>
			<h1 class="mt-10 ml-13">Site Users</h1>
		</v-col>

		<v-spacer></v-spacer>

		<v-col class="d-flex align-end flex-column">
			<v-btn 
				class="mt-10 mr-13"
				color="secondary"
				@click="handleClickAdmin()"
			>
				Create Admin
			</v-btn>
		</v-col>
	</v-row>

	<v-progress-circular
		v-if="loadingUsers"
		style="position: fixed; top: 40%; left: 50%"
		indeterminate
		size="64"
  ></v-progress-circular>

  <v-table v-else class="mt-4 mb-4 pl-10 pr-10">
    <thead>
      <tr>
        <th class="text-left">
          Name
        </th>
        <th class="text-left">
          Email
        </th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="user in usersList.slice(((page - 1) * perPage), ((page - 1) * perPage) + perPage)"
        :key="user.email"
				style="cursor: pointer"
				@click.stop="handleClickUser({...user})"
      >
        <td>{{ user.firstName }} {{user.lastName}}</td>
        <td>{{ user.email }}</td>
      </tr>
    </tbody>
  </v-table>

	<UserForm :isAdminCreation=isAdminCreation
		:showSelf="showForm" 
		@close="handleClose()" 
		@leftDialog="resetDefaults()" 
		:selectedUser="selectedUser" 
		@save="handleSaveEdit(selectedUser)"
		@delete="handleDeleteUser(selectedUser)"
		@createAdmin="handleCreateAdmin(selectsedUser)"	
	/>

	<v-pagination
		v-model="page"
		:length="Math.ceil(usersList.length / perPage)">
	</v-pagination>

	<v-snackbar
		v-model="snackbar"
		:timeout=2000
	>
		{{ snackbarText }}

		<template v-slot:actions>
			<v-btn
				color="pink"
				variant="text"
				@click="snackbar = false"
			>
				Close
			</v-btn>
		</template>
  </v-snackbar>

	<Footer></Footer>
	
	</div>
</template>

<script>
import UserForm from '../components/UserForm.vue'
import Footer from '../components/Footer.vue'
import {ref} from "vue";
import axios from "axios";

export default {

    name: "UsersPage",
    data() {
        return {
            // Replace with users from db
            usersList: [
                { firstName: "Test", lastName: "User", email: "testUser@example.com", dob: '2022-08-01' },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								{ firstName: "Test", lastName: "User", email: "testUser@example.com", dob: '2022-08-01' },	
            ],

            // Set to false when users have loading, default should be true
            loadingUsers: false,
						showForm: ref(false),
						isAdminCreation: ref(false),
						selectedUser: ref({
							firstName: "",
							lastName: "",
							email: "",
							dob: ""
						}),
						page: 1,
						perPage: 13,
						snackbar: false,
						snackbarText: "",

        };
    },
    components: { UserForm, Footer },

		methods: {

			// We don't have a users endpoint on the server yet
			async getUsers() {
				// this.loadingUsers = true;
				// await axios.get('http://127.0.0.1:8080/users')
				// 	.then(response => {
				// 		this.usersList = response.data[0];
				// 		this.loadingUsers = false
				// 	}).catch(err => {
				// 		console.log(err.response.data);
				// 		this.loadingUsers = false
				// 	});
			},

			handleClickUser(user) {
				this.showForm = true
				this.selectedUser = user
			},

			handleClickAdmin() {
				this.showForm = true
				this.isAdminCreation = true
			},

			resetDefaults() {
				this.showForm = false
				this.isAdminCreation = false
				this.selectedUser = {
					firstName: "",
					lastName: "",
					email: "",
					dob: ""
				}
			},

			handleClose() {
				this.showForm = false
			},

			handleSaveEdit(user, newPassword) {
				// Find user and update info
				// Need to check which items are being updated. New password may be blank.
				console.log(user)
				this.showForm = false

				// If successful:
				this.snackbar = true
				this.snackbarText = "Successfully updated user account!"
			},

			handleDeleteUser(user) {
				// Find user and delete
				console.log(user)
				this.showForm = false

				// If successful:
				this.snackbar = true
				this.snackbarText = "Successfully deleted user!"
			},

			handleCreateAdmin(user, newPassword) {
				// Create new admin user
				console.log(user)
				this.showForm = false
				
				// If successful:
				this.snackbar = true
				this.snackbarText = "Successfully created admin account!"
			}
		},
		
		created() {
			this.getUsers()
		}
}
</script>