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
        v-for="user in users"
        :key="user.email"
				style="cursor: pointer"
				@click="handleClickUser({...user})"
      >
        <td>{{ user.firstName }} {{user.lastName}}</td>
        <td>{{ user.email }}</td>
      </tr>
    </tbody>
  </v-table>

	<UserForm :isAdmin=isAdmin 
		:showSelf="showForm" 
		@close="handleClose()" 
		@leftDialog="resetDefaults()" 
		:selectedUser="selectedUser" 
		@save="handleSave(selectedUser)"
		@createAdmin="handleCreateAdmin(selectedUser)"	
	/>

	<Footer></Footer>
	
	</div>
</template>

<script>
import UserForm from '../components/UserForm.vue'
import Footer from '../components/Footer.vue'
import {ref} from "vue";
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
								
            ],
            // Set to false when users have loading, default should be true
            loading: false,
						showForm: ref(false),
						isAdmin: ref(false),
						selectedUser: ref({
							firstName: "",
							lastName: "",
							email: "",
							dob: ""
						})
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
				this.isAdmin = true
			},

			resetDefaults() {
				this.showForm = false
				this.isAdmin = false
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

			handleSaveEdit(user) {
				// Find user
				this.resetDefaults()
			},

			handleDeleteUser(user) {
				// Find user

			},

			handleCreateAdmin(user) {
				console.log(user)
			}
		}
}
</script>