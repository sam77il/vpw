<script>
	import { Trash } from "@lucide/svelte";

	const { data } = $props();

	function changeAmount(action, cartId) {
		fetch(`/api/cart/${cartId}/amount`, {
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
		fetch(`/api/cart/${cartId}`, {
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
		margin-top: 50px;
	}

	.cart-content {
		max-width: 70vw;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	.cart-content header {
		padding: 0 0 20px 0;
		border-bottom: 1px solid lightgray;
	}

	.cart-list {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.cart-item {
		display: flex;
		flex-direction: row;
		gap: 20px;
		background-color: whitesmoke;
		padding: 20px;
		border-radius: 15px;
		position: relative;
	}

	.cart-item-image img {
		width: 250px;
		height: 150px;
		object-fit: cover;
	}

	.cart-item-details {
		display: flex;
		flex-direction: column;
		gap: 0;
	}

	.cart-item-actions {
		position: absolute;
		right: 20px;
		bottom: 20px;
		display: flex;
		flex-direction: row;
		gap: 10px;
		margin-top: 10px;
	}

	.quantity-button {
		padding: 5px 10px;
		border: none;
		background-color: orange;
		color: white;
		cursor: pointer;
	}

	.quantity-button:first-child {
		border-top-left-radius: 5px;
		border-bottom-left-radius: 5px;
	}

	.quantity-button:last-child {
		border-top-right-radius: 5px;
		border-bottom-right-radius: 5px;
	}

	.cart-item-actions div {
		display: flex;
		flex-direction: row;
		gap: 0;
	}

	.cart-item-actions p {
		width: 50px;
		text-align: center;
		padding: 5px;
		border: none;
		outline: none;
		font-size: 1rem;
		text-align: center;
		background-color: white;
	}

	.remove-button {
		background-color: transparent;
		border: none;
		outline: none;
		cursor: pointer;
	}

	.remove-button:hover {
		color: red;
	}
</style>
