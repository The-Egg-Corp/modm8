<script lang="ts" setup>
import LanguageDropdown from '../components/settings/LanguageDropdown.vue'
import ThemeDropdown from '../components/settings/ThemeDropdown.vue'
</script>

<template>
<div class="settings">
    <h1 class="header">{{ $t('keywords.settings') }}</h1>

    <div class="setting">
        <div class="flex-item">
            <h3>{{ $t('keywords.language') }}</h3>
        </div>
        <div class="flex-item">
            <LanguageDropdown/>
        </div>
    </div>

    <div class="setting">
        <div class="flex-item">
            <h3>{{ $t('keywords.theme') }}</h3>
        </div>
        <div class="flex-item">
            <ThemeDropdown/>
        </div>
    </div>
</div>
</template>

<style scoped>
.settings {
    padding-top: 40px;
    padding-left: 60px;
}

.settings .header {
    user-select: none;
    font-size: 80px;
    font-weight: 480;
    margin-top: 0;
    margin-bottom: 10px;
}

.setting {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: first baseline;
    padding-left: 10px;
    gap: 80px;
}

.flex-item {
    text-align: center;
    user-select: none;
    font-size: 25px;
}

.flex-item h3 {
    font-weight: 420;
}
</style>