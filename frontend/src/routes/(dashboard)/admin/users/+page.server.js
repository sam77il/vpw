import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export async function load({ locals, cookies, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}

	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}
	const auth = cookies.get("auth");
	const res = await fetch(`${API_URL}/api/v1/users`, {
		headers: {
			Authorization: `Bearer ${auth}`
		}
	});

	if (!res.ok) {
		return {
			success: false,
			users: []
		};
	}
	const data = await res.json();

	return {
		success: true,
		users: data.users
	};
}
