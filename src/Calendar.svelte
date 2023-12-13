<script lang="ts">
    import Record from "./Record.svelte";

    let date = new Date();
    let select = { day: 0, row: 0, offset: 0 };
    let client = { width: 0, height: 0 };
    const cutoff = 1600;

    $: current = {
        day: new Date(date.getFullYear(), date.getMonth() + select.offset).getDay() - 1,
        month: new Date(date.getFullYear(), date.getMonth() + select.offset).getMonth(),
        year: new Date(date.getFullYear(), date.getMonth() + select.offset).getFullYear(),
        days: new Date(date.getFullYear(), date.getMonth() + select.offset + 1, 0).getDate(),
    };

    $: previous = {
        days: new Date(date.getFullYear(), date.getMonth() + select.offset, 0).getDate(),
    };

    $: weekdays = new Array(7).fill(0).map((_, i) => {
        return new Date(null, null, i + 1).toLocaleDateString("en-EN", { weekday: "long" });
    });

    $: week = {
        current: Math.floor((current.day + date.getDate()) / 7),
        month: Math.ceil((new Date(current.year, current.month, current.day).getTime() - new Date(current.year, 0, 1).getTime() + 86400000) / 86400000 / 7),
    };

    $: isWideScreen = client.width >= cutoff && select.row ? true : false;

    $: stringCutOff = (input: string) => {
        if (client.width <= cutoff / 3) return input.substring(0, 1);
        else if (client.width <= cutoff) return input.substring(0, 3);
        return input;
    };

    const onWheel = e => {
        if (e.deltaY > 0) select.offset++;
        else select.offset--;
    };

    const onClick = e => {
        if (select.day == e.target.innerText) {
            select.day = select.row = null;
        } else {
            select.day = Number.parseInt(e.target.innerText);
            select.row = Math.min(Math.floor(e.target.id / 7) + 3, 6);
        }
    };
</script>

0<!-- <svelte:window on:keydown={handler} /> -->

<div class="container" on:wheel={onWheel} bind:clientHeight={client.height} bind:clientWidth={client.width}>
    <div class="title">
        {#each [-2, -1, 0, 1, 2] as i}
            <div class={select.offset == -i && "current"} on:click={() => (i == 0 ? (select.offset = 0) : (select.offset += i))}>
                {stringCutOff(new Date(current.year, current.month + i).toLocaleDateString("en-EN", { month: "long" }))}
                {#if i == 0}{current.year}{/if}
            </div>
        {/each}
    </div>

    <div class="calendar-grid" style="grid-template-rows: 3cqh repeat({isWideScreen ? 7 : 5}, 1fr);">
        <div class="corner" />
        {#each weekdays as w, i}
            <div class="weekday">
                <div class={i == date.getDay() - 1 && !select.offset && "current"}>{stringCutOff(w)}</div>
            </div>
        {/each}
        {#each [2, 3, 4, 5, 6] as n, i}
            <div class="week" style="grid-row: {isWideScreen && select.row <= n ? n + 2 : n}/{isWideScreen && select.row <= n ? n + 2 : n};">
                <div class={i == week.current && !select.offset && "current"}>
                    {week.month + i}
                </div>
            </div>
        {/each}
        {#each Array(35) as _, i}
            <div class="day">
                {#if i < current.day}
                    <div class="prev">
                        {previous.days - current.day + i + 1}
                    </div>
                {:else if i < current.day + current.days}
                    {#if i + 1 - current.day == date.getDate() && i + 1 - current.day == select.day && !select.offset}
                        <div id={i.toString()} class="selected current" on:click={onClick}>
                            {i + 1 - current.day}
                        </div>
                    {:else if i + 1 - current.day == select.day}
                        <div id={i.toString()} class="selected" on:click={onClick}>
                            {i + 1 - current.day}
                        </div>
                    {:else if i + 1 - current.day == date.getDate() && !select.offset}
                        <div id={i.toString()} class="current" on:click={onClick}>
                            {i + 1 - current.day}
                        </div>
                    {:else}
                        <div id={i.toString()} class="day" on:click={onClick}>
                            {i + 1 - current.day}
                        </div>
                    {/if}
                {:else}
                    <div class="next">
                        {i + 1 - (current.day + current.days)}
                    </div>
                {/if}
            </div>
        {/each}
        {#if isWideScreen}
            <Record style="grid-row: {select.row} / {select.row + 2}; grid-column: 1 / 9;" />
        {/if}
    </div>
    {#if !isWideScreen && select.row}
        <Record style="flex: 1;" />
    {/if}
</div>

<style>
    .container {
        flex: 1;
        display: flex;
        flex-direction: column;
        background-color: var(--gray1);
        padding: 8cqw;
        gap: 2cqh;
    }

    .title {
        display: flex;
        flex-direction: row;
        align-items: flex-end;
        color: var(--inactive);

        & div {
            flex: 1;
            text-align: center;
            text-wrap: nowrap;
            font-size: 3em;
        }

        & :is(:nth-child(2), :nth-child(4)) {
            font-size: 2.5em;
        }

        & :is(:nth-child(1), :nth-child(5)) {
            font-size: 2em;
        }
    }

    .calendar-grid {
        flex: 1;
        display: grid;
        grid-template-columns: 2cqw repeat(7, 1fr);
        column-gap: min(0.4cqh, 0.4cqw);
        row-gap: min(0.4cqh, 0.4cqw);
        margin-left: -2cqw;

        & .day {
            flex: 1;
            display: flex;
            overflow: hidden;
            color: var(--inactive);

            & div {
                flex: 1;
                display: flex;
                justify-content: flex-end;
                font-size: 5em;
                background-color: var(--gray3);
                border-radius: 1cqw;
            }

            & .selected {
                font-weight: 600;
                text-shadow:
                    -1px 0 var(--inactive),
                    0 1px var(--inactive),
                    1px 0 var(--inactive),
                    0 -1px var(--inactive);
                color: var(--gray1);
                background-color: var(--gray4);
            }

            & :is(.prev, .next) {
                background-color: var(--gray2);
            }
        }

        & .weekday {
            flex: 1;
            display: flex;
            justify-content: center;
            align-items: flex-end;
            color: var(--inactive);

            & div {
                font-size: 2em;
            }
        }

        & .week {
            flex: 1;
            display: flex;
            justify-content: flex-end;
            align-items: center;
            color: var(--inactive);

            & div {
                font-size: 2em;
            }
        }
    }

    .current {
        font-weight: 200 !important;
        color: var(--blue) !important;
    }
</style>
