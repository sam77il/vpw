<div class="orders">
	<div class="orders-content">
		<header>
			<h2>Bestellungen</h2>
		</header>

		<div>
			<div class="orders-content-page">
				<header>
					<h3>Bestellübersicht</h3>
				</header>

				<div class="orders-content-page-list"></div>
			</div>
		</div>
	</div>
</div>

<style>
	.orders {
		display: flex;
		justify-content: center;
		width: 100%;
		padding: clamp(16px, 5vw, 24px);
	}

	.orders-content {
		width: 100%;
		max-width: 1100px;
		margin-top: 80px;
		display: flex;
		flex-direction: column;
		gap: clamp(24px, 4vw, 50px);
	}

	.orders-content > div {
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	@media (min-width: 768px) {
		.orders-content > div {
			flex-direction: row;
		}
	}

	.orders-content-page {
		background-color: whitesmoke;
		border-radius: 10px;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 10px;
		padding: 20px;
		height: max-content;
	}
</style>
