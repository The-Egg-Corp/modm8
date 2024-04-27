<script lang="ts" setup>
import { reactive } from 'vue'
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

import { GetUserPackages } from '../../wailsjs/go/backend/App'

interface PackageInfo {
  name: string
  resultText?: string
}

const data: PackageInfo = reactive({
  name: ""
})

function PackagesByUser() {
  if (data.name == "") {
    data.resultText = t('search-packages.empty-input')
    return
  }

  GetUserPackages(["lethal-company"], data.name).then(res => {
    data.resultText = res
  })
}
</script>

<template>
  <div class="container">
    <!--  <img id="logo" alt="modm8 icon" src="../assets/images/modm8-logo-transparent-white-donut.png"/> -->
    
    <h1>{{ t('search-packages.header') }}</h1>

    <div id="input" class="input-box no-drag">
      <FloatLabel>
        <InputText id="name" v-model="data.name" autocomplete="off"/>
        <!-- <label for="name">Owner</label> -->
        
        <Button class="btn" severity="help" outlined icon="pi pi-search" @click="PackagesByUser"/>
      </FloatLabel>
    </div>
    <div id="result" class="result no-drag">{{ data.resultText }}</div>
  </div>
</template>

<!-- #region Style -->
<style scoped>
.container {
  user-select: none;
  position: absolute;
  top: 30%;
  bottom: 0;
  left: 0;
  right: 0;
  text-align: center;
  display: flex;
  flex-direction: column;
}

.container h1 {
  margin-bottom: 35px;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  size-adjust: auto;
  user-select: all;
}

.p-inputtext {
  width: 300px;
}

.input-box .btn {
  height: 36px;
  line-height: 30px;
  margin: 0 0 0 15px;
  border: #c0c0c0 solid 1px;
  border-radius: 5px;
  cursor: pointer;
}

.input-box .btn:hover {
  border-radius: 3px;
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
  display: inline-block;
  max-width: 30%;
  min-width: 420px;
  align-self: center;
  object-fit: contain;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
<!-- #endregion -->