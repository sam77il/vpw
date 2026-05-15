import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export async function load({ fetch, cookies }) {
	const token = cookies.get("auth");
	if (!token) {
		throw redirect(301, "/");
	}

	const res = await fetch(`${API_URL}/api/v1/me`, {
		method: "GET",
		headers: {
			Authorization: `Bearer ${token}`
		}
	});
	const data = await res.json();

	if (data.success) {
		return {
			user: data.user
		};
	}
}

export const actions = {
	async personal({ request, fetch, cookies }) {
		const data = await request.formData();
		const newData = Object.fromEntries(data.entries());

		const token = cookies.get("auth");
		const res = await fetch(`${API_URL}/api/v1/me/update`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify(newData)
		});
		const answer = await res.json();
		return answer;
	},
	async delivery({ request, fetch, cookies }) {
		const data = await request.formData();
		const newData = Object.fromEntries(data.entries());

		const token = cookies.get("auth");
		const res = await fetch(`${API_URL}/api/v1/me/update`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify(newData)
		});
		const answer = await res.json();
		return answer;
	},
	async password({ request, fetch, cookies }) {
		const data = await request.formData();
		const newData = Object.fromEntries(data.entries());
		console.log(newData.new_password, newData.new_password_confirm);
		if (newData.new_password !== newData.new_password_confirm) {
			return { success: false, message: "Die neuen Passwörter stimmen nicht überein." };
		}
		const token = cookies.get("auth");
		const res = await fetch(`${API_URL}/api/v1/me/update`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify(newData)
		});
		const answer = await res.json();

		if (!answer.success) {
			if (answer.message === "wrong current password") {
				return { success: false, message: "Das aktuelle Passwort ist falsch." };
			}
		}
		return answer;
	},
	async companyinfo({ request, fetch, cookies }) {
		const data = await request.formData();
		const newData = Object.fromEntries(data.entries());
		newData.company = true;

		const token = cookies.get("auth");
		const res = await fetch(`${API_URL}/api/v1/me/update`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify(newData)
		});
		const answer = await res.json();
		console.log(answer);
		return answer;
	},
	async email({ request, fetch, cookies }) {
		const data = await request.formData();
		const newData = Object.fromEntries(data.entries());

		const token = cookies.get("auth");
		const res = await fetch(`${API_URL}/api/v1/me/update`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json"
			},
			body: JSON.stringify(newData)
		});
		const answer = await res.json();
		return answer;
	}
};
