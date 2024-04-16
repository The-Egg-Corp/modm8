<script lang="ts" setup>
import { reactive } from 'vue'
import { GetUserPackages } from '../../wailsjs/go/main/App'

interface PackageInfo {
  name: string
  resultText?: string
}

const data: PackageInfo = reactive({
  name: ""
})

function getUserPkgs() {
  GetUserPackages(data.name).then(res => {
    data.resultText = res
  })
}
</script>

<template>
  <main>
    <div class="container">
      <img id="logo" alt="modm8 icon" src="../assets/images/modm8-logo-transparent-white-donut.png"/>
      <div id="input" class="input-box">
        <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
        <button class="btn" @click="getUserPkgs">Get</button>
      </div>
      <div id="result" class="result">{{ data.resultText }}</div>
    </div>
  </main>
</template>

<!-- #region Style -->
<style scoped>
.container {
  user-select: none;
  position: absolute;
  left: 0;
  right: 0;
  text-align: center;
  display: block;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  size-adjust: auto;
}

.input-box .btn {
  width: 70px;
  height: 32px;
  line-height: 30px;
  margin: 2rem 0 0 20px;
  padding: 0 8px;
  cursor:auto;
  user-select: none;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #363636;
}

.input-box .input {
  border: inset;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgb(43, 43, 43);
  -webkit-font-smoothing: antialiased;
}

#logo {
  display: block;
  width: 25%;
  height: 25%;
  max-width: 50%;
  max-height: 50%;
  align-self: center;
  object-fit: contain;
  margin: auto;
  padding: 2% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
<!-- #endregion -->