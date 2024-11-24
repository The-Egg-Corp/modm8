<script lang="ts" setup>
import Listbox, { ListboxChangeEvent } from 'primevue/listbox' 

import InputGroup from 'primevue/inputgroup'
import Popover from 'primevue/popover'

import { NewProfile, DeleteProfile } from '@backend/profile/ProfileManager'
import { tooltipOpts } from '@frontend/src/util'

import { t } from '@i18n'
import { Profile } from '@types'
import { useGameStore, useProfileStore } from '@stores'

import { storeToRefs } from 'pinia'
import { onMounted, ref } from 'vue'

const gameStore = useGameStore()
const { selectedGame } = storeToRefs(gameStore)

const profileStore = useProfileStore()

const { setSelectedProfile, initProfiles } = profileStore
const { profiles, selectedProfile } = storeToRefs(profileStore)

const pp = ref()
const newProfNameInput = ref<string>('')

const getModsAmount = (prof: Profile) => {
    const tsAmt = prof.mods.thunderstore?.length ?? 0
    const nexusAmt = prof.mods.nexus?.length ?? 0

    return tsAmt + nexusAmt
}

const onChange = (e: ListboxChangeEvent) => {
    setSelectedProfile(e.value)
    emit('profileSelected', e)
}

const togglePopover = (e: MouseEvent) => {
    pp?.value.show(e)
}

const createNewProf = async (e: MouseEvent) => {
    await NewProfile(selectedGame.value.title, newProfNameInput.value!)
    await initProfiles()

    pp.value.hide(e)

    emit('profileCreated', e)
}

const renameProf = (_: MouseEvent, name: string) => {
    
}

const deleteProf = async (e: MouseEvent, name: string) => {
    e.stopPropagation()

    await DeleteProfile(selectedGame.value.title, name)
    await initProfiles()
}

const shouldDisableCreation = () => {
    if (newProfNameInput.value.length < 1) return true // No input yet

    // Profile already exists 
    return profiles.value.some(p => p.name == newProfNameInput.value)
}

const emit = defineEmits(['profileSelected', 'profileCreated'])

onMounted(async () => {
    await initProfiles()
})
</script>

<template>
<div class="profile-manager">
    <p class="header">{{ t('selected-game.profile-manager.header') }}</p>

    <div class="flex row gap-2">
        <InputText class="w-full" :placeholder="t('selected-game.profile-manager.search-placeholder')"/>

        <Dropdown :options="['hi']">

        </Dropdown>
    </div>

    <Listbox class="profile-list" 
        :modelValue="selectedProfile" :options="profiles" optionLabel="name" 
        @change="onChange"
    >
        <template #header>
            <div class="flex justify-content-between">
                <div class="flex row gap-1">
                    <Button
                        icon="pi pi-plus" severity="primary"
                        v-tooltip.top="tooltipOpts(t('selected-game.profile-manager.new-profile'))"
                        @click="togglePopover"
                    />

                    <Popover ref="pp">
                        <div class="flex flex-col gap-4 w-[25rem]">
                            <div>
                                <span class="font-medium block mb-2">Profile Name</span>
                                <InputGroup>
                                    <InputText v-model="newProfNameInput" v-keyfilter="/^[^<>*!\/]+$/"/>
                                    <Button 
                                        label="Create" severity="success" icon="pi pi-check"
                                        :disabled="shouldDisableCreation()" @click="createNewProf"
                                    />
                                </InputGroup>
                            </div>
                        </div>
                    </Popover>
                </div>

                <div class="flex row gap-1">
                    <!-- <Button 
                        icon="pi pi-download" severity="secondary"
                        v-tooltip.top="tooltipOpts(t('keywords.import'))"
                    />

                    <Button 
                        icon="pi pi-upload" severity="secondary"
                        v-tooltip.top="tooltipOpts(t('keywords.export'))"
                    /> -->

                    <Button 
                        icon="pi pi-refresh" severity="secondary"
                        v-tooltip.top="tooltipOpts(t('keywords.refresh'))"
                        @click="initProfiles()"
                    />
                </div>
            </div>
        </template>

        <template #empty>
            <div class="flex column justify-content-center align-items-center mt-1">
                <div style="font-weight: 500; color: var(--p-primary-color);">No profiles exist yet.</div>
                <div style="font-weight: 350; color: lightgrey;">Click the plus icon to get started!</div>

                <img class="mt-2" src="https://cdn.7tv.app/emote/01G7YR9X5G0003Z50SB3FM5WR4/3x.webp"></img>
            </div>
        </template>

        <template #option="{ option: profile }">
            <div class="flex row w-full justify-content-between align-items-center">
                <div>[{{getModsAmount(profile)}}] {{ profile.name }}</div>
                
                <div class="flex gap-1">
                    <Button
                        style="width: 34px; height: 32px;" icon="pi pi-pencil"
                        @click="e => renameProf(e, profile.name)"
                    />

                    <Button 
                        style="width: 34px; height: 32px;" icon="pi pi-trash" severity="danger"
                        @click="e => deleteProf(e, profile.name)"
                    />
                </div>
            </div>
        </template>
    </Listbox>
</div>
</template>

<style scoped>
.profile-manager {
    margin: 30px 0px 20px 0px;
    display: flex;
    flex-direction: column;
    flex: 1 0 auto;
}

.profile-list {
    display: flex;
    flex-direction: column;
    flex: 1 0 auto;
    margin-top: 10px;
    border: var(--border-faint);
}

:deep(.p-listbox-list) {
    padding: 0;
}

:deep(.p-listbox-list-container) {
    overflow: scroll;
    scrollbar-width: none;
}

:deep(.p-listbox-option) {
    padding: 10px 5px 10px 5px;
}

:deep(.p-listbox-header) {
    padding: 5px 0px 5px 0px;
}

.header {
    font-size: 32px;
    font-weight: 540;
    user-select: none;
    margin: 0px 0px 15px 0px;
}
</style>