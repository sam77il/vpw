import { json } from "@sveltejs/kit";

export async function DELETE({ cookies, fetch, params }) {
	const auth = cookies.get("auth");

	const res = await fetch("http://localhost:8080/api/v1/users/" + params.userId, {
		method: "DELETE",
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});

	if (!res.ok) {
		const data = await res.json();
		return json(data);
	}
	return json({ success: true });
}
