<script lang="ts" setup>
import CardOverlay from "./reusable/CardOverlay.vue"
import { useDialog } from "@composables"

const { 
    setVisible,
    visible, closable, draggable
} = useDialog('app-info')
</script>

<template>
<div :class="['app-info-container', { 'no-drag': visible }]">
    <CardOverlay 
        class="settings-dialog no-drag"
        v-model:visible="visible"
        v-model:closable="closable"
        v-model:draggable="draggable"
    >
        <template #cardContent>
            <div class="flex column align-items-center">
                <h1 class="header">App Info</h1>
                <p style="font-weight: 305; margin-bottom: 10px; margin-top: 5px;">
                    v0.0.1
                </p>

                <Button plain outlined severity="help" class="flex w-2" icon="pi pi-heart" label="Donate">
                </Button>
            </div>
        </template>

        <template #dialogContent>
            <div class="flex justify-content-end gap-3 mt-3">
                <Button class="w-full" type="button" :label="$t('keywords.close')" severity="secondary" @click="setVisible(false)"></Button>
            </div>
        </template>
    </CardOverlay>
</div>
</template>

<style scoped>
.app-info-container {
    display: flex;
}
</style>