<script lang="ts">
    export let name: string;
    var currentDate = new Date();
    var offSet = 0;
    var clientWidth = 0;
    var modal = true;

    var dayArray = new Array(7).fill(0).map((_, i) => {
        var date = new Date(null, null, i + 1);
        return date.toLocaleDateString("en-EN", { weekday: "long" });
    });

    $: daysInMonth = new Array(2).fill(0).map((_, i) => {
        return new Date(currentDate.getFullYear(), currentDate.getMonth() + i + offSet, 0).getDate();
    });

    $: firstDay = new Date(currentDate.getFullYear(), currentDate.getMonth() + offSet, 1).getDay() - 1;

    $: calendar = new Array(35).fill(0).map((_, i) => {
        if (i < firstDay) {
            return daysInMonth[0] - firstDay + i + 1;
        } else if (i >= firstDay && i < firstDay + daysInMonth[1]) {
            return i - firstDay + 1;
        } else {
            return i - (firstDay + daysInMonth[1]) + 1;
        }
    });

    function handleContextMenu(event) {
        event.preventDefault();
    }

    function handleClick(event) {
        event.preventDefault();
        modal = true;
        console.log(event.target.id);
    }
</script>

<main>
    {#if modal}
        <div class="modal">
            <div class="blur" on:click|self={() => (modal = false)} />
            <div class="modal-container">
                <h1>Modal</h1>
            </div>
        </div>
    {/if}
    <div class="container">
        <div class="title">
            <h1 on:click={() => offSet--}>{`<<`}</h1>
            <h1 on:click={() => (offSet = 0)}>
                {new Date(currentDate.getFullYear(), currentDate.getMonth() + offSet).toLocaleDateString("en-EN", { month: "long", year: "numeric" })}
            </h1>
            <h1 on:click={() => offSet++}>{`>>`}</h1>
        </div>

        <div bind:clientWidth class="days">
            {#each dayArray as day}
                <h3>
                    {clientWidth > 1000 ? day : day.slice(0, 3)}
                </h3>
            {/each}
        </div>
        <div class="calendar">
            {#each calendar as day, i}
                {#if i < firstDay || i > firstDay + daysInMonth[0]}
                    <div class="disabled">
                        <h2>{day}</h2>
                    </div>
                {:else}
                    <div
                        id={i.toString()}
                        class={offSet == 0 && i == firstDay + currentDate.getDate() - 1 ? "highlighted" : "day"}
                        on:contextmenu={handleContextMenu}
                        on:click={handleClick}
                    >
                        <h2>{day}</h2>
                    </div>
                {/if}
            {/each}
        </div>
    </div>
</main>

<style>
    main {
        display: flex;
        height: 100%;
        width: 100%;

        -webkit-touch-callout: none;
        -webkit-user-select: none;
        -khtml-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;

        --c0: #181a1c;
        --c1: #23282a;
        --c2: #313436;
        --c3: #3e4143;
        --c4: #3b627d;
        --c5: #d4dbde;
        --c6: #a2a7aa;
    }

    .modal {
        display: flex;
        position: absolute;
        height: 100%;
        width: 100%;
        justify-content: center;
        align-items: center;
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
        width: 50%;
        background-color: var(--c1);
        border: 1px solid var(--c2);
        border-radius: 1em;
    }

    .container {
        flex: 1;
        display: flex;
        flex-direction: column;
        padding: 4cqw;
        background-color: var(--c0);
    }

    .title {
        display: flex;
        flex-direction: row;
        gap: 1em;
        justify-content: space-around;
    }

    .title > h1 {
        color: var(--c4);
    }

    .calendar {
        display: grid;
        height: 100%;
        grid-template-columns: repeat(7, 1fr);
        column-gap: 10px;
        row-gap: 12px;
        padding: 0.5em;
        border: 1px solid transparent;

        overflow: hidden;
    }

    .calendar > div {
        display: flex;
        height: 100%;
        align-items: center;
        justify-content: center;
        position: relative;

        background-color: var(--c2);
        border: 1px solid var(--c2);
        transition: all 0.1s ease-in-out;

        border-radius: 0.2em;
        /* box-shadow: 1px 1px 6px black; */
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

    .calendar > div.disabled {
        background-color: var(--c1);
    }

    .calendar > div.highlighted {
        border-color: var(--c4);
    }

    .calendar > div.highlighted > h2 {
        color: var(--c4);
    }

    .calendar > div > h2 {
        position: absolute;
        pointer-events: none;

        font-size: 10cqw;
        color: var(--c6);
        opacity: 0.3;

        margin: 0;
        right: -1cqw;
        top: -3.5cqw;
    }

    .calendar > div:hover:not(.disabled) {
        background-color: var(--c3);

        & h2 {
            opacity: 0.5;
        }
    }

    h1 {
        color: var(--c4);
        text-align: center;
    }

    .days {
        display: flex;
        flex-direction: row;
        color: var(--c4);
    }

    .days > h3 {
        flex: 1;
        margin: 0.5cqw;

        text-align: center;
        text-transform: capitalize;
    }
</style>
