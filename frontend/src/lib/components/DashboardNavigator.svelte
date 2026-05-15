<script>
	import { resolve } from "$app/paths";
	import {
		Book,
		Home,
		Library,
		Menu,
		Settings,
		ShoppingBag,
		ShoppingCart,
		Users,
		X
	} from "@lucide/svelte";

	const { page } = $props();

	let mobileNavOpen = $state(false);

	function toggleMobileNav() {
		mobileNavOpen = !mobileNavOpen;
	}

	function closeMobileNav() {
		mobileNavOpen = false;
	}
</script>

<div class="navigator" class:open={mobileNavOpen}>
	<div class="nav-header">
		<h3>Menu</h3>
		<button class="close-btn" onclick={closeMobileNav} aria-label="Close menu">
			<X size={20} />
		</button>
	</div>
	<nav>
		<a href={resolve("/")} onclick={closeMobileNav}
			><Home color="orange" /><span>Startseite</span></a
		>
		{#if page === "admin"}
			<a href={resolve("/admin/users")} onclick={closeMobileNav}
				><Users color="orange" /><span>Benutzer</span></a
			>
			<a href={resolve("/admin/orders")} onclick={closeMobileNav}
				><ShoppingBag color="orange" /><span>Bestellungen</span></a
			>
			<a href={resolve("/admin/categories")} onclick={closeMobileNav}
				><Library color="orange" /><span>Kategorien</span></a
			>
			<a href={resolve("/admin/products")} onclick={closeMobileNav}
				><ShoppingCart color="orange" /><span>Produkte</span></a
			>
			<a href={resolve("/admin/pages")} onclick={closeMobileNav}
				><Book color="orange" /><span>Seiten</span></a
			>
		{:else}
			<a href={resolve("/account/orders")} onclick={closeMobileNav}
				><ShoppingCart color="orange" /><span>Bestellungen</span></a
			>
			<a href={resolve("/account/settings")} onclick={closeMobileNav}
				><Settings color="orange" /><span>Einstellungen</span></a
			>
		{/if}
	</nav>
</div>

<div class="dashboard-mobile-bar">
	<button
		class="menu-toggle"
		onclick={toggleMobileNav}
		aria-label={mobileNavOpen ? "Schließen" : "Menü"}
	>
		{#if mobileNavOpen}
			<X size={24} />
		{:else}
			<Menu size={24} />
		{/if}
	</button>
	<span>Dashboard</span>
</div>

{#if mobileNavOpen}
	<div class="drawer-overlay" onclick={closeMobileNav}></div>
{/if}

<style>
	.dashboard-mobile-bar {
		display: none;
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		height: 60px;
		padding: 0 16px;
		background-color: white;
		border-bottom: 1px solid var(--border);
		z-index: 60;
		align-items: center;
		display: flex;
		gap: 12px;
	}

	.menu-toggle {
		background: none;
		border: none;
		cursor: pointer;
		padding: 10px;
		color: var(--text);
		border-radius: 8px;
		transition:
			background-color 0.2s ease,
			color 0.2s ease;
	}

	.menu-toggle:hover {
		background-color: var(--accent-light);
		color: var(--accent);
	}

	.drawer-overlay {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-color: rgba(0, 0, 0, 0.35);
		z-index: 50;
	}

	.navigator {
		display: none;
		position: fixed;
		height: 100vh;
		width: max-content;
		padding: 25px;
		background-color: whitesmoke;
		overflow-y: auto;
		z-index: 40;
	}

	@media (min-width: 1024px) {
		.navigator {
			display: flex;
			flex-direction: column;
			width: auto;
		}
	}

	/* Mobile Navigation Drawer */
	@media (max-width: 1023px) {
		.dashboard-mobile-bar {
			display: flex;
		}

		.navigator {
			display: flex;
			flex-direction: column;
			position: fixed;
			left: -100%;
			top: 0;
			width: min(280px, 80vw);
			height: 100vh;
			background-color: whitesmoke;
			padding: 0;
			transition: left 0.3s ease;
			border-right: 1px solid var(--border);
			z-index: 55;
		}

		.navigator.open {
			left: 0;
		}

		.navigator nav {
			padding-top: 80px;
		}
	}

	.nav-header {
		display: none;
		align-items: center;
		justify-content: space-between;
		padding: 16px;
		border-bottom: 1px solid var(--border);
	}

	@media (max-width: 1023px) {
		.nav-header {
			display: flex;
		}
	}

	.nav-header h3 {
		margin: 0;
		font-size: 1rem;
		font-weight: 700;
	}

	.close-btn {
		background: none;
		border: none;
		cursor: pointer;
		padding: 4px;
		display: flex;
		align-items: center;
		justify-content: center;
		color: var(--text);
		transition: color 0.2s ease;
	}

	.close-btn:hover {
		color: var(--accent);
	}

	.navigator nav {
		display: flex;
		flex-direction: column;
		gap: 25px;
		justify-content: flex-start;
		align-items: center;
		padding: 25px;
	}

	@media (max-width: 1023px) {
		.navigator nav {
			gap: 0;
			padding: 0;
		}
	}

	.navigator nav a {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-decoration: none;
		font-size: 0.8rem;
		color: var(--text);
		gap: 5px;
		padding: 8px;
		border-radius: 6px;
		transition:
			background-color 0.2s ease,
			color 0.2s ease;
	}

	@media (max-width: 1023px) {
		.navigator nav a {
			flex-direction: row;
			justify-content: flex-start;
			align-items: center;
			padding: 12px 16px;
			width: 100%;
			border-radius: 0;
			gap: 12px;
		}
	}

	.navigator nav a:hover {
		color: var(--accent);
		background-color: var(--accent-light);
	}

	.navigator nav a span {
		display: none;
	}

	@media (max-width: 1023px) {
		.navigator nav a span {
			display: block;
			font-size: 0.95rem;
			flex: 1;
		}
	}
</style>
