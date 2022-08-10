<template>
    <div>
        <v-table>
            <thead>
                <tr>
                    <th v-for="header in headers" :key="header.value">{{ header.text }}</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in users" :key="user.ID">
                    <td>{{ user.id }}</td>
                    <td>{{ user.username }}</td>
                    <td>{{ user.displayname }}</td>
                    <td>{{ user.role }}</td>
                </tr>
            </tbody>
        </v-table>
    </div>
</template>

<script>
import axios from "axios";
export default {
    name: 'AdminAllUsers',

    data: () => ({
        users: [],
        loading: false,
        headers: [
            { text: 'ID', value: 'id' },
            { text: 'User Name', value: 'username' },
            { text: 'Name', value: 'displayname' },
            { text: 'Role', value: 'role' }
        ],
        users: []
    }),
    asyncComputed: {
        users: function () {
            return axios.get("http://localhost:8080/user")
                .then(resp => {
                    this.users = resp.data;
                })
        }
    },
}
</script>