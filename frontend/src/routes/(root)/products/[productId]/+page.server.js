export async function load({ params, fetch }) {
	const res = await fetch(`http://localhost:8080/api/v1/products/${params.productId}`);
	const api = await res.json();

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
