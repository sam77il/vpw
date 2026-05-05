<script>
	import { resolve } from "$app/paths";

	const { data } = $props();
</script>

<div class="category">
	<div class="category-content">
		<header>{data.categoryId}</header>

		{#if data.products.length > 0}
			<div class="category-content-items">
				{#each data.products as product (product.id)}
					<a class="category-content-items-item" href={resolve(`/products/${product.id}`)}>
						<!-- svelte-ignore a11y_img_redundant_alt -->
						<img src="/imgs/logo.png" alt="product image" />

						<h3>{product.label}</h3>
						<p>Preis: {product.price}€</p>
					</a>
				{/each}
			</div>
		{:else}
			<div class="category-content-items-notfound">
				<p>Keine Produkte gefunden</p>
			</div>
		{/if}
	</div>
</div>

<style>
	.category-content-items-notfound {
		text-align: center;
		font-size: 2rem;
	}

	.category {
		width: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		margin-top: 50px;
	}

	.category-content {
		max-width: 70%;
		width: 100%;
	}

	.category-content header {
		width: 100%;
		padding: 20px;
		text-align: center;
		color: white;
		background-color: orange;
		font-size: 2rem;
		text-transform: uppercase;
	}

	.category-content-items {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr;
		grid-template-rows: auto;
		gap: 25px;
		margin-top: 25px;
	}

	.category-content-items-item {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		border: 1px solid rgb(173, 173, 173);
		text-decoration: none;
		color: black;
		opacity: 0.8;
		transition: all 0.1s;
		width: 100%;
		height: 300px;
		padding: 15px;
	}

	.category-content-items-item:hover {
		opacity: 1;
	}

	.category-content-items-item:hover img {
		transform: scale(1.1);
	}

	.category-content-items-item img {
		width: 250px;
		height: 250px;
		object-fit: contain;
		z-index: -1;
		transition: transform 0.3s ease; /* <-- DAS fehlt */
	}

	.category-content-items-item h3 {
		color: black;
		text-align: center;
	}
</style>
