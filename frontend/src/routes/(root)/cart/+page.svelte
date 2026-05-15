<script>
	import { Trash } from "@lucide/svelte";
	import { PUBLIC_FRONTEND_URL } from "$env/static/public";

	const { data } = $props();

	function changeAmount(action, cartId) {
		fetch(`/sapi/cart/${cartId}/amount`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({ action })
		})
			.then((res) => res.json())
			.then((result) => {
				console.log(result);
				if (result.success) {
					const amountElement = document.querySelector(`[data-cart-amount="${cartId}"]`);
					if (amountElement) {
						amountElement.textContent = result.new_amount;
					}
				}
			});
	}

	function deleteCartItem(cartId) {
		fetch(`/sapi/cart/${cartId}`, {
			method: "DELETE"
		})
			.then((res) => res.json())
			.then((result) => {
				console.log(result);
				if (result.success) {
					const cartItemElement = document.querySelector(`[data-cart-item="${cartId}"]`);
					if (cartItemElement) {
						cartItemElement.remove();
					}
				}
			});
	}
</script>

<div class="cart">
	<div class="cart-content">
		<header>
			<h2>Warenkorb</h2>
		</header>

		{#if data.cart_items.length === 0}
			<p>Dein Warenkorb ist leer.</p>
		{:else}
			<div class="cart-list">
				{#each data.cart_items as cart (cart.id)}
					<div class="cart-item" data-cart-item={cart.id}>
						<div class="cart-item-image">
							<img src="/imgs/logo.png" alt="Produktbild" />
						</div>
						<div class="cart-item-details">
							<h3>{cart.label}</h3>
							{#if cart.metadata}
								{#each Object.entries(cart.metadata) as [key, value] (key)}
									<p>{key[0].toLocaleUpperCase() + key.slice(1)}: {value}</p>
								{/each}
							{/if}
							<p>Preis: {cart.price.toFixed(2)} €</p>
						</div>
						<div class="cart-item-actions">
							<button onclick={() => deleteCartItem(cart.id)} class="remove-button"
								><Trash /></button
							>
							<div>
								<button onclick={() => changeAmount("decrease", cart.id)} class="quantity-button"
									>-</button
								>
								<p data-cart-amount={cart.id}>{cart.amount}</p>
								<button onclick={() => changeAmount("increase", cart.id)} class="quantity-button"
									>+</button
								>
							</div>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>

<style>
	.cart {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		padding: clamp(16px, 5vw, 50px);
		min-height: calc(100vh - 120px);
	}

	.cart-content {
		max-width: 1200px;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: clamp(24px, 5vw, 50px);
	}

	.cart-content header {
		padding: 0 0 clamp(12px, 2vw, 20px) 0;
		border-bottom: 2px solid var(--border);
	}

	.cart-content header h2 {
		font-size: clamp(1.5rem, 4vw, 2rem);
		color: var(--text);
		margin: 0;
	}

	.cart-content > p {
		text-align: center;
		font-size: clamp(0.95rem, 1.5vw, 1.1rem);
		color: var(--text-light);
		padding: 32px 16px;
	}

	.cart-list {
		display: flex;
		flex-direction: column;
		gap: clamp(16px, 3vw, 20px);
	}

	.cart-item {
		display: flex;
		flex-direction: column;
		gap: clamp(16px, 3vw, 20px);
		background-color: var(--bg-light);
		padding: clamp(16px, 3vw, 20px);
		border-radius: 8px;
		border: 1px solid var(--border);
		position: relative;
		transition: box-shadow 0.2s ease;
	}

	@media (min-width: 768px) {
		.cart-item {
			flex-direction: row;
		}
	}

	.cart-item:hover {
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	}

	.cart-item-image {
		flex: 0 0 auto;
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 6px;
		background-color: white;
		padding: 8px;
		overflow: hidden;
	}

	.cart-item-image img {
		width: 100%;
		max-width: clamp(120px, 30vw, 200px);
		height: clamp(80px, 20vw, 150px);
		object-fit: cover;
		border-radius: 4px;
	}

	.cart-item-details {
		display: flex;
		flex-direction: column;
		gap: clamp(8px, 2vw, 12px);
		flex: 1;
		min-width: 0;
	}

	.cart-item-details h3 {
		font-size: clamp(1rem, 2vw, 1.3rem);
		font-weight: 600;
		color: var(--text);
		margin: 0;
		word-break: break-word;
	}

	.cart-item-details p {
		font-size: clamp(0.85rem, 1.5vw, 0.95rem);
		color: var(--text-light);
		margin: 0;
		line-height: 1.5;
	}

	.cart-item-actions {
		display: flex;
		flex-direction: column;
		gap: 12px;
		align-items: flex-start;
	}

	@media (min-width: 768px) {
		.cart-item-actions {
			position: absolute;
			right: clamp(12px, 3vw, 20px);
			bottom: clamp(12px, 3vw, 20px);
			align-items: flex-end;
		}
	}

	.cart-item-actions > div {
		display: flex;
		flex-direction: row;
		gap: 0;
		align-items: center;
	}

	.quantity-button {
		padding: clamp(6px, 1.5vw, 8px) clamp(10px, 2vw, 12px);
		border: none;
		background-color: var(--accent);
		color: white;
		cursor: pointer;
		font-weight: 600;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		min-height: 40px;
		transition: background-color 0.2s ease;
	}

	.quantity-button:hover {
		background-color: var(--accent-hover);
	}

	.quantity-button:first-child {
		border-top-left-radius: 6px;
		border-bottom-left-radius: 6px;
	}

	.quantity-button:last-child {
		border-top-right-radius: 6px;
		border-bottom-right-radius: 6px;
	}

	.cart-item-actions p {
		width: clamp(40px, 8vw, 60px);
		text-align: center;
		padding: clamp(6px, 1.5vw, 8px);
		border: 1px solid var(--border);
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		font-weight: 500;
		background-color: white;
		color: var(--text);
		margin: 0;
	}

	.remove-button {
		background-color: transparent;
		border: none;
		outline: none;
		cursor: pointer;
		padding: 8px;
		color: var(--text-light);
		transition: color 0.2s ease;
		display: flex;
		align-items: center;
		justify-content: center;
		border-radius: 4px;
	}

	.remove-button:hover {
		color: #dc2626;
		background-color: rgba(220, 38, 38, 0.1);
	}

	/* Mobile optimizations */
	@media (max-width: 767px) {
		.cart-item {
			padding-bottom: 120px;
		}

		.cart-item-actions {
			gap: 16px;
		}

		.quantity-button {
			min-width: 36px;
		}
	}

	@media (max-width: 639px) {
		.cart-item-image img {
			max-width: 100px;
			height: 100px;
		}
	}
</style>
