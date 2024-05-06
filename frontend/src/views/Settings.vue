<script lang="ts" setup>
import LanguageDropdown from '../components/LanguageDropdown.vue'
</script>

<template>
<div class="settings">
    <h1 class="header">{{ $t('keywords.settings') }}</h1>

    <div class="language-setting">
        <div class="flex-item">
            <h3>{{ $t('keywords.language') }}</h3>
        </div>
        <div class="flex-item">
            <LanguageDropdown/>
        </div>
    </div>
</div>
</template>

<style scoped>
.settings {
    padding-top: 50px;
    padding-left: 45px;
}

.settings .header {
    user-select: none;
    font-size: 52px;
    font-weight: 500;
    margin-top: 0;
    margin-bottom: 1rem;
}

.language-setting {
    display: flex;
    padding-left: 15px;
    align-items: center;
    gap: 40px;
}

.flex-item {
    text-align: center;
    user-select: none;
    font-size: 21px;
}

.flex-item h3 {
    font-weight: 420;
}
</style>