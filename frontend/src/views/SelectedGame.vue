<script lang="ts" setup>
import Breadcrumb from 'primevue/breadcrumb'

import { computed, ref } from "vue"
import { useGameStore } from "@stores"

const { selectedGame } = useGameStore()

const home = ref({
    icon: 'pi pi-home',
    route: '/game-selection'
})

const items = computed(() => [{ 
    label: selectedGame.title,
    route: '/selected-game'
}])
</script>

<template>
<div class="selected-game row">
    <Breadcrumb class="breadcrumb flex-full row" :home="home" :model="items"
            
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