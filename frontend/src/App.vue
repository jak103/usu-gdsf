<template>
<v-app>
  <v-navigation-drawer
    app
    color="primary"
  >
    <NavMenu/>
  </v-navigation-drawer>

<!-- Sizes your content based upon application components -->
<v-main>
<!-- Provides the application the proper gutter -->
<v-container fluid>
<router-view></router-view>
</v-container>
</v-main>

<v-footer color="primary" app
paddless
>
<v-row
justify="center"
no-gutters
>

<v-btn
color="white"
text
rounded
class="my-2"
>
    <router-link to="/">Home</router-link>
    </v-btn>

<v-btn
color="white"
text
rounded
class="my-2"
>
    <router-link to="/about">About Us</router-link>
    </v-btn>

<v-col
class="primary lighten-2 text-center black--text"
cols="12"
>
{{ new Date().getFullYear() }} — <strong>Vuetify</strong>
</v-col>
</v-row>
</v-footer>
</v-app>
</template>

<script>
import HomeView from './views/HomeView.vue'
import NavMenu from './components/NavMenu.vue'

export default {
name: 'App',

components: {
  HomeView,
  NavMenu
},

data: () => ({

}),

computed: {

}
}
</script>
