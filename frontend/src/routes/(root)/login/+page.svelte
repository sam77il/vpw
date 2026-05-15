<script>
	import { enhance } from "$app/forms";
	import { resolve } from "$app/paths";
</script>

<div class="login">
	<div class="login-content">
		<h2>Login</h2>

		<form method="post" use:enhance>
			<input type="email" name="email" placeholder="E-Mail" />
			<input type="password" name="password" placeholder="Passwort" />
			<input type="submit" value="Bestätigen" />
		</form>
		<a href={resolve("/register")}>Noch keinen Account? Hier registrieren</a>
	</div>
</div>

<style>
	.login {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		min-height: calc(100vh - 120px);
		padding: clamp(16px, 5vw, 50px);
	}

	.login-content {
		width: 100%;
		max-width: 500px;
		height: max-content;
		background-color: whitesmoke;
		padding: clamp(16px, 5vw, 32px);
		border-radius: 12px;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	}

	.login-content h2 {
		color: var(--text);
		font-size: clamp(1.5rem, 4vw, 1.7rem);
		text-align: center;
		margin-bottom: 24px;
	}

	.login-content form {
		display: flex;
		flex-direction: column;
		gap: clamp(16px, 3vw, 25px);
		margin-top: 24px;
	}

	.login-content form input {
		padding: clamp(10px, 2vw, 14px);
		background-color: rgb(235, 235, 235);
		border: 2px solid var(--border);
		outline: none;
		border-radius: 6px;
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		transition:
			border-color 0.2s ease,
			box-shadow 0.2s ease;
		touch-action: manipulation;
		-webkit-appearance: none;
		appearance: none;
	}

	.login-content form input::placeholder {
		color: var(--text-light);
	}

	.login-content form input:focus {
		border-color: var(--accent);
		box-shadow: 0 0 0 3px rgba(249, 115, 22, 0.1);
	}

	.login-content form input[type="submit"] {
		background-color: var(--accent);
		color: white;
		padding: clamp(12px, 2.5vw, 16px);
		cursor: pointer;
		font-weight: 600;
		font-size: clamp(0.95rem, 1.5vw, 1rem);
		border: none;
		border-radius: 6px;
		transition:
			background-color 0.2s ease,
			transform 0.1s ease;
		margin-top: 8px;
	}

	.login-content form input[type="submit"]:hover {
		background-color: var(--accent-hover);
	}

	.login-content form input[type="submit"]:active {
		transform: scale(0.98);
	}

	.login-content a {
		display: block;
		margin-top: clamp(12px, 3vw, 20px);
		color: var(--accent);
		text-decoration: none;
		font-size: clamp(0.9rem, 1.5vw, 1rem);
		text-align: center;
		transition: color 0.2s ease;
	}

	.login-content a:hover {
		color: var(--accent-hover);
	}

	/* Mobile improvements */
	@media (max-width: 639px) {
		.login-content form input {
			min-height: 44px;
		}

		.login-content form input[type="submit"] {
			min-height: 48px;
		}
	}
</style>
