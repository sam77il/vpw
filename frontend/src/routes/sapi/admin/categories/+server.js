import { json } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export async function GET({ fetch, locals }) {
	if (!locals?.user || locals?.user?.role !== "admin") {
		return json({ success: false, message: "unauthorized" });
	}

	// const auth = cookies.get("auth");
	const res = await fetch(`${API_URL}/api/v1/categories`);
	const data = await res.json();

	if (!res.ok) {
		return json(data);
	}
	return json(data);
}
