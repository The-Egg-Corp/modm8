<script lang="ts" setup>
import LanguageDropdown from '../components/LanguageDropdown.vue';
</script>

<template>
    <div class="container">
        <p class="header">{{ $t('keywords.settings') }}</p>

        <LanguageDropdown/>
    </div>
</template>

<style scoped>
.container {
    position: absolute;
    top: 30px;
    left: 50%;
    transform: translate(-50%);
}

.container .header {
    user-select: none;
    font-size: 45px;
}
</style>