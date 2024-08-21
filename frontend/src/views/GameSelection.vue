<script lang="ts" setup>
import { ref, onMounted, Ref, computed, ComputedRef } from 'vue'

import DataView from 'primevue/dataview'
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions'
import Skeleton from 'primevue/skeleton'

// TODO: Replace with real external service.
import { getGameList } from '../mocks/GameService'

import { Game, Layout, OptionItem, ValueItem, ValueItemLabeled } from '@types'
import { Nullable } from 'primevue/ts-helpers'

import { t } from '@i18n'
import { BepinexInstalled } from '@backend/backend/GameManager'
import { GetPersistence, OpenExternal } from '@backend/app/Application'
import { tooltipOpts } from "../../src/util"

import { useGameStore } from '@stores'
import { storeToRefs } from 'pinia'

import router from '../router'

const store = useGameStore()
const { 
    games,
    isGameInstalled, 
    isFavouriteGame,
    gamesAsArray
} = storeToRefs(store)

const {
    toggleFavouriteGame,
    setSelectedGame,
    gameByID
} = store

const loading = ref(true)
const searchInput: Ref<Nullable<string>> = ref(null)
const layout: Ref<Layout> = ref('grid')

const selectedFilter: Ref<ValueItemLabeled<string>> = ref({
    label: "ALL",
    value: t('keywords.all')
})

const filters: ComputedRef<ValueItemLabeled<string>[]> = computed(() => [{
    label: "ALL",
    value: t('keywords.all')
}, {
    label: "INSTALLED",
    value: t('keywords.installed')
}, {
    label: "FAVOURITES",
    value: t('keywords.favourites')
}])

const openLink = (path: string) => OpenExternal(path).catch(e => console.log(`Error opening folder: ${e}`))

const alphabetSort = (games: Game[]) => {
    if ((searchInput.value?.length ?? 0) < 1) return games
    return games.sort((g1, g2) => g1 > g2 ? 1 : (g1 === g2 ? 0 : -1))
}

const getThumbnail = (game: Game) => game.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${game.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

const filterBySearch = (games: Game[]) => {
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

const getGames = (sort = true, searchFilter = true) => {
    let out = gamesAsArray.value

    const filter = selectedFilter.value.label
    if (filter == "INSTALLED") out = out.filter(g => g.installed)
    else if (filter == "FAVOURITES") out = out.filter(g => g.favourited)

    if (searchFilter) out = filterBySearch(out)
    if (sort) out = alphabetSort(out)

    return out
}

const selectGame = (game: Game) => {
    setSelectedGame(game)
    router.push('/selected-game')
}

onMounted(async () => {
    loading.value = true

    const persistence = await GetPersistence()
    const gameList = getGameList()

    for (const g of gameList) {
        if (!!g.path) {
            // TODO: Make sure the game executable exists (call backend)
            g.installed = true
            g.bepinexSetup = await BepinexInstalled(g.path)
        }

        // TODO: Set `g.favourited` using backend
        g.favourited = persistence.favourite_games.includes(g.identifier)
    }

    games.value = new Map(gameList.map(g => {
        // Make sure the cache persists through mounts
        const modCache = gameByID(g.identifier)?.modCache
        if (modCache) g.modCache = modCache

        return [g.identifier, g]
    }))

    loading.value = false
})
</script>

<template>
    <div class="game-selection flex-span column">
        <h2 class="header no-select">{{ $t('game-selection.header') }}</h2>

        <div class="card game-container no-drag">

            <!-- While loading, show a skeleton of a grid. -->
            <DataView v-if="loading" data-key="game-selection-loading" layout="grid">
                <template #empty>
                    <div class="scrollable grid-nogutter pt-4">
                        <div v-for="i in 15" :key="i" class="grid-item col-6 sm:col-3 md:col-3 lg:col-2 xl:col-1">
                            <div class="flex flex-column p-3 border-1 surface-border border-round">
                                <div class="flex flex-column align-items-center interact-section pb-3">
                                    <Skeleton width="6rem" height="2rem" />
                                </div>
                
                                <div class="flex justify-content-center border-round">
                                    <Skeleton width="75%" height="10rem" />
                                </div>
                
                                <div class="flex flex-column align-items-center interact-section pt-2">
                                    <div class="flex flex-column gap-3">
                                        <div class="flex gap-2 justify-content-center align-items-baseline">
                                            <Skeleton width="10rem" height="1.5rem" />
                                            <Skeleton width="1.5rem" height="1.5rem" shape="circle" />
                                        </div>
                
                                        <div class="flex flex-row gap-2">
                                            <Skeleton width="4rem" height="2rem" />
                                            <Skeleton width="2rem" height="2rem" shape="circle" />
                                            <Skeleton width="2rem" height="2rem" shape="circle" />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </template>
            </DataView>

            <!-- Finished loading, render the DataView with proper info -->
            <DataView v-else lazy data-key="game-selection" :value="getGames()" :layout="layout">
                <template #empty>
                    <div v-if="selectedFilter.label == 'FAVOURITES'" class="dataview-empty">
                        <p>No games favourited!</p>
                    </div>
                    <div v-else-if="selectedFilter.label == 'INSTALLED'" class="dataview-empty">
                        <p>No games installed!</p>
                    </div>
                    <div v-else-if="searchInput && searchInput.length > 0" class="dataview-empty flex flex-column">
                        <p>{{ `${$t('game-selection.empty-results')}.` }}</p>

                        <!-- Sadge -->
                        <!-- <img class="pt-3" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/2x.png"> -->

                        <!-- modCheck -->
                        <img class="pt-3" src="https://cdn.7tv.app/emote/60abf171870d317bef23d399/2x.gif">
                    </div>
                    <div v-else class="dataview-empty flex flex-column">
                        <p>No games available! Something probably went wrong.</p>
                    </div>
                </template>

                <!-- Header (filter, search, layout) -->
                <template #header>
                    <div class="flex flex-row justify-content-between align-items-center">
                        <div>
                            <Dropdown 
                                checkmark
                                class="no-drag filter-dropdown"
                                :options="filters"
                                v-model="selectedFilter"
                            >
                                <template #option="selectedItem: OptionItem<ValueItemLabeled<string>>">
                                    <div class="flex no-drag align-items-center">
                                        <div class="no-select">{{ selectedItem.option.value }}</div>
                                    </div>
                                </template>
                                
                                <template #value="selectedItem: ValueItem<ValueItemLabeled<string>>">
                                    <div v-if="selectedItem.value" class="flex align-items-center">
                                        <div>{{ selectedItem.value.value }}</div>
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

                <!-- Grid layout -->
                <template #grid>
                    <div class="scrollable grid grid-nogutter">
                        <div v-for="(game, index) in getGames()" :key="index" class="grid-item col-6 sm:col-5 md:col-4 lg:col-3 xl:col-2">
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
                                            <p class="m-0" style="font-size: 16.5px">{{ t('game-selection.bepinex-setup') }}</p>
                                            <i
                                                :class="['pi', isGameInstalled(game.identifier) ? 'pi-check' : 'pi-times']" 
                                                :style="{ color: isGameInstalled(game.identifier) ? 'lime' : 'red' }"
                                            />
                                        </div>

                                        <div class="flex flex-row gap-2">
                                            <Button
                                                outlined plain 
                                                class="grid-select-game-btn"
                                                :label="$t('game-selection.select-button')" 
                                                @click="selectGame(game)"
                                            />

                                            <Button
                                                outlined plain
                                                icon="pi pi-folder"
                                                v-if="game.path"
                                                v-tooltip.top="tooltipOpts(t('tooltips.game-selection.open-folder-location'))"
                                                @click="openLink(`file://${game.path}`)"
                                            />
                                            
                                            <Button
                                                outlined plain
                                                v-tooltip.top="tooltipOpts(isFavouriteGame(game.identifier) ? t('keywords.unfavourite') : t('keywords.favourite'))"
                                                :icon="isFavouriteGame(game.identifier) ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                                :style="isFavouriteGame(game.identifier) ? { color: 'var(--primary-color)' } : {}"
                                                @click="toggleFavouriteGame(game.identifier)"
                                            />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </template>

                <!-- List layout -->
                <template #list>
                    <div class="scrollable list list-nogutter">
                        <div v-for="(game, index) in getGames()" :key="index" class="col-12">
                            <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-5" :class="{ 'border-top-1 surface-border': index !== 0 }">
                                <img class="game-list-thumbnail fadeinleft fadeinleft-thumbnail block xl:block mx-auto w-full" :src="getThumbnail(game)"/>

                                <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                                    <div class="fadeinleft fadeinleft-title flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                                        <div class="game-list-title">{{ game.title }}</div>

                                        <div class="flex gap-2 justify-content-center align-items-baseline">
                                            <p class="m-0" style="font-size: 16.5px">{{ t('game-selection.bepinex-setup') }}</p>
                                            <i
                                                :class="['pi', isGameInstalled(game.identifier) ? 'pi-check' : 'pi-times']" 
                                                :style="{ color: isGameInstalled(game.identifier) ? 'lime' : 'red' }"
                                            />
                                        </div>
                                    </div>

                                    <div class="flex flex-column md:align-items-end gap-5">
                                        <div class="flex flex-row md:flex-row gap-3">
                                            <Button
                                                outlined plain
                                                icon="pi pi-folder"
                                                v-if="game.path"
                                                v-tooltip.top="tooltipOpts(t('tooltips.game-selection.open-folder-location'))"
                                                @click="openLink(`file://${game.path}`)"
                                            />

                                            <Button
                                                outlined plain
                                                :icon="isFavouriteGame(game.identifier) ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                                :style="isFavouriteGame(game.identifier) ? { color: 'var(--primary-color)' } : {}"
                                                @click="toggleFavouriteGame(game.identifier)"
                                            />

                                            <Button
                                                outlined plain
                                                :label="$t('game-selection.select-button')"
                                                class="list-select-game-btn flex-auto md:flex-initial"
                                                @click="setSelectedGame(game)"
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

.scrollable {
    overflow-y: scroll;
    scrollbar-width: none;
    height: calc(100vh - 155px); /* 100vh alone causes issues */
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
    width: 185px;
    border-radius: 3px;
}

.game-list-title {
    font-size: 26px;
    font-weight: 380;
}

.game-grid-title {
    font-size: 23px;
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

:deep(.p-dataview-header) {
    background: transparent !important;
    padding: 10px 0px 10px 0px;
    margin: 0px 5px 0px 5px;
    border: none;
}

:deep(.p-dataview-layout-options .p-button) {
    background: none !important;
    border: none;
}

:deep(.p-dataview-content) {
    background: none !important;
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
    padding-top: 15px;
    font-size: 22px;
    margin: 0 auto;
}

.filter-dropdown {
    width: 8.5rem;
    margin-right: 5px;
}
</style>