import { json } from "@sveltejs/kit";

export async function DELETE({ fetch, cookies, locals, params }) {
	if (!locals?.user || locals?.user?.role !== "admin") {
		return json({ success: false, message: "unauthorized" });
	}

	const auth = cookies.get("auth");
	const res = await fetch(`http://localhost:8080/api/v1/categories/${params.categoryId}`, {
		method: "DELETE",
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});
	const data = await res.json();

	if (!res.ok) {
		return json(data);
	}
	return json(data);
}
