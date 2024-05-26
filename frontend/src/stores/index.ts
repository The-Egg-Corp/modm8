import { defineStore } from 'pinia'

// The global application state.
export interface AppState {
    
}

// Defaults. These will usually be different at runtime.
const state: AppState = {

}

export const useGlobalStore = defineStore({
    id: 'AppStore',
    state: () => state
})

export * from './settings'
export * from './version'