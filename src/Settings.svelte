<script lang="ts">
    import { fade } from "svelte/transition";

    export let settings;

    function handler(event) {}
</script>

<svelte:window on:keydown={handler} />

{#if settings}
    <div class="modal" on:contextmenu={event => event.preventDefault()} in:fade={{ duration: 50 }} out:fade={{ duration: 50 }}>
        <div class="blur" on:click|self={() => (settings = false)} />
        <div class="modal-container">
            <h1>Settings</h1>
        </div>
    </div>
{/if}

<style>
    .modal {
        display: flex;
        justify-content: center;
        align-items: center;
        position: absolute;
        height: 100%;
        width: 100%;
        z-index: 1;
    }

    .modal > .blur {
        position: absolute;
        height: 100%;
        width: 100%;
        background-color: var(--c1);
        opacity: 0.9;
        z-index: -1;
    }

    .modal-container {
        opacity: 1;
        height: 50%;
        width: 60%;
        background-color: var(--c1);
        border: 1px solid var(--c2);
        border-radius: 1em;
    }

    h1 {
        color: var(--blue);
        text-align: center;
    }
</style>
