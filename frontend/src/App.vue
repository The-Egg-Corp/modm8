<script lang="ts" setup>
import { onBeforeUnmount, onMounted } from 'vue'
import { changeLocale } from '@i18n'

import { GetSettings } from '@backend/app/Application'
import { NumCPU } from '@backend/appctx/Utils'

import { 
    Sidebar, ControlButtons,
    SettingsOverlay, AppInfoOverlay
} from '@components'

import { useAppStore } from '@stores'
import { getOpenDialogs } from '@composables'
import { storeToRefs } from 'pinia'

const appStore = useAppStore()
const { 
    sidebarWidthPx, 
    showBackButton
} = storeToRefs(appStore)

const { setMaxThreads } = useAppStore()

//const settingsDialog = useDialog('settings')

// Keydown event listener
const onKeydown = (event: KeyboardEvent) => {
    //console.log('Keydown fired!\n', event.key)

    if (event.key == "Escape") {
        const openDialogs = getOpenDialogs()
        if (openDialogs.length > 0) {
            for (const dialog of openDialogs) {
                dialog.visible.value = false
            }
        
            return
        }

        // No dialogs open + ESC -> open settings overlay.
        // settingsDialog.setVisible(!settingsDialog.visible.value)
    }
}

onMounted(async () => {
    const settings = await GetSettings()
    changeLocale(settings.general.locale)

    // We don't use `navigator.hardwareConcurrency` as it is known to be unreliable.
    setMaxThreads(await NumCPU())

    window.addEventListener('keydown', onKeydown)
})

onBeforeUnmount(() => {
    window.removeEventListener('keydown', onKeydown)
})
</script>

<template>
<div class="topbar drag">
    <Button text icon="pi pi-arrow-left" 
        class="back-btn"
        :style="{ visibility: showBackButton ? 'visible' : 'hidden' }"
    />

    <ControlButtons/>
</div>

<div class="app-container">
    <Sidebar/>

    <SettingsOverlay/>
    <AppInfoOverlay/>

    <RouterView/>
</div>
</template>

<style>
/* TODO: Are these even needed now?
:global(.p-dialog) {
    border: 1px solid #4d4d4d;
    border-radius: 9px;
}

:global(.p-dialog-content) {
    background: none;
    border-radius: 9px;
    padding: 1rem;
}

:global(.p-card-body) {
    padding-top: 0rem;
}

:global(.p-divider-content) {
    background: none;
}*/

.topbar {
    z-index: 9999;
    width: 100%;
    position: fixed;
    display: flex;
    justify-content: space-between;
    margin-left: v-bind(sidebarWidthPx);
    padding-right: v-bind(sidebarWidthPx);
    background-color: var(--topbar-bg-color);
}

.back-btn {
    color: rgba(201, 201, 201, 0.8) !important;
    width: 32px !important;
    margin-left: 20px;
}

.app-container {
    display: flex;
}

#logo {
    user-select: none;
    pointer-events: none;
}

.no-select {
    user-select: none;
}

.drag {
    --wails-draggable: drag;
}

.no-drag {
    --wails-draggable: no-drag;
}

.no-drag-all-children > * {
    --wails-draggable: no-drag;
}

:global(.p-tooltip-arrow) {
    border-left-color: transparent !important;
}
</style>