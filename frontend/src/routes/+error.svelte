<script>
	import { resolve } from "$app/paths";
	import { page } from "$app/state";
</script>

<svelte:head>
	<title>{page.status} - Fehler</title>
</svelte:head>

<div class="container">
	{#if page.status === 404}
		<h1>404 - Seite nicht gefunden</h1>
		<p>Die angeforderte Seite konnte nicht gefunden werden.</p>
	{:else if page.status === 500}
		<h1>500 - Interner Serverfehler</h1>
		<p>Es ist ein Fehler auf dem Server aufgetreten. Bitte versuchen Sie es später erneut.</p>
	{:else}
		<h1>{page.status} - Fehler</h1>
		<p>Ein unerwarteter Fehler ist aufgetreten.</p>
	{/if}
	<a href={resolve("/")}>Zur Startseite</a>
</div>

<style>
	.container {
		font-family: Arial, sans-serif;
		background-color: #f8f8f8;
		color: #333;
		margin: 0;
		padding: 0;
	}

	.container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		height: 100vh;
		text-align: center;
	}
</style>
