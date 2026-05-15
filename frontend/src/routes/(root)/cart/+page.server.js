import { API_URL } from "$env/static/private";

export async function load({ fetch, cookies, locals }) {
	if (!locals.user) {
		return {
			cart_items: [],
			message: "Bitte melde dich an, um deinen Warenkorb zu sehen."
		};
	}

	const auth = cookies.get("auth");
	const res = await fetch(`${API_URL}/api/v1/cart`, {
		method: "GET",
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});

	if (res.ok) {
		const api = await res.json();

		return {
			cart_items: api.cart_items || [],
			message: api.message || null
		};
	}

	return {
		cart_items: [],
		message: "Beim Laden des Warenkorbs ist ein Fehler aufgetreten."
	};
}
