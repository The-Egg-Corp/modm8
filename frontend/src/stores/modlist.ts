import { defineStore } from "pinia"
import { ref } from "vue"

import { 
    type Nullable,
    type ModListTab,
    ModListTabs
} from "../types/index.js"

export const useModListStore = defineStore('ModListStore', () => {
    //#region State
    const activeTab = ref<ModListTab>(ModListTabs.PROFILE)
    const searchInput = ref<Nullable<string>>(null)
    
    /** The {@link Element}s (that include mod info) shown on the current page.*/
    const modElements = ref<any[]>([])
    const scrollIndex = ref(0)
    //#endregion

    //#region Getters
    //#endregion

    //#region Actions
    //#endregion

    return {
        activeTab, searchInput,
        modElements, scrollIndex
    }
})