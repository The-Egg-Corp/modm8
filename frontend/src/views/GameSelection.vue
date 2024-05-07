<script lang="ts" setup>
import { ref, onMounted } from 'vue'

import DataView from 'primevue/dataview'
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions'

import { Game, getGameList } from '../mocks/GameService'

type Layout = 'grid' | 'list'

const games = ref()
const searchValue = ref()

const layout = ref('grid')
const getLayout = () => layout.value as Layout

const getThumbnail = (game: Game) => game.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${game.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

onMounted(() => {
    let data = getGameList()
    games.value = data
})

const sortMostPopular = () => {
    
}
</script>

<template>
    <div class="game-selection flex-span column">
        <h2 class="header no-select">{{ $t('game-selection.header') }}</h2>
        <div class="no-drag card game-container">
            <DataView lazy class="custom-dataview" data-key="game-list" :value="games" :layout="getLayout()">
                <template #header>
                    <div class="flex flex-row justify-content-between align-items-center">
                        <IconField iconPosition="left">
                            <InputIcon class="pi pi-search"></InputIcon>
                            <InputText :v-model="searchValue" placeholder="Search"/>
                        </IconField>

                        <DataViewLayoutOptions v-model="layout"/>
                    </div>
                </template>

                <!-- List layout -->
                <template #list="slotProps">
                    <div class="no-drag grid grid-nogutter">
                        <div v-for="(item, index) in slotProps.items" :key="index" class="col-12">
                            <div class="fadeinleft fadeinleft-props flex flex-column sm:flex-row sm:align-items-center p-2 gap-5" :class="{ 'border-top-1 surface-border': index !== 0 }">
                                <div class="relative">
                                    <img class="game-list-thumbnail block xl:block mx-auto w-full" :src="getThumbnail(item)" :alt="item.name" />
                                </div>

                                <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                                    <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                                        <div class="game-title">{{ item.title }}</div>
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
                <template #grid="slotProps">
                    <div class="grid grid-nogutter">
                        <div v-for="(item, index) in slotProps.items" :key="index" class="col-12 sm:col-6 md:col-6 xl:col-4 p-2">
                            <div class="flex flex-column p-4 border-1 surface-border border-round">
                                <div class="flex justify-content-center border-round">
                                    <div class="relative mx-auto">
                                        <img class="game-grid-thumbnail" :src="getThumbnail(item)" :alt="item.name"/>
                                    </div>
                                </div>

                                <div class="flex flex column align-items-center interact-section pt-3">
                                    <div class="game-title">{{ item.title }}</div>
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
    padding-top: 50px;
}

.game-selection .header {
    text-wrap: wrap;
    text-align: center;
    font-size: 30px;
    font-weight: 420;
    margin-top: 0;
    margin-bottom: 0;
}

.game-container {
    margin-left: 60px;
    margin-right: 60px;
}

.grid {
    display: sticky;
    overflow-y: scroll;
    scrollbar-width: none;
    height: calc(100vh - 160px);
}

.game-card {
    border: white 1px;
    max-width: 50px;
}

.game-list-thumbnail {
    border-radius: 2.5px;
    max-width: 90px;
    min-width: 75px;
}

.game-grid-thumbnail {
    user-select: none;
    width: 190px;
}

.game-title {
    font-size: 22px;
    font-weight: 380;
}

:deep(.p-dataview-header)  {
    background: transparent !important;
    padding: 10px 0px 10px 0px;
    margin: 0px 8px 0px 8px;
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
        transform: translateX(-100%);
        transition: transform .12s cubic-bezier(0, 0, 0.2, 1), opacity .12s cubic-bezier(0, 0, 0.2, 1);
    }
    100% {
        opacity: 1;
        transform: translateX(0%);
    }
}

.fadeinleft-props {
    animation-delay: 70ms;
    animation-duration: 580ms;
    animation-iteration-count: 1;
}
</style>