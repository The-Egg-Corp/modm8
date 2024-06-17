import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface AppState {
    maxThreads: number
}

export const useAppStore = defineStore('AppStore', () => {
    const maxThreads = ref(2)

    function setMaxThreads(num: number) {
        maxThreads.value = num
    }

    return {
        maxThreads,
        setMaxThreads
    }
})

export * from './settings'
export * from './version'
export * from './game'
export * from './profile'