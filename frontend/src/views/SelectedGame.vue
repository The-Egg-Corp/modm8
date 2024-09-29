<script lang="ts" setup>
import { onMounted, ref } from "vue"
import type { Ref } from "vue"

import Card from 'primevue/card'
import Button from 'primevue/button'
import Skeleton from 'primevue/skeleton'
import Tag from 'primevue/tag'

import { Nullable } from "primevue/ts-helpers"

import DataView, { DataViewPageEvent } from 'primevue/dataview'
import TabMenu, { TabMenuChangeEvent } from 'primevue/tabmenu'

import * as Steam from '@backend/steam/SteamRunner'
import { GetLatestPackageVersion, GetStrippedPackages } from '@backend/thunderstore/API'
import { thunderstore, v1 } from "@backend/models"
import { InstallByName } from '@backend/thunderstore/API.js'

import { useDialog } from '@composables'
import {
    ProfileManager,
    ModInstallationOverlay, 
    ConfigEditorOverlay
} from "@components"

import { Package } from "@types"
import { useGameStore } from "@stores"

import { debounce } from "../util"

const gameStore = useGameStore()
const { selectedGame, updateModCache } = gameStore

const configEditorDialog = useDialog('selected-game-config-editor')
const installingModDialog = useDialog('selected-game-installing-mod')

const ROWS = 25

const loading = ref(false)
const searchInput: Ref<Nullable<string>> = ref(null)

const mods: Ref<thunderstore.StrippedPackage[]> = ref([])
const currentPageMods: Ref<Package[]> = ref([])

const installing = ref(false)
const lastInstalledMod: Ref<Nullable<v1.PackageVersion>> = ref(null)

const activeTabIndex = ref(0)
const tabs = ref([
    { label: 'This Profile', icon: 'pi pi-box' },
    { label: 'Thunderstore', icon: 'pi pi-globe' },
    { label: 'Nexus', icon: 'pi pi-globe' }
])

const onTabChange = (e: TabMenuChangeEvent) => {
    activeTabIndex.value = e.index
}

onMounted(async () => {
    // Ensure no overlays are still shown when opening.
    configEditorDialog.setVisible(false)
    installingModDialog.setVisible(false)

    const t0 = performance.now()
    loading.value = true

    if (!selectedGame.modCache) {
        console.log(`Putting ${selectedGame.identifier} mods into cache...`)

        const pkgs = await GetStrippedPackages(selectedGame.identifier, false)
        updateModCache(pkgs)

        console.log(`Cached ${pkgs?.length} mods. Took: ${performance.now() - t0}ms`)
    } else {
        console.log(`Got ${selectedGame.modCache.length} mods from cache. Took: ${performance.now() - t0}ms`)
    }

    mods.value = getMods()
    await updatePage(0, ROWS)

    loading.value = false
})

const latestModVersion = (mod: thunderstore.StrippedPackage) => 
    GetLatestPackageVersion(selectedGame.identifier, mod.owner, mod.name)

const gameThumbnail = () => selectedGame.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${selectedGame.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

const startVanilla = () => Steam.LaunchGame(selectedGame.steamID, ["--doorstop-enable", "false"])
const startModded = () => Steam.LaunchGame(selectedGame.steamID, ["--doorstop-enable", "true"])

const onPageChange = (e: DataViewPageEvent) => updatePage(e.first, e.rows)
const updatePage = async (first: number, rows: number) => {
    const filtered = mods.value.slice(first, first + rows) as Package[]

    // TODO: This could be potentially VERY slow. Consider replacing with Map/Set instead.
    const promises = filtered.map(async mod => {
        mod.latestVersion = await latestModVersion(mod)
        return mod
    })

    currentPageMods.value = await Promise.all(promises)
}

const filterBySearch = (mods: thunderstore.StrippedPackage[]) => {
    if (!searchInput.value) return mods

    const input = searchInput.value.trim()
    if (input == "") return mods

    const lowerInput = input.toLowerCase()

    return mods.filter(mod => {
        const lowerTitle = mod.name?.toLowerCase() ?? ""

        // Necessary to not show irrelevent mods with only 1 letter input.
        if (input.length == 1 && !lowerTitle.startsWith(lowerInput)) {
            return false
        }

        return lowerTitle.includes(lowerInput)
    })
}

const getMods = (searchFilter = true, defaultSort = true) => {
    if (!selectedGame.modCache) return []
    const mods = searchFilter ? filterBySearch(selectedGame.modCache) : selectedGame.modCache

    if (defaultSort) {
        mods.sort((m1, m2) => m2.rating_score - m1.rating_score)
    }

    return mods
}

const hasSearchInput = () => searchInput.value ? searchInput.value.length > 0 : undefined
const onInputChange = async () => {
    // No input, no need to debounce.
    if (!searchInput.value?.trim()) {
        // Show all mods without filtering.
        mods.value = getMods(false)
        return await updatePage(0, ROWS)
    }

    debouncedSearch()
}

const debouncedSearch = debounce(async () => {
    mods.value = getMods()
    await updatePage(0, ROWS)
}, 250)

const installMod = async (fullName: string) => {
    installing.value = true

    installingModDialog.setClosable(false)
    installingModDialog.setVisible(true)

    try {
        const start = performance.now()
        lastInstalledMod.value = await InstallByName(selectedGame.title, selectedGame.identifier, fullName)
        console.info(`Installed mod: ${fullName}. Took ${((performance.now() - start) / 1000).toFixed(2)}s`)
    } catch(e: any) {
        console.error(`Failed to install mod.\n${e.message}`)
    }

    installing.value = false
    installingModDialog.setClosable(true)
}
</script>

<template>
<div :class="['selected-game', { 'no-drag': configEditorDialog.visible || installingModDialog.visible }]">
    <div class="flex row">
        <div class="flex column">
            <Card class="selected-game-card no-drag">
                <template #title>
                    <p class="selected-game-card-header mt-0 mb-1">
                        {{ $t('selected-game.currently-selected') }}
                    </p>
                </template>
        
                <template #content>
                    <div class="flex no-drag">
                        <div class="game-thumbnail-container">
                            <img class="selected-game-thumbnail-background" :src="gameThumbnail()"/>
                            <img class="selected-game-thumbnail-foreground" :src="gameThumbnail()"/>
                        </div>
    
                        <div class="flex column ml-3 no-drag">
                            <p class="selected-game-title mt-0 mb-0">{{ selectedGame.title }}</p>
                            <div class="flex column gap-2 mt-3">
                                <Button plain class="btn" 
                                    icon="pi pi-caret-right"
                                    :label="$t('selected-game.start-modded-button')"
                                    @click="startModded"
                                />
        
                                <Button plain class="btn" severity="secondary"
                                    icon="pi pi-caret-right"
                                    :label="$t('selected-game.start-vanilla-button')"
                                    @click="startVanilla"
                                />
        
                                <Button plain class="btn mt-4"
                                    icon="pi pi-file-edit"
                                    :label="$t('selected-game.config-button')"
                                    @click="configEditorDialog.setVisible(true)"
                                />
                            </div>
                        </div>
                    </div>
                </template>
            </Card>

            <ProfileManager/>
        </div>
    
        <div class="flex mod-list-container">
            <!-- Show skeleton of mod list while loading -->
            <DataView v-if="loading" data-key="mod-list-loading" layout="list">
                <template #empty>
                    <div class="list-nogutter pt-4">
                        <div v-for="i in 6" :key="i" class="list-item">
                            <div style="width: 1200px;" class="flex flex-row p-3 border-1 surface-border border-round">
                                <Skeleton size="6.4rem"/>
                                <div class="flex column gap-1 ml-2">
                                    <Skeleton height="1.5rem" width="20rem"/>
                                    <Skeleton width="65rem"/>

                                    <div class="flex row gap-2">
                                        <Skeleton class="mt-3" width="6.8rem" height="2.2rem"/>
                                        <div class="flex row gap-1">
                                            <Skeleton class="mt-3" width="2.8rem" height="2.2rem"/>
                                            <Skeleton class="mt-4" width="1.8rem" height="1.3rem"/>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </template>
            </DataView>

            <DataView
                v-else lazy stripedRows
                layout="list" data-key="mod-list"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink"
                :paginator="mods.length > 0" :rows="ROWS"
                :value="mods" @page="onPageChange"
            >
                <template #empty>
                    <div v-if="hasSearchInput()" class="pl-2">
                        <h2>{{ $t('selected-game.empty-results') }}.</h2>

                        <!-- Sadge -->
                        <img class="pt-3" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/3x.png">
                    </div>

                    <!-- TODO: If failed, make this show regardless of search input. --> 
                    <div v-else>
                        <h2 style="padding-top: 15px; font-size: 22px; margin: 0 auto;">
                            No mods available! Something probably went wrong.
                        </h2>
                    </div>
                </template>
        
                <template #header>
                    <div class="flex row align-items-center">
                        <div class="searchbar no-drag">
                            <IconField iconPosition="left">
                                <InputIcon class="pi pi-search"></InputIcon>
                                <InputText type="text" :placeholder="$t('selected-game.search-mods')" 
                                    v-model="searchInput" @input="onInputChange"
                                />
                            </IconField>
                        </div>

                        <div class="ml-2 justify-self-end">
                            <TabMenu :model="tabs" @tab-change="onTabChange"/>
                        </div>

                        <!-- <div class="flex row">
                            <ModListDropdown>
                                
                            </ModListDropdown>
                        </div> -->
                    </div>
                </template>
        
                <template #list>
                    <div class="scrollable list-nogutter no-drag">
                        <div v-for="(mod, index) in currentPageMods" :key="index" class="list-item col-12">
                            <div class="flex-grow-1 flex column sm:flex-row align-items-center pt-2 gap-3" :class="{ 'border-top-1 surface-border': index != 0 }">
                                <img class="mod-list-thumbnail block xl:block" :src="mod.latestVersion?.icon || ''"/>
                                
                                <div class="flex-grow-1 flex column md:flex-row md:align-items-center">
                                    <div class="flex-grow-1 flex column justify-content-between">
                                        <div class="flex row align-items-baseline">
                                            <div class="mod-list-title">{{ mod.name }}</div>
                                            <div class="mod-list-author">({{ mod.owner }})</div>
                                        </div>

                                        <div class="mod-list-description mb-1">{{ mod.latestVersion.description }}</div>

                                        <!--
                                            :icon="isFavouriteGame(game.identifier) ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                            :style="isFavouriteGame(game.identifier) ? { color: 'var(--primary-color)' } : {}"
                                            @click="toggleFavouriteGame(game.identifier)"
                                        /> -->

                                        <div class="mod-list-bottom-row"> 
                                            <div class="flex row gap-2">
                                                <Button plain class="btn w-full" icon="pi pi-download"
                                                    :label="$t('keywords.install')" @click="installMod(mod.full_name)"
                                                />

                                                <div class="flex row align-items-center">
                                                    <Button outlined plain 
                                                        style="margin-right: 6.5px;"
                                                        :icon="'pi pi-thumbs-up'"
                                                    />
                                                    
                                                    <div class="mod-list-rating">{{ mod.rating_score }}</div>
                                                </div>
                                            </div>

                                            <!-- TODO: Ensure the tags flex to the end of the DataView and not the item content. -->
                                            <div class="flex row flex-shrink-0 gap-1">
                                                <div v-for="category in mod.categories.filter(c => c.toLowerCase() != 'mods')">
                                                    <Tag :value="category"></Tag>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </template>
            </DataView>

        </div>

        <ModInstallationOverlay :dialog="installingModDialog" 
            :installing="installing" :lastInstalledMod="lastInstalledMod!"
        />

        <ConfigEditorOverlay :dialog="configEditorDialog" :selectedGame="selectedGame"/>
    </div>
</div>
</template>

<style scoped>
.material-symbols-sharp {
    font-size: 26.5px;
    font-variation-settings:
    'FILL' 0,
    'wght' 270,
    'GRAD' 100,
    'opsz' 40
}

.selected-game {
    display: flex;
    flex-direction: column;
    margin-left: 80px; /* Account for sidebar */
    margin-top: 30px;
}

.selected-game-card {
    background: none;
    width: max-content;
    flex-shrink: 0;
}

.selected-game-card-header {
    font-size: 30px;
    font-weight: 540;
    user-select: none;
}

.selected-game-card :deep(.p-card-body) {
    margin: 0px 15px 0px 30px;
    padding: 0px;
}

.game-thumbnail-container {
    position: relative;
}

.selected-game-thumbnail-foreground {
    position: relative;
    min-width: 160px;
    max-width: 40%;
    max-height: 200px;
    border-radius: 3px;
    user-select: none;
    border: 1px ridge rgba(255, 255, 255, 0.45);
}

.selected-game-thumbnail-background {
    position: absolute;
    z-index: -1;
    filter: blur(6px);
    width: 100%;
    height: 100%;
}

.selected-game-title {
    font-size: 25px;
    font-weight: 330;
    text-shadow: 0px 0px 10px rgba(255, 255, 255, 0.45);
}

.btn {
    border-radius: 4px;
    text-align: left;
}

.mod-list-container {
    max-width: 100vw;
    width: 100vw;
}

.mod-list-thumbnail {
    min-width: 80px;
    width: 4%;
    max-width: 140px;
    border-radius: 2.5px;
    user-select: none;
}

.mod-list-title {
    font-size: 20px;
    font-weight: 460;
    padding-right: 5px;
}

.mod-list-author {
    font-size: 16px;
    font-weight: 310;
}

.mod-list-description {
    font-size: 16px;
    font-weight: 220;
    padding-bottom: 5px;
    text-wrap: pretty;
}

.mod-list-rating {
    font-size: 17px;
    font-weight: 370;
}

.mod-list-bottom-row {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    flex-grow: 1;
}

.scrollable {
    overflow-y: scroll;
    scrollbar-width: none;
    height: calc(100vh - 150px);
    margin-bottom: 0.25rem;
}

:deep(.p-dataview-layout-options .p-button) {
    background: none !important;
    border: none;
}

:deep(.p-dataview-header) {
    background: none !important;
    padding: 10px 0px 10px 0px;
    margin: 0px 5px 0px 5px;
    border: none;
}

:deep(.p-dataview-content) {
    background: none !important;
}

:deep(.searchbar .p-inputtext) {
    background: rgba(0, 0, 0, 0.2);
    margin-left: auto;
    margin-right: auto;
    width: 350px;
    min-width: 200px;
}

:deep(.p-paginator) {
    background: none !important;
    padding: 0 0 0 0;
    justify-content: start;
}

.list-item {
    display: flex;
    width: 100vw;
    padding-bottom: 15px;
    padding-top: 0px;
}

:deep(.p-tabmenu-nav) {
    background: none;
}

:deep(.p-menuitem-link) {
    background: none;
}
</style>