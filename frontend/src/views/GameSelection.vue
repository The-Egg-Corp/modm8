<script lang="ts" setup>
import { 
    ref, computed, 
    nextTick, onMounted,
    onBeforeMount
} from 'vue'

import { Viewport } from '@components'

import DataView from 'primevue/dataview'
import SelectButton from 'primevue/selectbutton'

// TODO: Replace with real external service. Use ecosystem logic we implemented in the backend.
import { mockGameList } from '../mocks/GameService'

import type { 
    ThunderstoreGame, 
    Layout, OptionItem, 
    ValueItem, ValueItemLabeled,
    Nullable,
} from '@types'

import { t } from '@i18n'
import { tooltipOpts, openLink } from "../../src/util"

import { useGameStore, useGameStoreTS } from '@stores'
import { storeToRefs } from 'pinia'

import router from '../router'

const gameStore = useGameStore()
const gameStoreTS = useGameStoreTS()
const { gamesAsArray } = storeToRefs(gameStoreTS)

const loading = ref(true)
const searchInput = ref<Nullable<string>>(null)

const layout = ref<Layout>('grid')
const options = ref(['list', 'grid'])

const scrollableList = ref<HTMLElement | null>(null)

//#region Game Filter (Dropdown)
type FilterValueItem = ValueItemLabeled<string, GameFilterType>
type GameFilterType = keyof typeof GameFilter
const GameFilter = {
    ALL: "ALL",
    FAVOURITES: "FAVOURITES",
    INSTALLED: "INSTALLED"
} as const

const selectedFilter = ref<FilterValueItem>({
    label: GameFilter.ALL,
    value: t('keywords.all')
})

const filters = computed<FilterValueItem[]>(() => [{
    label: GameFilter.ALL,
    value: t('keywords.all')
}, {
    label: GameFilter.INSTALLED,
    value: t('keywords.installed')
}, {
    label: GameFilter.FAVOURITES,
    value: t('keywords.favourites')
}])
//#endregion

//#region Search Filter (Search bar input box)
const alphabetSort = (games: ThunderstoreGame[]) => {
    const searchInputLen = searchInput.value?.length ?? 0
    return searchInputLen < 1 ? games : games.sort((g1, g2) => g1.title.localeCompare(g2.title)) // TODO: .toLowerCase() the titles?
}

const filterBySearch = (games: ThunderstoreGame[]) => {
    if (!searchInput.value) return games

    const input = searchInput.value.trim()
    if (input == "") return games

    const lowerInput = input.toLowerCase()
    const inputWords = lowerInput.split(" ")

    return games.filter(g => {
        const lowerTitle = g.title?.toLowerCase() ?? ""

        // Only show relevent games that start with this (only) letter.
        if (input.length == 1 && !lowerTitle.startsWith(lowerInput)) {
            return false
        }

        const wordMatch = inputWords.some(word => lowerTitle.includes(word))
        if (wordMatch) return true

        //#region Check if each character in the input appears in order in the title.
        let inputIdx = 0
        let matchFound = false

        const titleLen = lowerTitle.length

        for (let i = 0; i < titleLen; i++) {
            if (lowerTitle[i] !== input[inputIdx]) {
                inputIdx = 0
                continue
            }

            inputIdx++
            if (inputIdx === input.length) {
                matchFound = true
                break
            }
        }
        //#endregion

        // Match found - title matches all characters in the input.
        return matchFound
    })
}

const getThunderstoreGames = (sort = true, searchFilter = true) => {
    let games = gamesAsArray.value

    const filter = selectedFilter.value.label
    if (filter == GameFilter.INSTALLED) games = games.filter(g => g.installed)
    else if (filter == GameFilter.FAVOURITES) games = games.filter(g => g.favourited)

    if (searchFilter) games = filterBySearch(games)
    if (sort) games = alphabetSort(games)

    return games
}

const selectThunderstoreGame = (game: ThunderstoreGame) => {
    const tsGame = gameStoreTS.gameByID(game.identifier)
    if (!tsGame) {
        console.warn(`Failed to select Thunderstore game. ${game.title} not found within GameStoreTS game cache.`)
        return
    }

    gameStore.setSelectedGame('THUNDERSTORE', tsGame)
    router.push('/selected-game')
}

const gameElements = ref<any[]>([])
const scrollIndex = ref(0)

/**
 * Scrolls to a game in the list using the specified index.\
 * For the game to successfully scroll into view, index must be valid and not OOB of `gameElements`.
 * 
 * @returns Whether we successfully scrolled to the game.
 */
function scrollToGame(idx: number) {
    // Because ! check considers 0 falsy.
    if (idx == null || idx == undefined) {
        console.warn(`Failed to scroll to game. Specified index is ${idx}`)
        return false
    }
    
    const games = gameElements.value
    if (idx < 0 || idx >= games.length) {
        console.warn(`Prevented OOB access of 'gameElements' [${games.length}] with index: ${idx}`)
        return false
    }

    const game: Element = games[idx]
    if (!game) return false

    // We found the game element, scroll to it.
    game.scrollIntoView(true)
    return true
}

// NOTE: Might want to make this a ref in future so it can be user defined.
const SCROLL_STEP = 1

const handleScroll = (e: WheelEvent) => {
    // The vertical scroll amount. Positive = down | Negative = up
    if (e.deltaY > 0) {     
        const lastGameIndex = gameElements.value.length - 1
        scrollIndex.value = Math.min(scrollIndex.value + SCROLL_STEP, lastGameIndex)
    } else {
        scrollIndex.value = Math.max(scrollIndex.value - SCROLL_STEP, 0)
    }

    // On DOM update complete
    nextTick(() => scrollToGame(scrollIndex.value))
}

// TODO: Replace with better alternative. These could change at any time.
const getThumbnail = (game: ThunderstoreGame) => game.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${game.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

onBeforeMount(async () => {
    const t0 = performance.now()

    // NOTE: We currently need to initialize the game store cache every mount in-case properties (path, installed) change.
    //       This may be an issue in future, ignore it until we transition from our static mock list.
    const tsGames = mockGameList.filter(g => g.platform == 'THUNDERSTORE').map(g => g.value)
    const size = await gameStoreTS.initGames(tsGames)

    loading.value = false

    console.info(`Populated GameStoreTS cache with ${size} games. Took: ${performance.now() - t0}ms`)
    if (size != mockGameList.length) {
        console.warn("Size of GameStoreTS cache does not match original input.")
    }
})
</script>

<template>
<Viewport class="flex-full column">
    <h2 class="header no-select">{{ t('game-selection.header') }}</h2>

    <div class="card no-drag">
        <!-- While loading, show a skeleton of a grid. -->
        <!-- <DataView v-if="loading" data-key="game-selection-loading" layout="grid">
            <template #empty>
                <div class="grid grid-nogutter pt-1">
                    <div v-for="i in 15" :key="i" class="grid-item col-2 sm:col-6 md:col-5 lg:col-2 xl:col-2">
                        <div class="flex flex-column p-3">
                            <div class="flex flex-column align-items-center pb-2">
                                <Skeleton width="8rem" height="2.3rem" />
                            </div>
                            
                            <div class="flex justify-content-center border-round">
                                <Skeleton width="75%" height="16.5rem" />
                            </div>

                            <div class="flex flex-column align-items-center pt-2">
                                <div class="flex flex-column gap-2">
                                    <div class="flex flex-row gap-2 justify-content-center">
                                        <Skeleton width="2.5rem" height="2rem" />
                                        <Skeleton width="2.5rem" height="2rem" />
                                    </div>

                                    <div class="flex gap-2 justify-content-center align-items-baseline">
                                        <Skeleton width="16rem" height="2.5rem" />
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </template>
        </DataView> -->

        <!-- Finished loading, render the DataView with proper info -->
        <DataView lazy data-key="game-selection" :value="getThunderstoreGames()" :layout="layout">
            <template #empty>
                <div v-if="selectedFilter.label == GameFilter.FAVOURITES" class="dataview-empty">
                    <p>No games favourited!</p>
                </div>
                <div v-else-if="selectedFilter.label == GameFilter.INSTALLED" class="dataview-empty">
                    <p>No games installed!</p>
                </div>
                <div v-else-if="searchInput && searchInput.length > 0" class="dataview-empty flex flex-column">
                    <p>{{ `${t('game-selection.empty-results')}.` }}</p>

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
                    <div class="flex row gap-2">
                        <div class="searchbar">
                            <!-- <IconField iconPosition="left">
                                <InputIcon class="pi pi-search"/>
                                <InputText :placeholder="t('game-selection.search-placeholder')" v-model="searchInput"/>
                            </IconField> -->

                            <FloatLabel variant="on">
                                <IconField>
                                    <InputIcon class="pi pi-search"/>
                                    <InputText id="search_title" v-model="searchInput"/>
                                </IconField>
                    
                                <label for="search_title">{{ t('game-selection.search-placeholder') }}</label>
                            </FloatLabel>
                        </div>

                        <Dropdown checkmark
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
                    
                    <div class="flex row layout-btns">
                        <SelectButton v-model="layout" :options="options" :allowEmpty="false">
                            <template #option="{ option }">
                                <i :class="[option === 'list' ? 'pi pi-list' : 'pi pi-th-large']" />
                            </template>
                        </SelectButton>
                    </div>
                </div>
            </template>

            <!-- Grid layout -->
            <template #grid>
                <div class="scrollable-grid grid grid-nogutter">
                    <div 
                        v-for="(game, index) in getThunderstoreGames()" :key="index" 
                        class="grid-item border-faint col-6 sm:col-6 md:col-4 lg:col-3 xl:col-2"
                    >
                        <div class="flex column game-card gap-2">
                            <div class="flex column align-items-center">
                                <div class="game-grid-title">{{ game.title }}</div>
                            </div>

                            <div class="flex column relative mx-auto">
                                <img class="game-grid-thumbnail" :src="getThumbnail(game)"/>

                                <!-- <div class="flex gap-2 justify-content-center align-items-baseline mt-2 mb-3">
                                    <p class="m-0" style="font-size: 16px">{{ t('game-selection.bepinex-setup') }}</p>
                                    <i
                                        :class="['pi', game.bepinexSetup ? 'pi-check' : 'pi-times']" 
                                        :style="{ color: game.bepinexSetup  ? 'lime' : 'red' }"
                                    />
                                </div> -->
                            </div>

                            <div class="flex row justify-content-center gap-2">
                                <Button outlined plain
                                    v-tooltip.top="tooltipOpts(game.favourited ? t('keywords.unfavourite') : t('keywords.favourite'))"
                                    :icon="game.favourited ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                    :style="game.favourited ? { color: 'var(--primary-color)' } : {}"
                                    @click="gameStoreTS.toggleFavouriteGame(game.identifier)"
                                />

                                <Button outlined plain
                                    v-if="game.installed"
                                    v-tooltip.top="tooltipOpts(t('tooltips.game-selection.open-folder-location'))"
                                    icon="pi pi-folder" @click="openLink(`file://${game.path}`)"
                                />
                            </div>

                            <div class="grid-item-bottom-row justify-content-center">
                                <Button severity="primary"
                                    class="grid-select-game-btn"
                                    :label="t('game-selection.select-button')" 
                                    @click="selectThunderstoreGame(game)"
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </template>

            <!-- List layout -->
            <template #list>
                <div class="scrollable-list list-nogutter" ref="scrollableList" tabindex="0" @wheel.prevent="handleScroll">
                    <div 
                        v-for="(game, index) in getThunderstoreGames()" class="snap-top col-12" 
                        :key="index" :ref="el => gameElements[index] = el"
                    >
                        <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-4" :class="{ 'border-top-faint': index !== 0 }">
                            <img class="game-list-thumbnail fadeinleft fadeinleft-thumbnail block xl:block mx-auto w-full" :src="getThumbnail(game)"/>

                            <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1">
                                <div class="fadeinleft fadeinleft-title flex flex-row md:flex-column justify-content-between align-items-start gap-1">
                                    <div class="game-list-title">{{ game.title }}</div>

                                    <div class="flex gap-2 justify-content-center align-items-baseline">
                                        <p class="m-0" style="font-size: 16px">{{ t('game-selection.bepinex-setup') }}</p>
                                        <i
                                            :class="['pi', game.installed ? 'pi-check' : 'pi-times']" 
                                            :style="{ color: game.installed ? 'lime' : 'red' }"
                                        />
                                    </div>
                                </div>

                                <div class="flex flex-column md:align-items-end gap-5">
                                    <div class="flex flex-row md:flex-row gap-2">
                                        <Button
                                            outlined plain
                                            icon="pi pi-folder"
                                            v-if="game.installed"
                                            v-tooltip.top="tooltipOpts(t('tooltips.game-selection.open-folder-location'))"
                                            @click="openLink(`file://${game.path}`)"
                                        />

                                        <Button
                                            outlined plain
                                            :icon="game.favourited ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                            :style="game.favourited ? { color: 'var(--primary-color)' } : {}"
                                            @click="gameStoreTS.toggleFavouriteGame(game.identifier)"
                                        />

                                        <Button
                                            :label="t('game-selection.select-button')"
                                            class="list-select-game-btn flex-auto md:flex-initial"
                                            @click="selectThunderstoreGame(game)"
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
</Viewport>
</template>

<style scoped>
.header {
    text-wrap: nowrap;
    text-align: left;
    font-size: 36px;
    font-weight: 450;
    margin: 0px 0px 10px 0px;
}

.snap-top {
    height: 160px;
}

.scrollable-list {
    overflow-y: scroll; /* Enable vertical scrolling */
    scrollbar-width: none;
    height: calc(160px * 5); /* Height of single item * amount of items */
}

.scrollable-grid {
    overflow-y: scroll;
    scrollbar-width: none;
    height: calc(100vh - 140px); /* TODO: Investigate why 100vh alone prevents scrolling to bottom. */
}

/* Container for the game-card */
.grid-item {
    display: flex;
    flex: 1 0 auto;
    min-width: fit-content;
    border-radius: 3px;
    margin: 5px;
}

.game-card {
    padding: 10px 25px 10px 25px;
    justify-content: center;
    flex: 1 0 auto;
}

.game-grid-thumbnail {
    user-select: none;
    width: 193px;
    border-radius: 4px;
}

.game-list-thumbnail {
    user-select: none;
    max-width: 105px;
    min-width: 40px;
    opacity: 0;
    border-radius: 1px;
}

.game-list-title {
    font-size: 26px;
    font-weight: 360;
}

.game-grid-title {
    font-size: 24px;
    font-weight: 360;
    text-shadow: 0px 0px 10px rgba(255, 255, 255, 0.45);
    text-wrap: nowrap;
}

/* Should contain at least the select button */
.grid-item-bottom-row {
    display: flex;
}

.grid-select-game-btn {
    white-space: nowrap;
    font-size: 18px;
    width: 100%;
    max-width: 500px;
}

.list-select-game-btn {
    display: flex;
    white-space: nowrap;
    font-size: 18px;
}

.layout-btns {
    border: 1px solid var(--p-inputtext-border-color);
    border-radius: 6.5px;
}

.searchbar {
    width: calc(22vw);
    min-width: 335px;
}

:deep(.searchbar .p-inputtext) {
    margin-left: auto;
    margin-right: auto;
    width: 100%;
}

:deep(.p-dataview-header) {
    background: transparent !important;
    padding: 5px 0px 5px 0px;
    margin: 0px 5px 0px 5px;
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