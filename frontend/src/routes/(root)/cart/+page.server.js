export async function load({ fetch, cookies }) {
	const auth = cookies.get("auth");
	const res = await fetch(`http://localhost:8080/api/v1/cart`, {
		method: "GET",
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});

	if (res.ok) {
		const api = await res.json();
		if (!api.cart_items) {
			return {
				cart_items: []
			};
		}

		return {
			cart_items: api.cart_items
		};
	}
}
