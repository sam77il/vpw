import { json } from "@sveltejs/kit";

export function POST({ cookies }) {
	cookies.delete("auth", { path: "/" });

	return json({ success: true });
}
