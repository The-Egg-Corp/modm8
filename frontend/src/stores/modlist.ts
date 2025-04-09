import { defineStore } from "pinia"

import { useModListStoreTS } from "./thunderstore/modlist"

export const useModListStore = defineStore('ModListStore', () => {
    //#region State
    const modListStoreTs = useModListStoreTS()
    //const modListStoreNexus = useModListStoreNexus()
    //#endregion

    //#region Getters

    //#endregion

    //#region Actions

    //#endregion

    return {
        thunderstore: modListStoreTs
        //nexus: modListStoreNexus
    }
})