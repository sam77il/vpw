import { json } from "@sveltejs/kit";

export async function POST({ request, fetch, cookies }) {
	const authToken = cookies.get("auth");
	const { id, price, amount, metadata } = await request.json();
	console.log("Received cart item:", { id, price, amount, metadata });
	const res = await fetch("http://localhost:8080/api/v1/cart", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			Authorization: `Bearer ${authToken}`
		},
		body: JSON.stringify({
			product_id: id,
			amount,
			price,
			metadata
		})
	});

	if (!res.ok) {
		const errorData = await res.json();
		return json(errorData);
	}
	return json({ success: true });
}
