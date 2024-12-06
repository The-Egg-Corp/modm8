<script lang="ts" setup>
import { openLink } from "@frontend/src/util"

import { version } from "@frontend/package.json" with { type: "json" }

import { CardOverlay } from "@components"
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
            <div class="flex column">
                <div class="flex row justify-content-between align-items-center" style="padding-top: 1rem;">
                    <h1 class="heading">App Information</h1>
                    <Button 
                        style="height: 40px;" icon="pi pi-times" severity="secondary"
                        @click="setVisible(false)" 
                    />
                </div>

                <div class="flex column">
                    <div>
                        <h3 class="subheading mb-0 mt-2">Version</h3>
                        <p style="margin-top: 5px; font-weight: 250;">v{{ version }}</p>
                    </div>

                    <div>
                        <h3 class="subheading">Links</h3>
                        <div class="flex row gap-1">
                            <Button class="w-full" severity="secondary" raised
                                icon="pi pi-github" label="View source code"
                                @click="openLink('https://github.com/The-Egg-Corp/modm8')"
                            />

                            <Button class="discord-btn w-full" severity="secondary" raised
                                icon="pi pi-discord" label="Join the community"
                                @click="openLink('https://discord.gg/64Vq7cpdGV')"
                            />

                            <Button class="donate-btn btn w-7" severity="secondary" raised
                                icon="pi pi-heart-fill" label="Donate"
                                @click="openLink('https://github.com/sponsors/Owen3H')"
                            />
                        </div>
                    </div>
                </div>
            </div>
        </template> 
    </CardOverlay>
</div>
</template>

<style scoped>
.heading {
    font-weight: 590;
    font-size: 34px;
    margin: 0;
}

.subheading {
    font-weight: 440;
    font-size: 22px;
    margin-top: 5px;
    margin-bottom: 10px;
}

.app-info-container {
    display: flex;
}

.donate-btn :deep(.p-button-icon) {
    color: var(--p-red-400);
}

.donate-btn :deep(.p-button-label) {
    color: var(--p-red-400);
}

.discord-btn :deep(.p-button-icon) {
    color: var(--p-blue-400);
}

.discord-btn :deep(.p-button-label) {
    color: var(--p-blue-400);
}
</style>