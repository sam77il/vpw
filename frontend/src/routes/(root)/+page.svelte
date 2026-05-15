<script lang="ts">
	import { resolve } from "$app/paths";
	import {
		Check,
		Scissors,
		Truck,
		ShieldCheck,
		Ruler,
		Factory,
		PanelsTopLeft,
		ArrowRight
	} from "@lucide/svelte";

	const advantages = [
		"Präziser Zuschnitt nach Wunschmaß",
		"Hochwertige Aluverbundplatten in verschiedenen Ausführungen",
		"Einfache Online-Bestellung",
		"Schnelle Bearbeitung und zuverlässige Lieferung",
		"Fachhandel mit Erfahrung und persönlichem Service",
		"Ideal für Werbung, Schilder, Verkleidungen, Druckplatten und vieles mehr"
	];

	const applications = [
		"Werbeschilder und Firmenschilder",
		"Fassadenverkleidungen",
		"Messe- und Ladenbau",
		"Innenausbau",
		"Druck- und Präsentationsflächen",
		"Technische Anwendungen"
	];
</script>

<svelte:head>
	<title>Aluverbundplatten nach Maß</title>
	<meta
		name="description"
		content="Aluverbundplatten nach Maß online bestellen. Präziser Zuschnitt, schnelle Lieferung und hochwertige Qualität direkt vom Fachhändler."
	/>
</svelte:head>

<div class="page">
	<!-- HERO -->
	<section class="hero">
		<div class="container hero-grid">
			<div class="hero-content">
				<div class="badge">
					<Factory size={16} />
					<span>Direkt vom Fachhändler</span>
				</div>

				<h1>Aluverbundplatten nach Maß</h1>

				<p class="hero-text">
					Hochwertige Aluverbundplatten mit individuellem Zuschnitt für Industrie, Handwerk,
					Werbung, Messebau, Innenausbau und viele weitere Einsatzbereiche.
				</p>

				<div class="hero-buttons">
					<a href={resolve("/categories/aluverbundplatten3mm")} class="btn btn-primary">
						Jetzt konfigurieren
						<ArrowRight size={18} />
					</a>

					<a href="#vorteile" class="btn btn-secondary"> Vorteile ansehen </a>
				</div>
			</div>

			<div class="hero-card">
				<div class="info-grid">
					<div class="info-box">
						<Ruler size={32} />
						<h3>Millimetergenau</h3>
						<p>Exakter Zuschnitt nach Ihren Wunschmaßen.</p>
					</div>

					<div class="info-box">
						<Scissors size={32} />
						<h3>Individuelle Fertigung</h3>
						<p>Einzelstücke oder größere Mengen schnell gefertigt.</p>
					</div>

					<div class="info-box">
						<Truck size={32} />
						<h3>Schnelle Lieferung</h3>
						<p>Zuverlässiger Versand direkt zu Ihnen.</p>
					</div>

					<div class="info-box">
						<ShieldCheck size={32} />
						<h3>Fachhandel Qualität</h3>
						<p>Hochwertige Materialien und persönlicher Service.</p>
					</div>
				</div>
			</div>
		</div>
	</section>

	<!-- INTRO -->
	<section class="section">
		<div class="container narrow">
			<div class="badge light">Maßanfertigung einfach online bestellen</div>

			<h2>Keine Standardlösungen von der Stange</h2>

			<p>
				Sie bestimmen die gewünschten Maße – wir übernehmen den präzisen Zuschnitt. Über unseren
				Onlineshop können Sie Ihre Aluverbundplatten schnell und unkompliziert konfigurieren und
				direkt bestellen.
			</p>

			<p>
				Ob Einzelstück oder größere Mengen: Wir fertigen Ihre Bestellung millimetergenau nach Ihren
				Vorgaben.
			</p>
		</div>
	</section>

	<!-- ADVANTAGES -->
	<section class="section section-light" id="vorteile">
		<div class="container">
			<div class="badge light">Ihre Vorteile</div>

			<h2>Alles aus einer Hand</h2>

			<div class="advantages-grid">
				{#each advantages as item, i (i)}
					<div class="advantage-card">
						<div class="icon-circle">
							<Check size={18} />
						</div>

						<span>{item}</span>
					</div>
				{/each}
			</div>
		</div>
	</section>

	<!-- APPLICATIONS -->
	<section class="section">
		<div class="container applications-grid">
			<div>
				<div class="badge light">Vielseitig einsetzbar</div>

				<h2>Für zahlreiche Anwendungen geeignet</h2>

				<p>
					Aluverbundplatten überzeugen durch ihr geringes Gewicht, hohe Stabilität und moderne
					Optik.
				</p>

				<div class="application-list">
					{#each applications as item, i (i)}
						<div class="application-item">
							<div class="icon-circle dark">
								<PanelsTopLeft size={18} />
							</div>

							<span>{item}</span>
						</div>
					{/each}
				</div>
			</div>

			<div class="dark-box">
				<h3>Qualität und Service vom Profi</h3>

				<p>
					Als spezialisierter Fachhändler legen wir großen Wert auf Qualität, saubere Verarbeitung
					und zuverlässigen Service.
				</p>

				<p>Jede Platte wird mit Sorgfalt zugeschnitten und versandbereit vorbereitet.</p>

				<p>
					Unser Ziel ist es, Ihnen den Einkauf von maßgefertigten Aluverbundplatten so einfach wie
					möglich zu machen – schnell, professionell und unkompliziert.
				</p>
			</div>
		</div>
	</section>

	<!-- CTA -->
	<section class="cta">
		<div class="container cta-content">
			<h2>Jetzt online konfigurieren und bestellen</h2>

			<p>
				Geben Sie einfach Ihre Wunschmaße ein und bestellen Sie Ihre Aluverbundplatten bequem
				online. Wir kümmern uns um den Zuschnitt und die sichere Lieferung direkt zu Ihnen.
			</p>

			<a href={resolve("/categories/aluverbundplatten3mm")} class="btn btn-light">
				Jetzt bestellen
				<ArrowRight size={18} />
			</a>
		</div>
	</section>
</div>

<style>
	:global(body) {
		margin: 0;
		font-family: Inter, system-ui, sans-serif;
		background: #ffffff;
		color: #18181b;
	}

	/* ORANGE THEME */
	:root {
		--accent: #f97316;
		--accent-hover: #ea580c;
		--accent-light: #fff7ed;
		--accent-border: #fdba74;

		--text: #18181b;
		--text-light: #52525b;

		--bg-light: #f8fafc;
		--border: #e4e4e7;
	}

	.page {
		width: 100%;
		overflow-x: hidden;
	}

	.container {
		width: min(1200px, calc(100% - 48px));
		margin: 0 auto;
	}

	.narrow {
		max-width: 850px;
	}

	.section {
		padding: 100px 0;
	}

	.section-light {
		background: linear-gradient(to bottom, #fff7ed, #ffffff);
	}

	.hero {
		padding: 120px 0;
		border-bottom: 1px solid var(--border);
		background:
			radial-gradient(circle at top left, #ffedd5, transparent 40%),
			radial-gradient(circle at bottom right, #fed7aa, transparent 30%), white;
	}

	.hero-grid {
		display: grid;
		grid-template-columns: 1.2fr 1fr;
		gap: 64px;
		align-items: center;
	}

	.hero h1 {
		font-size: clamp(3rem, 6vw, 5rem);
		line-height: 1;
		margin: 24px 0;
		font-weight: 900;
		letter-spacing: -0.04em;
		color: var(--text);
	}

	.hero-text {
		font-size: 1.2rem;
		line-height: 1.8;
		color: var(--text-light);
		max-width: 700px;
	}

	.hero-buttons {
		display: flex;
		gap: 16px;
		margin-top: 40px;
		flex-wrap: wrap;
	}

	.btn {
		display: inline-flex;
		align-items: center;
		gap: 10px;
		padding: 16px 28px;
		border-radius: 18px;
		text-decoration: none;
		font-weight: 700;
		transition:
			0.25s ease,
			transform 0.2s ease;
	}

	.btn:hover {
		transform: translateY(-2px);
	}

	.btn-primary {
		background: var(--accent);
		color: white;
		box-shadow: 0 10px 25px rgba(249, 115, 22, 0.25);
	}

	.btn-primary:hover {
		background: var(--accent-hover);
	}

	.btn-secondary {
		border: 1px solid var(--accent-border);
		color: var(--accent);
		background: white;
	}

	.btn-secondary:hover {
		background: var(--accent-light);
	}

	.btn-light {
		background: white;
		color: var(--accent);
	}

	.badge {
		display: inline-flex;
		align-items: center;
		gap: 8px;
		padding: 10px 16px;
		border-radius: 999px;
		background: white;
		border: 1px solid var(--accent-border);
		color: var(--accent);
		font-size: 0.9rem;
		font-weight: 700;
	}

	.badge.light {
		background: var(--accent-light);
		border: none;
	}

	.hero-card {
		background: white;
		border-radius: 32px;
		padding: 32px;
		border: 1px solid var(--border);
		box-shadow: 0 20px 50px rgba(249, 115, 22, 0.12);
	}

	.info-grid {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		gap: 20px;
	}

	.info-box {
		background: linear-gradient(to bottom right, #fff7ed, white);
		border-radius: 24px;
		padding: 28px;
		border: 1px solid #fed7aa;
	}

	.info-box h3 {
		margin: 18px 0 10px;
		font-size: 1.1rem;
	}

	.info-box p {
		margin: 0;
		color: var(--text-light);
		line-height: 1.7;
	}

	h2 {
		font-size: clamp(2.2rem, 4vw, 3.5rem);
		line-height: 1.1;
		margin: 24px 0;
		font-weight: 900;
		letter-spacing: -0.03em;
		color: var(--text);
	}

	p {
		font-size: 1.1rem;
		line-height: 1.9;
		color: var(--text-light);
	}

	.advantages-grid {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 24px;
		margin-top: 60px;
	}

	.advantage-card {
		background: white;
		border-radius: 24px;
		padding: 24px;
		border: 1px solid #fed7aa;
		display: flex;
		gap: 18px;
		align-items: flex-start;
		transition:
			transform 0.2s ease,
			box-shadow 0.2s ease;
	}

	.advantage-card:hover {
		transform: translateY(-4px);
		box-shadow: 0 12px 30px rgba(249, 115, 22, 0.12);
	}

	.advantage-card span {
		line-height: 1.7;
		font-weight: 500;
	}

	.icon-circle {
		min-width: 42px;
		width: 42px;
		height: 42px;
		border-radius: 14px;
		display: flex;
		align-items: center;
		justify-content: center;
		background: var(--accent);
		color: white;
		box-shadow: 0 8px 20px rgba(249, 115, 22, 0.25);
	}

	.icon-circle.dark {
		background: var(--accent);
	}

	.applications-grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 64px;
		align-items: start;
	}

	.application-list {
		margin-top: 40px;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.application-item {
		display: flex;
		align-items: center;
		gap: 18px;
	}

	.application-item span {
		font-size: 1.1rem;
		font-weight: 600;
	}

	.dark-box {
		background: linear-gradient(135deg, #ea580c, #f97316);
		color: white;
		padding: 48px;
		border-radius: 32px;
		box-shadow: 0 20px 50px rgba(249, 115, 22, 0.25);
	}

	.dark-box h3 {
		font-size: 2rem;
		margin-top: 0;
		margin-bottom: 24px;
	}

	.dark-box p {
		color: rgba(255, 255, 255, 0.88);
	}

	.cta {
		padding: 120px 0;
		background: linear-gradient(135deg, #ea580c, #f97316);
		color: white;
		text-align: center;
	}

	.cta-content {
		max-width: 900px;
	}

	.cta h2 {
		color: white;
	}

	.cta p {
		color: rgba(255, 255, 255, 0.85);
		margin-bottom: 40px;
	}

	@media (max-width: 900px) {
		.hero-grid,
		.applications-grid,
		.advantages-grid,
		.info-grid {
			grid-template-columns: 1fr;
		}

		.hero {
			padding: 80px 0;
		}

		.section {
			padding: 80px 0;
		}

		.dark-box {
			padding: 32px;
		}

		.hero-buttons {
			flex-direction: column;
			align-items: stretch;
		}

		.btn {
			justify-content: center;
		}
	}
</style>
