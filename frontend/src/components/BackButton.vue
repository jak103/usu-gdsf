<template>
    <v-row class="pa-3 pb-3" justify="left">
        <v-btn color="primary" @click="goBack">BACK
        </v-btn>
    </v-row>
</template>
<script>
    export default {
        data() {
            return {
               
            }
        },
         methods: {
           goBack(){
              this.$router.go(-1)
           }
         }
    }
</script>