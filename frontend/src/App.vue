<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <HelloWorld/>
</template>

<style>
#logo {
  display: flex;
  width: 40%;
  height: 40%;
  max-width: 90%;
  max-height: 90%;
  align-self: center;
  object-fit: cover;
  margin: auto;
  padding: 12% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
