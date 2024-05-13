<script>
	const BASEURL = 'http://localhost:9001';

	import { onMount } from 'svelte';

	import { setModeCurrent } from '@skeletonlabs/skeleton';
	import { clipboard } from '@skeletonlabs/skeleton';
	import { ProgressRadial } from '@skeletonlabs/skeleton';
	import { Icon } from 'svelte-icons-pack';
	import {
		AiOutlineCopy,
		AiOutlineCheck,
		AiOutlineCloudUpload,
		AiOutlineReload
	} from 'svelte-icons-pack/ai';

	onMount(() => {
		setModeCurrent(false);

		const urlParams = new URLSearchParams(window.location.search);
		if (urlParams.has('goto')) {
			redirect(urlParams.get('goto'));
		}
	});

	$: URL = `${window.location.href}?goto=${data.message}`;

	let data, copied, fetched, loading;
	const reset = () => {
		data = '';
		copied = false;
		fetched = false;
		loading = true;
	};
	reset();

	const getShort = () => {
		if (data !== '') {
			fetched = true;
			fetch(`${BASEURL}/p?v=${data}`, { method: 'PUT' })
				.then((res) => res.json())
				.then((res) => {
					loading = false;
					data = res;
				})
				.catch((err) => console.error(err));
		}
	};

	const getLong = (key, callback) => {
		fetch(`${BASEURL}/g?k=${key}`)
			.then((res) => res.json())
			.then((res) => {
				console.log(res);
				callback(res);
			})
			.catch((err) => console.error(err));
	};

	const redirect = (key) => {
		getLong(key, (res) => {
			window.location.replace(res.message);
		});
	};

	const handleCopy = () => {
		copied = true;
		setTimeout(() => {
			copied = false;
		}, 1000);
	};

	const handleEnter = (evt) => {
		if (evt.key === 'Enter') getShort();
	};
</script>

<div class="w-screen">
	<div class="grid w-screen md:grid-cols-3">
		<div class="mt-64 pl-2 pr-2 md:col-start-2">
			{#if fetched}
				{#if loading}
					<div class="grid grid-cols-5 content-center p-3">
						<div class="col-span-2"></div>
						<div class="col-span-1 content-center">
							<ProgressRadial width="w-8" class="m-auto" />
						</div>

						<div class="col-span-2"></div>
					</div>
				{:else}
					<input type="text" class="input p-4" readonly="true" placeholder={URL} />
				{/if}
			{:else}
				<input
					type="text"
					class="input col-span-11 p-4"
					bind:value={data}
					on:keypress={handleEnter}
					maxlength="2048"
					placeholder="paste your link"
				/>
			{/if}
			<div class="grid grid-cols-3">
				{#if fetched}
					<button
						use:clipboard={URL}
						on:click={handleCopy}
						class="variant-filled-surface btn col-start-2 m-2 pt-4"
					>
						{#if !copied}
							<Icon src={AiOutlineCopy} size="32" />
						{:else}
							<Icon src={AiOutlineCheck} size="32" />
						{/if}
					</button>
					<button class="variant-filled-surface btn col-start-3 m-2 pt-4" on:click={reset}>
						<Icon src={AiOutlineReload} size="32" />
					</button>
				{:else}
					<button
						class="variant-filled-surface btn col-start-2 m-2 pt-4"
						on:click={() => getShort()}
					>
						<Icon src={AiOutlineCloudUpload} size="32" />
					</button>
				{/if}
			</div>
		</div>
	</div>
</div>

<!-- <style>
    div{
        border: 1px solid red;
    }
</style> -->
