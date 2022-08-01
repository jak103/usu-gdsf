<template>
<div>
	<h1 class="mt-10 ml-13">Site Users</h1>

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
			<!-- TODO: Add onclick event -->
      <tr
        v-for="user in users"
        :key="user.email"
				style="cursor: pointer"
				@click="handleClick"
      >
        <td>{{ user.name }}</td>
        <td>{{ user.email }}</td>
      </tr>
    </tbody>
  </v-table>

	<UserForm v-if="showForm"></UserForm>
	</div>
</template>

<script>
import UserForm from '../components/UserForm.vue'
export default {
    name: "UsersPage",
    data() {
        return {
            // Get list of users from database
            users: [
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
                { name: "Hailey Dennis", email: "haileydennis@example.com" },
                { name: "Test Account", email: "test@example.com" },
            ],
            // Set to false when users have loading, default should be true
            loading: false,
						showForm: false,
						selectedUser: {
							name: "",
							email: ""
						}
        };
    },
    components: { UserForm },

		methods: {
			handleClick() {
				console.log("clicked")
				this.showForm = true
			}
		}
}
</script>