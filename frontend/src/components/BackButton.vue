<template>
    <v-row justify="left">
        <v-btn class="ma-3 mb-6" color="primary" @click="goBack">BACK
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