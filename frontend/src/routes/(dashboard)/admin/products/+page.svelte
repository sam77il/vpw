<script>
	import { enhance } from "$app/forms";
	import { goto } from "$app/navigation";
	import { resolve } from "$app/paths";
	import { X } from "@lucide/svelte";
	const { data, form } = $props();
	import { PUBLIC_FRONTEND_URL } from "$env/static/public";

	let creating = $state(false);
	let categories = $state([]);

	$effect(() => {
		if (creating) {
			async function loadCategories() {
				const res = await fetch(`${PUBLIC_FRONTEND_URL}/sapi/admin/categories`);
				const data = await res.json();
				if (data.success) {
					categories = data.categories;
				}
			}
			loadCategories();
		}
	});
</script>

<div class="admin">
	<div class="admin-content">
		<div>
			<div class="admin-content-page">
				<header>
					<h3>Produkte</h3>
					<div>
						<button onclick={() => (creating = !creating)}>Produkt erstellen</button>
						<input type="text" placeholder="Produkt suchen..." />
					</div>
				</header>

				{#if data.success}
					<table class="admin-content-page-list">
						<thead>
							<tr>
								<th>Nr.</th>
								<th>ID</th>
								<th>Label</th>
								<th>Preis (Aktuell, Alt)</th>
								<th>Lagerzustand</th>
								<th>Kategorie</th>
							</tr>
						</thead>
						<tbody>
							{#each data.products as product, i (product.id)}
								<tr
									class="admin-content-page-list-item"
									onclick={() => goto(resolve(`/admin/products/${product.id}`))}
								>
									<td>{i + 1}</td>
									<td>{product.id}</td>
									<td>{product.label}</td>
									<td>{product.price} | {product.old_price}</td>
									<td>{product.stock}</td>
									<td
										><a href={resolve(`/admin/categories/${product.category_id}`)}
											>{product.category_id}</a
										></td
									>
								</tr>
							{/each}
						</tbody>
					</table>
				{:else}
					<p>Fehler beim Laden der Produkte: {data.message}</p>
				{/if}
			</div>
		</div>
	</div>
</div>

{#if creating}
	<div class="creating">
		<div class="creating-content">
			<header>
				<h3>Kategorie erstellen</h3>
				<button onclick={() => (creating = false)}><X /></button>
			</header>

			<form
				method="POST"
				action="?/create"
				use:enhance={() => {
					return async ({ update }) => {
						await update({ reset: false });
					};
				}}
			>
				<input type="text" name="id" placeholder="ID" />
				<input type="text" name="label" placeholder="Label" />
				<textarea name="description" placeholder="Beschreibung"></textarea>
				<input name="price" placeholder="Preis" />
				<input name="old_price" placeholder="Alter Preis (Rabatt?)" />
				<input name="stock" placeholder="Lagerzustand" type="number" />
				<select name="category_id">
					<option>Kategorie auswählen</option>
					{#each categories as category (category.id)}
						<option value={category.id}>{category.label}</option>
					{/each}
				</select>
				{#if form?.success}
					<p style="color: green;">Produkt erfolgreich hinzugefügt</p>
				{:else}
					<p style="color: red;">{form?.message}</p>
				{/if}
				<button type="submit">Erstellen</button>
			</form>
		</div>
	</div>
{/if}

<style>
	.creating {
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: rgba(0, 0, 0, 0.5);
		position: fixed;
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
	}

	.admin {
		display: flex;
		justify-content: center;
		height: 100vh;
		width: 100%;
	}

	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
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

	.admin-content-page-list {
		width: 100%;
	}

	.admin-content-page-list-item {
		border-top: 1px solid black;
	}

	thead {
		background-color: orange;
		color: white;
	}

	.admin-content-page-list-item {
		cursor: pointer;
		background-color: white;
	}

	.admin-content-page-list-item td {
		padding: 10px;
	}

	.admin-content-page-list-item td a {
		color: black;
	}

	.admin-content-page-list-item:hover {
		background-color: lightgray;
	}

	header input {
		padding: 10px;
		border-radius: 10px;
		border: 1px solid lightgray;
		outline: none;
		width: 250px;
	}

	header button {
		padding: 10px;
		border-radius: 10px;
		background-color: rgb(1, 172, 86);
		border: none;
		color: white;
		outline: none;
		width: max-content;
		cursor: pointer;
	}

	.creating-content {
		background-color: whitesmoke;
		padding: 20px;
		border-radius: 10px;
		width: 400px;
	}

	.creating-content header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 20px;
	}

	.creating-content header > button {
		background-color: red;
	}

	.creating-content form {
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.creating-content form input {
		padding: 10px;
		border-radius: 5px;
		background-color: white;
		outline: none;
		border: 1px solid lightgray;
	}

	.creating-content form select {
		padding: 10px;
		border-radius: 5px;
		background-color: white;
		outline: none;
		border: 1px solid lightgray;
	}

	.creating-content form textarea {
		padding: 10px;
		border-radius: 5px;
		background-color: white;
		outline: none;
		border: 1px solid lightgray;
	}

	.creating-content form button {
		padding: 10px;
		border-radius: 5px;
		background-color: orange;
		color: white;
		border: none;
		cursor: pointer;
		width: max-content;
		font-size: 1rem;
		align-self: flex-end;
	}
</style>
