import { API_URL } from "$env/static/private";

export async function load({ params, fetch }) {
	const res = await fetch(`${API_URL}/api/v1/products/${params.productId}`);
	const api = await res.json();
	console.log(API_URL);
	if (api.product?.id) {
		return {
			productId: params.productId,
			product: api.product
		};
	} else {
		return {
			productId: params.productId,
			product: {}
		};
	}
}
