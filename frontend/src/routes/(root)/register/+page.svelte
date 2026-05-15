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
		min-height: auto;
		padding: clamp(16px, 5vw, 50px);
	}

	.register-content {
		display: flex;
		flex-direction: column;
		width: 100%;
		max-width: 800px;
	}

	.register-content h2 {
		font-size: clamp(1.5rem, 4vw, 2rem);
		text-align: center;
		margin-bottom: 32px;
		color: var(--text);
	}

	.register-content a {
		color: var(--accent);
		margin-top: clamp(16px, 3vw, 20px);
		text-align: center;
		text-decoration: none;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		transition: color 0.2s ease;
	}

	.register-content a:hover {
		color: var(--accent-hover);
	}

	.register-box {
		margin-top: clamp(32px, 5vw, 50px);
		display: flex;
		flex-direction: column;
		padding: clamp(16px, 3vw, 24px);
		background-color: var(--bg-light);
		border-radius: 8px;
	}

	.register-box header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		border-bottom: 2px solid var(--accent);
		margin-bottom: clamp(16px, 3vw, 25px);
		padding-bottom: 12px;
	}

	.register-box header h3 {
		text-transform: uppercase;
		font-size: clamp(0.95rem, 2vw, 1.1rem);
		font-weight: 700;
		color: var(--text);
		margin: 0;
	}

	.register-box > div {
		display: flex;
		flex-direction: column;
		gap: clamp(16px, 3vw, 25px);
	}

	.register-box > div > div {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}

	.register-box > div > div > p {
		margin: 0;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		font-weight: 500;
		color: var(--text);
	}

	/* Radio button groups */
	.register-box > div > div > div {
		display: flex;
		flex-direction: row;
		gap: clamp(16px, 3vw, 24px);
		flex-wrap: wrap;
	}

	.register-box > div > div > div > div {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 8px;
	}

	.register-box > div > div input[type="radio"],
	.register-box > div > div input[type="checkbox"] {
		cursor: pointer;
		width: 18px;
		height: 18px;
		accent-color: var(--accent);
	}

	.register-box > div > div label {
		cursor: pointer;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		user-select: none;
	}

	/* Input fields */
	.register-box input[type="text"],
	.register-box input[type="email"],
	.register-box input[type="password"],
	.register-box input[type="tel"],
	.register-box select {
		padding: clamp(10px, 2vw, 14px);
		background-color: white;
		border: 2px solid var(--border);
		outline: none;
		border-radius: 6px;
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		font-family: inherit;
		transition:
			border-color 0.2s ease,
			box-shadow 0.2s ease;
		touch-action: manipulation;
		-webkit-appearance: none;
		appearance: none;
	}

	.register-box input::placeholder {
		color: var(--text-light);
	}

	.register-box input:focus,
	.register-box select:focus {
		border-color: var(--accent);
		box-shadow: 0 0 0 3px rgba(249, 115, 22, 0.1);
	}

	/* Postal code and city row */
	.postal-city {
		display: grid;
		grid-template-columns: 1fr 2fr;
		gap: 12px;
	}

	@media (max-width: 639px) {
		.postal-city {
			grid-template-columns: 1fr;
		}

		.register-box input[type="text"],
		.register-box input[type="email"],
		.register-box input[type="password"],
		.register-box input[type="tel"],
		.register-box select {
			min-height: 44px;
		}
	}

	/* Submit button */
	form input[type="submit"] {
		padding: clamp(12px, 2.5vw, 16px);
		outline: none;
		border: none;
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		width: 100%;
		background-color: var(--accent);
		color: white;
		cursor: pointer;
		margin-top: clamp(32px, 5vw, 50px);
		border-radius: 6px;
		font-weight: 600;
		transition:
			background-color 0.2s ease,
			transform 0.1s ease;
	}

	form input[type="submit"]:hover {
		background-color: var(--accent-hover);
	}

	form input[type="submit"]:active {
		transform: scale(0.98);
	}

	/* Error message */
	form > p {
		color: #dc2626;
		text-align: center;
		margin-top: 16px;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
	}

	/* Mobile optimizations */
	@media (max-width: 639px) {
		form input[type="submit"] {
			min-height: 48px;
		}

		.register-box {
			padding: 12px;
		}

		.register-box > div > div > div {
			flex-direction: column;
			gap: 12px;
		}
	}
</style>
