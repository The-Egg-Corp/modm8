import { defineStore } from "pinia"
import { ref } from "vue"

import type { GameContainer, ThunderstoreGame } from "@types"

const defaultGame: GameContainer = { 
    platform: 'THUNDERSTORE',
    value: {
        title: 'Placeholder',
        identifier: 'tf2',
        steamID: 0
    } satisfies ThunderstoreGame
}

type GameType<T> = Extract<GameContainer, { platform: T }>['value']

export const useGameStore = defineStore('GameStore', () => {
    //#region State
    const selectedGame = ref<GameContainer>(defaultGame)
    //#endregion

    //#region Getters

    //#endregion

    //#region Actions
    /** 
     * Sets the currently selected game. Usually called before navigating to the Selected Game page.\
     * If the function complains, ensure the specified `game` type is allowed by the specified `platform`.
     * @see {@link GameContainer} for platform <-> game type definitions.
     */
    function setSelectedGame<T extends GameContainer['platform']>(platform: T, game: GameType<T>) {
        // We could just allow GameContainer, but the error msg and param typing is more useful this way.
        selectedGame.value = { platform, value: game } as GameContainer
    }
    //#endregion

    return {
        selectedGame,
        setSelectedGame
    }
})