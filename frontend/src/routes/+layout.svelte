<script>
	import { navigating } from "$app/state";
	import { onDestroy } from "svelte";
	import Footer from "$lib/components/Footer.svelte";

	const { children } = $props();

	let visible = $state(false);
	let timeout;

	const MIN_DURATION = 500;

	$effect(() => {
		if (navigating.to) {
			visible = true;

			clearTimeout(timeout);
		} else {
			timeout = setTimeout(() => {
				visible = false;
			}, MIN_DURATION);
		}
	});

	onDestroy(() => {
		clearTimeout(timeout);
	});
</script>

{#if visible}
	<div class="loading">
		<div class="loading-bar"></div>
	</div>
{/if}
{@render children()}
<Footer />

<style>
	.loading {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 4px;
		background-color: transparent;
		z-index: 9999;
	}

	.loading-bar {
		width: 0%;
		height: 100%;
		background-color: orange;
		animation: loadingAnimation 0.5s linear infinite;
	}

	@keyframes loadingAnimation {
		0% {
			width: 0%;
			left: 0;
			right: auto;
		}
		50% {
			width: 100%;
			left: auto;
			right: auto;
		}
		100% {
			width: 0%;
			left: auto;
			right: 0;
		}
	}
</style>
