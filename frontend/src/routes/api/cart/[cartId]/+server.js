import { json } from "@sveltejs/kit";

export async function DELETE({ params, fetch, cookies }) {
	const auth = cookies.get("auth");
	const res = await fetch(`http://localhost:8080/api/v1/cart/${params.cartId}`, {
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
