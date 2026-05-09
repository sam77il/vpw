<script>
	import "$lib/styles/app.css";
	import DashboardHeader from "$lib/components/DashboardNavigator.svelte";
	import { page } from "$app/state";
	const { children } = $props();
</script>

<DashboardHeader page={page.url.pathname.includes("/admin") ? "admin" : "user"} />
{@render children()}
