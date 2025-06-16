import { defineStore, storeToRefs } from "pinia"
import { ref } from "vue"

import { 
    type Nullable, 
    type ThunderstoreGame
} from "@types"

import { 
    useGameStore,
    useGameStoreTS,
    useModListStore,
    useProfileStore
} from '@stores'

import type { Dialog } from "@composables"

import { 
    GetStrippedPackages,
    InstallByName
} from "@backend/thunderstore/API"

import type { thunderstore, v1 } from "@backend/models"

export const useModListStoreTS = defineStore('ModListStoreTS', () => {
    //#region Stores
    const modListStore = useModListStore()
    const { searchInput } = storeToRefs(modListStore)

    const profileStore = useProfileStore()
    const { selectedProfile } = storeToRefs(profileStore)

    const gameStore = useGameStore()
    const { selectedGame } = storeToRefs(gameStore)

    const gameStoreTS = useGameStoreTS()
    const { updateModCache } = gameStoreTS
    //#endregion

    //#region State
    const loading = ref(false)
    const installing = ref(false)
    
    const PAGE_ROWS = 50
    const pageFirstRecordIdx = ref(0) // Index of the first record on the current page.
    const currentPageMods = ref<thunderstore.StrippedPackage[]>([])

    const mods = ref<thunderstore.StrippedPackage[]>([])
    const lastInstalledMod = ref<Nullable<v1.PackageVersion>>(null)
    //#endregion

    //#region Getters
    //const thunderstoreMods = computed(() => activeTab.value == ModListTabType.TS ? filterByProfile(allMods.value) : allMods.value)
    //#endregion

    //#region Actions
    /**
     * For the given selected game, this function will attempt to install the mod specified by {@link fullName},
     * with all of its dependencies, if they do not already exist.
     * @param fullName The name of the mod in the format: "Author-ModName". Version must be excluded.
     * @param selectedGame The instance of the currently selected Thunderstore game.
     * @param installingDialog The dialog UI to show during mod installation.
     */
    async function installMod(fullName: string, selectedGame: ThunderstoreGame, installingDialog?: Dialog) {
        installing.value = true
    
        if (installingDialog) {
            installingDialog.setClosable(false)
            installingDialog.setVisible(true)
        }
    
        let success = false
        try {
            const start = performance.now()

            lastInstalledMod.value = await InstallByName(selectedGame.title, selectedGame.identifier, fullName)
            success = !!lastInstalledMod.value

            console.info(`Installed mod: ${fullName}. Took ${((performance.now() - start) / 1000).toFixed(2)}s`)
        } catch(e: any) {
            console.error(`[${selectedGame.identifier}] Failed to install mod.\n${e.message}`)
        }
    
        installing.value = false
        
        if (installingDialog) {
            installingDialog.setClosable(true)
        }

        return success
    }

    /** 
     * Refreshes the page with the currently cached mods.
     * @param fetch Whether to fetch mods from the API and update the cache (if isnt populated already) before refreshing the page.
     */
    async function refreshMods(fetchIfEmpty: boolean) {
        if (selectedGame.value.platform != 'THUNDERSTORE') {
            throw new Error('[TS/modlist] Could not refresh mods. Selected game is not of type `THUNDERSTORE`.')
        }

        // The actual instance from the game container.
        const selectedGameVal: ThunderstoreGame = selectedGame.value.value
        if (!selectedGameVal.modCache && fetchIfEmpty) {
            loading.value = true // Fetching might take a while, let the user know we are doing something.

            try {
                const t0 = performance.now()

                const pkgs = await GetStrippedPackages(selectedGameVal.identifier, false)
                console.log(`[${selectedGameVal.identifier}] Putting mods into cache...`)

                updateModCache(pkgs)
                console.log(`Cached ${pkgs.length} mods. Took: ${performance.now() - t0}ms`)
            } catch(e: any) {
                console.error(`[${selectedGameVal.identifier}] Failed to update mod cache!\n${e.message}`)
            }
        }

        mods.value = getMods()
        await refreshPage()

        // We are done regardless of outcome, stop loading.
        loading.value = false
    }

    function getMods(searchFilter = true, defaultSort = true) {
        const cache = selectedGame.value.value.modCache
        if (!cache) return []

        const filteredMods = searchFilter ? filterBySearch(cache) : cache
        return !defaultSort ? filteredMods : filteredMods.sort((m1, m2) => m2.rating_score - m1.rating_score)
    }

    function filterBySearch(mods: thunderstore.StrippedPackage[]) {
        if (!searchInput.value) return mods
    
        const input = searchInput.value.trim()
        if (input == "") return mods
    
        return mods.filter(mod => {
            const lowerTitle = mod.name?.toLowerCase() ?? ""
    
            // Necessary to not show irrelevent mods with only 1 letter input.
            if (input.length == 1 && !lowerTitle.startsWith(input.toLowerCase())) {
                return false
            }
    
            return lowerTitle.includes(input.toLowerCase())
        })
    }

    function filterByProfile(mods: thunderstore.StrippedPackage[]) {
        if (!selectedProfile.value?.name) return mods // Invalid or no selected profile.

        const tsProfMods = selectedProfile.value.mods.thunderstore
        const filtered = mods.filter(mod => {
            const latestVerName = mod.latest_version.full_name.toLowerCase()
            return tsProfMods?.some(name => name.toLowerCase() == latestVerName)
        })

        return filtered
    }

    /**
     * Updates the current page to show mods between `firstIdx` and `firstIdx` + `rows`.
     * @param firstIdx Index of the first record on the page.
     * @param rows The number of rows/mods to display on the page.
     */
    const updatePage = async (firstIdx: number, rows: number) => {
        if (firstIdx < 0) {
            console.warn(`[ModListStoreTS] Cannot update page. Param 'firstIdx' is: ${firstIdx}. Must be >= 0.`)
            return
        }

        pageFirstRecordIdx.value = firstIdx
        currentPageMods.value = mods.value.slice(firstIdx, firstIdx + rows) // If firstIdx=0 & rows=40, mods[0] - mods[39] will be shown.
    }

    /**
     * Updates the current page of mods to display using default parameters.\
     * Shorthand for calling {@link updatePage}(0, PAGE_ROWS). 
     */
    const refreshPage = () => updatePage(0, PAGE_ROWS)
    //#endregion

    return {
        loading,
        installing,
        PAGE_ROWS,
        mods,
        currentPageMods,
        lastInstalledMod,
        pageFirstRecordIdx,
        installMod,
        getMods,
        refreshMods,
        filterBySearch,
        filterByProfile,
        refreshPage,
        updatePage
    }
})