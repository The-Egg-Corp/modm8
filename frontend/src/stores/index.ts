import { defineStore } from 'pinia'

type Game = null | {}

// The global application state.
export interface AppState {
    selectedGame: Game
}

// Defaults. These will usually be different at runtime.
const state: AppState = {
    selectedGame: null
}

export const useGlobalStore = defineStore({
    id: 'AppStore',
    state: () => state
})

export * from './settings.js'