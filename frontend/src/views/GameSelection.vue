<script lang="ts" setup>
import { ref, onMounted, Ref } from 'vue'

import DataView from 'primevue/dataview'
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions'

// TODO: Replace with real external service.
import { getGameList } from '../mocks/GameService'

import { Nullable } from 'primevue/ts-helpers'
import { Game, GameProps, Layout } from '@types'

const games: Ref<Game[]> = ref([])
const searchInput: Ref<Nullable<string>> = ref(null)
const layout: Ref<Layout> = ref('grid')

const getThumbnail = (game: Game) => game.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${game.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

onMounted(() => games.value = getGameList())

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
</script>

<template>
    <div class="game-selection flex-span column">
        <h2 class="header no-select">{{ $t('game-selection.header') }}</h2>
        <div class="no-drag card game-container">
            <DataView lazy data-key="game-list" :value="filterBySearch(games)" :layout="layout">
                <template #empty>
                    <div class="dataview-empty">
                        <p>No games match the search query! 😔</p>
                    </div>
                </template>

                <template #header>
                    <div class="flex flex-row justify-content-between align-items-center">
                        <div class="flex flex-row">
                            <DataViewLayoutOptions v-model="layout"/>
                        </div>

                        <div class="searchbar">
                            <IconField iconPosition="left">
                                <InputIcon class="pi pi-search"></InputIcon>
                                <InputText type="text" :placeholder="$t('game-selection.search-placeholder')" v-model="searchInput"/>
                            </IconField>
                        </div>

                        <div class="flex flex-row">
                            <DataViewLayoutOptions v-model="layout"/>
                        </div>
                    </div>
                </template>

                <!-- List layout -->
                <template #list="slotProps: GameProps">
                    <div class="grid grid-nogutter">
                        <div v-for="(item, index) in slotProps.items" :key="index" class="col-12">
                            <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-5" :class="{ 'border-top-1 surface-border': index !== 0 }">
                                <img class="game-list-thumbnail fadeinleft fadeinleft-thumbnail block xl:block mx-auto w-full" :src="getThumbnail(item)"/>

                                <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                                    <div class="fadeinleft fadeinleft-title flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                                        <div class="game-list-title">{{ item.title }}</div>
                                    </div>

                                    <div class="flex flex-column md:align-items-end gap-5">
                                        <div class="flex flex-row md:flex-row gap-3">
                                            <Button outlined plain :label="$t('game-selection.select-button')" class="flex-auto md:flex-initial white-space-nowrap"></Button>
                                            <Button outlined plain icon="pi pi-star"></Button>
                                        </div>
                                    </div>
                                </div>
                            </div>  
                        </div>
                    </div>
                </template>

                <!-- Grid layout -->
                <template #grid="slotProps: GameProps">
                    <div class="grid grid-nogutter">
                        <div v-for="(item, index) in slotProps.items" :key="index" class="grid-item col-6 sm:col-2 md:col-3 lg:col-4 xl:col-2">
                            <div class="flex flex-column p-3 border-1 surface-border border-round">
                                <div class="flex justify-content-center border-round">
                                    <div class="relative mx-auto">
                                        <img class="game-grid-thumbnail" :src="getThumbnail(item)"/>
                                    </div>
                                </div>

                                <div class="flex flex-column align-items-center interact-section pt-3">
                                    <div class="game-grid-title">{{ item.title }}</div>
                                    <div class="flex gap-4 mt-3">
                                        <div class="flex gap-2">
                                            <Button outlined plain :label="$t('game-selection.select-button')" class="flex flex-grow-1 white-space-nowrap"></Button>
                                            <Button outlined plain icon="pi pi-star-fill" class="star-btn"></Button>
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
    margin: 45px 0px 5px 0px;
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
    max-width: 90px;
    min-width: 30px;
    opacity: 0;
    border-radius: 3px;
}

.game-grid-thumbnail {
    user-select: none;
    width: 200px;
    border-radius: 4px;
}

.game-list-title {
    font-size: 24px;
    font-weight: 380;
}

.game-grid-title {
    font-size: 22px;
    font-weight: 380;
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