<template>
    <v-container :style="containerStyle">
        <v-row
            class="fill-height"
            align-content="center"
            justify="center"
        >
            <v-col
                class="text-subtitle-1 text-center"
                cols="12"
            >
                {{ text }}
            </v-col>
            <v-col cols="6">
                <v-progress-linear
                    color="primary"
                    indeterminate
                    rounded
                    height="6"
                ></v-progress-linear>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    export default {
        name: 'Loading',
        props: {
            text:String,
            containerStyle:String
        }
    }
</script>