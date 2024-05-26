import { defineStore } from 'pinia'

// Game-related state type.
export interface GameState {
    
}

// Defaults. These will usually be different at runtime.
const state: GameState = {
    
}

export const useGlobalStore = defineStore({
    id: 'AppStore',
    state: () => state
})