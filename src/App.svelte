<script lang="ts">
    import Calendar from "./Calendar.svelte";
    import Settings from "./Settings.svelte";

    let settings = false;
</script>

<main>
    <Settings bind:settings />
    <Calendar />
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

        --gray1: #1b1b1b;
        --gray2: #2b2b2b;
        --gray3: #3b3b3b;
        --gray4: #4b4b4b;

        --inactive: #70757a;
        --text: #3c4043;

        --red: #ea4335;
        --yellow: #fbbc05;
        --green: #34a853;
        --blue: #4285f4;
        --highlight: #1a73e8;
    }
</style>
