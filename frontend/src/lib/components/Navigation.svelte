<script>
	import { resolve } from "$app/paths";
	// import { ArrowRight, X } from "@lucide/svelte";
	// import ChevronRight from "@lucide/svelte/icons/chevron-right";

	const { categories } = $props();
	// let hoveredNav = $state({});

	// function changeHoveredNav(newNav) {
	// 	hoveredNav = newNav;
	// }
</script>

<nav>
	{#each categories as category (category.id)}
		<a href={resolve(`/categories/${category.id}`)}>{category.label}</a>
	{/each}
</nav>

<!-- {#if hoveredNav.id}
	<div class="nav-hover">
		<div class="nav-hover-content">
			<header>
				<a href={resolve(`/categories/${hoveredNav.id}`)}
					>Zur Kategorie {hoveredNav.label} <ArrowRight /></a
				>
				<button onclick={() => changeHoveredNav({})}><X /></button>
			</header>
			<div>
				{#each hoveredNav.children as hovChild, i (i)}
					<a href={resolve(hovChild.href)}><ChevronRight /> {hovChild.label}</a>
				{/each}
			</div>
		</div>
	</div>
{/if} -->

<style>
	/* .nav-hover {
		position: absolute;
		width: 100%;
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
	}

	.nav-hover-content {
		width: 50%;
		background-color: whitesmoke;
		padding: 20px;
	}

	.nav-hover-content header {
		width: 100%;
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	.nav-hover-content header a {
		text-decoration: none;
		color: black;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: transform 0.1s;
	}

	.nav-hover-content header a:hover {
		transform: translateX(5px);
	}

	.nav-hover-content header button {
		background-color: transparent;
		border: none;
		outline: none;
		cursor: pointer;
	}

	.nav-hover-content div {
		display: flex;
		flex-direction: column;
		flex-wrap: wrap;
		max-height: 500px;
	} */

	/* .nav-hover-content div > a {
		color: black;
		display: flex;
		align-items: center;
		text-decoration: none;
		width: max-content;
	}

	.nav-hover-content div > a:hover {
		color: orange;
	} */

	nav {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		background-color: orange;
	}

	nav a {
		text-transform: uppercase;
		padding: 10px 15px;
		text-decoration: none;
		color: white;
	}

	nav a:hover {
		background-color: white;
		color: orange;
	}
</style>
