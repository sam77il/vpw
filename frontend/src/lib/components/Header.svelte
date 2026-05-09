<script>
	import { resolve } from "$app/paths";
	import { ShoppingCart, Star, User, UserStar } from "@lucide/svelte";

	const { user, categories } = $props();
</script>

<header>
	<div class="header-content">
		<div class="header-left">
			<a href={resolve("/")}><img src="/imgs/logo.png" alt="Logo" /></a>
		</div>
		<div class="header-center">
			<nav>
				{#each categories as category (category.id)}
					<a href={resolve(`/categories/${category.id}`)}>{category.label}</a>
				{/each}
			</nav>
		</div>
		<div class="header-right">
			{#if user?.role === "admin"}
				<a href={resolve(user?.role === "admin" && "/admin")}><UserStar /><span>Admin</span></a>
			{/if}
			<a href={resolve(user ? `/account` : "/login")}><User /><span>Mein Konto</span></a>
			<a href={resolve("/favorites")}><Star /><span>Favoriten</span></a>
			<a href={resolve("/cart")}><ShoppingCart /><span>Warenkorb</span></a>
		</div>
	</div>
</header>

<style>
	header {
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		border-bottom: 1px solid lightgray;
	}

	.header-content {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		padding: 25px 25px 0 25px;
	}

	.header-left img {
		height: 75px;
	}

	.header-right {
		display: flex;
		flex-direction: row;
		gap: 15px;
		align-items: center;
		justify-content: center;
	}

	.header-right a {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-decoration: none;
		color: black;
		font-size: 0.8rem;
		transition: opacity 0.1s;
	}

	.header-right a:hover {
		opacity: 0.7;
	}

	nav {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		gap: 25px;
	}

	nav a {
		text-decoration: none;
		color: black;
		font-size: 1rem;
	}

	nav a:hover {
		background-color: white;
		color: orange;
	}
</style>
