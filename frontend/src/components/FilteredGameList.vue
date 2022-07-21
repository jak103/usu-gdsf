<template>
  <v-container>
    {{ something }}
  </v-container>
</template>

<script>
export default {
  name: 'FilteredGameList',
  
  data: () => ({
    something: 'This is where we will implement a filtered game list... We will need to pass in props.'
  }),
  
  methods: {
    
  }
}
</script>