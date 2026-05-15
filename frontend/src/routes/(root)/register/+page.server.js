import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export const actions = {
	async default({ request, fetch, cookies }) {
		const data = await request.formData();
		const formData = Object.fromEntries(data.entries());

		if (
			!formData.gender ||
			!formData.first_name ||
			!formData.last_name ||
			!formData.email ||
			!formData.email_confirmation ||
			!formData.password ||
			!formData.password_confirmation ||
			!formData.street ||
			!formData.postal_code ||
			!formData.city ||
			!formData.country ||
			!formData.phone_number ||
			!formData.company
		) {
			return { success: false, message: "Bitte füllen Sie alle Pflichtfelder aus." };
		} else if (formData.email !== formData.email_confirmation) {
			return { success: false, message: "Die E-Mail-Adressen stimmen nicht überein." };
		} else if (formData.password !== formData.password_confirmation) {
			return { success: false, message: "Die Passwörter stimmen nicht überein." };
		} else if (
			formData.company === "true" &&
			(!formData.company_name || !formData.company_ustidnr)
		) {
			return {
				success: false,
				message: "Bitte füllen Sie alle Pflichtfelder für die Firmendaten aus."
			};
		}
		if (formData.company === "true") {
			formData.company = true;
		} else {
			formData.company = false;
		}

		const res = await fetch(`${API_URL}/api/v1/auth/register`, {
			method: "POST",
			body: JSON.stringify(formData)
		});
		const answer = await res.json();

		if (answer.success) {
			cookies.set("auth", answer.token, {
				path: "/",
				httpOnly: true,
				secure: true,
				sameSite: "lax"
			});
			throw redirect(303, "/");
		}
		return answer;
	}
};
