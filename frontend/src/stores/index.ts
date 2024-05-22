import { defineStore } from 'pinia'

// import ProfileModule from './modules/profile.js'
// import VersionModule from './modules/version.js'

// Defaults. These will usually be different at runtime.
const state = {
    selectedGame: null
}

// The root state of all modules.
export type AppState = typeof state

export const useGlobalStore = defineStore({
    id: 'AppStore',
    state: () => state
})

export * from './settings.js'