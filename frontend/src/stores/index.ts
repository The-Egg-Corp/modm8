import { defineStore } from 'pinia'

// Defaults. These will usually be different at runtime.
const state = {
    maxThreads: 2
}

const actions = {
    setMaxThreads(num: number) {
        state.maxThreads = num
    }
}

export type AppState = typeof state

export const useAppStore = defineStore({
    id: 'AppStore',
    state: () => state,
    actions
})

export * from './settings'
export * from './version'
export * from './game'
export * from './profile'