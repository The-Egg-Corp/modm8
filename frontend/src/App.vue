<script lang="ts" setup>
import Sidebar from './components/Sidebar.vue'
import ControlButtons from './components/ControlButtons.vue'
import SettingsOverlay from './components/settings/SettingsOverlay.vue'

import { onMounted } from 'vue'
import { changeLocale } from '@i18n'
import { GetSettings, NumCPU } from '@backend/core/App'
import { core } from '@backend/models'
import { useAppStore } from '@stores'

let settings: core.AppSettings

onMounted(async () => {
    settings = await GetSettings()
    console.log(settings)

    changeLocale(settings.general.locale)

    const store = useAppStore()
    store.setMaxThreads(await NumCPU())
})
</script>

<template>
<ControlButtons/>
  
<div class="app-container">
  <Sidebar/>
  <SettingsOverlay/>
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
</style>