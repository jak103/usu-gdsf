<template>
    <v-container data-test="container" :style="containerStyle">
        <v-row
            class="fill-height"
            align-content="center"
            justify="center"
        >
            <v-col
                data-test="text-wrapper"
                class="text-subtitle-1 text-center"
                cols="12"
            >
                Oops! There was no data for: {{ text }}
            </v-col>
            <v-col cols="1">
                <v-btn
                    data-test="home-button"
                    color="secondary"
                    @click.stop="goHome()"
                    prepend-icon="mdi-home"
                >
                    Go Home
                </v-btn>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    export default {
        name: 'NoData',
        props: {
            text:String,
            containerStyle:String
        },
        methods: {
            goHome() {
                this.$router.push("/")
            }
        }
    }
</script>