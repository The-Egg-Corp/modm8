<script lang="ts" setup>
import { reactive } from 'vue'
import { GetUserPackages } from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "",
})

function getUserPkgs() {
  GetUserPackages(data.name).then((res: string) => {
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
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  background-color: rgba(255, 255, 255, 1);
}
</style>
<!-- #endregion -->