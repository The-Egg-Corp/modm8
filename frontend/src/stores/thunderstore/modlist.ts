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
    
    const PAGE_ROWS = 40
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
    async function refreshMods(fetch: boolean) {
        if (selectedGame.value.type != 'THUNDERSTORE') {
            throw new Error('[TS/modlist] Could not refresh mods. Selected game is not of type `THUNDERSTORE`.')
        }

        if (fetch && !selectedGame.value.value.modCache) {
            loading.value = true // Fetching might take a while, let the user know we are doing something.

            try {
                const t0 = performance.now()

                const pkgs = await GetStrippedPackages(selectedGame.value.value.identifier, false)
                console.log(`[${selectedGame.value.value.identifier}] Putting mods into cache...`)

                updateModCache(pkgs)
                console.log(`Cached ${pkgs.length} mods. Took: ${performance.now() - t0}ms`)
            } catch(e: any) {
                console.error(`[${selectedGame.value.value.identifier}] Failed to update mod cache!\n${e.message}`)
            }
        }

        mods.value = getMods()
        await refreshPage()

        // We are done regardless of outcome, stop loading.
        loading.value = false
    }

    function getMods(searchFilter = true, defaultSort = true) {
        if (!selectedGame.value.value.modCache) return []
        const filteredMods = searchFilter ? filterBySearch(selectedGame.value.value.modCache) : selectedGame.value.value.modCache

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

    async function filterByProfile(mods: thunderstore.StrippedPackage[]) {
        if (!selectedProfile.value?.name) return mods // Invalid or no selected profile.

        const tsProfMods = selectedProfile.value.mods.thunderstore
        const filtered = await mods.filter(mod => {
            const latestVerName = mod.latest_version.full_name.toLowerCase()
            return tsProfMods?.some(name => name.toLowerCase() == latestVerName)
        })

        return filtered
    }

    const refreshPage = () => updatePage(0, PAGE_ROWS)
    const updatePage = async (newFirst: number, rows: number) => {
        pageFirstRecordIdx.value = newFirst
        currentPageMods.value = mods.value.slice(newFirst, newFirst + rows)
    }
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