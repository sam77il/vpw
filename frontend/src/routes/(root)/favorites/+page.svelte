<script>
	const pageTitle = "Favoriten";
</script>

<div class="placeholder-page">
	<header>
		<h1>{pageTitle}</h1>
		<p>Favoriten sind derzeit noch nicht verfügbar. Diese Funktion erscheint bald.</p>
	</header>

	<div class="placeholder-card">
		<strong>SOON</strong>
		<p>Wir arbeiten daran, deine Favoriten hier anzuzeigen.</p>
	</div>
</div>

<style>
	.placeholder-page {
		max-width: 740px;
		margin: 0 auto;
		padding: clamp(24px, 5vw, 48px);
		display: flex;
		flex-direction: column;
		gap: clamp(20px, 4vw, 28px);
		text-align: center;
	}

	.placeholder-page header h1 {
		font-size: clamp(2rem, 5vw, 2.8rem);
		margin: 0;
	}

	.placeholder-page header p {
		font-size: clamp(1rem, 2.5vw, 1.1rem);
		color: var(--text-light);
		margin: 0;
	}

	.placeholder-card {
		background: var(--bg-light);
		border: 1px solid var(--border);
		border-radius: 16px;
		padding: clamp(24px, 4vw, 32px);
		box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
	}

	.placeholder-card strong {
		display: block;
		font-size: clamp(2.5rem, 6vw, 3.5rem);
		color: var(--accent);
		margin-bottom: 16px;
	}

	.placeholder-card p {
		font-size: clamp(1rem, 2.5vw, 1.1rem);
		color: var(--text-light);
		margin: 0;
	}
</style>
