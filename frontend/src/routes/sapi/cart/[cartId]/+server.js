import { json } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export async function DELETE({ params, fetch, cookies }) {
	const auth = cookies.get("auth");
	const res = await fetch(`${API_URL}/api/v1/cart/${params.cartId}`, {
		method: "DELETE",
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});

	if (!res.ok) {
		const errorData = await res.json();
		return json(errorData);
	}

	return json({ success: true });
}
