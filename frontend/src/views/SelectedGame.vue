<script lang="ts" setup>
import { nextTick, onMounted, ref } from "vue"

import Skeleton from 'primevue/skeleton'
import Tag from 'primevue/tag'

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

import type { Package, Nullable } from "@types"
import { useAppStore, useGameStore } from "@stores"

import { debounce } from "../util"
import { storeToRefs } from "pinia"

const appStore = useAppStore()
const { sidebarWidth } = storeToRefs(appStore)

const gameStore = useGameStore()
const { selectedGame, updateModCache } = gameStore

const configEditorDialog = useDialog('selected-game-config-editor')
const installingModDialog = useDialog('selected-game-installing-mod')

const ROWS = 40

const loading = ref(false)
const searchInput = ref<Nullable<string>>(null)
const first = ref(0) // Starting index of the current page

const modElements = ref<any[]>([])
const scrollIndex = ref(0)

const mods = ref<thunderstore.StrippedPackage[]>([])
const currentPageMods = ref<Package[]>([])

const installing = ref(false)
const lastInstalledMod = ref<Nullable<v1.PackageVersion>>(null)

const activeTabIndex = ref(0)
const tabs = ref([
    { label: 'This Profile', icon: 'pi pi-box' },
    { label: 'Thunderstore', icon: 'pi pi-globe' },
    //{ label: 'Nexus', icon: 'pi pi-globe' } // Uncomment when Nexus is implemented.
])

const onTabChange = (e: TabMenuChangeEvent) => {
    activeTabIndex.value = e.index

    if (e.index == 1) refreshMods(false)
    else mods.value = []
}

/** 
 * Refreshes the page with the currently cached mods.
 * @param fetch Whether to fetch mods from the API and update the cache (if isnt populated already) before refreshing the page.
*/
const refreshMods = async (fetch: boolean) => {
    if (fetch && !selectedGame.modCache) {
        loading.value = true // Fetching might take a while, let the user know we are doing something.

        try {
            const t0 = performance.now()

            const pkgs = await GetStrippedPackages(selectedGame.identifier, false)
            console.log(`[${selectedGame.identifier}] Putting mods into cache...`)

            updateModCache(pkgs)
            console.log(`Cached ${pkgs.length} mods. Took: ${performance.now() - t0}ms`)
        } catch(e: any) {
            console.error(`[${selectedGame.identifier}] Failed to update mod cache!\n${e.message}`)
        }
    }

    mods.value = getMods()
    await updatePage(0, ROWS)

    // We are done regardless of outcome, stop loading.
    loading.value = false
}

onMounted(async () => {
    // Ensure no overlays are still shown.
    configEditorDialog.setVisible(false)
    installingModDialog.setVisible(false)

    await refreshMods(true)
})

const latestModVersion = (mod: thunderstore.StrippedPackage) => 
    GetLatestPackageVersion(selectedGame.identifier, mod.owner, mod.name)
    
const gameThumbnail = () => selectedGame.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${selectedGame.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

const startVanilla = () => Steam.LaunchGame(selectedGame.steamID, ["--doorstop-enable", "false"])
const startModded = () => Steam.LaunchGame(selectedGame.steamID, ["--doorstop-enable", "true"])

const onPageChange = (e: DataViewPageEvent) => updatePage(e.first, e.rows)
const updatePage = async (newFirst: number, rows: number) => {
    first.value = newFirst

    const filtered = mods.value.slice(newFirst, newFirst + rows) as Package[]

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

    return mods.filter(mod => {
        const lowerTitle = mod.name?.toLowerCase() ?? ""

        // Necessary to not show irrelevent mods with only 1 letter input.
        if (input.length == 1 && !lowerTitle.startsWith(input.toLowerCase())) {
            return false
        }

        return lowerTitle.includes(input.toLowerCase())
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

const debouncedSearch = debounce(() => refreshMods(false), 250)

const installMod = async (fullName: string) => {
    installing.value = true

    installingModDialog.setClosable(false)
    installingModDialog.setVisible(true)

    try {
        const start = performance.now()
        lastInstalledMod.value = await InstallByName(selectedGame.title, selectedGame.identifier, fullName)
        console.info(`Installed mod: ${fullName}. Took ${((performance.now() - start) / 1000).toFixed(2)}s`)
    } catch(e: any) {
        console.error(`[${selectedGame.identifier}] Failed to install mod.\n${e.message}`)
    }

    installing.value = false
    installingModDialog.setClosable(true)
}

const scrollToMod = () => {
    const i = scrollIndex.value
    if (i < 0 || i >= modElements.value.length) return // Ensure no OOB

    const game = modElements.value[i]
    if (!game) return

    game.scrollIntoView({ block: 'start' })
}

const handleScroll = (e: WheelEvent) => {
    if (e.deltaY > 0) {
        // Scrolling down
        if (scrollIndex.value < modElements.value.length - 1) {
            scrollIndex.value++
        }
    } else if (scrollIndex.value > 0) {
        scrollIndex.value--
    }

    // Scroll to the corresponding section
    nextTick(() => scrollToMod())
}
</script>

<template>
<div :class="['selected-game', { 'no-drag': configEditorDialog.visible || installingModDialog.visible }]">
    <div class="flex row">
        <div class="flex column">
            <Card class="selected-game-card no-drag">
                <template #title>
                    <p class="selected-game-card-header mt-0 mb-2">
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
                                <Button plain class="btn justify-left" 
                                    icon="pi pi-caret-right"
                                    :label="$t('selected-game.start-modded-button')"
                                    @click="startModded"
                                />
        
                                <Button plain class="btn justify-left" severity="secondary"
                                    icon="pi pi-caret-right"
                                    :label="$t('selected-game.start-vanilla-button')"
                                    @click="startVanilla"
                                />
        
                                <Button plain class="btn justify-left mt-4"
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
                        <div v-for="i in 6" :key="i" class="loading-list-item">
                            <div style="width: 1280px;" class="flex flex-row ml-1 p-3 border-faint border-round">
                                <Skeleton size="6.5rem"/> <!-- Thumbnail -->
                                
                                <div class="flex column gap-1 ml-2">
                                    <Skeleton height="1.5rem" width="20rem"/> <!-- Title -->
                                    <Skeleton width="65rem"/> <!-- Description -->

                                    <div class="flex row gap-2">
                                        <Skeleton class="mt-3" width="6.8rem" height="2.2rem"/> <!-- Install Button -->
                                        
                                        <div class="flex row gap-1 align-items-center">
                                            <Skeleton class="mt-3" width="2.8rem" height="2.2rem"/> <!-- Like button -->
                                            <Skeleton class="mt-3" width="1.8rem" height="1.6rem"/> <!-- Likes -->
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
                :value="mods" @page="onPageChange" :first="first"
            >
                <template #empty>
                    <div v-if="hasSearchInput()" class="pl-2">
                        <h2 class="m-0 mt-1">{{ $t('selected-game.empty-results') }}.</h2>

                        <!-- Sadge -->
                        <img class="mt-2" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/3x.png">
                    </div>

                    <!-- TODO: If failed, make this show regardless of search input. --> 
                    <div v-else>
                        <h2 v-if="activeTabIndex == 0" class="ml-1" style="color: orange; font-size: 24px; margin: 0;">
                            No mods installed.
                        </h2>

                        <div v-else class="ml-1">
                            <h2 class="mb-2" style="color: red; font-size: 24px; margin: 0 auto;">
                                No mods available! Something probably went wrong.
                            </h2>

                            <Button class="mt-1" :label="$t('keywords.refresh')" icon="pi pi-refresh" @click="refreshMods(true)"/>
                        </div>
                    </div>
                </template>
        
                <template #header>
                    <div class="flex row align-items-center gap-2">
                        <div class="searchbar no-drag">
                            <IconField iconPosition="left">
                                <InputIcon class="pi pi-search"></InputIcon>
                                <InputText type="text" :placeholder="$t('selected-game.search-mods')" 
                                    v-model="searchInput" @input="onInputChange"
                                />
                            </IconField>
                        </div>

                        <TabMenu :model="tabs" @tab-change="onTabChange"/>
                        <!-- <div class="flex row">
                            <ModListDropdown>
                                
                            </ModListDropdown>
                        </div> -->
                    </div>
                </template>
        
                <template #list>
                    <div class="scrollable-list list-nogutter no-drag" @wheel.prevent="handleScroll">
                        <div 
                            v-for="(mod, index) in currentPageMods" class="list-item col-12"
                            :key="index" :ref="el => modElements[index] = el"
                        >
                            <div class="flex-grow-1 flex column sm:flex-row align-items-center pt-2 gap-3" :class="{ 'border-top-faint': index != 0 }">
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
    margin-left: v-bind(sidebarWidth); /* Account for sidebar */
    margin-top: 30px;
}

.selected-game-card {
    background: none;
    box-shadow: none;
    width: max-content;
    flex-shrink: 0;
}

.selected-game-card-header {
    font-size: 34px;
    font-weight: 540;
    user-select: none;
    justify-self: start;
    text-align: start;
}

.selected-game-card :deep(.p-card-body) {
    margin: 0px 20px 0px 30px;
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
    font-size: 26px;
    font-weight: 330;
    text-shadow: 0px 0px 10px rgba(255, 255, 255, 0.45);
}

.btn {
    border-radius: 4px;
    text-align: left;
}

.justify-left {
    justify-content: left !important;
}

.mod-list-container {
    max-width: 100vw;
    width: 100vw;
}

.mod-list-thumbnail {
    min-width: 85px;
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

.scrollable-list {
    overflow-y: scroll;
    scrollbar-width: none;
    max-height: calc(100vh - 150px);
}

/*:deep(.p-dataview-layout-options .p-button) {
    background: none !important;
    border: none;
}*/

:deep(.p-tabmenu-tablist) {
    background: none !important;
}

/* TODO: Investigate why this padding affects profile manager. */
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
    justify-content: start;
    padding: 0;
}

:deep(.p-dataview-paginator-bottom) {
    border: none;
}

.list-item {
    display: flex;
    width: 1300px; /* TODO: Make this calc from right edge minus 30px. 100vw and 100% dont work? */
    padding-bottom: 15px;
    padding-top: 0px;
}

.loading-list-item {
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