export const handle = async ({ event, resolve }) => {
	const token = event.cookies.get("auth");

	if (token) {
		const res = await event.fetch("http://localhost:8080/api/v1/auth/validate", {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});
		if (res.ok) {
			const user = await res.json();
			event.locals.user = {
				role: user.role,
				user_id: user.user_id
			};
		} else {
			event.locals.user = null;
			event.cookies.delete("auth", { path: "/" });
		}
	}

	return resolve(event);
};
