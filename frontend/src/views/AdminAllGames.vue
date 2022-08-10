<template>
    <div>
        <v-table>
            <thead>
                <tr>
                    <th v-for="header in headers" :key="header.value">{{ header.text }}</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="game in games" :key="game.ID">
                    <td>{{ game.ID }}</td>
                    <td>{{ game.title }}</td>
                    <td>{{ game.description }}</td>
                    <td>{{ game.userid }}</td>
                    <td>{{ game.version }}</td>
                    <td>{{ game.timestamp }}</td>
                    <td>{{ game.tags }}</td>
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
    loading: false,
    games: [],
    headers: [
      { text: 'ID', value: 'ID' },
      { text: 'Title', value: 'title' },
      { text: 'Description', value: 'description' },
      { text: 'UserId', value: 'userid' },
      { text: 'VersionNumber', value: 'version' },
      { text: 'PublishTimestamp', value: 'timestamp' },
      { text: 'Tags', value: 'tags' },
    ],
    testGames: [
      {
        ID: 1,
        title: 'test game 1',
        description: 'test description 1',
        userid: 1,
        version: 1,
        timestamp: '2020-01-01',
        tags: 'test, tags'
      },
      {
        ID: 2,
        title: 'test game 2',
        description: 'test description 2',
        userid: 2,
        version: 2,
        timestamp: '2020-01-02',
        tags: 'test, tags'
      },
      {
        ID: 3,
        title: 'test game 3',
        description: 'test description 3',
        userid: 3,
        version: 3,
        timestamp: '2020-01-03',
        tags: 'test, tags'
      }
    ]
  }),
  asyncComputed: {
    games: function () {
      return axios.get("http://localhost:8080/game")
        .then(resp => {
          this.games = resp.data;
        })
    }
  },
}
</script> 