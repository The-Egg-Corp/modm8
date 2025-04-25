import { defineStore } from "pinia"

//import type { AnyGame } from "@types"

export const useGameStore = defineStore('GameStore', () => {
    //#region State
    // const selectedGame = ref<AnyGame>({
    //     title: 'Placeholder',
    //     identifier: 'tf2',
    //     steamID: 0
    // })

    // const thunderstore = ref({
    //     ROWS: 40,
    //     allMods: null,
    //     currentPageMods: null,
    //     lastInstalledMod: null,
    // })

    // const nexus = ref({
    //     ROWS: 40,
    //     allMods: null,
    //     currentPageMods: null,
    //     lastInstalledMod: null,
    // })
    //#endregion

    //#region Getters

    //#endregion

    //#region Actions
    // TODO: Possibly need to take in ID, then get from `games` instead ?
    // function setSelectedGame(game: AnyGame) {
    //     selectedGame.value = _gameByID(game.identifier) ?? game
    // }
    //#endregion

    return {

    }
})