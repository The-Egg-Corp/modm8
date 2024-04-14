<script lang="ts" setup>
import PackageInfo from './components/PackageInfo.vue'
</script>

<template>
  
<div class="container">
  <img id="logo" alt="modm8 icon" src="./assets/images/modm8-logo-transparent-white-donut.png"/>
  <PackageInfo/>
</div>
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

.container {
  user-select: none;
  position: absolute;
  left: 0;
  right: 0;
}
</style>
