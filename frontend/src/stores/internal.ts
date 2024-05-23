import { defineStore } from "pinia"

export interface InternalState {
    appVersion: string
}

const state: InternalState = {
    appVersion: '1.0'
}

export const useInternalStore = defineStore({
    id: 'InternalStore',
    state: () => state
})