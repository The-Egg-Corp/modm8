<script lang="ts" setup>
import LanguageDropdown from '../components/LanguageDropdown.vue'
</script>

<template>
<div class="container">
    <h1 class="header">{{ $t('keywords.settings') }}</h1>

    <div class="language-setting">
        <div class="flex-item">
            <h2>{{ $t('keywords.language') }}</h2>
        </div>
        <div class="flex-item">
            <LanguageDropdown/>
        </div>
    </div>
</div>
</template>

<style scoped>
.container {
    position: absolute;
    top: 50px;
    left: 50%;
    transform: translate(-50%);
}

.container .header {
    user-select: none;
    font-size: 48px;
    font-weight: 500;
    text-align: center;
    margin-top: 0;
    margin-bottom: 0.5rem;
}

.language-setting {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    gap: 25px;
}

.flex-item {
    text-align: center;
    user-select: none;
}

.flex-item h2 {
    font-weight: 400;
    font-size: 27px;
}
</style>