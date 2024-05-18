<script lang="ts" setup>
import Sidebar from './components/Sidebar.vue'
import ControlButtons from './components/ControlButtons.vue'

import { onMounted } from 'vue'
import { changeLocale } from '@i18n'
import { GetSettings } from '@backend/core/App'
import { core } from '@backend/models'

let settings: core.AppSettings

onMounted(async () => {
    settings = await GetSettings()
    console.log(settings)

    await changeLocale(settings.locale)
})
</script>

<template>
<ControlButtons/>
  
<div class="app-container">
  <Sidebar/>
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