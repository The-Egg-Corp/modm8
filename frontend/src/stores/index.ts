import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export interface AppState {
    maxThreads: number
}

export const useAppStore = defineStore('AppStore', () => {
    const maxThreads = ref(2)

    const expanded = ref(false)
    const sidebarWidth = computed(() => (expanded.value ? '180px' : '75px'))
  
    function toggleSidebar() {
        expanded.value = !expanded.value
    }

    /**
     * Updates value of `maxThreads` ref to make frontend aware of max logical CPUs.
     */
    async function setMaxThreads(count: number) {
        maxThreads.value = count
    }

    return { 
        maxThreads,
        expanded, 
        sidebarWidth, 
        toggleSidebar,
        setMaxThreads
    }
})

export * from './settings'
export * from './version'
export * from './game'
export * from './profile'