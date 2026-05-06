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
		height: 100vh;
		width: 100%;
	}

	.orders-content {
		width: 50vw;
		height: 50%;
		margin-top: 50px;
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	.orders-content > div {
		display: flex;
		flex-direction: row;
		gap: 50px;
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
