<script lang="ts" setup>
import Listbox, { ListboxChangeEvent } from 'primevue/listbox' 
import Popover from 'primevue/popover'

import { NewProfile, DeleteProfile } from '@backend/profile/ProfileManager'
import { tooltipOpts } from '@frontend/src/util'

import { t } from '@i18n'
import { Nullable, Profile } from '@types'
import { useProfileStore, useGameStoreTS } from '@stores'

import { storeToRefs } from 'pinia'
import { onMounted, ref } from 'vue'

const profileStore = useProfileStore()
const gameStoreTS = useGameStoreTS()

const { setSelectedProfile, initProfiles } = profileStore
const { profiles, selectedProfile } = storeToRefs(profileStore)

const searchInput = ref<Nullable<string>>(null)

const pp = ref()
const newProfNameInput = ref<string>('')

const MAX_PROFILE_NAME_LEN = 32

const getModsAmount = (prof: Profile) => {
    const tsAmt = prof.mods.thunderstore?.length ?? 0
    const nexusAmt = prof.mods.nexus?.length ?? 0

    return tsAmt + nexusAmt
}

const onChange = (e: ListboxChangeEvent) => {
    setSelectedProfile(e.value)
    emit('profileSelected', e)
}

const onClickPlus = (e: MouseEvent) => {
    if (pp.value) newProfNameInput.value = ''

    // Show the popover with creation options.
    pp?.value.show(e)
}

const createNewProf = async (e: MouseEvent) => {
    await NewProfile(gameStoreTS.selectedGame.title, newProfNameInput.value)
    await initProfiles()

    pp.value.hide(e)

    emit('profileCreated', e)
}

const renameProf = (_: MouseEvent, name: string) => {
    
}

const deleteProf = async (e: MouseEvent, name: string) => {
    e.stopPropagation()

    await DeleteProfile(gameStoreTS.selectedGame.title, name)
    await initProfiles()
}

const shouldDisableCreation = () => {
    if (newProfNameInput.value.length < 1) return true // No input yet
    if (newProfNameInput.value.length > MAX_PROFILE_NAME_LEN) return true // Too long

    // Profile already exists 
    return profiles.value.some(p => p.name == newProfNameInput.value)
}

const onSearchInputChange = () => {
    if (newProfNameInput.value.length > 24) {
        
    }
}

const onNameInputChange = () => {
    // Display toast
}

const emit = defineEmits(['profileSelected', 'profileCreated'])

onMounted(async () => {
    await initProfiles()
})
</script>

<template>
<div class="profile-manager">
    <p class="header">{{ t('selected-game.profile-manager.header') }}</p>

    <Listbox class="profile-list" 
        :options="profiles" optionLabel="name" 
        :modelValue="selectedProfile" @change="onChange"
    >
        <template #header>
            <div class="profile-manager-toolbars">
                <div class="flex row top-bar">
                    <div class="flex row">
                        <Button
                            icon="pi pi-plus" severity="primary"
                            v-tooltip.top="tooltipOpts(t('selected-game.profile-manager.new-profile'))"
                            @click="onClickPlus"
                        />
    
                        <Popover ref="pp">
                            <div class="profile-creator">
                                <h3 class="block mb-1 mt-0">{{ $t('selected-game.profile-manager.profile-creator.header') }}</h3>
    
                                <div>
                                    <span class="font-medium block mb-2">{{ $t('selected-game.profile-manager.profile-name') }}</span>
                                    
                                    <FloatLabel variant="on">
                                        <IconField>
                                            <InputIcon class="pi pi-pencil"/>
                                            <InputText class="profile-creator-name-input"
                                                v-model="newProfNameInput" v-keyfilter="/^[^<>*!\/]+$/"
                                                :invalid="newProfNameInput.length > MAX_PROFILE_NAME_LEN" 
                                                @input="onNameInputChange"
                                            />
                                        </IconField>
    
                                        <label></label>
                                    </FloatLabel>
                                </div>
    
                                <!-- <div class="flex column">
                                    <span class="font-medium block mb-2">{{ $t('keywords.favourite') }}</span>
                                    <Checkbox/>
                                </div> -->
    
                                <Button class="mt-2"
                                    :label="$t('keywords.create')" icon="pi pi-check"
                                    :disabled="shouldDisableCreation()" @click="createNewProf"
                                />
                            </div>
                        </Popover>
                    </div>
    
                    <div class="flex row gap-1 align-items-center">
                        <Button
                            icon="pi pi-download" severity="secondary"
                            v-tooltip.top="tooltipOpts(t('keywords.import'))"
                        />
    
                        <Button
                            icon="pi pi-upload" severity="secondary"
                            v-tooltip.top="tooltipOpts(t('keywords.export'))"
                        />
    
                        <Button
                            icon="pi pi-clone" severity="secondary"
                            v-tooltip.top="tooltipOpts(t('keywords.merge'))"
                        />
                    </div>
    
                    <div class="flex row">
                        <Button text severity="secondary" icon="pi pi-refresh"
                            v-tooltip.top="tooltipOpts(t('keywords.refresh'))"
                            @click="initProfiles()"
                        />
                    </div>
                </div>

                <div class="flex row gap-2 bottom-bar">
                    <!-- <IconField iconPosition="left">
                        <InputIcon class="pi pi-search"></InputIcon>
                        <InputText type="text" :placeholder="$t('selected-game.search-mods')" 
                            v-model="searchInput" @input="onSearchInputChange"
                        />
                    </IconField> -->
            
                    <div class="profile-searchbar w-full">
                        <FloatLabel variant="on">
                            <IconField>
                                <InputIcon class="pi pi-search"/>
                                <InputText id="search_profiles" v-model="searchInput"/>
                            </IconField>
                
                            <label for="search_profiles">{{ t('selected-game.profile-manager.search-placeholder') }}</label>
                        </FloatLabel>
                    </div>
            
                    <Dropdown class="w-5" :options="['All', 'Favourited', 'Most mods', 'Least mods']">
                        
                    </Dropdown>
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
                <div class="flex row gap-1">
                    <div>[{{getModsAmount(profile)}}]</div>
                    <div>{{ profile.name }}</div>    
                </div>

                <div class="flex row">
                    <Button text
                        style="width: 30px; height: 32px;" icon="pi pi-pencil"
                        @click="e => renameProf(e, profile.name)"
                    />

                    <Button text severity="danger"
                        style="width: 30px; height: 32px;" icon="pi pi-trash"
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
    display: flex;
    flex: 1 0 auto;
    flex-direction: column;
    align-items: center;
    margin: 20px 0px 20px 0px;
}

.header {
    font-size: 36px;
    font-weight: 540;
    user-select: none;
    margin: 0;
}

.profile-manager-toolbars .top-bar {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 10px;
}

.profile-manager-toolbars .bottom-bar {
    margin-bottom: 5px;
}

.profile-list {
    display: flex;
    flex-direction: column;
    flex: 1 0 auto;
    margin-top: 10px;
    border: none;
    background: transparent !important;
    width: 100%;
}

.profile-creator {
    display: flex;
    flex-direction: column;
    /* gap: 10px; Is this even needed? */
}

.profile-creator-name-input {
    width: 350px;
}

:deep(.profile-searchbar .p-inputtext) {
    background: rgba(0, 0, 0, 0.2);
    margin-left: auto;
    margin-right: auto;
    width: 100%;
    min-width: 200px;
}

:deep(.p-listbox-list) {
    padding: 0;
}

:deep(.p-listbox-list-container) {
    overflow-y: auto;
    overflow-x: hidden;
    padding-right: 5px; /* Make this only occur when scrollbar is visible */
    max-height: 19.3rem !important;
}

:deep(.p-listbox-option) {
    padding: 5px 5px 5px 5px;
}

:deep(.p-listbox-header) {
    padding: 0px 0px 5px 0px;
}
</style>