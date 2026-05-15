import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export async function load({ locals, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}

	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}

	const res = await fetch(`${API_URL}/api/v1/categories`);

	if (!res.ok) {
		return {
			success: false,
			categories: []
		};
	}
	const data = await res.json();

	return {
		success: true,
		categories: data.categories
	};
}

export const actions = {
	async create({ cookies, fetch, locals, request }) {
		if (locals?.user?.role !== "admin") {
			return {
				success: false,
				message: "unauthorized"
			};
		}

		const auth = cookies.get("auth");
		const formData = await request.formData();
		const catId = formData.get("id");
		const catLabel = formData.get("label");
		const res = await fetch(`${API_URL}/api/v1/categories`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${auth}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify({ id: catId, label: catLabel })
		});

		if (!res.ok) {
			const data = await res.json();
			return {
				success: false,
				message: data.message
			};
		}

		return {
			success: true
		};
	}
};
