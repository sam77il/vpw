import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export async function load({ params, locals, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}
	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}

	const res = await fetch(`${API_URL}/api/v1/categories/data/${params.categoryId}`);

	if (!res.ok) {
		return {
			success: false,
			category: null
		};
	}
	const data = await res.json();

	return {
		success: true,
		category: data.category
	};
}

export const actions = {
	async update({ request, fetch, cookies }) {
		const formData = await request.formData();
		const formDataObj = Object.fromEntries(formData.entries());
		const res = await fetch(`${API_URL}/api/v1/categories/${formDataObj.id}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${cookies.get("auth")}`
			},
			body: JSON.stringify(formDataObj)
		});

		if (!res.ok) {
			return {
				success: false,
				message: "Failed to update category"
			};
		}

		return {
			success: true
		};
	}
};
