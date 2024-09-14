<script lang="ts" setup>
import { onMounted, ref } from "vue"
import type { Ref } from "vue"

import { Nullable } from "primevue/ts-helpers"
import DataView, { DataViewPageEvent } from 'primevue/dataview'

import TabMenu from 'primevue/tabmenu'
// import Splitter from 'primevue/splitter'
// import SplitterPanel from 'primevue/splitterpanel'

import Card from 'primevue/card'
import Button from 'primevue/button'
import Skeleton from 'primevue/skeleton'
import Tag from 'primevue/tag'

import * as Steam from '@backend/steam/SteamRunner'
import { BepinexConfigFiles, ParseBepinexConfig } from '@backend/game/GameManager'
import { GetLatestPackageVersion, GetPackagesStripped } from '@backend/thunderstore/API'
import { thunderstore, game } from "@backend/models"
import { InstallWithDependencies } from '@backend/thunderstore/API.js'

import { useDialog } from '@composables'
import { CardOverlay, ConfigEditLayout } from "@components"
import { Package } from "@types"
import { useGameStore } from "@stores"

import { debounce, openLink } from "../util"

const gameStore = useGameStore()
const { selectedGame, updateModCache } = gameStore

const configEditorDialog = useDialog('selected-game-config-editor')
const installingModDialog = useDialog('selected-game-installing-mod')

const loading = ref(false)
const searchInput: Ref<Nullable<string>> = ref(null)

const selectedConfig: Ref<Nullable<game.BepinexConfig>> = ref(null)
const selectedConfigName: Ref<Nullable<string>> = ref(null)

const mods: Ref<thunderstore.StrippedPackage[]> = ref([])
const currentPageMods: Ref<Package[]> = ref([])

const installing = ref(false)
const lastInstalledMod: Ref<Nullable<string>> = ref(null)

const tabMenuItems = ref([
    { label: 'Current Profile', icon: 'pi pi-box' },
    { label: 'Thunderstore', icon: 'pi pi-globe' }
])

const ROWS = 25
let configFiles: string[] = []

onMounted(async () => {
    loading.value = true

    configFiles = await getConfigFiles()
    
    const t0 = performance.now()
    if (!selectedGame.modCache) {
        console.log(`Putting ${selectedGame.identifier} mods into cache...`)

        const pkgs = await GetPackagesStripped(selectedGame.identifier, false)
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

const getConfigFiles = async () => {
    if (!selectedGame.bepinexSetup || !selectedGame.path) return []
    return await BepinexConfigFiles([selectedGame.path, "BepInEx", "config"])
}

const getRelativeConfigPath = (absPath: string) => {
    const configPath = 'BepInEx/config'

    const normalizedFile = absPath.replace(/\\/g, '/')
    const startIndex = normalizedFile.indexOf(configPath) + configPath.length + 1

    return normalizedFile.substring(startIndex)
}

const closeConfigEditor = () => {
    configEditorDialog.setVisible(false)
    resetSelectedConfig()
}

const resetSelectedConfig = () => {
    selectedConfig.value = null
    selectedConfigName.value = null
}

const editConfig = async (path: string) => {
    // Read file contents from backend
    try {
        const config = await ParseBepinexConfig(path)
        selectedConfig.value = config
        selectedConfigName.value = getRelativeConfigPath(path)
    } catch(err: any) {
        // Show error toast
    }
}

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
const onInputChanged = async () => {
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
    lastInstalledMod.value = fullName

    installingModDialog.setClosable(false)
    installingModDialog.setVisible(true)

    const start = performance.now()
    await InstallWithDependencies(selectedGame.title, selectedGame.identifier, fullName)
    console.log(`Installed mod: ${fullName}. Took ${((performance.now() - start) / 1000).toFixed(2)}s`)

    installing.value = false
    installingModDialog.setClosable(true)
}
</script>

<template>
<div :class="['selected-game', { 'no-drag': configEditorDialog.visible || installingModDialog.visible }]">
    <!-- <Breadcrumb class="breadcrumb flex-full row" :home="homePage" :model="pages">
        <template #item="{ item, props }">
            <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                <a :href="href" v-bind="props.action" @click="navigate">
                    <span v-if="item.icon" :class="item.class" :style="item.style">{{ item.icon }}</span>
                    <span v-else-if="item.label" :class="item.class" :style="item.style">{{ item.label }}</span>
                </a>
            </router-link>
        </template>
        <template #separator>/</template>
    </Breadcrumb> -->

    <div class="flex row">
        <Card class="selected-game-card no-drag">
            <template #title>
                <p style="font-size: 30px; font-weight: 520; user-select: none;" class="mt-0 mb-1">{{ $t('selected-game.currently-selected') }}</p>
            </template>
    
            <template #content>
                <div class="flex no-drag">
                    <img class="selected-game-thumbnail" :src="gameThumbnail()"/>
                    
                    <div class="flex column ml-3 no-drag">
                        <p style="font-size: 25px; font-weight: 330" class="mt-0 mb-0">{{ selectedGame.title }}</p>
                        <div class="flex column gap-2 mt-3">
                            <Button 
                                plain
                                class="btn" 
                                icon="pi pi-caret-right"
                                :label="$t('selected-game.start-modded-button')"
                                @click="startModded"
                            />
    
                            <Button 
                                plain outlined
                                class="btn" 
                                icon="pi pi-caret-right"
                                :label="$t('selected-game.start-vanilla-button')"
                                @click="startVanilla"
                            />
    
                            <Button 
                                plain outlined
                                class="btn mt-4" 
                                icon="pi pi-file-edit"
                                :label="$t('selected-game.config-button')"
                                @click="configEditorDialog.setVisible(true)"
                            />
                        </div>
                    </div>
                </div>
            </template>
        </Card>
    
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
                    <div v-if="hasSearchInput()">
                        <p>{{ $t('selected-game.empty-results') }}.</p>

                        <!-- Sadge -->
                        <img class="pt-3" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/2x.png">
                    </div>

                    <!-- TODO: If failed, make this show regardless of search input. --> 
                    <div v-else class="dataview-empty flex flex-column">
                        <p>No mods available! Something probably went wrong.</p>
                    </div>
                </template>
        
                <template #header>
                    <div class="flex row align-items-center">
                        <div class="searchbar no-drag">
                            <IconField iconPosition="left">
                                <InputIcon class="pi pi-search"></InputIcon>
                                <InputText type="text" :placeholder="$t('selected-game.search-mods')" 
                                    v-model="searchInput" 
                                    @input="onInputChanged"
                                />
                            </IconField>
                        </div>

                        <div class="ml-2 justify-self-end">
                            <TabMenu class="header-tab-menu" :model="tabMenuItems">

                            </TabMenu>
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
                                                <Button plain
                                                    class="btn w-full" 
                                                    icon="pi pi-download"
                                                    :label="$t('keywords.install')"
                                                    @click="installMod(mod.full_name)"
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

        <CardOverlay
            class="installing-mod-card no-drag"
            v-model:visible="installingModDialog.visible"
            v-model:closable="installingModDialog.closable"
            v-model:draggable="installingModDialog.draggable"
        >
            <template #cardContent>
                <div v-if="installing" class="flex column justify-content-center align-items-baseline">
                    <h1 class="mb-1 mt-2">Installing...</h1>
                    <p class="mt-1" style="font-size: 18px">{{ lastInstalledMod }}</p>
                </div>
                <div v-else>
                    <h1 class="mb-1 mt-2">Installed</h1>
                    <p class="mt-1" style="font-size: 18px">{{ lastInstalledMod }}</p>
                </div>
            </template>

            <template #dialogContent>
                <div style="position: sticky; bottom: 0;" class="flex justify-content-center w-full">
                    <div v-if="!installing" class="flex row gap-1 flex-grow-1">
                        <Button class="w-full"
                            type="button" severity="secondary"
                            :label="$t('keywords.close')" @click="installingModDialog.setVisible(false)"
                        />
                        <Button class="w-6"
                            type="button" severity="danger" icon="pi pi-trash"
                            label="Uninstall"
                        />
                    </div>
                    <!-- <Button v-else class="flex flex-grow-1" 
                        type="button" severity="danger" icon="pi pi-ban"
                        :label="$t('keywords.cancel')" 
                        @click=""
                    /> -->
                </div>
            </template>
        </CardOverlay>

        <CardOverlay 
            class="no-drag"
            v-model:visible="configEditorDialog.visible"
            v-model:closable="configEditorDialog.closable"
            v-model:draggable="configEditorDialog.draggable"
        >
            <template #cardContent>
                <div v-if="!selectedConfig" class="flex flex-column" style="max-height: calc(100vh - 180px);">
                    <!-- #region Heading & Subheading -->
                    <h1 class="header">{{ $t('selected-game.config-editor.title') }}</h1>
                    <p style="font-weight: 340; margin-bottom: 15px; margin-top: 3px; padding-left: 5px; user-select: none;">
                        Choose the configuration file you would like to edit values for.
                    </p>
        
                    <div v-if="configFiles.length < 1" class="flex justify-content-center align-items-center">
                        <p class="mb-3 mt-2" style="font-size: 18.5px; font-weight: 450; user-select: none;">No config files found!</p>
                    </div>
    
                    <div style="overflow-y: auto;">
                        <div 
                            v-if="configFiles.length > 0"
                            v-for="(path, index) in configFiles" :key="index" 
                            class="flex flex-row pl-2 pr-2 justify-content-between align-items-center"
                            style="height: 58.5px"
                        >
                            <p style="font-size: 18.5px; font-weight: 285; user-select: none;">{{ getRelativeConfigPath(path) }}</p>
    
                            <div class="flex gap-2">
                                <Button outlined plain
                                    icon="pi pi-folder"
                                    style="font-size: 17px; width: 3rem;"
                                    v-tooltip.top="'Open in Explorer'"
                                    @click="openLink(`file://${path}`)"
                                />
    
                                <Button plain
                                    class="justify-content-center"
                                    style="font-size: 17px; width: 5rem; height: 2.5rem;"
                                    :label="$t('selected-game.config-editor.edit-button')"
                                    @click="editConfig(path)"
                                />
                            </div>
                        </div>
                    </div>
                    <!-- #endregion -->
                </div>

                <ConfigEditLayout v-else :config="selectedConfig" :fileName="selectedConfigName">
                    
                </ConfigEditLayout>
            </template>
        
            <template #dialogContent>
                <div style="position: sticky; bottom: 0;" class="flex justify-content-end gap-2">
                    <Button v-if="!selectedConfig" 
                        class="w-full" type="button" severity="secondary" 
                        :label="$t('keywords.close')" @click="closeConfigEditor()"
                    />

                    <Button v-else 
                        class="w-5" type="button" severity="secondary" 
                        icon="pi pi-arrow-left" :label="$t('keywords.back')"
                        @click="resetSelectedConfig"
                    />

                    <Button v-if="selectedConfig"
                        class="w-full" type="button" 
                        :label="$t('keywords.apply')" @click=""
                    />
                </div>
            </template>
        </CardOverlay>
    </div>

    <!-- <Splitter style="height: 350px; background: none; border: none;" class="mb-9 no-drag mx-auto">
        <SplitterPanel class="flex" :minSize="33">

        </SplitterPanel>

        <SplitterPanel class="flex align-items-center"> 
            <DataView data-key="profile-selection">

            </DataView>
        </SplitterPanel>
    </Splitter> -->
</div>
</template>

<style scoped>
/*.breadcrumb-container {
    width: 140%;
    height: 140%;
}*/

.breadcrumb {
    margin-top: 20px;
    padding: 5px;
    justify-content: center;
    height: auto;
    background: none;
}

:deep(.breadcrumb *) {
    font-size: 19px;
    user-select: none;
    --wails-draggable: none;
}

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

.selected-game-card :deep(.p-card-body) {
    margin: 0px 30px 0px 30px;
    padding: 0px;
}

.selected-game-thumbnail {
    user-select: none;
    min-width: 160px;
    max-width: 40%;
    max-height: 200px;
    border-radius: 4px;
    border: 2px outset var(--primary-color);
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
    user-select: none;
    max-width: 85px;
    min-width: 60px;
    border-radius: 2.5px;
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

.dataview-empty {
    justify-content: center;
    align-items: center;
    display: flex;
}

.dataview-empty p {
    padding-top: 15px;
    font-size: 22px;
    margin: 0 auto;
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

/*.searchbar {
    
}*/

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

.installing-mod-card {
    padding: 0 1rem 1rem 1rem;
}

:deep(.p-tabmenu-nav) {
    background: none;
}

:deep(.p-menuitem-link) {
    background: none;
}
</style>