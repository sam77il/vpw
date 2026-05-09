import { redirect } from "@sveltejs/kit";

export async function load({ locals }) {
	if (!locals.user) {
		throw redirect(302, "/");
	}

	if (locals.user?.role !== "admin") {
		throw redirect(302, "/");
	}

	throw redirect(302, "/admin/users");
}
