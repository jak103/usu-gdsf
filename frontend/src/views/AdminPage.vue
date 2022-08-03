<template>
	<div>
		<!-- File uploads: FormData -->
		<h1>This is where you login. Thanks for stopping by.</h1>
		<router-link :to="{ name: 'newaccount', params: { type: 'admin' }}">
			Need an account? Create one here
		</router-link>

	</div>
</template>

<script>
export default {
	name: 'AdminPage',

}
</script>