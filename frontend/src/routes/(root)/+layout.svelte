<script>
	import favicon from "$lib/assets/favicon.svg";
	import Header from "$lib/components/Header.svelte";
	import Navigation from "$lib/components/Navigation.svelte";
	import "$lib/styles/app.css";
	let { children, data } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Header user={data.user} />
<Navigation categories={data.categories} />
{@render children()}
