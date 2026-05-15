import { API_URL } from "$env/static/private";

export async function load({ params, fetch }) {
	const res = await fetch(`${API_URL}/api/v1/categories/${params.categoryId}`);
	const products = await res.json();

	if (products.success) {
		return {
			categoryId: params.categoryId,
			products: products.products
		};
	}

	return {
		categoryId: params.categoryId,
		products: []
	};
}
