import { redirect } from "@sveltejs/kit";

export async function load({ locals, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}

	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}

	const res = await fetch("http://localhost:8080/api/v1/products");
	if (!res.ok) {
		return {
			success: false,
			products: []
		};
	}
	const data = await res.json();

	return {
		success: true,
		products: data.products
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
		const dataObj = Object.fromEntries(formData.entries());
		dataObj.price = Number(dataObj.price);
		dataObj.old_price = Number(dataObj.old_price);
		dataObj.stock = Number(dataObj.stock);
		const res = await fetch("http://localhost:8080/api/v1/products", {
			method: "POST",
			headers: {
				Authorization: `Bearer ${auth}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify(dataObj)
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
