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
		margin-top: 50px;
	}

	.login-content {
		width: 500px;
		height: max-content;
		background-color: whitesmoke;
		padding: 20px;
	}

	.login-content h2 {
		color: black;
		font-size: 1.7rem;
		text-align: center;
	}

	.login-content form {
		display: flex;
		flex-direction: column;
		gap: 25px;
		margin-top: 25px;
	}

	.login-content form input {
		padding: 10px;
		background-color: rgb(235, 235, 235);
		border: 1px solid lightgray;
		outline: none;
	}

	.login-content form input:focus {
		border: 1px solid orange;
		box-shadow: 1px 1px 5px rgba(255, 165, 0, 0.5);
	}

	.login-content form input[type="submit"] {
		background-color: orange;
		color: white;
		padding: 10px;
		cursor: pointer;
	}

	.login-content a {
		display: block;
		margin-top: 5px;
		color: black;
		text-decoration: none;
	}
</style>
