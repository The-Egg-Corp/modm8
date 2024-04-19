<script lang="ts" setup>
import GameSelection from './views/GameSelection.vue'
import Settings from './views/Settings.vue'

import Sidebar from './components/Sidebar.vue'
import Topbar from './components/Topbar.vue'
</script>

<template>
  <GameSelection/>
  <Settings/>

  <Topbar/>
  <Sidebar/>
</template>

<style>
#logo {
  user-select: none;
  pointer-events: none;
}

.drag {
  --wails-draggable: drag;
}

.no-drag {
  --wails-draggable: no-drag;
}
</style>