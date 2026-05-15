import { API_URL } from "$env/static/private";

export async function load({ fetch, locals }) {
	const res = await fetch(`${API_URL}/api/v1/categories`);
	const api = await res.json();

	return {
		categories: api.categories,
		user: locals.user ?? null
	};
}
