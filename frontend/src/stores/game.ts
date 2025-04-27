import { defineStore } from "pinia"

// import type { GameContainer, ThunderstoreGame } from "@types"

// const defaultGame: GameContainer = { 
//     type: 'THUNDERSTORE',
//     value: {
//         title: 'Placeholder',
//         identifier: 'tf2',
//         steamID: 0
//     } satisfies ThunderstoreGame
// }

export const useGameStore = defineStore('GameStore', () => {
    //#region State
    //const selectedGame = ref<GameContainer>(defaultGame)

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