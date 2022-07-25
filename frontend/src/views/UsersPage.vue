<template>
	<h1 class="mt-10 ml-13">Site Users</h1>
  <v-table class="mt-4 mb-4 pl-10 pr-10">
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
			<!-- TODO: Add onclick event to go to users page -->
      <tr
        v-for="user in users"
        :key="user.email"
				style="cursor: pointer"
      >
        <td>{{ user.name }}</td>
        <td>{{ user.email }}</td>
      </tr>
    </tbody>
  </v-table>
</template>

<script>
export default {
	name: 'UsersPage',
	
	data () {
		return {
			// Get list of users from database
			users: [
				{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},
								{name: 'Hailey Dennis', email: 'haileydennis@example.com'},
				{name: 'Test Account', email: 'test@example.com'},

			]
		}
	}
}
</script>