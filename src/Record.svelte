<script lang="ts">
    import Icon from "./Icon.svelte";

    export let style: string;
</script>

<div {style} class="container">
    <div class="manage">
        <Icon fill="var(--inactive)" />
        <div class="work">Add work</div>
        <div class="work">Add study</div>
        <div class="work">Add custom</div>
    </div>
    <div class="records">
        {#each [1, 2, 3, 4, 5, 6, 7, 8, 9] as i}
            <div>1</div>
        {/each}
    </div>
</div>

<style>
    .container {
        flex: 1;
        display: flex;
        flex-direction: row;
        align-items: center;
        color: var(--inactive);
        border-radius: 1cqw;
        gap: 4cqw;
    }

    .manage {
        flex: 1;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 4cqh;

        & .work {
            font-size: 2em;
            color: var(--inactive);
        }
    }

    .records {
        flex: 2;
        flex-direction: column;
        background-color: var(--gray3);
        font-size: 2em;
        gap: 4cqw;
    }
</style>
