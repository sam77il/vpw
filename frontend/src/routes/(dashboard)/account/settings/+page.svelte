<script>
	import { enhance } from "$app/forms";
	import { AtSign, Key, ShoppingBag, Truck, User } from "@lucide/svelte";

	const { data, form } = $props();
	let currentPage = $state("personal");
</script>

<div class="settings">
	<div class="settings-content">
		<header>
			<h2>Einstellungen</h2>
		</header>

		<div>
			<div class="left">
				<nav class="settings-content-nav">
					<div class="divider">
						<p>Persönliche Einstellungen</p>
					</div>
					<button onclick={() => (currentPage = "personal")}
						><User size={20} /><span>Persönliche Daten</span></button
					>
					<button onclick={() => (currentPage = "email")}
						><AtSign size={20} /><span>E-Mail Einstellungen</span></button
					>
					<button onclick={() => (currentPage = "password")}
						><Key size={20} /><span>Passwort ändern</span></button
					>
					<button onclick={() => (currentPage = "delivery")}
						><Truck size={20} /><span>Lieferung</span></button
					>
				</nav>

				{#if data.user.company}
					<nav class="settings-content-nav">
						<div class="divider">
							<p>Firma Einstellungen</p>
						</div>
						<button onclick={() => (currentPage = "companyinfo")}
							><ShoppingBag size={20} /><span>Firmeninformationen</span></button
						>
					</nav>
				{/if}
			</div>
			{#if currentPage === "personal"}
				<div class="settings-content-page">
					<header>
						<h3>Persönliche Daten</h3>
					</header>
					<form
						method="post"
						action="?/personal"
						use:enhance={() => {
							return async ({ update }) => {
								await update({ reset: false });
							};
						}}
					>
						<div>
							<label for="first_name">Vorname</label>
							<input id="first_name" type="text" name="first_name" value={data.user.first_name} />
						</div>
						<div>
							<label for="last_name">Nachname</label>
							<input id="last_name" type="text" name="last_name" value={data.user.last_name} />
						</div>
						<div>
							<label for="phone_number">Telefonnummer</label>
							<input
								id="phone_number"
								type="text"
								name="phone_number"
								value={data.user.phone_number}
							/>
						</div>
						{#if form?.success}
							<p style="color: green">Erfolgreich gespeichert.</p>
						{:else if !form?.success}
							<p style="color: red">{form?.message}</p>
						{/if}
						<input type="submit" value="Speichern" />
					</form>
				</div>
			{:else if currentPage === "email"}
				<div class="settings-content-page">
					<header>
						<h3>E-Mail Einstellungen</h3>
					</header>
					<form
						method="post"
						action="?/email"
						use:enhance={() => {
							return async ({ update }) => {
								await update({ reset: false });
							};
						}}
					>
						<div>
							<label for="email">E-Mail</label>
							<input id="email" type="text" name="email" value={data.user.email} />
						</div>
						{#if form?.success}
							<p style="color: green">Erfolgreich gespeichert.</p>
						{:else if !form?.success}
							<p style="color: red">{form?.message}</p>
						{/if}
						<input type="submit" value="Speichern" />
					</form>
				</div>
			{:else if currentPage === "password"}
				<div class="settings-content-page">
					<header>
						<h3>Passwort ändern</h3>
					</header>
					<form
						method="post"
						action="?/password"
						use:enhance={() => {
							return async ({ update }) => {
								await update({ reset: false });
							};
						}}
					>
						<div>
							<label for="current_password">Aktuelles Passwort</label>
							<input id="current_password" name="current_password" type="password" />
						</div>
						<div>
							<label for="new_password">Neues Passwort</label>
							<input id="new_password" name="new_password" type="password" />
						</div>
						<div>
							<label for="new_password_confirm">Neues Passwort bestätigen</label>
							<input id="new_password_confirm" name="new_password_confirm" type="password" />
						</div>
						{#if form?.success}
							<p style="color: green">Erfolgreich gespeichert.</p>
						{:else if !form?.success}
							<p style="color: red">{form?.message}</p>
						{/if}
						<input type="submit" value="Speichern" />
					</form>
				</div>
			{:else if currentPage === "delivery"}
				<div class="settings-content-page">
					<header>
						<h3>Lieferung</h3>
					</header>
					<form
						method="post"
						action="?/delivery"
						use:enhance={() => {
							return async ({ update }) => {
								await update({ reset: false });
							};
						}}
					>
						<div>
							<label for="country">Land</label>
							<select id="country" name="country" value={data.user.country}>
								<option value="DE">Deutschland</option>
							</select>
						</div>
						<div>
							<label for="street">Straße / Nr</label>
							<input id="street" type="text" name="street" value={data.user.street} />
						</div>
						<div>
							<label for="postal_code">Postleitzahl</label>
							<input
								id="postal_code"
								type="text"
								name="postal_code"
								value={data.user.postal_code}
							/>
						</div>
						<div>
							<label for="city">Stadt</label>
							<input id="city" type="text" name="city" value={data.user.city} />
						</div>
						{#if form?.success}
							<p style="color: green">Erfolgreich gespeichert.</p>
						{:else if !form?.success}
							<p style="color: red">{form?.message}</p>
						{/if}
						<input type="submit" value="Speichern" />
					</form>
				</div>
			{:else if currentPage === "companyinfo"}
				<div class="settings-content-page">
					<header>
						<h3>Firmeninformationen</h3>
					</header>
					<form
						method="post"
						action="?/companyinfo"
						use:enhance={() => {
							return async ({ update }) => {
								await update({ reset: false });
							};
						}}
					>
						<div>
							<label for="company_name">Firmenname</label>
							<input
								id="company_name"
								type="text"
								name="company_name"
								value={data.user.company_name}
							/>
						</div>
						<div>
							<label for="company_ustidnr">Firmen-UstId Nr.</label>
							<input
								id="company_ustidnr"
								type="text"
								name="company_ustidnr"
								value={data.user.company_ustidnr}
							/>
						</div>
						{#if form?.success}
							<p style="color: green">Erfolgreich gespeichert.</p>
						{:else if !form?.success}
							<p style="color: red">{form?.message}</p>
						{/if}
						<input type="submit" value="Speichern" />
					</form>
				</div>
			{/if}
		</div>
	</div>
</div>

<style>
	.left {
		display: flex;
		flex-direction: column;
		gap: 25px;
	}
	.settings {
		display: flex;
		justify-content: center;
		height: 100vh;
		width: 100%;
	}

	.settings-content {
		width: 50vw;
		height: 50%;
		margin-top: 50px;
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	.settings-content > div {
		display: flex;
		flex-direction: row;
		gap: 50px;
	}

	.settings-content-nav {
		background-color: whitesmoke;
		border-radius: 10px;
		width: 300px;
		display: flex;
		flex-direction: column;
		gap: 10px;
		padding: 20px;
		height: max-content;
	}

	.settings-content-page {
		background-color: whitesmoke;
		border-radius: 10px;
		width: 60%;
		display: flex;
		flex-direction: column;
		gap: 10px;
		padding: 20px;
		height: max-content;
	}

	.settings-content-page > form {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}

	.settings-content-page > form > input[type="submit"] {
		align-self: flex-end;
		width: max-content;
		background-color: orange;
		outline: none;
		border: none;
		border-radius: 10px;
		padding: 10px 20px;
		color: white;
		cursor: pointer;
	}

	.settings-content-page > form > div {
		display: flex;
		flex-direction: column;
		gap: 2px;
	}

	.settings-content-page > form > div > label {
		font-size: 0.9rem;
	}

	.settings-content-page > form > div > input {
		background-color: white;
		border-radius: 10px;
		padding: 10px;
		border: 1px solid lightgray;
		font-size: 1rem;
		outline: none;
	}

	.settings-content-page > form > div > select {
		background-color: white;
		border-radius: 10px;
		padding: 10px;
		border: 1px solid lightgray;
		font-size: 1rem;
		outline: none;
	}

	.settings-content-page > form > div > input:focus {
		border-color: orange;
	}

	.settings-content-nav button {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 5px;
		color: black;
		text-decoration: none;
		width: 100%;
		padding: 10px;
		border-radius: 5px;
		transition: background-color 0.1s;
		font-size: 0.9rem;
		outline: none;
		border: none;
		cursor: pointer;
		background-color: transparent;
	}

	.settings-content-nav button:hover {
		background-color: white;
	}

	.divider {
		border-bottom: 1px solid lightgray;
		padding: 5px 0;
	}

	.divider p {
		text-transform: uppercase;
		font-size: 0.7rem;
		font-weight: bolder;
	}
</style>
