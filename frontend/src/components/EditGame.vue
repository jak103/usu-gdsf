<template>
  
</template>

<script setup>
import Game from '../models/Game.js'
  defineProps({
    showEdit: Boolean,
    game: {
      name: String,
      developer: String,
      version: String,
      description: String,
      imagePath: String,
	  downloadLink: String
    }
  })
</script>

<style>

</style>