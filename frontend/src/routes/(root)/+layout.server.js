export async function load({ fetch, locals }) {
	const res = await fetch("http://localhost:8080/api/v1/categories");
	const api = await res.json();

	return {
		categories: api.categories,
		user: locals.user ?? null
	};
}
