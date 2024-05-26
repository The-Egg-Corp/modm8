import { Game } from '@types'
import { defineStore } from 'pinia'

// Game-related state type.
export interface GameState {
    selectedGame: Game
}

// Defaults. These will usually be different at runtime.
const state: GameState = {
    selectedGame: {
        identifier: ''
    }
}

export const useGlobalStore = defineStore({
    id: 'GameStore',
    state: () => state
})