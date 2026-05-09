<script>
  import { goto } from '$app/navigation';
  import { resolve } from '$app/paths';
	const { data } = $props();
</script>

<div class="admin">
	<div class="admin-content">
		<div>
			<div class="admin-content-page">
				<header>
					<h3>Registrierte Nutzer</h3>
          <input type="text" placeholder="Nutzer suchen..." />
				</header>

        {#if data.success}
          <table class="admin-content-page-list">
            <thead>
              <tr>
                <th>Nr.</th>
                <th>E-Mail</th>
                <th>Vorname</th>
                <th>Nachname</th>
                <th>Rolle</th>
              </tr>
            </thead>
            <tbody>
            {#each data.users as user, i (user.id)}
              <tr class="admin-content-page-list-item" onclick={() => goto(resolve(`/admin/users/${user.id}`))}>
                <td>{i + 1}</td>
                <td>{user.email}</td>
                <td>{user.first_name}</td>
                <td>{user.last_name}</td>
                <td>{user.role}</td>
              </tr>
            {/each}
            </tbody>
          </table>
        {:else}
          <p>Fehler beim Laden der Nutzer: {data.message}</p>
        {/if}
			</div>
		</div>
	</div>
</div>

<style>
	.admin {
		display: flex;
		justify-content: center;
		height: 100vh;
		width: 100%;
	}

  header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

	.admin-content {
		width: 50vw;
		height: 50%;
		margin-top: 50px;
		display: flex;
		flex-direction: column;
		gap: 50px;
	}

	.admin-content > div {
		display: flex;
		flex-direction: row;
		gap: 50px;
	}

	.admin-content-page {
		background-color: whitesmoke;
		border-radius: 10px;
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 10px;
		padding: 20px;
		height: max-content;
	}

  .admin-content-page-list {
    width: 100%;
  }

  .admin-content-page-list-item {
    border-top: 1px solid black;
  }

  thead {
    background-color: orange;
    color: white;
  }

  .admin-content-page-list-item {
    cursor: pointer;
    background-color: white;
  }

  .admin-content-page-list-item td {
    padding: 10px;
    }

  .admin-content-page-list-item:hover {
    background-color: lightgray;
  }

  header > input {
    padding: 10px;
    border-radius: 10px;
    border: 1px solid lightgray;
    outline: none;
    width: 250px;
  }
</style>
