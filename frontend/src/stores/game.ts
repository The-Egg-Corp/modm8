import { Game } from '@types'
import { defineStore } from 'pinia'

// Stores state for anything game-related.
const state = {
    selectedGame: {
        identifier: ''
    }
}

const actions = {
    setSelectedGame: (game: Game) => {
        state.selectedGame = game
    }
}

export type GameState = typeof state

export const useGlobalStore = defineStore({
    id: 'GameStore',
    state: () => state,
    actions
})