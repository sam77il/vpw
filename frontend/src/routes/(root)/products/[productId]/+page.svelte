<script>
	const { data } = $props();
	import { PUBLIC_FRONTEND_URL } from "$env/static/public";

	let productAmount = $state(1);
	let productPrice = $derived(
		data.product?.price * productAmount + (zuschnittState === "with" ? productAmount * 5.95 : 0)
	);
	let zuschnittState = $state("without");
	let typState = $state("matt");
	let colorState = $state("weiss");
	let sizeState = $state("");
	let resultText = $state("");

	$effect(() => {
		if (resultText) {
			const timeout = setTimeout(() => {
				resultText = "";
			}, 3000);
			return () => clearTimeout(timeout);
		}
	});

	function updateProductAmount(action) {
		if (action === "decrease") {
			if (productAmount - 1 <= 0) {
				return;
			}
			productAmount--;
		} else if (action === "increase") {
			productAmount++;
		}
	}

	async function addToCart() {
		const cartProduct = {
			id: data.product.id,
			price: productPrice,
			metadata: {
				size: sizeState,
				zuschnitt: zuschnittState,
				typ: typState,
				color: colorState
			},
			amount: productAmount
		};

		const res = await fetch(`/sapi/cart`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify(cartProduct)
		});
		const result = await res.json();
		if (result.success) {
			resultText = "Produkt erfolgreich zum Warenkorb hinzugefügt!";
		} else {
			resultText = "Fehler beim Hinzufügen des Produkts zum Warenkorb.";
		}
	}
</script>

{#if data.product?.id}
	<div class="product">
		<div class="product-content">
			<div class="product-top">
				<div class="product-image">
					<img src="/imgs/logo.png" alt="product picturee" />
				</div>
				<div class="product-details">
					<h2>{data.product.label}</h2>
					<h3>{productPrice.toFixed(2)} €</h3>
					<div>
						<p class="smallprice">Nettopreis: {(productPrice * 0.81).toFixed(2)} €</p>
						<p class="smallprice">Mehrwertsteuer {(productPrice * 0.19).toFixed(2)} €</p>
					</div>

					{#if data.product.options.product_type === "aluverbundplatte"}
						<div class="product-details-aluverbundplatte">
							<header>
								<h4>Zuschnitt:</h4>
								<h4>+ 5,95 €</h4>
							</header>
							<div class="product-details-aluverbundplatte-content">
								<div class="product-details-aluverbundplatte-zuschnitt">
									<input
										id="without"
										type="radio"
										name="zuschnitt"
										value="without"
										bind:group={zuschnittState}
										checked
									/>
									<label for="without">Ohne Zuschnitt</label>
								</div>
								<div class="product-details-aluverbundplatte-zuschnitt">
									<input
										id="with"
										type="radio"
										name="zuschnitt"
										value="with"
										bind:group={zuschnittState}
									/>
									<label for="with">Mit Zuschnitt</label>
								</div>
							</div>
						</div>
						<div class="product-details-aluverbundplatte">
							<header>
								<h4>Typ:</h4>
							</header>
							<div class="product-details-aluverbundplatte-content">
								<div class="product-details-aluverbundplatte-typ">
									<input
										id="matt"
										type="radio"
										name="typ"
										value="matt"
										bind:group={typState}
										checked
									/>
									<label for="matt">Matt</label>
								</div>
								<div class="product-details-aluverbundplatte-typ">
									<input id="glanz" type="radio" name="typ" value="glanz" bind:group={typState} />
									<label for="glanz">Glanz</label>
								</div>
							</div>
						</div>
						<div class="product-details-aluverbundplatte">
							<header>
								<h4>Farbe:</h4>
							</header>
							<div class="product-details-aluverbundplatte-content">
								<div class="product-details-aluverbundplatte-color">
									<input
										id="weiss"
										type="radio"
										name="color"
										value="weiss"
										bind:group={colorState}
										checked
									/>
									<label for="weiss">Weiß</label>
								</div>
								<div class="product-details-aluverbundplatte-color">
									<input
										id="schwarz"
										type="radio"
										name="color"
										value="schwarz"
										bind:group={colorState}
									/>
									<label for="schwarz">Schwarz</label>
								</div>
							</div>
						</div>
						<div class="product-details-aluverbundplatte-size">
							<p>Ideale Schreibweise: z.B. 3 Strk. 100 x 1100 mm</p>
							<textarea bind:value={sizeState}></textarea>
						</div>
					{:else if data.product.options.product_type === "digitaldruck"}
						<p>Digitaldruck</p>
					{/if}
					<div class="product-detail-amount">
						<button onclick={() => updateProductAmount("decrease")}>-</button>
						<input type="number" bind:value={productAmount} />
						<button onclick={() => updateProductAmount("increase")}>+</button>
					</div>
					<button class="button" onclick={addToCart}>In den Warenkorb</button>
					{#if resultText}
						<p>{resultText}</p>
					{/if}
				</div>
			</div>
			<div class="product-bottom"></div>
		</div>
	</div>
{:else}
	<div class="product-notfound">
		<p>Produkt nicht gefunden</p>
	</div>
{/if}

<style>
	.product-notfound {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		height: 50vh;
	}

	.product-notfound p {
		font-size: clamp(1.2rem, 3vw, 1.5rem);
		font-weight: bolder;
	}

	.product {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		padding: clamp(16px, 5vw, 50px);
		min-height: calc(100vh - 120px);
	}

	.product-content {
		max-width: 100%;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: clamp(32px, 5vw, 50px);
	}

	.product-top {
		display: flex;
		flex-direction: column;
		gap: clamp(24px, 5vw, 50px);
		width: 100%;
	}

	@media (min-width: 768px) {
		.product-top {
			flex-direction: row;
		}
	}

	.product-image {
		flex: 0 0 auto;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.product-image img {
		width: 100%;
		max-width: 500px;
		height: auto;
		object-fit: contain;
	}

	.product-details {
		flex: 1;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: clamp(16px, 3vw, 20px);
		background-color: var(--bg-light);
		padding: clamp(16px, 4vw, 24px);
		border-radius: 8px;
	}

	.product-details h2 {
		font-size: clamp(1.5rem, 4vw, 2rem);
		font-weight: bold;
		color: var(--text);
		margin: 0;
	}

	.product-details h3 {
		font-size: clamp(1.3rem, 3vw, 2rem);
		color: var(--accent);
		margin: 0;
	}

	.smallprice {
		color: var(--text-light);
		font-size: clamp(0.8rem, 1.5vw, 0.9rem);
		margin: 0;
	}

	/* Product configuration sections */
	.product-details-aluverbundplatte {
		display: flex;
		flex-direction: column;
		gap: 8px;
		padding: clamp(12px, 2vw, 16px);
		background-color: white;
		border-radius: 6px;
		border: 1px solid var(--border);
	}

	.product-details-aluverbundplatte header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 8px;
	}

	.product-details-aluverbundplatte header h4 {
		font-size: clamp(0.95rem, 1.5vw, 1.1rem);
		font-weight: 600;
		color: var(--text);
		margin: 0;
	}

	.product-details-aluverbundplatte-content {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}

	.product-details-aluverbundplatte-zuschnitt,
	.product-details-aluverbundplatte-typ,
	.product-details-aluverbundplatte-color {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 8px;
	}

	.product-details-aluverbundplatte input[type="radio"] {
		cursor: pointer;
		width: 18px;
		height: 18px;
		accent-color: var(--accent);
	}

	.product-details-aluverbundplatte label {
		cursor: pointer;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		user-select: none;
	}

	.product-details-aluverbundplatte-size {
		display: flex;
		flex-direction: column;
		padding: clamp(12px, 2vw, 16px);
		border: 2px solid var(--border);
		border-radius: 6px;
		background-color: white;
		gap: 8px;
	}

	.product-details-aluverbundplatte-size p {
		margin: 0;
		font-size: clamp(0.85rem, 1.5vw, 0.95rem);
		color: var(--text-light);
	}

	.product-details-aluverbundplatte-size textarea {
		resize: vertical;
		min-height: clamp(75px, 20vw, 120px);
		padding: clamp(8px, 1.5vw, 12px);
		border: 1px solid var(--border);
		border-radius: 4px;
		font-family: inherit;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		outline: none;
		transition: border-color 0.2s ease;
	}

	.product-details-aluverbundplatte-size textarea:focus {
		border-color: var(--accent);
		box-shadow: 0 0 0 3px rgba(249, 115, 22, 0.1);
	}

	/* Quantity controls */
	.product-detail-amount {
		display: flex;
		flex-direction: row;
		gap: 0;
		align-items: center;
	}

	.product-detail-amount button {
		background-color: var(--accent);
		border: none;
		outline: none;
		padding: clamp(8px, 2vw, 12px);
		cursor: pointer;
		color: white;
		font-size: clamp(1rem, 2vw, 1.2rem);
		font-weight: bold;
		width: clamp(40px, 10vw, 50px);
		min-height: 44px;
		transition: background-color 0.2s ease;
	}

	.product-detail-amount button:hover {
		background-color: var(--accent-hover);
	}

	.product-detail-amount button:first-child {
		border-top-left-radius: 6px;
		border-bottom-left-radius: 6px;
	}

	.product-detail-amount button:last-child {
		border-top-right-radius: 6px;
		border-bottom-right-radius: 6px;
	}

	.product-detail-amount input {
		border: none;
		border-radius: 0;
		outline: none;
		background-color: white;
		padding: clamp(8px, 2vw, 12px);
		text-align: center;
		flex: 1;
		min-width: 50px;
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		border-top: 1px solid var(--border);
		border-bottom: 1px solid var(--border);
		-webkit-appearance: none;
		appearance: none;
	}

	.product-detail-amount input:focus {
		outline: 2px solid var(--accent);
		outline-offset: -2px;
	}

	/* Add to cart button */
	.button {
		background-color: var(--accent);
		padding: clamp(12px, 2.5vw, 16px) clamp(16px, 3vw, 24px);
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		border: none;
		outline: none;
		color: white;
		border-radius: 6px;
		cursor: pointer;
		font-weight: 600;
		width: 100%;
		min-height: 48px;
		transition:
			background-color 0.2s ease,
			transform 0.1s ease;
	}

	.button:hover {
		background-color: var(--accent-hover);
	}

	.button:active {
		transform: scale(0.98);
	}

	/* Result message */
	.product-details p {
		margin: 0;
		text-align: center;
		padding: 8px;
		border-radius: 4px;
	}

	.product-details p:not(.smallprice) {
		background-color: var(--accent-light);
		color: var(--accent);
		font-weight: 500;
	}

	@media (max-width: 639px) {
		.product-image {
			max-height: 300px;
		}

		.product-details {
			min-height: auto;
		}
	}
</style>
