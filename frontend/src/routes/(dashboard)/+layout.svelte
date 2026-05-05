<script>
	import "$lib/styles/app.css";
	import DashboardHeader from "$lib/components/DashboardNavigator.svelte";
	const { children } = $props();
</script>

<DashboardHeader />
{@render children()}
