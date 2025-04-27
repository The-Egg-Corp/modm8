<script lang="ts" setup>
import { storeToRefs } from "pinia"
import { useAppStore } from "@stores"
import { PlatformSelectPanels } from "@components"

// enum Platform {
//     THUNDERSTORE = 1,
//     NEXUS
// }

// const selectedPlatform = ref(0)
// const setSelectedPlatform = (platform: Platform) => {
//     selectedPlatform.value = platform
//     console.log(`[ModDevTools] Selected platform set to '${Platform[platform]}'. Index: ${platform}`)
// }

// onUnmounted(() => {
//     selectedPlatform.value = 0
// })

const appStore = useAppStore()
const { sidebarWidth } = storeToRefs(appStore)
</script>

<template>
<div class="platform-selection">
    <PlatformSelectPanels
        @selectThunderstore="$router.push('/mod-dev-tools/thunderstore')" 
        @selectNexus="$router.push('/mod-dev-tools/nexus')"
    />
</div>
</template>

<style scoped>
.platform-selection {
    margin-left: v-bind(sidebarWidth);
}
</style>