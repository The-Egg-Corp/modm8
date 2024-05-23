import { defineStore } from "pinia"

export interface InternalState {
    appVersion: string
}

const state: InternalState = {
    appVersion: ''
}

export const useInternalStore = defineStore({
    id: 'InternalStore',
    state: () => state
})