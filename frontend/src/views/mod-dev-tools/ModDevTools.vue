<script lang="ts" setup>
import { PlatformSelectPanels } from "@components"
import { Nullable } from "primevue/ts-helpers"
import { ref } from "vue"

enum Platform {
    THUNDERSTORE,
    NEXUS
}

const selectedPlatform = ref<Nullable<Platform>>(null)
</script>

<template>
    <PlatformSelectPanels v-if="!selectedPlatform" 
        @selectThunderstore="selectedPlatform = Platform.THUNDERSTORE" 
        @selectNexus="selectedPlatform = Platform.NEXUS"
    />
</template>

<style scoped>

</style>