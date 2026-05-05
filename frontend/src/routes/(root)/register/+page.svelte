<script>
	import { enhance } from "$app/forms";
	import { resolve } from "$app/paths";

	const { form } = $props();

	let isCompany = $state("false");
	let creating = $state(false);
</script>

<div class="register">
	<div class="register-content">
		<h2>Konto erstellen</h2>

		<form
			method="post"
			use:enhance={() => {
				creating = true;

				return async ({ update }) => {
					await update({ reset: false });
					creating = false;
				};
			}}
		>
			<div class="register-box">
				<header>
					<h3>Ihre persönlichen Daten</h3>
				</header>
				<div>
					<div>
						<p>Anrede*</p>
						<div>
							<div>
								<input type="radio" name="gender" id="male" value="male" />
								<label for="male">Herr</label>
							</div>
							<div>
								<input type="radio" name="gender" id="female" value="female" />
								<label for="female">Frau</label>
							</div>
						</div>
					</div>
					<div>
						<p>Vorname*</p>
						<input type="text" name="first_name" placeholder="Vorname" />
					</div>
					<div>
						<p>Nachname*</p>
						<input type="text" name="last_name" placeholder="Nachname" />
					</div>
					<div>
						<p>E-Mail*</p>
						<input type="email" name="email" placeholder="E-Mail" />
					</div>
					<div>
						<p>E-Mail bestätigen*</p>
						<input type="email" name="email_confirmation" placeholder="E-Mail bestätigen" />
					</div>
					<div>
						<p>Passwort*</p>
						<input type="password" name="password" placeholder="Passwort" />
					</div>
					<div>
						<p>Passwort bestätigen*</p>
						<input type="password" name="password_confirmation" placeholder="Passwort bestätigen" />
					</div>
				</div>
			</div>
			<div class="register-box">
				<header>
					<h3>Firmendaten</h3>
				</header>
				<div>
					<div>
						<p>Gewerbetreibend*</p>
						<div>
							<div>
								<input
									type="radio"
									name="company"
									id="company_yes"
									bind:group={isCompany}
									value="true"
								/>
								<label for="company_yes">Ja</label>
							</div>
							<div>
								<input
									type="radio"
									name="company"
									id="company_no"
									bind:group={isCompany}
									value="false"
								/>
								<label for="company_no">Nein</label>
							</div>
						</div>
					</div>
					{#if isCompany === "true"}
						<div>
							<p>Firmenname*</p>
							<input type="text" name="company_name" placeholder="Firmenname" />
						</div>
						<div>
							<p>USt-IdNr*</p>
							<input type="text" name="company_ustidnr" placeholder="USt-IdNr" />
						</div>
					{/if}
				</div>
			</div>
			<div class="register-box">
				<header>
					<h3>Ihre Adresse</h3>
				</header>
				<div>
					<div>
						<p>Straße/Nr.*</p>
						<input type="text" name="street" placeholder="Straße" />
					</div>
					<div>
						<p>Postleitzahl / Ort*</p>
						<div class="postal-city">
							<input type="text" name="postal_code" placeholder="Postleitzahl" />
							<input type="text" name="city" placeholder="Ort" />
						</div>
					</div>
					<div>
						<p>Land*</p>
						<select name="country">
							<option value="DE">Deutschland</option>
						</select>
					</div>
				</div>
			</div>
			<div class="register-box">
				<header>
					<h3>Ihre Kontaktinformationen</h3>
				</header>
				<div>
					<div>
						<p>Telefonnummer*</p>
						<input type="text" name="phone_number" placeholder="Telefonnummer" />
					</div>
				</div>
			</div>
			{#if !form?.success}
				<p style="color: red">{form?.message}</p>
			{/if}
			<input type="submit" value={creating ? "Wird erstellt..." : "Bestätigen"} />
		</form>
		<a href={resolve("/login")}>Haben Sie bereits einen Account? Hier einloggen</a>
	</div>
</div>

<style>
	.register {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		margin: 50px 0;
	}

	.register-content {
		display: flex;
		flex-direction: column;
		width: 50vw;
	}

	.register-content a {
		color: black;
		margin-top: 20px;
	}

	.register-box {
		margin-top: 50px;
		display: flex;
		flex-direction: column;
	}

	.register-box header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		border-bottom: 1px solid orange;
		margin-bottom: 25px;
		padding-bottom: 10px;
	}

	.register-box header h3 {
		text-transform: uppercase;
	}

	.register-box div {
		display: flex;
		flex-direction: column;
		gap: 25px;
	}

	.register-box div > div {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		height: 25px;
	}

	.register-box div > div {
		display: flex;
		flex-direction: row;
		gap: 5px;
	}

	.register-box div > div input {
		padding: 10px;
		background-color: whitesmoke;
		border: 1px solid lightgray;
		width: 50%;
		outline: none;
	}

	.register-box div > div input[type="radio"] {
		width: auto;
		padding: 5px;
	}

	.register-box div > div select {
		padding: 10px;
		background-color: whitesmoke;
		border: 1px solid lightgray;
		width: 50%;
		outline: none;
	}

	.register-box div > div input[type="text"]:focus,
	.register-box div > div input[type="email"]:focus {
		border-color: orange;
		box-shadow: 1px 1px 5px rgba(255, 165, 0, 0.5);
	}

	form input[type="submit"] {
		padding: 5px;
		outline: none;
		border: none;
		font-size: 1rem;
		width: 100%;
		border: 1px solid transparent;
		background-color: orange;
		color: white;
		padding: 10px;
		cursor: pointer;
		margin-top: 50px;
	}
</style>
