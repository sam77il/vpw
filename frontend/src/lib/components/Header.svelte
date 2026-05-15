<script>
	import { resolve } from "$app/paths";
	import { ShoppingCart, Star, User, UserStar, Menu, X } from "@lucide/svelte";

	const { user, categories } = $props();

	let mobileMenuOpen = $state(false);

	function toggleMobileMenu() {
		mobileMenuOpen = !mobileMenuOpen;
	}

	function closeMobileMenu() {
		mobileMenuOpen = false;
	}
</script>

<header>
	<div class="header-content">
		<div class="header-left">
			<a href={resolve("/")}><img src="/imgs/logo.png" alt="Logo" /></a>
		</div>
		<div class="header-center">
			<nav class="desktop-nav">
				{#each categories as category (category.id)}
					<a href={resolve(`/categories/${category.id}`)}>{category.label}</a>
				{/each}
			</nav>
		</div>
		<div class="header-right">
			<div class="desktop-actions">
				{#if user?.role === "admin"}
					<a href={resolve(user?.role === "admin" && "/admin")}><UserStar /><span>Admin</span></a>
				{/if}
				<a href={resolve(user ? `/account` : "/login")}><User /><span>Mein Konto</span></a>
				<a href={resolve("/favorites")}><Star /><span>Favoriten</span></a>
				<a href={resolve("/cart")}><ShoppingCart /><span>Warenkorb</span></a>
			</div>
			<button class="hamburger-menu" onclick={toggleMobileMenu} aria-label="Menu">
				{#if mobileMenuOpen}
					<X size={24} />
				{:else}
					<Menu size={24} />
				{/if}
			</button>
		</div>
	</div>

	<!-- Mobile Menu -->
	{#if mobileMenuOpen}
		<div class="mobile-menu">
			<nav class="mobile-nav">
				<h3>Kategorien</h3>
				{#each categories as category (category.id)}
					<a href={resolve(`/categories/${category.id}`)} onclick={closeMobileMenu}
						>{category.label}</a
					>
				{/each}

				<h3 style="margin-top: 20px;">Konto</h3>
				{#if user?.role === "admin"}
					<a href={resolve("/admin")} onclick={closeMobileMenu}>
						<UserStar size={18} />
						<span>Admin</span>
					</a>
				{/if}
				<a href={resolve(user ? `/account` : "/login")} onclick={closeMobileMenu}>
					<User size={18} />
					<span>Mein Konto</span>
				</a>
				<a href={resolve("/favorites")} onclick={closeMobileMenu}>
					<Star size={18} />
					<span>Favoriten</span>
				</a>
				<a href={resolve("/cart")} onclick={closeMobileMenu}>
					<ShoppingCart size={18} />
					<span>Warenkorb</span>
				</a>
			</nav>
		</div>
	{/if}
</header>

<style>
	header {
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		border-bottom: 1px solid var(--border);
		position: relative;
		background-color: white;
		z-index: 100;
	}

	.header-content {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		padding: clamp(12px, 3vw, 25px);
		max-width: 1400px;
		margin: 0 auto;
	}

	.header-left img {
		height: clamp(40px, 10vw, 75px);
		width: auto;
	}

	/* Desktop Navigation */
	.header-center {
		display: none;
	}

	@media (min-width: 1024px) {
		.header-center {
			display: flex;
			flex: 1;
		}
	}

	.desktop-nav {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		gap: clamp(15px, 3vw, 25px);
	}

	.desktop-nav a {
		text-decoration: none;
		color: var(--text);
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		font-weight: 500;
		transition: color 0.2s ease;
		padding: 8px 12px;
		border-radius: 6px;
	}

	.desktop-nav a:hover {
		color: var(--accent);
		background-color: var(--accent-light);
	}

	/* Desktop Actions */
	.desktop-actions {
		display: none;
		flex-direction: row;
		gap: clamp(10px, 2vw, 15px);
		align-items: center;
		justify-content: center;
	}

	@media (min-width: 1024px) {
		.desktop-actions {
			display: flex;
		}
	}

	.desktop-actions a {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-decoration: none;
		color: var(--text);
		font-size: 0.8rem;
		transition:
			opacity 0.2s ease,
			color 0.2s ease;
		gap: 4px;
		padding: 8px;
	}

	.desktop-actions a:hover {
		opacity: 0.7;
		color: var(--accent);
	}

	/* Hamburger Menu */
	.header-right {
		display: flex;
		align-items: center;
		gap: 15px;
	}

	.hamburger-menu {
		display: flex;
		align-items: center;
		justify-content: center;
		background: none;
		border: none;
		cursor: pointer;
		padding: 8px;
		color: var(--text);
		transition: color 0.2s ease;
	}

	.hamburger-menu:hover {
		color: var(--accent);
	}

	@media (min-width: 1024px) {
		.hamburger-menu {
			display: none;
		}
	}

	/* Mobile Menu */
	.mobile-menu {
		display: flex;
		position: absolute;
		top: 100%;
		left: 0;
		right: 0;
		background-color: white;
		border-bottom: 1px solid var(--border);
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
		z-index: 50;
		animation: slideDown 0.3s ease;
	}

	@keyframes slideDown {
		from {
			opacity: 0;
			transform: translateY(-10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	@media (min-width: 1024px) {
		.mobile-menu {
			display: none;
		}
	}

	.mobile-nav {
		display: flex;
		flex-direction: column;
		gap: 8px;
		padding: 16px;
		width: 100%;
	}

	.mobile-nav h3 {
		font-size: 0.9rem;
		font-weight: 700;
		color: var(--text);
		margin: 12px 0 8px 0;
		padding-bottom: 8px;
		border-bottom: 1px solid var(--border);
	}

	.mobile-nav a {
		display: flex;
		align-items: center;
		gap: 12px;
		padding: 12px 8px;
		text-decoration: none;
		color: var(--text);
		font-size: 0.95rem;
		border-radius: 6px;
		transition:
			background-color 0.2s ease,
			color 0.2s ease;
	}

	.mobile-nav a:hover {
		background-color: var(--accent-light);
		color: var(--accent);
	}

	.mobile-nav a span {
		flex: 1;
	}
</style>
