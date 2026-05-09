import { redirect } from "@sveltejs/kit";

export async function load({ params, locals, cookies, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}
	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}
	const auth = cookies.get("auth");
	const res = await fetch(`http://localhost:8080/api/v1/users/${params.userId}`, {
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});

	if (!res.ok) {
		return {
			success: false,
			user: null
		};
	}
	const data = await res.json();

	return {
		success: true,
		user: data.user
	};
}

export const actions = {
	async update({ request, fetch, cookies }) {
		const formData = await request.formData();
		const formDataObj = Object.fromEntries(formData.entries());
		formDataObj.company = formDataObj.company === "true" ? true : false;
		const res = await fetch(`http://localhost:8080/api/v1/users/${formDataObj.id}`, {
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
				message: "Failed to update user"
			};
		}

		return {
			success: true
		};
	}
};
