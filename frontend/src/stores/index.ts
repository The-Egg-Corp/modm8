import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

// export interface AppState {
//     maxThreads: number
//     sidebarExpanded: boolean
//     sidebarWidth: string
// }

// TODO: Instead of hardcoding height/width refs, they should all be
//       computed and point to their respective element in the DOM.
export const useAppStore = defineStore('AppStore', () => {
    const maxThreads = ref(2)

    const sidebarExpanded = ref(false)

    const topbarHeight = ref(30)
    const topbarHeightPx = computed(() => `${topbarHeight.value}px`)

    const sidebarMargin = ref(20)
    const sidebarMarginPx = computed(() => `${sidebarMargin.value}px`)

    const sidebarWidth = computed(() => sidebarExpanded.value ? 170 : 75)
    const sidebarWidthPx = computed(() => `${sidebarWidth.value}px`)

    const sidebarOffset = computed(() => sidebarWidth.value + sidebarMargin.value)
    const sidebarOffsetPx = computed(() => `${sidebarOffset.value}px`)

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
        topbarHeight, topbarHeightPx,
        sidebarMargin, sidebarMarginPx,
        sidebarWidth, sidebarWidthPx,
        sidebarOffset, sidebarOffsetPx,
        toggleSidebar,
        setMaxThreads
    }
})

export * from './settings'
export * from './version'
export * from './game'
export * from './profile'