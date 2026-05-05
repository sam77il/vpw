import { redirect } from "@sveltejs/kit";

export function load({ url }) {
	if (url.pathname === "/account") {
		throw redirect(302, "/account/orders");
	}
}
