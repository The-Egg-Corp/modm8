<script lang="ts" setup>
import Dashboard from './views/Dashboard.vue'
import Settings from './views/Settings.vue'

import Sidebar from './components/Sidebar.vue'
import Topbar from './components/Topbar.vue'
</script>

<template>
  <Dashboard/>
  <Settings/>

  <Sidebar/>
  <Topbar/>
</template>

<style>
#logo {
  user-select: none;
  pointer-events: none;
}
</style>