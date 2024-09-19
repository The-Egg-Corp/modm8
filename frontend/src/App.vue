<script lang="ts" setup>
import { 
    Sidebar, ControlButtons,
    SettingsOverlay, AppInfoOverlay
} from '@components'

import { onMounted } from 'vue'
import { changeLocale } from '@i18n'

import { 
    NumCPU,
    GetSettings
} from '@backend/app/Application'

import { useAppStore } from '@stores'

const appStore = useAppStore()
const {
    setMaxThreads
} = appStore

onMounted(async () => {
    const settings = await GetSettings()
    changeLocale(settings.general.locale)

    setMaxThreads(await NumCPU())
})
</script>

<template>
<ControlButtons/>
  
<div class="app-container">
    <Sidebar/>

    <SettingsOverlay/>
    <AppInfoOverlay/>

    <RouterView/>
</div>
</template>

<style>
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