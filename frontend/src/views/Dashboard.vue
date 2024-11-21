<script lang="ts" setup>
import { GetPackagesByUser } from '@backend/thunderstore/API'

import { Viewport } from '@components'

import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';

import FloatLabel from 'primevue/floatlabel'

import { reactive } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

interface PackageInfo {
    name: string
    resultText?: string
}

const data: PackageInfo = reactive({
    name: ""
})

async function getPackages() {
    if (data.name == "") {
        data.resultText = t('search-packages.empty-input')
        return
    }

    const pkgs = await GetPackagesByUser(["lethal-company"], data.name)
    data.resultText = pkgs
}
</script>

<template>
<Viewport class="dashboard flex-full column">
    <Tabs :value="0">
        <TabList>
            <Tab :value="0">Thunderstore</Tab>
            <Tab :value="1">Nexus</Tab>
        </TabList>

        <TabPanels>
            <TabPanel :value="0">
                <div class="search-user-packages">
                    <h1>{{ t('search-packages.header') }}</h1>
    
                    <div id="input" class="input-box no-drag">
                        <FloatLabel>
                            <InputText id="name" v-model="data.name" autocomplete="off"/>
                            <!-- <label for="name">Owner</label> -->
                        
                            <Button class="search-btn" severity="help" outlined icon="pi pi-search" @click="getPackages"/>
                        </FloatLabel>
                    </div>
    
                    <div id="result" class="result no-drag">{{ data.resultText }}</div>
                </div>
            </TabPanel>

            <TabPanel :value="1">
                <p class="m-0">buh</p>
            </TabPanel>
        </TabPanels>
    </Tabs>
</Viewport>
</template>

<!-- #region Style -->
<style scoped>
.dashboard {
    position: relative;
    text-align: center;
    user-select: none;
    top: 35px;
}

.dashboard h1 {
    margin-bottom: 35px;
}

.search-user-packages {
    align-items: center;
    justify-content: center;
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

.input-box .search-btn {
    height: 36px;
    line-height: 30px;
    margin: 0 0 0 15px;
    border: #c0c0c0 solid 1px;
    border-radius: 5px;
    cursor: pointer;
}

.input-box .search-btn:hover {
    border-radius: 2.5px;
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

:deep(.p-tabpanels) {
    background: none;
}

:deep(.p-tablist-tab-list) {
    background: none;
    border: none;
    justify-content: center;
}

:deep(.p-tab) {
    font-size: 24px;
}
</style>
<!-- #endregion -->