<!-- Followed very similar format to the "all users" page that others created, since these two things are fairly similar. -->

<template>
    <div>
        <v-table>
            <thead>
                <tr>
                    <th v-for="header in headers" :key="header.value">{{ header.text }}</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="game in testGames" :key="game.ID">
                    <td>{{ game.ID }}</td>
                    <td>{{ game.author }}</td>
                    <td>{{ game.description }}</td>
                    <td>{{ game.timeStamp }}</td>
                </tr>
            </tbody>
        </v-table>
    </div>
</template>


<script>
export default {
    name: 'AdminAllUsers',
    
    data: () => ({
        users: [],
        loading: false,
        headers: [
            { text: 'ID', value: 'ID' },
            { text: 'author', value: 'author' },
            { text: 'description', value: 'description' },
            { text: 'timeStamp', value: 'timeStamp' }
        ],
        testGames: [
            {
                ID: 1,
                author: 'testAuthor',
                description: 'a very nice game',
                timeStamp: 'today'
            },
            {
                ID: 2,
                author: 'test2',
                description: 'Test User 2',
                timeStamp: 'yesterday'
            }
        ],
        games: []
    }),

    async mounted() {
        await this.getAllGames();
    },

    methods: {
        async getAllGames() {
            try {
                //const response = await axios.get('/api/games');
                this.games = response.data;
            } catch (error) {
                console.log(error);
            }
        }
    }
}
</script> 