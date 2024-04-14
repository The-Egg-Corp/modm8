<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <HelloWorld/>
</template>

<style>
#logo {
  display: block;
  width: 45%;
  height: 45%;
  min-width: 30%;
  min-height: 30%;
  align-self: center;
  object-fit: contain;
  margin: auto;
  padding: 8% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
