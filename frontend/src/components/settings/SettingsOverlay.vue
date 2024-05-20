<script lang="ts" setup>
import Dialog from 'primevue/dialog'
import LanguageDropdown from './LanguageDropdown.vue'
import ThemeDropdown from '@components/settings/ThemeDropdown.vue'

import { useDialog } from '@composables'
const { visible, draggable, closable } = useDialog()

const dialogStyle = {
    "margin-left": "70px",
    "width": '40rem',
    "background": 'none'
}

const cardStyle = {
    "box-shadow": 'none',
    "overflow": 'hidden',
    "margin": '-15px'
}
</script>

<template>
    <div class="settings-container">
        <Dialog 
            v-model:visible="visible"
            v-model:draggable="draggable" 
            v-model:closable="closable"
            :block-scroll="true"
            :dismissable-mask="true"
            modal class="settings-dialog no-drag"
            :style="dialogStyle"
        >
            <Card :style=cardStyle>
                <template #content>
                    <div class="card-content">
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
            </Card>
        </Dialog>
    </div>
</template>

<style scoped>
:deep(.p-dialog-title) {
    user-select: none;
}

.settings-container {
    display: flex;
}

.settings-dialog {
    width: 40rem;
}

.card-content {
    justify-content: center;
}

.card-content .header {
    user-select: none;
    font-size: 45px;
    font-weight: 520;
    margin-top: -15px;
    margin-bottom: 10px;
}

.setting {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: first baseline;
    padding-left: 5px;
}

.flex-item {
    text-align: center;
    user-select: none;
    font-size: 20px;
}

.flex-item h3 {
    font-weight: 380;
}
</style>