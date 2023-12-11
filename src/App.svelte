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

        --c0: #222222;
        --c1: #282828;
        --c2: #383838;
        --c3: #3e4143;
        --white: #ffffff;
        --c6: #949494;

        --inactive: #70757a;
        --text: #3c4043;

        --red: #ea4335;
        --yellow: #fbbc05;
        --green: #34a853;
        --blue: #4285f4;
        --highlight: #1a73e8;

        /* font-family: "Google Sans", Roboto, Arial, sans-serif; */
    }
</style>
