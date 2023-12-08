<script lang="ts">
    import Record from "./Record.svelte";

    let newDate = new Date();
    let select = { day: 0, row: 0, offset: 0 };

    $: current = {
        day: new Date(newDate.getFullYear(), newDate.getMonth() + select.offset).getDay() - 1,
        month: new Date(newDate.getFullYear(), newDate.getMonth() + select.offset).getMonth(),
        year: new Date(newDate.getFullYear(), newDate.getMonth() + select.offset).getFullYear(),
        days: new Date(newDate.getFullYear(), newDate.getMonth() + select.offset + 1, 0).getDate(),
    };

    $: previous = {
        days: new Date(newDate.getFullYear(), newDate.getMonth() + select.offset, 0).getDate(),
    };

    $: weekdays = new Array(7).fill(0).map((_, i) => {
        return new Date(null, null, i + 1).toLocaleDateString("en-EN", { weekday: "long" });
    });

    $: week = Math.ceil((new Date(current.year, current.month).getTime() - new Date(current.year, 0, 1).getTime() + 86400000) / 86400000 / 7);

    const onWheel = e => {
        if (e.deltaY > 0) select.offset++;
        else select.offset--;
    };

    const onClick = e => {
        if (select.day == e.target.innerText) {
            select.day = null;
            select.row = null;
        } else {
            select.day = Number.parseInt(e.target.innerText);
            select.row = Math.min(Math.floor(e.target.id / 7) + 3, 6);
        }
    };
</script>

<!-- <svelte:window on:keydown={handler} /> -->

<div class="container" on:wheel={onWheel}>
    <div class="title">
        {#each [-2, -1, 0, 1, 2] as i}
            <div on:click={() => (i == 0 ? (select.offset = 0) : (select.offset += i))}>
                {new Date(current.year, current.month + i).toLocaleDateString("en-EN", { month: "long", year: i == 0 ? "numeric" : undefined })}
            </div>
        {/each}
    </div>

    <div class="calendar-grid" style="grid-template-rows: 3cqh repeat({select.row ? 7 : 5}, 1fr);">
        <div />
        <!-- Add weedays to top -->
        {#each weekdays as weekday}
            <div class="weekday">{weekday}</div>
        {/each}

        <!-- Add weeks to left -->
        {#each [2, 3, 4, 5, 6] as n, i}
            <div class="week" style="grid-row: {select.row && select.row <= n ? n + 2 : n}/{select.row && select.row <= n ? n + 2 : n};">{week + i}</div>
        {/each}

        <!-- Add days -->
        {#each Array(35) as _, i}
            {#if i < current.day}
                <div class="prev">{previous.days - current.day + i + 1}</div>
            {:else if i < current.day + current.days}
                <div id={i.toString()} class={select.day == i + 1 - current.day ? "curr" : "day"} on:click={onClick}>{i + 1 - current.day}</div>
            {:else}
                <div class="next">{i + 1 - (current.day + current.days)}</div>
            {/if}
        {/each}

        {#if select.row}
            <div style="grid-row: {select.row} / {select.row + 2}; grid-column: 1 / 9;"><Record /></div>
        {/if}
    </div>
</div>

<style>
    .container {
        flex: 1;
        display: flex;
        flex-direction: column;
        padding: 8cqw;
        background-color: var(--c0);
    }

    .title {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
        align-items: flex-end;
        color: var(--inactive);

        & div:nth-child(2) {
            font-size: 20px;
        }

        & div:nth-child(3) {
            font-size: 30px;
        }

        & div:nth-child(4) {
            font-size: 20px;
        }
    }

    .calendar-grid {
        flex: 1;
        display: grid;
        grid-template-columns: 2cqw repeat(7, 1fr);
        column-gap: min(0.4cqh, 0.4cqw);
        row-gap: min(0.4cqh, 0.4cqw);
        margin-left: -2cqw;

        & div {
            display: flex;
            justify-content: flex-end;
            color: var(--inactive);
            overflow: hidden;
        }

        & .weekday {
            display: flex;
            justify-content: center;
            align-items: flex-end;
        }

        & .week {
            display: flex;
            justify-content: center;
            align-items: center;
            grid-column: 1/1;
        }

        & .day {
            color: var(--blue);
            background-color: var(--c2);
            font-variant: tabular-nums;
        }

        & .prev {
            background-color: var(--c1);
        }

        & .next {
            background-color: var(--c1);
        }
    }

    /* .calendar-container {
        flex: 1;
        display: flex;
        flex-direction: row;
    }

    .calendar {
        flex: 1;
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        column-gap: 0.4cqh;
        row-gap: 0.4cqh;
        padding: 0.5cqh;
        overflow: hidden;
    }

    .calendar > div {
        display: flex;
        align-items: center;
        justify-content: center;
        position: relative;
        background-color: var(--c2);
        border: 1px solid transparent;

        transition: all 0.05s ease-in-out;

        border-radius: 0.2em;
        overflow: hidden;
    }

    .calendar > div:nth-child(1) {
        border-top-left-radius: 1em;
    }

    .calendar > div:nth-child(7) {
        border-top-right-radius: 1em;
    }

    .calendar > div:nth-child(29) {
        border-bottom-left-radius: 1em;
    }

    .calendar > div:nth-child(35) {
        border-bottom-right-radius: 1em;
    }

    .calendar > div.highlight > h2 {
        opacity: 0.8;
        color: var(--highlight);
    }

    .calendar > div > h2 {
        position: absolute;
        pointer-events: none;

        font-size: 10cqw;
        color: var(--white);
        opacity: 0.6;

        margin: 0;
        right: -0.5cqw;
        top: -2.5cqw;
    }

    .calendar > div.disabled {
        background-color: var(--c1);

        & h2 {
            color: var(--inactive);
        }
    }

    .calendar > div:hover:not(.disabled):not(.selected) {
        background-color: var(--c3);

        & h2 {
            opacity: 0.5;
        }
    }

    .calendar > .selected {
        background-color: var(--blue);
    }

    h1 {
        color: var(--blue);
        text-align: center;
    }

    .days {
        display: flex;
        flex-direction: row;
        color: var(--c6);
    }

    .week {
        display: grid;
        row-gap: 1cqh;
        align-items: center;

        width: 1.5cqw;
        /* height: calc(100% - 1cqh);
        margin-left: -1.5cqw;

        color: var(--c6);
        font-size: 1cqw;
        text-align: right;

        overflow: hidden;
        padding-top: 0.5cqh;
        padding-bottom: 0.5cqh;
    }

    .week > div {
        display: none;
    }

    .week > h3 {
        margin: 0;
        font-variant: tabular-nums;
        font-weight: bold;
    }

    .week > .highlight {
        color: var(--highlight);
    }

    .days > h3 {
        flex: 1;
        margin: 0;
        text-align: center;
        text-transform: capitalize;
    } */
</style>
