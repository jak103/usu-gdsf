<template>
    <div>
        <v-table>
            <thead>
                <tr>
                    <th v-for="header in headers" :key="header.value">{{ header.text }}</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in testUsers" :key="user.ID">
                    <td>{{ user.ID }}</td>
                    <td>{{ user.Username }}</td>
                    <td>{{ user.Displayname }}</td>
                    <td>{{ user.Role }}</td>
                </tr>
            </tbody>
        </v-table>
    </div>
</template>


<script>
//TODO: errors pop up when trying to import this, I think it's a backend issue?
// import axios from "axios"; 
export default {
    name: 'AdminAllUsers',

    data: () => ({
        // users: [],
        loading: false,
        headers: [
            { text: 'ID', value: 'ID' },
            { text: 'Username', value: 'Username' },
            { text: 'Name', value: 'Displayname' },
            { text: 'Role', value: 'Role' }
        ],
        testUsers: [
            {
                ID: 1,
                Username: 'test',
                Displayname: 'Test User',
                Role: 'Admin'
            },
            {
                ID: 2,
                Username: 'test2',
                Displayname: 'Test User 2',
                Role: 'User'
            }
        ]
    }),

    methods: {
        async users() {
            this.loading = true;
            await axios.get('http://127.0.0.1:8080/allusers')
                .then(response => {
                    this.users = response.data[0];
                    this.loading = false
                }).catch(error => {
                    console.log(error.response.data);
                    this.dataLoading = false
                });
        },
    },
    created() {
        this.users();
    }
}
</script>