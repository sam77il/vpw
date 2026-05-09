import { redirect } from "@sveltejs/kit";

export async function load({ params, locals, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}
	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}

	const res = await fetch(`http://localhost:8080/api/v1/products/${params.productId}`);

	if (!res.ok) {
		return {
			success: false,
			product: null,
			categories: null
		};
	}
	const data = await res.json();
	const res2 = await fetch("/api/admin/categories");
	if (!res2.ok) {
		return {
			success: false,
			product: null,
			categories: null
		};
	}
	const data2 = await res2.json();

	return {
		success: true,
		product: data.product,
		categories: data2.categories
	};
}

export const actions = {
	async update({ request, fetch, cookies }) {
		const formData = await request.formData();
		const formDataObj = Object.fromEntries(formData.entries());
		formDataObj.price = Number(formDataObj.price);
		formDataObj.old_price = Number(formDataObj.old_price);
		formDataObj.stock = Number(formDataObj.stock);
		const res = await fetch(`http://localhost:8080/api/v1/products/${formDataObj.id}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${cookies.get("auth")}`
			},
			body: JSON.stringify(formDataObj)
		});
		const data = await res.json();
		console.log(data);
		if (!res.ok) {
			return {
				success: false,
				message: "Failed to update product"
			};
		}
		return data;
	}
};
