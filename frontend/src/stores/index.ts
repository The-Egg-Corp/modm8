import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

// export interface AppState {
//     maxThreads: number
//     sidebarExpanded: boolean
//     sidebarWidth: string
// }

// TODO: 
// Instead of hardcoding height/width refs, they should all be
// computed and point to their respective element in the DOM.
export const useAppStore = defineStore('AppStore', () => {
    //#region State
    const maxThreads = ref(2)

    const topbarMargin = ref(0)

    const sidebarMargin = ref(25)
    const sidebarExpanded = ref(false)

    const showBackButton = ref(false)
    //#endregion

    //#region Getters
    const topbarHeightPx = computed(() => {
        const el = document.getElementsByClassName("topbar")
        return el.item(0)?.clientHeight + 'px';
    })

    const collapsibleContentPx = computed(() => {
        const el = document.getElementsByClassName("collapsible-content")
        return el.item(0)?.clientHeight + 'px';
    })

    const topbarMarginPx = computed(() => `${topbarMargin.value}px`)
    const sidebarMarginPx = computed(() => `${sidebarMargin.value}px`)

    const sidebarWidth = computed(() => sidebarExpanded.value ? 180 : 80)
    const sidebarWidthPx = computed(() => `${sidebarWidth.value}px`)

    const sidebarOffset = computed(() => sidebarWidth.value + sidebarMargin.value)
    const sidebarOffsetPx = computed(() => `${sidebarOffset.value}px`)
    //#endregion

    //#region Actions
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
    //#endregion

    return {
        maxThreads,
        sidebarExpanded,
        showBackButton, 
        topbarHeightPx, 
        topbarMargin, topbarMarginPx,
        sidebarMargin, sidebarMarginPx,
        sidebarWidth, sidebarWidthPx,
        sidebarOffset, sidebarOffsetPx,
        collapsibleContentPx,
        toggleSidebar,
        setMaxThreads
    }
})

export * from './settings'
export * from './version'
export * from './profile'

export * from './game'
export * from './modlist'