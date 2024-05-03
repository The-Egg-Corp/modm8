<script lang="ts" setup>
import Sidebar from './components/Sidebar.vue'
import ControlButtons from './components/ControlButtons.vue'
</script>

<template>
  <ControlButtons/>
  <Sidebar/>

  <RouterView />
</template>

<style>
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