<script lang="ts" setup>

</script>

<template>
    <div class="game-selection flex-span">
        <h2 class="header">{{ $t('game-selection.header') }}</h2>
    </div>
</template>

<style scoped>
.game-selection {
    padding-top: 65px;
    justify-content: center; /* Center horizontally */
}

.game-selection .header {
    text-wrap: nowrap;
    user-select: none;
    font-size: 30px;
    font-weight: 420;
    margin-top: 0;
    margin-bottom: 0.5rem;
    text-align: center;
}
</style>