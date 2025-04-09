<script lang="ts" setup>
import { ref, nextTick, computed } from 'vue'
import { storeToRefs } from 'pinia'

import { TabMenuChangeEvent } from 'primevue/tabmenu'
import DataView, { DataViewPageEvent } from 'primevue/dataview'
import Tag from 'primevue/tag'

import {
    useModListStore,
    useGameStore,
} from '@stores'

import { ModListTabType } from '@types'
import { Dialog } from '@composables'

import { debounce } from "../../util"

const gameStore = useGameStore()

const modListStore = useModListStore()
const { 
    installMod, refreshMods,
    updatePage, refreshPage,
    getMods, ROWS
} = modListStore.thunderstore

const {
    loading,
    activeTab,
    searchInput,
    modElements, scrollIndex, first,
    mods, currentPageMods,
} = storeToRefs(modListStore.thunderstore)

// Choose which mods to show based on tab type.
const dataViewMods = computed(() => {
    if (activeTab.value == ModListTabType.NEXUS) {
        // TODO: Implement state for nexus mods.
        return []
    }
    
    if (activeTab.value == ModListTabType.TS) {
        return mods.value
    }

    // Not a store, must be profile tab.
    return []
})

const tabs = ref([
    { type: ModListTabType.PROFILE, label: 'This Profile', icon: 'pi pi-box' },
    { type: ModListTabType.TS, label: 'Thunderstore', icon: 'pi pi-globe' },
    { type: ModListTabType.NEXUS, label: 'Nexus', icon: 'pi pi-globe' } // Uncomment when Nexus is implemented.
])

const onTabChange = async (e: TabMenuChangeEvent) => {
    const newTab = tabs.value[e.index]
    activeTab.value = newTab.type

    await refreshPage()
}

const onPageChange = (e: DataViewPageEvent) => updatePage(e.first, e.rows)

const hasSearchInput = () => searchInput.value ? searchInput.value.length > 0 : undefined
const onSearchInputChange = async () => {
    // No input, no need to debounce.
    if (!searchInput.value?.trim()) {
        // Show all mods without filtering.
        mods.value = getMods(false)
        return await refreshPage()
    }

    debouncedSearch()
}

const debouncedSearch = debounce(() => refreshMods(false), 250)

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

const props = defineProps<{ installingModDialog: Dialog }>()

// TODO: Implement Thunderstore login using GitHub/Discord for things like rating mods.
// This may require an update to Wails V3 so we can make OAuth easier as it has plugins and multi-window support.
const openLoginPage = () => {
    // const url = 'https://auth.thunderstore.io/auth/login/github'
    // const loginWindow = window.open(url)

    // if (!loginWindow) return

    // loginWindow.onload = () => {
        
    // }
}
</script>

<template>
<!-- Show skeleton of mod list while loading -->
<DataView v-if="loading" data-key="mod-list-loading" layout="list">
    <template #empty>
        <div class="list-nogutter pt-4">
            <div v-for="i in 6" :key="i" class="loading-list-item">
                <div style="width: 1280px;" class="flex flex-row ml-1 p-3 border-top-faint border-round">
                    <Skeleton size="6.5rem"/> <!-- Thumbnail -->
                    
                    <div class="flex column gap-1 ml-2">
                        <Skeleton height="1.5rem" width="20rem"/> <!-- Title -->
                        <Skeleton width="65rem"/> <!-- Description -->

                        <div class="flex row gap-2">
                            <Skeleton class="mt-3" width="6.8rem" height="2.2rem"/> <!-- Install Button -->
                            
                            <div class="flex row gap-1 align-items-center">
                                <Skeleton class="mt-3" width="2.8rem" height="2.2rem"/> <!-- Rate button -->
                                <Skeleton class="mt-3" width="1.8rem" height="1.6rem"/> <!-- Rate Count -->
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
    :paginator="dataViewMods.length > ROWS" :rows="ROWS"
    :value="dataViewMods" @page="onPageChange" :first="first"
>
    <template #header>
        <div class="flex row align-items-center gap-2">
            <div class="mod-searchbar no-drag">
                <FloatLabel variant="on">
                    <IconField>
                        <InputIcon class="pi pi-search" />
                        <InputText id="search_mods" v-model="searchInput" @input="onSearchInputChange"/>
                    </IconField>
        
                    <label for="search_mods">{{ $t('selected-game.search-mods') }}</label>
                </FloatLabel>
            </div>

            <TabMenu :model="tabs" @tab-change="onTabChange"/>
            <!-- <div class="flex row">
                <ModListDropdown>
                    
                </ModListDropdown>
            </div> -->
        </div>
    </template>

    <template #empty>
        <div v-if="hasSearchInput()" class="pl-2">
            <h2 class="m-0 mt-1">{{ $t('selected-game.empty-results') }}.</h2>

            <!-- Sadge -->
            <img class="mt-2" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/3x.png">
        </div>

        <!-- TODO: If failed, make this show regardless of search input. --> 
        <div v-else>
            <h2 v-if="activeTab == ModListTabType.PROFILE" class="empty-profile">
                {{ $t('selected-game.no-mods-installed') }}
            </h2>

            <div v-if="activeTab == ModListTabType.TS" class="ml-1">
                <h2 class="mb-2" style="color: red; font-size: 24px; margin: 0 auto;">
                    No mods available! Something probably went wrong.
                </h2>

                <Button class="mt-1" :label="$t('keywords.refresh')" icon="pi pi-refresh" @click="refreshMods(true)"/>
            </div>

            <div v-if="activeTab == ModListTabType.NEXUS" class="ml-1">
                <h2 class="mb-2" style="color: red; font-size: 24px; margin: 0 auto;">
                    Nexus Mods support is not implemented yet.
                </h2>
            </div>
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
                                    <!-- <Button v-if="activeTab == ModListTabType.PROFILE" 
                                        class="btn w-full" severity="danger" icon="pi pi-trash"
                                        :label="$t('keywords.uninstall')"
                                    />
                                    <Button v-if="activeTab == ModListTabType.TS" 
                                        class="btn w-full" icon="pi pi-download"
                                        :label="$t('keywords.install')" @click="installMod(mod.full_name, gameStore.thunderstore.selectedGame, props.installingModDialog)"
                                    /> -->

                                    <Button
                                        class="btn w-full" icon="pi pi-download"
                                        :label="$t('keywords.install')" @click="installMod(mod.full_name, gameStore.thunderstore.selectedGame, props.installingModDialog)"
                                    />

                                    <div class="flex row align-items-center">
                                        <Button outlined plain 
                                            style="margin-right: 6.5px;"
                                            :icon="'pi pi-thumbs-up'"
                                            @click="openLoginPage"
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
</template>

<style scoped>
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
    padding-bottom: 6px;
    padding-top: 3px;
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
    height: calc(100vh - 160px); /* This seems arbitrary and won't work on different resolutions? */
    /* width: calc(100vw - v-bind(sidebarOffsetPx) - 420px); TODO: Replace 420px. Get it dynamically from selected game card width. */
}

.empty-profile {
    color: rgba(235, 235, 235, 0.95);
    font-size: 25px;
    margin: 0;
}

/* TODO: Investigate why this padding affects profile manager. */
:deep(.p-dataview-header) {
    background: none !important;
    padding: 0px 0px 10px 0px;
    margin: 0;
    border: none;
}

:deep(.p-dataview-content) {
    background: none !important;
}

:deep(.mod-searchbar .p-inputtext) {
    background: rgba(0, 0, 0, 0.2);
    margin-left: auto;
    margin-right: auto;
    width: 350px;
    min-width: 200px;
}

:deep(.p-paginator) {
    background: none !important;
    padding: 0;
}

:deep(.p-dataview-paginator-bottom) {
    border: none;
    justify-self: center;
    min-width: 500px;
    margin-top: 10px;
}

.list-item {
    display: flex;
    width: 100%;
    padding: 0px 0px 10px 0px;
}

.loading-list-item {
    display: flex;
    width: 100vw;
    padding-bottom: 15px;
    padding-top: 0px;
}
</style>