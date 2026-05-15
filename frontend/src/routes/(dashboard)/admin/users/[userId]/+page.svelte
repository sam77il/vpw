<script>
	import { enhance } from "$app/forms";
	import { goto } from "$app/navigation";
	import { resolve } from "$app/paths";
	import { ArrowLeft } from "@lucide/svelte";
	const { data, form } = $props();
	let company = $derived(data.user.company);
	import { PUBLIC_FRONTEND_URL } from "$env/static/public";

	async function closeAccount() {
		const res = await fetch(`${PUBLIC_FRONTEND_URL}/sapi/admin/users/` + data.user.id, {
			method: "DELETE"
		});

		if (res.ok) {
			goto(resolve("/admin/users"));
		}
	}
</script>

<svelte:head>
	<title>{data.user.first_name} {data.user.last_name} - Admin</title>
</svelte:head>

<div class="admin">
	<div class="admin-content">
		<div>
			<div class="admin-content-page">
				<header>
					<h3>
						<a href={resolve("/admin/users")}><ArrowLeft /></a>
						{data.user.last_name} - {data.user.id}
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
									<h3>Persönliche Informationen</h3>
								</header>

								<label for="id">ID</label>
								<input id="id" name="id" value={data.user.id} />

								<label for="email">E-Mail</label>
								<input
									id="email"
									type="text"
									placeholder="E-Mail"
									name="email"
									value={data.user.email}
								/>

								<label for="first_name">Vorname</label>
								<input
									id="first_name"
									type="text"
									placeholder="Vorname"
									name="first_name"
									value={data.user.first_name}
								/>

								<label for="last_name">Nachname</label>
								<input
									id="last_name"
									type="text"
									placeholder="Nachname"
									name="last_name"
									value={data.user.last_name}
								/>

								<label for="role">Rolle</label>
								<select name="role" id="role" value={data.user.role}>
									<option value="user">Kunde</option>
									<option value="admin">Admin</option>
								</select>

								<label for="gender">Geschlecht</label>
								<select name="gender" id="gender" value={data.user.gender}>
									<option value="male">Männlich</option>
									<option value="female">Weiblich</option>
								</select>

								<label for="phone_number">Telefonnummer</label>
								<input
									id="phone_number"
									type="text"
									placeholder="Telefonnummer"
									name="phone_number"
									value={data.user.phone_number}
								/>
							</div>
							<div>
								<header>
									<h3>Adresse</h3>
								</header>

								<label for="street">Straße</label>
								<input
									id="street"
									type="text"
									placeholder="Straße"
									name="street"
									value={data.user.street}
								/>

								<label for="postal_code">Postleitzahl</label>
								<input
									type="text"
									placeholder="Postleitzahl"
									name="postal_code"
									value={data.user.postal_code}
								/>

								<label for="city">Stadt</label>
								<input
									id="city"
									type="text"
									placeholder="Stadt"
									name="city"
									value={data.user.city}
								/>

								<label for="country">Land</label>
								<input
									id="country"
									type="text"
									placeholder="Land"
									name="country"
									value={data.user.country}
								/>
							</div>
							<div>
								<header>
									<h3>Firmeninformationen</h3>
								</header>

								<label for="company">Gewerbetreibend</label>
								<select
									name="company"
									id="company"
									onchange={(e) => (company = e.target.value === "true" ? true : false)}
									value={company.toString()}
								>
									<option value="false">Nein</option>
									<option value="true">Ja</option>
								</select>
								{#if company}
									<label for="company_name">Firmenname</label>
									<input
										id="company_name"
										type="text"
										placeholder="Firmenname"
										name="company_name"
										value={data.user?.company_name}
									/>

									<label for="company_ustidnr">UstId Nr.</label>
									<input
										id="company_ustidnr"
										type="text"
										placeholder="USt-IdNr."
										name="company_ustidnr"
										value={data.user?.company_ustidnr}
									/>
								{/if}
							</div>
							{#if form?.success}
								<p style="color: green;">Nutzer erfolgreich aktualisiert</p>
							{:else}
								<p style="color: red;">{form?.message}</p>
							{/if}
							<div>
								<button type="button" onclick={closeAccount} class="delete">Konto schließen</button>
								<input type="submit" value="Änderungen speichern" />
							</div>
						</form>
					</div>
				{:else}
					<p>Fehler beim Laden der Nutzer: {data.message}</p>
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

	.admin-content-user-content > form > div > input,
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
