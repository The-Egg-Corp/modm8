<script lang="ts" setup>
import { ref, onMounted, Ref, computed, ComputedRef, reactive } from 'vue'

import DataView from 'primevue/dataview'
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions'

// TODO: Replace with real external service.
import { getGameList } from '../mocks/GameService'

import { Game, Layout, OptionItem, ValueItem } from '@types'
import { Nullable } from 'primevue/ts-helpers'

import { t } from '@i18n'
import { BepinexInstalled } from '@backend/backend/GameManager'
import { OpenWindowAtLocation } from '@backend/core/App'
import { tooltipOpts } from "../../src/util"

const searchInput: Ref<Nullable<string>> = ref(null)
const layout: Ref<Layout> = ref('grid')

const getThumbnail = (game: Game) => game.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${game.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

function filterBySearch(games: Game[]) {
    if (!searchInput.value) return games

    const input = searchInput.value.trim()
    if (input == "") return games

    const lowerInput = input.toLowerCase()
    const inputWords = lowerInput.split(" ")

    return games.filter(g => {
        const lowerTitle = g.title?.toLowerCase() ?? ""

        // Necessary to not show irrelevent games with only 1 letter input.
        if (input.length == 1 && !lowerTitle.startsWith(lowerInput)) {
            return false
        }

        const wordMatch = inputWords.some(word => lowerTitle.includes(word))
        if (wordMatch) return true

        // Check if each character in the input appears in order in the title
        let inputIndex = 0
        let matchFound = false

        const titleLen = lowerTitle.length

        for (let titleIndex = 0; titleIndex < titleLen; titleIndex++) {
            if (lowerTitle[titleIndex] !== input[inputIndex]) {
                inputIndex = 0
                continue
            }

            inputIndex++
            if (inputIndex === input.length) {
                matchFound = true
                break
            }
        }

        // If we've matched all characters in the input, it's a match
        return matchFound
    })
}

const alphabetSort = (games: Game[]) => {
    if ((searchInput.value?.length ?? 0) < 1) return games
    return games.sort((g1, g2) => g1 > g2 ? 1 : (g1 === g2 ? 0 : -1))
}

const selectedFilter: Ref<Nullable<string>> = ref(t('keywords.all'))
const filters: ComputedRef<string[]> = computed(() => [
    t('keywords.all'), 
    t('keywords.installed')
])

const installedStatuses: Map<string, boolean> = reactive(new Map())

const games: Ref<Game[]> = ref([])
const getGames = (sort = true, filter = true) => {
    let out = games.value

    if (filter) out = filterBySearch(out)
    if (sort) out = alphabetSort(out)

    return out
}

const openWindowAtLoc = (path: string) => OpenWindowAtLocation(path).catch(console.log)

onMounted(() => {
    games.value = getGameList()

    games.value.forEach(async g => {
        const installed = await (g.path ? BepinexInstalled(g.path) : false)
        installedStatuses.set(g.identifier, installed)
    })
})
</script>

<template>
    <div class="game-selection flex-span column">
        <h2 class="header no-select">{{ $t('game-selection.header') }}</h2>

        <div class="card game-container no-drag">
            <DataView lazy data-key="game-list" :value="getGames()" :layout="layout">
                <template #empty>
                    <div class="dataview-empty">
                        <p>{{ `${$t('game-selection.empty-results')}. 😔` }}</p>
                    </div>
                </template>

                <!-- Header (filter, search, layout) -->
                <template #header>
                    <div class="flex flex-row justify-content-between align-items-center">
                        <div>
                            <Dropdown 
                                class="no-drag w-full w-8rem" checkmark
                                :options="filters" v-model="selectedFilter"
                            >
                                <template #option="selectedItem: OptionItem<string>">
                                    <div class="flex no-drag align-items-center">
                                        <div class="no-select">{{ selectedItem.option }}</div>
                                    </div>
                                </template>
                                
                                <template #value="selectedItem: ValueItem<string>">
                                    <div v-if="selectedItem.value" class="flex align-items-center">
                                        <div>{{ selectedItem.value }}</div>
                                    </div>
                                </template>
                            </Dropdown>
                        </div>

                        <div class="searchbar">
                            <IconField iconPosition="left">
                                <InputIcon class="pi pi-search"></InputIcon>
                                <InputText type="text" :placeholder="$t('game-selection.search-placeholder')" v-model="searchInput"/>
                            </IconField>
                        </div>
                        
                        <div class="flex flex-row">
                            <DataViewLayoutOptions v-model="layout">
                                <template #listicon>
                                    <div class="pi pi-list"></div>
                                </template>
                            </DataViewLayoutOptions>
                        </div>
                    </div>
                </template>

                <!-- List layout -->
                <template #list>
                    <div class="grid grid-nogutter">
                        <div v-for="(game, index) in getGames()" :key="index" class="col-12">
                            <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-5" :class="{ 'border-top-1 surface-border': index !== 0 }">
                                <img class="game-list-thumbnail fadeinleft fadeinleft-thumbnail block xl:block mx-auto w-full" :src="getThumbnail(game)"/>

                                <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                                    <div class="fadeinleft fadeinleft-title flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                                        <div class="game-list-title">{{ game.title }}</div>
                                    </div>

                                    <div class="flex flex-column md:align-items-end gap-5">
                                        <div class="flex flex-row md:flex-row gap-3">
                                            <Button 
                                                outlined plain
                                                icon="pi pi-folder"
                                                v-if="game.path"
                                                v-tooltip.top="tooltipOpts(t('tooltips.game-selection.open-folder-location'))"
                                                @click="openWindowAtLoc(game.path)"
                                            />

                                            <Button outlined plain icon="pi pi-star"></Button>
                                            <Button 
                                                outlined plain 
                                                :label="$t('game-selection.select-button')" 
                                                class="list-select-game-btn flex-auto md:flex-initial"
                                            />
                                        </div>
                                    </div>
                                </div>
                            </div>  
                        </div>
                    </div>
                </template>

                <!-- Grid layout -->
                <template #grid>
                    <div class="grid grid-nogutter">
                        <div v-for="(game, index) in getGames()" :key="index" class="grid-item col-6 sm:col-3 md:col-3 lg:col-2 xl:col-1">
                            <div class="flex flex-column p-3 border-1 surface-border border-round">
                                <div class="flex flex-column align-items-center interact-section pb-3">
                                    <div class="game-grid-title">{{ game.title }}</div>
                                </div>

                                <div class="flex justify-content-center border-round">
                                    <div class="relative mx-auto">
                                        <img class="game-grid-thumbnail" :src="getThumbnail(game)"/>
                                    </div>
                                </div>

                                <div class="flex flex-column align-items-center interact-section pt-2">
                                    <div class="flex flex-column gap-3">
                                        <div class="flex gap-2 justify-content-center align-items-baseline">
                                            <p class="m-0" style="font-size: 16.5px">BepInEx Installed</p>
                                            <i
                                                :class="['pi', 'pi-spin', installedStatuses.get(game.identifier) ? 'pi-check' : 'pi-times']" 
                                                :style="{ color: installedStatuses.get(game.identifier) ? 'lime' : 'red' }"
                                            />
                                        </div>

                                        <div class="flex flex-row gap-2">
                                            <Button 
                                                outlined plain 
                                                :label="$t('game-selection.select-button')" 
                                                class="grid-select-game-btn"
                                                @click=""
                                            />

                                            <Button 
                                                outlined plain 
                                                icon="pi pi-folder" 
                                                v-if="game.path"
                                                v-tooltip.top="tooltipOpts(t('tooltips.game-selection.open-folder-location'))"
                                                @click=""
                                            />
                                            
                                            <Button 
                                                outlined plain 
                                                icon="pi pi-star"
                                                v-tooltip.top="tooltipOpts(t('keywords.favourite'))"
                                                @click=""
                                            />
                                        </div>


                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </template>
            </DataView>
        </div>
    </div>
</template>

<style scoped>
.game-selection {
    padding-top: 0px;
}

.game-selection .header {
    text-wrap: wrap;
    text-align: center;
    font-size: 33px;
    font-weight: 420;
    margin: 40px 0px 10px 0px;
    padding: 0px 20px 0px 20px; /* Makes text get wrapped earlier */
}

.game-container {
    margin-left: 50px;
    margin-right: 50px;
}

.grid {
    overflow-y: scroll;
    scrollbar-width: none;
    height: calc(100vh - 155px); /* 100vh alone causes spacing issues */
}

.grid-item {
    min-width: fit-content;
    flex: 1 0 auto;
    padding: 5px;
    margin: 0;
}

.game-list-thumbnail {
    user-select: none;
    max-width: 111px;
    min-width: 30px;
    opacity: 0;
    border-radius: 3px;
}

.game-grid-thumbnail {
    user-select: none;
    width: 210px;
    border-radius: 4px;
}

.game-list-title {
    font-size: 26px;
    font-weight: 380;
}

.game-grid-title {
    font-size: 24.5px;
    font-weight: 400;
}

.grid-select-game-btn {
    display: flex;
    flex-grow: 1;
    white-space: nowrap;
    font-size: 16px;
}

.list-select-game-btn {
    display: flex;
    white-space: nowrap;
    font-size: 18px;
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

:deep(.p-dataview-header)  {
    background: transparent !important;
    padding: 10px 0px 10px 0px;
    margin: 0px 5px 0px 5px;
    border: none;
}

:deep(.p-dataview-layout-options .p-button)  {
    background: transparent !important;
    border: none;
}

:deep(.p-dataview-content)  {
    background: transparent !important;
}

@keyframes fadeinleft {
    0% {
        opacity: 0;
        transform: translateX(-200px);
        transition: transform .12s cubic-bezier(0, 0, 0.2, 1), opacity .12s cubic-bezier(0, 0, 0.2, 1);
    }
    100% {
        opacity: 1;
        transform: translateX(0%);
    }
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

.dataview-empty {
    justify-content: center;
    align-items: center;
    display: flex;
}

.dataview-empty p {
    font-size: 22.5px;
    margin: 0 auto;
}
</style>