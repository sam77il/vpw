<div class="legal-page">
	<h1>Impressum</h1>
	<p>VerbundplattenWelt GmbH</p>
	<p>Musterstraße 1</p>
	<p>12345 Berlin</p>
	<p>Telefon: +49 123 456 789</p>
	<p>Email: info@verbundplattenwelt.example</p>

	<h2>Geschäftsführer</h2>
	<p>Max Mustermann</p>

	<h2>Handelsregister</h2>
	<p>Amtsgericht Berlin, HRB 123456</p>

	<p>Alle Angaben sind Beispielinhalte und dienen der Demonstration.</p>
</div>

<style>
	.legal-page {
		max-width: 800px;
		margin: 0 auto;
		padding: clamp(24px, 5vw, 48px);
		display: flex;
		flex-direction: column;
		gap: 18px;
	}

	.legal-page h1,
	.legal-page h2 {
		color: var(--text);
	}

	.legal-page p {
		color: var(--text-light);
		margin: 0;
	}
</style>
