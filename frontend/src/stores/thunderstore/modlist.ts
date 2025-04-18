import { defineStore, storeToRefs } from "pinia"
import { ref, computed } from "vue"

import { 
    type Nullable, 
    type ThunderstoreGame,
    type Package,
    type ModListTabType,
    ModListTabs
} from "@types"

import { useGameStore, useProfileStore } from '@stores'
import type { Dialog } from "@composables"

import { 
    GetStrippedPackages,
    GetLatestPackageVersion, 
    InstallByName
} from "@backend/thunderstore/API"

import type { thunderstore, v1 } from "@backend/models"

export const useModListStoreTS = defineStore('ModListStoreTS', () => {
    //#region Stores
    const profileStore = useProfileStore()
    const { selectedProfile } = storeToRefs(profileStore)

    const gameStore = useGameStore()
    const { updateModCache } = gameStore.thunderstore
    //#endregion

    //#region State
    const loading = ref(false)
    const installing = ref(false)

    const activeTab = ref<ModListTabType>(ModListTabs.PROFILE)
    const searchInput = ref<Nullable<string>>(null)

    const modElements = ref<any[]>([])
    const scrollIndex = ref(0)
    
    const ROWS = 40
    const first = ref(0) // Starting index of the current page
    const currentPageMods = ref<Package[]>([])

    const mods = ref<thunderstore.StrippedPackage[]>([])
    const lastInstalledMod = ref<Nullable<v1.PackageVersion>>(null)
    //#endregion

    //#region Getters
    const selectedGame = computed(() => gameStore.thunderstore.selectedGame) // TODO: Computed may not be needed?
    //const thunderstoreMods = computed(() => activeTab.value == ModListTabType.TS ? filterByProfile(allMods.value) : allMods.value)
    //#endregion

    //#region Actions
    async function installMod(fullName: string, selectedGame: ThunderstoreGame, installingDialog?: Dialog) {
        installing.value = true
    
        if (installingDialog) {
            installingDialog.setClosable(false)
            installingDialog.setVisible(true)
        }
    
        try {
            const start = performance.now()
            lastInstalledMod.value = await InstallByName(selectedGame.title, selectedGame.identifier, fullName)
            console.info(`Installed mod: ${fullName}. Took ${((performance.now() - start) / 1000).toFixed(2)}s`)
        } catch(e: any) {
            console.error(`[${selectedGame.identifier}] Failed to install mod.\n${e.message}`)
        }
    
        installing.value = false
        
        if (installingDialog) {
            installingDialog.setClosable(true)
        }
    }

    /** 
     * Refreshes the page with the currently cached mods.
     * @param fetch Whether to fetch mods from the API and update the cache (if isnt populated already) before refreshing the page.
     */
    async function refreshMods(fetch: boolean) {
        if (fetch && !selectedGame.value.modCache) {
            loading.value = true // Fetching might take a while, let the user know we are doing something.

            try {
                const t0 = performance.now()

                const pkgs = await GetStrippedPackages(selectedGame.value.identifier, false)
                console.log(`[${selectedGame.value.identifier}] Putting mods into cache...`)

                updateModCache(pkgs)
                console.log(`Cached ${pkgs.length} mods. Took: ${performance.now() - t0}ms`)
            } catch(e: any) {
                console.error(`[${selectedGame.value.identifier}] Failed to update mod cache!\n${e.message}`)
            }
        }

        mods.value = getMods()
        await refreshPage()

        // We are done regardless of outcome, stop loading.
        loading.value = false
    }
    
    async function latestModVersion(mod: thunderstore.StrippedPackage) {
        return await GetLatestPackageVersion(selectedGame.value.identifier, mod.owner, mod.name)
    }

    function getMods(searchFilter = true, defaultSort = true) {
        if (!selectedGame.value.modCache) return []
        const filteredMods = searchFilter ? filterBySearch(selectedGame.value.modCache) : selectedGame.value.modCache

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
        return mods.filter(mod => selectedProfile.value?.mods.thunderstore?.includes(mod.full_name))
    }

    const refreshPage = () => updatePage(0, ROWS)
    const updatePage = async (newFirst: number, rows: number) => {
        first.value = newFirst

        const filtered = mods.value.slice(newFirst, newFirst + rows) as Package[]

        // TODO: This could be potentially VERY slow. Consider replacing with Map/Set instead.
        const promises = filtered.map(async mod => {
            mod.latestVersion = await latestModVersion(mod)
            return mod
        })

        currentPageMods.value = await Promise.all(promises)
    }
    //#endregion

    return {
        loading,
        installing,
        activeTab,
        searchInput,
        modElements,
        scrollIndex,
        ROWS,
        mods,
        currentPageMods,
        lastInstalledMod,
        first,
        installMod,
        getMods,
        refreshMods,
        latestModVersion,
        filterByProfile,
        filterBySearch,
        refreshPage,
        updatePage
    }
})