<script lang="ts" setup>
import Dashboard from './views/Dashboard.vue'
import Settings from './views/Settings.vue'
import Sidebar from './components/Sidebar.vue'
</script>

<template>
  <Sidebar/>
  <Dashboard/>
  <Settings/>
</template>

<style>
#logo {
  display: block;
  width: 25%;
  height: 25%;
  max-width: 50%;
  max-height: 50%;
  align-self: center;
  object-fit: contain;
  margin: auto;
  padding: 3% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
  user-select: none;
  pointer-events: none;
}
</style>