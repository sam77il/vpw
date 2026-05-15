<script>
	import { resolve } from "$app/paths";
	import { Book, Home, Library, Settings, ShoppingBag, ShoppingCart, Users } from "@lucide/svelte";
	// import { page = p } from "$app/state";

	const { page } = $props();
</script>

<div class="navigator">
	<nav>
		<a href={resolve("/")}><Home color="orange" /><span>Startseite</span></a>
		{#if page === "admin"}
			<a href={resolve("/admin/users")}><Users color="orange" /><span>Benutzer</span></a>
			<a href={resolve("/admin/orders")}><ShoppingBag color="orange" /><span>Bestellungen</span></a>
			<a href={resolve("/admin/categories")}><Library color="orange" /><span>Kategorien</span></a>
			<a href={resolve("/admin/products")}><ShoppingCart color="orange" /><span>Produkte</span></a>
			<a href={resolve("/admin/pages")}><Book color="orange" /><span>Seiten</span></a>
		{:else}
			<a href={resolve("/account/orders")}
				><ShoppingCart color="orange" /><span>Bestellungen</span></a
			>
			<a href={resolve("/account/settings")}
				><Settings color="orange" /><span>Einstellungen</span></a
			>
		{/if}
	</nav>
</div>

<style>
	.navigator {
		position: fixed;
		height: 100vh;
		width: max-content;
		padding: 25px;
		background-color: whitesmoke;
	}

	.navigator nav {
		display: flex;
		flex-direction: column;
		gap: 25px;
		justify-content: flex-start;
		align-items: center;
	}

	.navigator nav a {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-decoration: none;
		font-size: 0.8rem;
		color: black;
		gap: 5px;
	}

	.navigator nav a:hover {
		color: orange;
	}
</style>
