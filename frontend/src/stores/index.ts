import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

// export interface AppState {
//     maxThreads: number
//     sidebarExpanded: boolean
//     sidebarWidth: string
// }

export const useAppStore = defineStore('AppStore', () => {
    const maxThreads = ref(2)

    const sidebarExpanded = ref(false)
    const sidebarWidth = computed(() => sidebarExpanded.value ? '180px' : '75px')
  
    function toggleSidebar() {
        sidebarExpanded.value = !sidebarExpanded.value
    }

    /**
     * Updates value of `maxThreads` ref to make frontend aware of max logical CPUs.
     * @count Number of threads to set as the limit.
     */
    async function setMaxThreads(count: number) {
        maxThreads.value = count
    }

    return { 
        maxThreads,
        sidebarExpanded,
        sidebarWidth,
        toggleSidebar,
        setMaxThreads
    }
})

export * from './settings'
export * from './version'
export * from './game'
export * from './profile'