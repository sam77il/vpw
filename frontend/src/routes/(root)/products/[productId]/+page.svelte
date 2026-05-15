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

		const res = await fetch(`${PUBLIC_FRONTEND_URL}/sapi/cart`, {
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
		font-size: 1.5rem;
		font-weight: bolder;
	}

	.product-details-aluverbundplatte header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}

	.product-details-aluverbundplatte-zuschnitt {
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
		gap: 5px;
	}

	.product-details-aluverbundplatte-content {
		display: flex;
		flex-direction: column;
		gap: 5px;
	}

	.product-details-aluverbundplatte-size {
		display: flex;
		flex-direction: column;
		padding: 5px;
		border: 1px solid black;
	}

	.product-details-aluverbundplatte-size textarea {
		resize: none;
		height: 75px;
		padding: 5px;
	}

	.product {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		margin-top: 50px;
	}

	.product-content {
		max-width: 50vw;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	.product-top {
		display: flex;
		flex-direction: row;
		gap: 50px;
		width: 100%;
	}

	.product-image img {
		width: 500px;
	}

	.product-details {
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 20px;
		background-color: whitesmoke;
		padding: 20px;
	}

	.product-details h2 {
		font-size: 2rem;
		font-weight: bolder;
	}

	.product-details h3 {
		font-size: 2rem;
	}

	.smallprice {
		color: gray;
		font-size: 0.7rem;
	}

	.product-detail-amount {
		display: flex;
		flex-direction: row;
	}

	.product-detail-amount input {
		border: none;
		border-radius: 0;
		outline: none;
		background-color: rgb(235, 235, 235);
		padding: 10px;
		text-align: center;
		width: 150px;
		font-size: 1.2rem;
	}

	.product-detail-amount button {
		background-color: orange;
		border: none;
		outline: none;
		padding: 10px;
		cursor: pointer;
		color: white;
		font-size: 1.5rem;
	}

	.product-detail-amount button:first-child {
		border-top-left-radius: 5px;
		border-bottom-left-radius: 5px;
	}

	.product-detail-amount button:last-child {
		border-top-right-radius: 5px;
		border-bottom-right-radius: 5px;
	}

	.button {
		background-color: orange;
		padding: 10px 20px;
		font-size: 1rem;
		border: none;
		outline: none;
		color: white;
		border-radius: 5px;
		cursor: pointer;
	}
</style>
