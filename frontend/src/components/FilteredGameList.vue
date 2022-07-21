<template>
  <v-container>
    <v-list v-for="item in getFilteredList()">
        <v-card>
            {{item.id}}
        </v-card>
    </v-list>
  </v-container>
</template>

<script>
export default {
  name: 'FilteredGameList',
  
  data: () => ({
    something: 'This is where we will implement a filtered game list... We will need to pass in props.',
    // filteredList: getFilteredList()
  }),
  
  methods: {
    getFilteredList: function() {
        // TODO get lists of games based on determined filter
        return [{id: 1}, {id: 2}]
    }

  }
}
</script>