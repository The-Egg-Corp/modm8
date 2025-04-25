import { defineStore } from "pinia"
import { ref } from "vue"

import { 
    type Nullable,
    type ModListTabType,
    ModListTabs
} from "../types/index.js"

export const useModListStore = defineStore('ModListStore', () => {
    //#region State
    const activeTab = ref<ModListTabType>(ModListTabs.PROFILE)
    const searchInput = ref<Nullable<string>>(null)
    //#endregion

    //#region Getters
    //#endregion

    //#region Actions
    //#endregion

    return {
        activeTab, searchInput 
    }
})