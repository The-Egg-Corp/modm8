<script lang="ts" setup>
import { ref, nextTick, computed } from 'vue'
import { storeToRefs } from 'pinia'

import { TabMenuChangeEvent } from 'primevue/tabmenu'
import DataView, { DataViewPageEvent } from 'primevue/dataview'
import Tag from 'primevue/tag'

import {
    useGameStore,
    useProfileStore,
    useModListStore,
    useModListStoreTS,
} from '@stores'

import { ModListTabs, Package } from '@types'
import { Dialog } from '@composables'

import { debounce } from "../../util"

const profileStore = useProfileStore()
const { selectedProfile } = storeToRefs(profileStore)

const gameStore = useGameStore()
const { selectedGame } = storeToRefs(gameStore)

// Shared mod list state for all mod sites/platforms.
const modListStore = useModListStore()
const { activeTab, searchInput } = storeToRefs(modListStore)

// Mod list store for Thunderstore only.
const modListStoreTS = useModListStoreTS()
const { 
    refreshMods,
    updatePage, refreshPage,
    getMods, ROWS
} = modListStoreTS

const {
    loading,
    modElements, scrollIndex, first,
    mods, currentPageMods,
} = storeToRefs(modListStoreTS)

/** Choose which mods to show based on tab type. */
const dataViewMods = computed(() => {
    if (activeTab.value == ModListTabs.NEXUS) {
        // TODO: Implement state for nexus mods.
        return []
    }
    
    if (activeTab.value == ModListTabs.TS) {
        return mods.value
    }

    // Not a store, must be profile tab - return all installed mods that are in this profile.
    // TODO: Add nexus support here, otherwise they won't show up.
    return []
})

const tabs = ref([
    { type: ModListTabs.PROFILE, label: 'This Profile', icon: 'pi pi-box' },
    { type: ModListTabs.TS, label: 'Thunderstore', icon: 'pi pi-globe' },
    { type: ModListTabs.NEXUS, label: 'Nexus', icon: 'pi pi-globe' } // Uncomment when Nexus is implemented.
])

// TODO: Store index in `activeTab` itself?
const activeTabIdx = computed(() => {
    const index = tabs.value.findIndex(tab => tab.type == activeTab.value)
    return index < 0 ? 0 : index
})

const switchTab = async (e: TabMenuChangeEvent) => {
    const newTab = tabs.value[e.index]
    activeTab.value = newTab.type

    await refreshPage()
}

const onPageChange = (e: DataViewPageEvent) => {
    updatePage(e.first, e.rows)

    const scrolled = scrollToMod(0) // Attempt to scroll to first mod.
    if (scrolled) scrollIndex.value = 0 // Update cur index if successful.
}

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

/**
 * Scrolls to a mod in the list using the specified index.\
 * For the mod to successfully scroll into view, index must be valid and not OOB of `modElements`.
 * 
 * @returns Whether we successfully scrolled to the mod.
 */
function scrollToMod(idx: number) {
    // Because ! check considers 0 falsy.
    if (idx == null || idx == undefined) {
        console.warn(`Failed to scroll to mod. Specified index is ${idx}`)
        return false
    }
    
    const mods = modElements.value
    if (idx < 0 || idx >= mods.length) {
        console.warn(`Prevented OOB access of 'modElements' [${mods.length}] with index: ${idx}`)
        return false
    }

    const mod: Element = mods[idx]
    if (!mod) return false

    // We found the mod element, scroll to it.
    mod.scrollIntoView(true)
    return true
}

// TODO: Should we be manipulating `scrollIndex` directly before attempting to scroll?
//       If we can't scroll to it then it should remain its old value like how `onPageChange` does it.
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
    nextTick(() => scrollToMod(scrollIndex.value))
}

async function installTsMod(mod: Package) {
    if (selectedGame.value.type != 'THUNDERSTORE') {
        throw new Error('Could not install Thunderstore mod. Selected game is not of type `THUNDERSTORE`.')
    }

    return modListStoreTS.installMod(mod.full_name, selectedGame.value.value, props.installingModDialog)
}

const props = defineProps<{ installingModDialog: Dialog }>()

// TODO: Implement Thunderstore login using GitHub/Discord for things like rating mods.
// This may require an update to Wails V3 so we can make OAuth easier as it has plugins and multi-window support.
// const TS_GITHUB_LOGIN = 'https://auth.thunderstore.io/auth/login/github'
// const openLoginPage = () => {
//     const loginWindow = window.open(TS_GITHUB_LOGIN)
//     if (!loginWindow) return

//     loginWindow.onload = () => {
        
//     }
// }
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

            <TabMenu :model="tabs" :activeIndex="activeTabIdx" @tab-change="switchTab"/>
            <!-- <div class="flex row">
                <ModListDropdown>
                    
                </ModListDropdown>
            </div> -->
        </div>
    </template>

    <template #empty>
        <div v-if="selectedProfile == null">
            <h2 v-if="activeTab == ModListTabs.PROFILE" class="empty-profile">
                No profile selected.
            </h2>
        </div>
        <div v-else>
            <!-- TODO: If failed, make this show regardless of search input. --> 
            <div v-if="!hasSearchInput() /*|| failed*/">
                <h2 v-if="activeTab == ModListTabs.PROFILE" class="empty-profile">
                    {{ $t('selected-game.no-mods-installed') }}
                </h2>

                <div v-if="activeTab == ModListTabs.TS" class="ml-1">
                    <h2 class="mb-2" style="color: red; font-size: 24px; margin: 0 auto;">
                        No mods available! Something probably went wrong.
                    </h2>

                    <Button class="mt-1" :label="$t('keywords.refresh')" icon="pi pi-refresh" @click="refreshMods(true)"/>
                </div>

                <div v-if="activeTab == ModListTabs.NEXUS" class="ml-1">
                    <h2 class="mb-2" style="color: red; font-size: 24px; margin: 0 auto;">
                        Nexus Mods support is not implemented yet.
                    </h2>
                </div>
            </div>
            <div v-else class="pl-2">
                <h2 class="m-0 mt-1">{{ $t('selected-game.empty-results') }}.</h2>

                <!-- Sadge -->
                <img class="mt-2" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/3x.png">
            </div>
        </div>
    </template>

    <template #list>
        <div v-if="activeTab == ModListTabs.NEXUS" class="ml-1">
            <h2 class="mb-2" style="color: red; font-size: 24px; margin: 0 auto;">
                Nexus Mods support is not implemented yet.
            </h2>
        </div>
        <div v-else class="scrollable-list list-nogutter no-drag" @wheel.prevent="handleScroll">
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
                                        :label="$t('keywords.install')" @click="installMod(mod.full_name, gameStoreTS.selectedGame, props.installingModDialog)"
                                    /> -->

                                    <Button
                                        class="btn w-full" icon="pi pi-download"
                                        :label="$t('keywords.install')"
                                        :disabled="selectedProfile == null"
                                        @click="installTsMod(mod)"
                                    />

                                    <div class="flex row align-items-center">
                                        <!-- TODO: On click, call openLoginPage when implemented. -->
                                        <Button outlined plain
                                            style="margin-right: 6.5px;"
                                            :icon="'pi pi-thumbs-up'"
                                            @click=""
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