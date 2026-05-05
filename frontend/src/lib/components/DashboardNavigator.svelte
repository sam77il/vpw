<script>
	import { resolve } from "$app/paths";
	import { Home, Settings, ShoppingCart } from "@lucide/svelte";
</script>

<div class="navigator">
	<nav>
		<a href={resolve("/")}><Home color="orange" /><span>Startseite</span></a>
		<a href={resolve("/account/orders")}><ShoppingCart color="orange" /><span>Bestellungen</span></a
		>
		<a href={resolve("/account/settings")}><Settings color="orange" /><span>Einstellungen</span></a>
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
