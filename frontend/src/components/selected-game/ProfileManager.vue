<script lang="ts" setup>
import DataView from 'primevue/dataview'
//import Toolbar from 'primevue/toolbar'

import { Ref, ref } from 'vue'
import { t } from '@i18n'

interface Profile {
    name: string
    mods: {
        thunderstore?: string[]
        nexus?: string[]
    }
}

const selectedProfile: Ref<Profile> = ref({
    name: "Default", 
    mods: {}
})

const profiles: Ref<Profile[]> = ref([
    { name: 'Default', mods: {} },
    { name: 'testing', mods: {} },
    { name: 'MEGALOPHOBIA', mods: {} },
    { name: 'big futa cock mods', mods: {} },
])
</script>

<template>
<div class="profile-manager">
    <p class="header mt-4 mb-3">{{ t('selected-game.profile-manager.header') }}</p>
    
    <DataView data-key="profile-manager-list" layout="list">
        <template #header>
            <div class="flex row justify-content-betwen">
                <Button icon="pi pi-plus" class="mr-2" severity="primary"/>
                <InputText :placeholder="t('selected-game.profile-manager.search-placeholder')"/>
            </div>
        </template>
    </DataView>

    <!-- <Card class="selector-card">
        <template #title>
            <Toolbar>
                <template #start>
                    <Button :label="$t('keywords.new')" icon="pi pi-plus" class="mr-2" severity="primary"/>
                </template>
            
                <template #end>
                    <div class="flex row gap-1">
                        <Button 
                            icon="pi pi-download" severity="secondary"
                            v-tooltip.top="tooltipOpts($t('keywords.import'))"
                        />

                        <Button 
                            icon="pi pi-upload" severity="secondary"
                            v-tooltip.top="tooltipOpts($t('keywords.export'))"
                        />
                    </div>
                    
                </template>
            </Toolbar>
        </template>

        <template #content>
            <List filter class="selector-list w-full" v-model="selectedProfile" 
                :options="profiles" optionLabel="name"
            />
        </template>
    </Card> -->
</div>
</template>

<style scoped>
.profile-manager {
    margin: 0px 15px 0px 30px;
}

.header {
    font-size: 28px;
    font-weight: 540;
    user-select: none;
}

.selector-card {
    height: 100% !important;
    background: none;
}

:deep(.p-toolbar) {
    background: none;
}

.selector-card :deep(.p-card-body) {
    padding: 0rem;
    gap: 10px;
}

.selector-list {
    background: none; 
}

:deep(.selector-list > *) {
    background: none; 
}
</style>