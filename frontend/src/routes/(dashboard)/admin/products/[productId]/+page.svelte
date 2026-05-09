<script>
	import { enhance } from "$app/forms";
	import { goto } from "$app/navigation";
	import { resolve } from "$app/paths";
	import { ArrowLeft } from "@lucide/svelte";
	const { data, form } = $props();

	async function deleteProduct() {
		const res = await fetch("/api/admin/products/" + data.product.id, {
			method: "DELETE"
		});

		if (res.ok) {
			goto(resolve("/admin/products"));
		}
	}

	$effect(() => {
		console.log(data.product.options);
	});
</script>

<svelte:head>
	<title>{data.product.label} - Admin</title>
</svelte:head>

<div class="admin">
	<div class="admin-content">
		<div>
			<div class="admin-content-page">
				<header>
					<h3>
						<a href={resolve("/admin/products")}><ArrowLeft /></a>
						{data.product.label}
					</h3>
				</header>

				{#if data.success}
					<div class="admin-content-user-content">
						<form
							action="?/update"
							method="POST"
							use:enhance={() => {
								return async ({ update }) => {
									await update({ reset: false });
								};
							}}
						>
							<div>
								<header>
									<h3>Produkt Daten</h3>
								</header>

								<label for="id">ID (einzigartig)</label>
								<input type="text" placeholder="ID" name="id" value={data.product.id} />

								<label for="label">Label</label>
								<input
									id="label"
									type="text"
									placeholder="Label"
									name="label"
									value={data.product.label}
								/>

								<label for="description">Beschreibung</label>
								<textarea id="description" type="text" placeholder="Beschreibung" name="description"
									>{data.product.description}</textarea
								>

								<label for="price">Preis</label>
								<input
									id="price"
									type="text"
									placeholder="Preis"
									name="price"
									value={data.product.price}
								/>

								<label for="old_price">Alter Preis (Rabatt?)</label>
								<input
									id="old_price"
									type="text"
									placeholder="Alter Preis"
									name="old_price"
									value={data.product.old_price}
								/>

								<label for="category_id">Alter Preis (Rabatt?)</label>
								<select id="category_id" name="category_id" value={data.product.category_id}>
									{#each data.categories as category (category.id)}
										<option value={category.id}>{category.label}</option>
									{/each}
								</select>

								<label for="stock">Lagerzustand</label>
								<input
									id="stock"
									type="text"
									placeholder="Alter Preis"
									name="stock"
									value={data.product.stock}
								/>
							</div>
							{#if form?.success}
								<p style="color: green;">Produkt erfolgreich aktualisiert</p>
							{:else}
								<p style="color: red;">{form?.message}</p>
							{/if}
							<div>
								<button type="button" onclick={deleteProduct} class="delete">Produkt löschen</button
								>
								<input type="submit" value="Änderungen speichern" />
							</div>
						</form>
					</div>
				{:else}
					<p>Fehler beim Laden des Produkts: {data.message}</p>
				{/if}
			</div>
		</div>
	</div>
</div>

<style>
	.admin {
		display: flex;
		justify-content: center;
		height: 100vh;
		width: 100%;
		margin-bottom: 100px;
	}

	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	header h3 {
		display: flex;
		align-items: center;
		gap: 10px;
	}

	header h3 a {
		color: black;
	}

	.admin-content {
		width: 50vw;
		height: 50%;
		margin-top: 50px;
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	.admin-content > div {
		display: flex;
		flex-direction: row;
		gap: 50px;
	}

	.admin-content-page {
		background-color: whitesmoke;
		border-radius: 10px;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 10px;
		padding: 20px;
		height: max-content;
	}

	.admin-content-user-content > form {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.admin-content-user-content > form > div {
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.admin-content-user-content > form > div > header {
		border-bottom: 1px solid lightgray;
		padding-bottom: 5px;
	}

	.admin-content-user-content > form > div > input {
		padding: 10px;
		border-radius: 5px;
		background-color: white;
		outline: none;
		border: 1px solid lightgray;
	}

	.admin-content-user-content > form > div > textarea {
		padding: 10px;
		border-radius: 5px;
		background-color: white;
		outline: none;
		border: 1px solid lightgray;
	}

	.admin-content-user-content > form > div > select {
		padding: 10px;
		border-radius: 5px;
		background-color: white;
		outline: none;
		border: 1px solid lightgray;
	}

	.admin-content-user-content > form > div:last-child {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}

	.admin-content-user-content > form > div > input[type="submit"] {
		padding: 15px;
		border-radius: 5px;
		background-color: orange;
		color: white;
		border: none;
		cursor: pointer;
		width: max-content;
		font-size: 1.2rem;
		align-self: flex-end;
	}

	.admin-content-user-content > form > div > button {
		padding: 15px;
		border-radius: 5px;
		background-color: red;
		color: white;
		border: none;
		cursor: pointer;
		width: max-content;
		font-size: 1.2rem;
		align-self: flex-end;
	}
</style>
