<script lang="ts" setup>
import { ref, onMounted } from 'vue'

import DataView from 'primevue/dataview'
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions'

import { Game, getGameList } from '../mocks/GameService'

type Layout = 'grid' | 'list'

const games = ref()

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
        <div class="card game-container no-drag">
            <DataView class="custom-dataview" data-key="game-list" :value="games" :layout="getLayout()">
                <template #header>
                    <div class="flex justify-content-start">
                        <DataViewLayoutOptions v-model="layout" />
                    </div>
                </template>

                <!-- List layout -->
                <template #list="slotProps">
                    <div class="grid grid-nogutter">
                        <div v-for="(item, index) in slotProps.items" :key="index" class="col-12">
                            <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-4" :class="{ 'border-top-1 surface-border': index !== 0 }">
                                <div class="md:w-4rem relative">
                                    <img class="block xl:block mx-auto border-round w-full" :src="getThumbnail(item)" :alt="item.name" />
                                </div>

                                <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                                    <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                                        <div>
                                            <span class="font-medium text-secondary text-sm">{{ item.category }}</span>
                                            <div class="text-lg font-medium text-900 mt-2">{{ item.title }}</div>
                                        </div>
                                    </div>
                                    <div class="flex flex-column md:align-items-end gap-5">
                                        <div class="flex flex-row-reverse md:flex-row gap-2">
                                            <Button outlined plain icon="pi pi-star"></Button>
                                            <Button outlined plain label="Select Game" class="flex-auto md:flex-initial white-space-nowrap"></Button>
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
                            <div class="flex flex-column p-4 border-1 surface-border surface-card border-round">
                                <div class="flex justify-content-center border-round">
                                    <div class="relative mx-auto">
                                        <img class="grid-img" :src="getThumbnail(item)" :alt="item.name"/>
                                    </div>
                                </div>

                                <div class="flex flex column align-items-center interact-section pt-3">
                                    <div>
                                        <div class="text-xl font-medium text-900">{{ item.title }}</div>
                                    </div>

                                    <div class="flex gap-4 mt-3">
                                        <div class="flex gap-2">
                                            <Button outlined plain label="Select Game" class="flex flex-grow-1 white-space-nowrap"></Button>
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
}

:deep(.p-dataview-header)  {
    background: transparent !important;
    border: none;
}

:deep(.p-dataview-content)  {
    background: transparent !important;
}

.grid-img {
    user-select: none;
    width: 180px;
}
</style>