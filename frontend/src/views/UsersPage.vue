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
			color="secondary">
			Create Admin</v-btn>
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
				@click="handleClick(user)"
      >
        <td>{{ user.firstName }} {{user.lastName}}</td>
        <td>{{ user.email }}</td>
      </tr>
    </tbody>
  </v-table>

	<EditUserForm :showSelf="showForm" @close="handleClose()" :selectedUser="selectedUser" @save="handleSave(selectedUser)"/>
	
	</div>
</template>

<script>
import EditUserForm from '../components/EditUserForm.vue'
import {ref} from "vue";
export default {

    name: "UsersPage",
    data() {
        return {
            // Get list of users from database
            users: [
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
								{ firstName: "Test", lastName: "User", email: "testUser@example.com" },
								
            ],
            // Set to false when users have loading, default should be true
            loading: false,
						showForm: ref(false),
						selectedUser: {
							firstName: "",
							lastName: "",
							email: ""
						}
        };
    },
    components: { EditUserForm },

		methods: {
			handleClick(user) {
				this.showForm = true
				this.selectedUser = user
			},

			handleClose() {
				this.showForm = false
				this.selectedUser = {
					firstName: "",
					lastName: "",
					email: "",
				}
			},

			handleSave(user) {

			}
		}
}
</script>