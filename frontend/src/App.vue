<script lang="ts" setup>
import Sidebar from './components/Sidebar.vue'
import Topbar from './components/Topbar.vue'
</script>

<template>
  <RouterView />

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

.no-drag-all-children > * {
  --wails-draggable: no-drag;
}
</style>