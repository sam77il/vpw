import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export const actions = {
	async default({ request, cookies, fetch }) {
		const data = await request.formData();
		const email = data.get("email");
		const password = data.get("password");

		const res = await fetch(`${API_URL}/api/v1/auth/login`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				email,
				password
			})
		});
		const answer = await res.json();
		console.log(answer);
		if (answer.success) {
			cookies.set("auth", answer.token, {
				httpOnly: true,
				path: "/",
				sameSite: "lax"
			});
			throw redirect(302, "/account");
		}
		return { success: false };
	}
};
