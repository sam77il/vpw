import { redirect } from "@sveltejs/kit";

export async function load({ locals, cookies, fetch }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}

	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}
	const auth = cookies.get("auth");
	const res = await fetch("http://localhost:8080/api/v1/users", {
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
