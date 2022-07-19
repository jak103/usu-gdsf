<template>
  <v-container>
    {{ something }}
  </v-container>
</template>

<script>
export default {
  name: 'GameList',
  
  data: () => ({
    something: 'This is where we will implement a dynamic view of different games. We will need to pass in props.'
  }),
  
  methods: {
    
  }
}
</script>