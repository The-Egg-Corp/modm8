<script lang="ts" setup>
import { computed, onMounted, ref } from "vue"
import type { ComputedRef, Ref } from "vue"

import DataView, { DataViewPageEvent } from 'primevue/dataview'
import BlockUI from 'primevue/blockui'

import Breadcrumb from 'primevue/breadcrumb'
import CardOverlay from '../components/reusable/CardOverlay.vue'

// import Splitter from 'primevue/splitter'
// import SplitterPanel from 'primevue/splitterpanel'

import Card from 'primevue/card'
import Button from 'primevue/button'

import { GetLatestPackageVersion, GetPackagesStripped } from '@backend/thunderstore/API'
import { LaunchSteamGame } from '@backend/steam/SteamRunner'
import { FindCfgFiles } from '@backend/backend/GameManager'
import { useGameStore } from "@stores"
import { useDialog } from '@composables'
import { thunderstore } from "@backend/models"
import { BreadcrumbPage, Package } from "@types"
import { Nullable } from "primevue/ts-helpers"

const gameStore = useGameStore()
const { selectedGame, updateModCache } = gameStore

const { 
    setVisible,
    visible, closable, draggable 
} = useDialog('selected-game')

const loading = ref(false)
const searchInput: Ref<Nullable<string>> = ref(null)
const currentPageMods: Ref<Package[]> = ref([])

const homePage: Ref<BreadcrumbPage> = ref({
    home: true,
    route: '/game-selection',
    icon: "stadia_controller",
    class: "material-symbols-sharp text-color-secondary"
})

const pages: ComputedRef<BreadcrumbPage[]> = computed(() => [{ 
    label: selectedGame.title,
    route: '/selected-game',
    class: "text-primary"
}])

const ROWS = 15
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

    await updatePage(0, ROWS)
    loading.value = false
})

const latestModVersion = (mod: thunderstore.StrippedPackage) => 
    GetLatestPackageVersion(selectedGame.identifier, mod.owner, mod.name)

const gameThumbnail = () => selectedGame.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${selectedGame.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

const startVanilla = () => LaunchSteamGame(selectedGame.id, ["--doorstop-enable", "false"])
const startModded = () => LaunchSteamGame(selectedGame.id, ["--doorstop-enable", "true"])

const getConfigFiles = async () => {
    if (!selectedGame.bepinexSetup || !selectedGame.path) return []
    
    const files = await FindCfgFiles([selectedGame.path, "BepInEx", "config"])
    const configPath = 'BepInEx/config'

    return files.map(file => {
        const normalizedFile = file.replace(/\\/g, '/')
        const startIndex = normalizedFile.indexOf(configPath) + configPath.length + 1

        return normalizedFile.substring(startIndex)
    })
}

const editConfig = (path: string) => {
    // Read file contents from backend
}

const updatePage = async (first: number, rows: number) => {
    const mods = (selectedGame.modCache || []).slice(first, first + rows) as Package[]
    const promises = mods.map(async mod => {
        mod.latestVersion = await latestModVersion(mod)
        return mod
    })

    currentPageMods.value = await Promise.all(promises)
}

const onPageChange = async (e: DataViewPageEvent) => {
    await updatePage(e.first, e.rows)
    //console.log(currentPageMods.value)
}
</script>

<template>
<div :class="['selected-game', { 'no-drag': visible }]">
    <Breadcrumb class="breadcrumb flex-full row" :home="homePage" :model="pages">
        <template #item="{ item, props }">
            <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                <a :href="href" v-bind="props.action" @click="navigate">
                    <span v-if="item.icon" :class="item.class" :style="item.style">{{ item.icon }}</span>
                    <span v-else-if="item.label" :class="item.class" :style="item.style">{{ item.label }}</span>
                </a>
            </router-link>
        </template>
        <template #separator>/</template>
    </Breadcrumb>

    <BlockUI fullScreen v-if="loading"></BlockUI>
    
    <div class="flex row gap-8 no-drag">
        <Card class="selected-game-card">
            <template #title>
                <p style="font-size: 30px; font-weight: 520; user-select: none;" class="no-drag mt-0 mb-1">{{ $t('selected-game.currently-selected') }}</p>
            </template>
    
            <template #content>
                <div class="flex no-drag sm:flex-column md:flex-column lg:flex-row xl:flex-row">
                    <img class="selected-game-thumbnail" :src="gameThumbnail()"/>
                    
                    <div class="flex column">
                        <p style="font-size: 25px; font-weight: 330" class="mt-0 mb-0 ml-3">{{ selectedGame.title }}</p>
                        <div class="flex column gap-2 mt-3">
                            <Button 
                                plain
                                class="btn ml-3" 
                                icon="pi pi-caret-right"
                                :label="$t('selected-game.start-modded-button')"
                                @click="startModded"
                            />
    
                            <Button 
                                plain outlined
                                class="btn ml-3" 
                                icon="pi pi-caret-right"
                                :label="$t('selected-game.start-vanilla-button')"
                                @click="startVanilla"
                            />
    
                            <Button 
                                plain outlined
                                class="btn ml-3 mt-4" 
                                icon="pi pi-file-edit"
                                :label="$t('selected-game.config-button')"
                                @click="setVisible(true)"
                            />
                        </div>
                    </div>
                </div>
            </template>
        </Card>
    
        <DataView lazy paginator stripedRows 
            v-if="!loading" class="w-8"
            style="margin: 70px 0px 0px 1.5%;"
            layout="list" data-key="mod-list"
            :value="currentPageMods"
            :totalRecords="selectedGame.modCache?.length"
            :rows="ROWS"
            @page="onPageChange"
        >
            <template #empty>
                <div class="dataview-empty flex flex-column">
                    <p>No mods available! Something probably went wrong.</p>
                </div>
            </template>
    
            <template #header>
                <div class="searchbar">
                    <IconField iconPosition="left">
                        <InputIcon class="pi pi-search"></InputIcon>
                        <InputText type="text" :placeholder="$t('selected-game.search-mods')" v-model="searchInput"/>
                    </IconField>
                </div>
            </template>
    
            <template #list>
                <div class="scrollable list list-nogutter">
                    <div v-for="(mod, index) in currentPageMods" :key="index" class="list-item col-10">
                        <div class="flex flex-column sm:flex-row sm:align-items-center pt-2 gap-3" :class="{ 'border-top-1 surface-border': index != 0 }">
                            <img class="mod-list-thumbnail fadeinleft fadeinleft-thumbnail block xl:block mx-auto w-full" :src="mod.latestVersion?.icon || ''"/>
                            
                            <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                                <div class="fadeinleft fadeinleft-title flex column justify-content-between align-items-start gap-2">
                                    
                                    <div class="flex row align-items-baseline">
                                        <div class="mod-list-title">{{ mod.name }}</div>
                                        <div class="mod-list-author">({{ mod.owner }})</div>
                                    </div>

                                    <div class="mod-list-description">{{ mod.latestVersion.description }}</div>

                                    <!--
                                        :icon="isFavouriteGame(game.identifier) ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                        :style="isFavouriteGame(game.identifier) ? { color: 'var(--primary-color)' } : {}"
                                        @click="toggleFavouriteGame(game.identifier)"
                                    /> -->

                                    <div class="flex row align-items-baseline">
                                        <Button outlined plain 
                                            class="mt-1" style="margin-right: 10px;"
                                            :icon="'pi pi-thumbs-up'"
                                        />

                                        <div class="mod-list-rating">{{ mod.rating_score }}</div>
                                    </div>
                                </div>

                                <div>
                                    <Button plain
                                        class="btn ml-3" 
                                        icon="pi pi-download"
                                        :label="$t('keywords.install')"
                                        @click=""
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </template>
        </DataView>
    
        <CardOverlay 
            class="no-drag"
            v-model:visible="visible"
            v-model:closable="closable"
            v-model:draggable="draggable"
        >
            <template #cardContent>
                <div class="flex flex-column" style="max-height: calc(100vh - 180px);">
                    <!-- #region Heading & Subheading -->
                    <h1 class="header">{{ $t('selected-game.config-editor.title') }}</h1>
                    <p style="font-weight: 340; margin-bottom: 15px; margin-top: 3px; padding-left: 5px; user-select: none;">
                        Choose the configuration file you would like to edit values for.
                    </p>
        
                    <div v-if="configFiles.length < 1" class="flex justify-content-center align-items-center">
                        <p class="mb-1 mt-2" style="font-size: 18.5px; font-weight: 450; user-select: none;">No config files found!</p>
                    </div>
    
                    <div style="overflow-y: auto;">
                        <div 
                            v-if="configFiles.length > 0"
                            v-for="(path, index) in configFiles" :key="index" 
                            class="flex flex-row pl-2 pr-2 justify-content-between align-items-center"
                            style="height: 58.5px"
                        >
                            <p style="font-size: 18.5px; font-weight: 285; user-select: none;">{{ path }}</p>
    
                            <div class="flex gap-2">
                                <Button outlined plain
                                    icon="pi pi-folder"
                                    style="font-size: 17px; width: 3rem;"
                                    v-tooltip.top="'Open in Explorer'"
                                    @click=""
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
            </template>
        
            <template #dialogContent>
                <div style="position: sticky; bottom: 0;" class="flex justify-content-end gap-3">
                    <Button class="w-full" type="button" :label="$t('keywords.close')" severity="secondary" @click="setVisible(false)"></Button>
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

/*.selected-game {
    
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

.selected-game-card {
    background: none;
    width: fit-content;
}

:deep(.selected-game-card .p-card-body) {
    padding: 15px 0px 0px 40px;
}

.selected-game-thumbnail {
    user-select: none;
    width: 162.5px;
    max-height: 100%;
    border-radius: 4px;
    border: 2px outset var(--primary-color);
}

.btn {
    border-radius: 4px;
    text-align: left;
}

.mod-list-title {
    font-size: 21px;
    font-weight: 460;
    padding-right: 5px;
}

.mod-list-author {
    font-size: 16px;
    font-weight: 310;
}

.mod-list-description {
    font-size: 16.5px;
    font-weight: 220;
}

.mod-list-rating {
    font-size: 17px;
    font-weight: 370;
}

.mod-list-thumbnail {
    user-select: none;
    max-width: 80px;
    min-width: 60px;
    opacity: 0;
    border-radius: 3px;
}

.fadeinleft {
    --title-duration: 450ms;
}

.fadeinleft-thumbnail {
    animation-duration: 300ms;
    animation-delay: calc(var(--title-duration) - 100ms);
    animation-fill-mode: forwards;
}

.fadeinleft-title {
    animation-duration: var(--title-duration);
    animation-delay: 50ms;
}

.scrollable {
    overflow-y: scroll;
    scrollbar-width: none;
    height: calc(100vh - 265px);
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

.searchbar {
    margin-left: auto;
    margin-right: auto;
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
}

.list-item {
    padding-bottom: 15px;
    padding-top: 0px;
}
</style>