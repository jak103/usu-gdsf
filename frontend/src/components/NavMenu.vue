<template>
  <v-list>
    <v-list-item
      prepend-avatar="https://randomuser.me/api/portraits/women/2.jpg"
      title="Sandra Adams"
      subtitle="sandra_a88@gmailcom"
    ></v-list-item>
  </v-list>

  <v-divider></v-divider>

  <v-list density="compact" nav>
    <v-list-item prepend-icon="mdi-folder" title="My Files" value="myfiles"></v-list-item>
    <v-list-item prepend-icon="mdi-account-multiple" title="Shared with me" value="shared"></v-list-item>
    <v-list-item prepend-icon="mdi-star" title="Starred" value="starred"></v-list-item>
  </v-list>
</template>

<script>
import { defineComponent } from 'vue';

// Components


export default defineComponent({
  name: 'NavMenu',

  components: {

  },
});
</script>
