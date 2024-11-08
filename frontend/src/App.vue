<script lang="ts" setup>
import { onBeforeUnmount, onMounted } from 'vue'
import { changeLocale } from '@i18n'

import { 
    NumCPU,
    GetSettings
} from '@backend/app/Application'

import { 
    Sidebar, ControlButtons,
    SettingsOverlay, AppInfoOverlay
} from '@components'

import { useAppStore } from '@stores'
import { getOpenDialogs } from '@composables'

const { setMaxThreads } = useAppStore()

//const settingsDialog = useDialog('settings')

// Keydown event listener
const onKeydown = (event: KeyboardEvent) => {
    console.log('Keydown fired!\n', event.key)

    const openDialogs = getOpenDialogs()
    if (openDialogs.length > 0) {
        for (const dialog of openDialogs) {
            dialog.visible.value = false
        }

        return
    }

    // No dialogs open + pressing ESC -> open settings overlay.
    // if (event.key == 'Escape') {
    //     settingsDialog.setVisible(!settingsDialog.visible.value)
    // }
}

onMounted(async () => {
    const settings = await GetSettings()
    changeLocale(settings.general.locale)

    // We don't use `navigator.hardwareConcurrency` as it is known to be unreliable.
    setMaxThreads(await NumCPU())

    // Add keydown event listener when the app mounts
    window.addEventListener('keydown', onKeydown)
})

onBeforeUnmount(() => {
    window.removeEventListener('keydown', onKeydown)
})
</script>

<template>
<div class="topbar drag">
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
.topbar {
    z-index: 9999;
    position: fixed;
    display: flex;
    justify-content: end;
    width: 100%;
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