<script lang="ts">
    export let style: string;
</script>

<div {style}></div>

<style>
    div {
        flex: 1;
        display: flex;
        flex-direction: row;
        background-color: var(--gray3);
        border-radius: 1cqw;
    }
</style>
