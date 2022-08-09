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
		v-if="loading"
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
        v-for="user in users.slice(((page - 1) * perPage), ((page - 1) * perPage) + perPage)"
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
		data-test="pagination" 
		v-model="page"
		:length="Math.ceil(users.length / perPage)">
	</v-pagination>

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
            // Get list of users from database
            users: [
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
								
            ],
            // Set to false when users have loading, default should be true
            loading: false,
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
        };
    },
    components: { UserForm, Footer },

		methods: {
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
				console.log("Saved user edit!")
				console.log(user)
				this.showForm = false
			},

			handleDeleteUser(user) {
				// Find user and delete
				console.log(user)
				console.log("User deleted")
				this.showForm = false
			},

			handleCreateAdmin(user, newPassword) {
				// Create new admin user
				console.log("Admin created")
				console.log(user)
				this.showForm = false
			}
		}
}
</script>