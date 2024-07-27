<script lang="ts" setup>
import Breadcrumb from 'primevue/breadcrumb'

import { computed, ref } from "vue"
import type { ComputedRef, Ref } from "vue"

import { useGameStore } from "@stores"

const { selectedGame } = useGameStore()

const gameSelection: Ref<BreadcrumbPage> = ref({
    route: '/game-selection',
    icon: 'pi pi-home'
})

interface BreadcrumbPage {
    label?: string
    icon?: string
    route: string
}

const pages: ComputedRef<BreadcrumbPage[]> = computed(() => [{ 
    label: selectedGame.title,
    route: '/selected-game'
}])
</script>

<template>
<div class="selected-game row">
    <Breadcrumb class="breadcrumb flex-full row" :home="gameSelection" :model="pages"
        
    />
</div>
</template>

<style scoped>
/*.breadcrumb-container {
    width: 140%;
    height: 140%;
}*/

.breadcrumb {
    margin-top: 20px;
    padding: 5px;
    justify-content: center;
    height: auto;
    background: none;
}

:deep(.breadcrumb *) {
    font-size: 21px;
    user-select: none;
    --wails-draggable: none;
}
</style>