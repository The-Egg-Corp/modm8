<script lang="ts" setup>

</script>

<template>
    <div class="container">
        <p class="header">{{ $t('game-selection.header') }}</p>
    </div>
</template>

<style scoped>
.container {
    position: absolute;
    top: 40px;
    left: 50%;
    transform: translate(-50%);
}

.container .header {
    text-wrap: nowrap;
    user-select: none;
    font-size: 26px;
}
</style>