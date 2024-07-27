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
<div class="selected-game flex-full column">
    <Breadcrumb class="breadcrumb flex row" :home="home" :model="items"
            
    />
</div>
</template>

<style scoped>
/*.breadcrumb-container {
    width: 140%;
    height: 140%;
}*/

.breadcrumb {
    margin-top: 10px;
    align-items: flex-start;
    flex-grow: 2;
    justify-content: center;
    background: none;
    height: auto
}
</style>