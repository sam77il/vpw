<script>
	import { resolve } from "$app/paths";
	import { ShoppingCart, User } from "@lucide/svelte";

	const { user } = $props();
</script>

<header>
	<div class="header-content">
		<div class="header-left">
			<a href={resolve("/")}><img src="/imgs/logo.png" alt="Logo" /></a>
		</div>
		<div class="header-right">
			<a href={resolve(user ? `/account` : "/login")}><User /> Mein Konto</a>
			<a href={resolve("/cart")}><ShoppingCart /></a>
		</div>
	</div>
</header>

<style>
	header {
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.header-content {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		max-width: 70vw;
		width: 100%;
	}

	.header-left img {
		height: 150px;
	}

	.header-right {
		display: flex;
		flex-direction: row;
		gap: 25px;
		align-items: center;
		justify-content: center;
	}

	.header-right a {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		text-decoration: none;
		color: black;
	}
</style>
