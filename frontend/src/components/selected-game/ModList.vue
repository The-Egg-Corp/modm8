<script lang="ts" setup>
import { ref, nextTick, computed, watch } from 'vue'
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

import { ModListTabs, Nullable, Profile } from '@types'
import { Dialog } from '@composables'

import {
    AddThunderstoreModToProfile,
    RemoveThunderstoreModFromProfile
} from '@backend/profile/ProfileManager'

import {
    LinkModToProfile,
    UnlinkModFromProfile
} from '@backend/game/GameManager'

import { loaders, thunderstore } from '@backend/models'
import { debounce } from "../../util"

const profileStore = useProfileStore()
const { selectedProfile } = storeToRefs(profileStore)

const gameStore = useGameStore()
const { selectedGame } = storeToRefs(gameStore)

// Shared mod list state for all mod sites/platforms.
const modListStore = useModListStore()
const { 
    activeTab, searchInput, 
    modElements, scrollIndex
} = storeToRefs(modListStore)

// Mod list store for Thunderstore only.
const modListStoreTS = useModListStoreTS()
const { 
    refreshMods, filterByProfile,
    updatePage, PAGE_ROWS
} = modListStoreTS

const {
    loading,
    pageFirstRecordIdx,
    mods, currentPageMods,
} = storeToRefs(modListStoreTS)

const tsProfileMods = ref<thunderstore.StrippedPackage[]>([])

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

    await refreshPageWithMods()
}

const onPageChange = (e: DataViewPageEvent) => {
    updatePage(e.first, e.rows)

    const scrolled = scrollToModElement(0) // Attempt to scroll to first mod.
    if (scrolled) scrollIndex.value = 0 // Update cur index if successful.
}

const debouncedSearch = debounce(() => refreshMods(false, { searchFilter: true }), 250)
const hasSearchInput = () => searchInput.value ? searchInput.value.length > 0 : undefined

async function onSearchInputChange() {
    await refreshPageWithMods()
}

// TODO: This won't work in the profile tab since we use `tsProfileMods` there.
async function refreshPageWithMods() {
    // No input, no need to debounce.
    if (!searchInput.value?.trim()) {
        // Don't filter by search, show all mods.
        return await refreshMods(true, { searchFilter: false })
    }

    debouncedSearch()
}

const profileHasMod = (prof: Nullable<Profile>, modVerFullName: string) => {
    return prof?.mods.thunderstore?.some(m => m.toLowerCase() == modVerFullName.toLowerCase())
}

async function initTsProfileMods() {
    const cache = selectedGame.value.value.modCache ?? []
    tsProfileMods.value = cache.length > 0 ? await filterByProfile(cache) : []
}

async function installTsMod(mod: thunderstore.StrippedPackage) {
    if (selectedGame.value.platform != 'THUNDERSTORE') {
        throw new Error('Cannot install Thunderstore mod. Selected game platform is not `THUNDERSTORE`.')
    }

    if (!selectedProfile.value?.name) {
        throw new Error('Cannot install Thunderstore mod. Selected profile is not valid.')
    }

    const selectedGameTitle = selectedGame.value.value.title
    const selectedProfName = selectedProfile.value.name

    // The user wants this mod. Update manifest regardless of whether mod will successfully install.
    // We can notify them in the profile tab if something is wrong and provide the option to start without it.
    await AddThunderstoreModToProfile(selectedGameTitle, selectedProfName, mod.latest_version.full_name)
    
    // We installed the mod to ../GameName/ModCache. 
    const success = await modListStoreTS.installMod(mod.full_name, selectedGame.value.value, props.installingModDialog)
    if (success) {
        // Create a symlink/junction for the mod and every one of its dependencies in the loader's mod path.
        // Ex: ../GameName/Profiles/SelectedProfile/BepInEx/plugins
        
        // TODO: Replace ModLoader.BEPINEX with one the game actually uses.
        const loader = loaders.ModLoaderType.BEPINEX

        await LinkModToProfile(loader, selectedGameTitle, selectedProfName, mod.latest_version.full_name)
        for (const dep of mod.latest_version.dependencies) {
            await LinkModToProfile(loader, selectedGameTitle, selectedProfName, dep)
        }
    }

    // Keep profiles refreshed and in line with manifest.
    await initTsProfileMods()
    await profileStore.initProfiles()
}

async function uninstallTsMod(mod: thunderstore.StrippedPackage) {
    if (selectedGame.value.platform != 'THUNDERSTORE') {
        throw new Error('Cannot uninstall Thunderstore mod. Selected game platform is not `THUNDERSTORE`.')
    }

    if (!selectedProfile.value?.name) {
        throw new Error('Cannot uninstall Thunderstore mod. Selected profile is not valid.')
    }

    const selectedGameTitle = selectedGame.value.value.title
    const selectedProfName = selectedProfile.value.name

    await RemoveThunderstoreModFromProfile(selectedGameTitle, selectedProfName, mod.latest_version.full_name)
    // TODO: Delete from ModCache if no profile uses it?
    

    // TODO: Replace ModLoader.BEPINEX with one the game actually uses.
    const loader = loaders.ModLoaderType.BEPINEX

    // Drop every symlink/junction for the mod and its dependencies.
    await UnlinkModFromProfile(loader, selectedGameTitle, selectedProfName, mod.latest_version.full_name)
    for (const dep in mod.latest_version.dependencies) {
        await UnlinkModFromProfile(loader, selectedGameTitle, selectedProfName, dep)
    }

    // Keep profiles refreshed and in line with manifest.
    await initTsProfileMods()
    await profileStore.initProfiles()
}

watch(selectedProfile, async (newVal, oldVal) => {
    if (newVal == null) return
    if (newVal == oldVal) return

    await initTsProfileMods()
})

/**
 * Scrolls to a mod in the list using the specified index.\
 * For the mod to successfully scroll into view, index must be valid and not OOB of `modElements`.
 * 
 * @returns Whether we successfully scrolled to the mod.
 */
function scrollToModElement(idx: number) {
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
    nextTick(() => scrollToModElement(scrollIndex.value))
}

const props = defineProps<{
    installingModDialog: Dialog
}>()

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
<div class="flex column">
    <!-- #region HEADER (search, tabs, filter) -->
    <div class="flex row gap-2">
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
    </div>
    <!-- #endregion -->

    <!-- #region PROFILE TAB -->
    <DataView v-if="activeTab == ModListTabs.PROFILE" 
        layout="list" data-key="mod-list-profile"
        :value="tsProfileMods"
    >
        <template #empty>
            <h2 v-if="selectedProfile == null" class="empty-profile">No profile selected.</h2>
            <div v-else>
                <div v-if="!hasSearchInput() /*|| failed*/">
                    <h2 class="empty-profile">{{ $t('selected-game.no-mods-installed') }}</h2>
                </div>
                <div v-else class="pl-2">
                    <h2 class="m-0 mt-1">{{ $t('selected-game.empty-results') }}.</h2>

                    <!-- Sadge -->
                    <img class="mt-2" src="https://cdn.7tv.app/emote/603cac391cd55c0014d989be/3x.png">
                </div>
            </div>
        </template>

        <template #list>
            <div class="scrollable-list list-nogutter no-drag" @wheel.prevent="handleScroll">
                <div 
                    v-for="(mod, index) in tsProfileMods" class="list-item col-12"
                    :key="index" :ref="el => modElements[index] = el"
                >
                    <div class="flex-grow-1 flex column sm:flex-row align-items-center pt-2 gap-3" :class="{ 'border-top-faint': index != 0 }">
                        <img class="mod-list-thumbnail block xl:block" :src="mod.latest_version?.icon || ''"/>
                        
                        <div class="flex-grow-1 flex column md:flex-row md:align-items-center">
                            <div class="flex-grow-1 flex column justify-content-between">
                                <div class="flex row align-items-baseline">
                                    <div class="mod-list-title">{{ mod.name }}</div>
                                    <div class="mod-list-author">({{ mod.owner }})</div>
                                </div>

                                <div class="mod-list-description mb-1">{{ mod.latest_version.description }}</div>

                                <!--
                                    :icon="isFavouriteGame(game.identifier) ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                    :style="isFavouriteGame(game.identifier) ? { color: 'var(--primary-color)' } : {}"
                                    @click="toggleFavouriteGame(game.identifier)"
                                /> -->

                                <div class="mod-list-bottom-row"> 
                                    <div class="flex row gap-2">
                                        <Button
                                            class="btn w-full" icon="pi pi-trash" severity="danger"
                                            :label="$t('keywords.uninstall')"
                                            :disabled="selectedProfile == null"
                                            @click="uninstallTsMod(mod)"
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
    <!-- #endregion -->

    <!-- #region THUNDERSTORE TAB -->
    <DataView v-if="activeTab == ModListTabs.TS" 
        lazy layout="list" data-key="mod-list-ts"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink"
        :paginator="mods.length > PAGE_ROWS" :rows="PAGE_ROWS"
        :value="mods" @page="onPageChange" :first="pageFirstRecordIdx"
    >
        <template #empty>
            <div v-if="loading" layout="list">
                <h2 style="font-size: 40px; margin: 10px auto;">Loading mods...</h2>
            </div>
            <div v-else>
                <!-- TODO: If failed, make this show regardless of search input. --> 
                <div v-if="!hasSearchInput() /*|| failed*/">
                    <div class="ml-1">
                        <h1 class="error-text">No mods available!</h1>
                        <h2 style="margin: 5px 0px;">Either Thunderstore is down or something went terribly wrong :(</h2>

                        <Button class="mt-2" icon="pi pi-refresh" 
                            :label="$t('keywords.refresh')" 
                            @click="refreshMods(true)"
                            size="large"
                        />
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
            <div class="scrollable-list list-nogutter no-drag" @wheel.prevent="handleScroll">
                <div 
                    v-for="(mod, index) in currentPageMods" class="list-item col-12"
                    :key="index" :ref="el => modElements[index] = el"
                >
                    <div class="flex-grow-1 flex column sm:flex-row align-items-center pt-2 gap-3" :class="{ 'border-top-faint': index != 0 }">
                        <img class="mod-list-thumbnail block xl:block" :src="mod.latest_version?.icon || ''"/>
                        
                        <div class="flex-grow-1 flex column md:flex-row md:align-items-center">
                            <div class="flex-grow-1 flex column justify-content-between">
                                <div class="flex row align-items-baseline">
                                    <div class="mod-list-title">{{ mod.name }}</div>
                                    <div class="mod-list-author">({{ mod.owner }})</div>
                                </div>

                                <div class="mod-list-description mb-1">{{ mod.latest_version.description }}</div>

                                <!--
                                    :icon="isFavouriteGame(game.identifier) ? 'pi pi-heart-fill' : 'pi pi-heart'"
                                    :style="isFavouriteGame(game.identifier) ? { color: 'var(--primary-color)' } : {}"
                                    @click="toggleFavouriteGame(game.identifier)"
                                /> -->

                                <div class="mod-list-bottom-row"> 
                                    <div class="flex row gap-2">
                                        <Button v-if="profileHasMod(selectedProfile, mod.latest_version.full_name)"
                                            class="btn w-full" icon="pi pi-trash" severity="danger"
                                            :label="$t('keywords.uninstall')"
                                            :disabled="selectedProfile == null"
                                            @click="uninstallTsMod(mod)"
                                        />
                                        <Button v-else
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
    <!-- #endregion -->

    <!-- #region NEXUS TAB -->
    <DataView v-if="activeTab == ModListTabs.NEXUS" 
        lazy layout="list" data-key="mod-list-nexus"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink"
    >
        <template #empty>
            <div class="ml-1">
                <h2 class="error-text">Nexus Mods support is not implemented yet.</h2>
            </div>
        </template>
    </DataView>
    <!-- #endregion -->
</div>
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
    /*scrollbar-width: none;*/
    overflow-x: hidden;
    overflow-y: scroll;
    height: calc(100vh - 160px); /* This seems arbitrary and won't work on different resolutions? */
    /* width: calc(100vw - v-bind(sidebarOffsetPx) - 420px); TODO: Replace 420px. Get it dynamically from selected game card width. */
}

.empty-profile {
    color: rgba(235, 235, 235, 0.95);
    font-size: 25px;
    margin: 0;
}

.error-text {
    color: red;
    font-size: 28px;
    margin: 0 auto;
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