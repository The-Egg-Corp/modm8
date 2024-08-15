<script lang="ts" setup>
import { computed, onMounted, ref } from "vue"
import type { ComputedRef, Ref } from "vue"

import Breadcrumb from 'primevue/breadcrumb'
import CardOverlay from '../components/reusable/CardOverlay.vue'

// import Splitter from 'primevue/splitter'
// import SplitterPanel from 'primevue/splitterpanel'

import Card from 'primevue/card'
import Button from 'primevue/button'

import { GetPackagesStripped } from '@backend/thunderstore/API'
import { LaunchSteamGame } from '@backend/steam/SteamRunner'
import { FindCfgFiles } from '@backend/backend/GameManager'
import { useGameStore } from "@stores"
import { useDialog } from '@composables'

const { 
    selectedGame
} = useGameStore()

interface BreadcrumbPage {
    route: string
    home?: boolean
    label?: string
    icon?: string
    class?: string
    style?: string
}

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

const getThumbnail = () => selectedGame.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${selectedGame.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

const startVanilla = () => LaunchSteamGame(selectedGame.id, ["--doorstop-enable", "false"])
const startModded = () => LaunchSteamGame(selectedGame.id, ["--doorstop-enable", "true"])

const { 
    setVisible,
    visible, closable, draggable 
} = useDialog('selected-game')

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

let configFiles: string[] = []

onMounted(async () => {
    configFiles = await getConfigFiles()
    
    if (!selectedGame.modCache) {
        console.log(`Putting ${selectedGame.identifier} mods into cache...`)

        const t0 = performance.now()
        const pkgs = await GetPackagesStripped(selectedGame.identifier, false)
        //updateModCache(pkgs)

        console.log(`Cached ${pkgs?.length} mods. Took: ${performance.now() - t0}ms`)
    }
})

const editConfig = (path: string) => {
    // Read file contents from backend
}
</script>

<template>
<div :class="['selected-game row', { 'no-drag': visible }]">
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

    <Card class="current-game-card">
        <template #title>
            <p style="font-size: 30px; font-weight: 520; user-select: none;" class="no-drag mt-0 mb-1">{{ $t('selected-game.currently-selected') }}</p>
        </template>

        <template #content>
            <div class="no-drag ml-1">
                <img class="current-game-thumbnail" :src="getThumbnail()"/>
                <div class="flex column" style="float: right">
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

    <!-- <DataView lazy data-key="mod-list" :value="selectedGame.modCache" layout="list">

    </DataView> -->

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

.current-game-card {
    background: none;
    width: fit-content;
}

:deep(.current-game-card .p-card-body) {
    padding: 15px 0px 0px 40px;
}

.current-game-thumbnail {
    user-select: none;
    width: 162.5px;
    border-radius: 4px;
    border: 2px outset var(--primary-color);
}

.btn {
    position: relative;
    border-radius: 5px;
    text-align: left;
}
</style>