<script lang="ts">
    function handler(event) {}
</script>

<svelte:window on:keydown={handler} />

<div>
    <div></div>
    <div></div>
</div>

<style>
    div {
        flex: 1;
        display: flex;
        flex-direction: row;
        background-color: var(--c1);
    }
</style>
