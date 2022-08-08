<template>
</template>

<script setup>
  defineProps({
    showCreate: Boolean,
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