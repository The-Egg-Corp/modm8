<script lang="ts" setup>
import Breadcrumb from 'primevue/breadcrumb'
import Dialog from 'primevue/dialog'

// import Splitter from 'primevue/splitter'
// import SplitterPanel from 'primevue/splitterpanel'

import Card from 'primevue/card'
import Button from 'primevue/button'

import { computed, ref } from "vue"
import type { ComputedRef, Ref } from "vue"

import { LaunchSteamGame } from '@backend/steam/SteamRunner'
import { useGameStore } from "@stores"
import { useDialog } from '@composables'

const { selectedGame } = useGameStore()

interface BreadcrumbPage {
    route: string
    home?: boolean
    label?: string
    icon?: string
    class?: string
    style?: string
}

const homePage: Ref<BreadcrumbPage> = ref({
    home: true,
    route: '/game-selection',
    icon: "stadia_controller",
    class: "material-symbols-sharp text-color-secondary"
})

const pages: ComputedRef<BreadcrumbPage[]> = computed(() => [{ 
    label: selectedGame.title,
    route: '/selected-game',
    class: "text-primary"
}])

const getThumbnail = () => selectedGame.image
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${selectedGame.image}` 
    : "https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/Titanfall2.jpg"

const startVanilla = () => LaunchSteamGame(selectedGame.id, ["--doorstop-enable", "false"])
const startModded = () => LaunchSteamGame(selectedGame.id, [])

const { 
    setVisible,
    visible, closable, draggable 
} = useDialog('selected-game')

const dialogStyle = {
    "margin-left": "70px",
    "width": 'auto',
    "min-width": '45rem'
}
</script>

<template>
<div class="selected-game row">
    <Breadcrumb class="breadcrumb flex-full row" :home="homePage" :model="pages">
        <template #item="{ item, props }">
            <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                <a :href="href" v-bind="props.action" @click="navigate">
                    <span v-if="item.icon" :class="item.class" :style="item.style">{{ item.icon }}</span>
                    <span v-else-if="item.label" :class="item.class" :style="item.style">{{ item.label }}</span>
                </a>
            </router-link>
        </template>
        <template #separator>➥</template>
    </Breadcrumb>

    <Card class="current-game-card">
        <template #title>
            <p style="font-size: 30px; font-weight: 520; user-select: none;" class="no-drag mt-0 mb-1">Currently Selected</p>
        </template>

        <template #content>
            <div class="no-drag ml-1">
                <img class="current-game-thumbnail" :src="getThumbnail()"/>
                <div class="flex column" style="float: right">
                    <p style="font-size: 25px; font-weight: 330" class="mt-0 mb-0 ml-3">{{ selectedGame.title }}</p>
                    <div class="flex column gap-2 mt-3">
                        <Button 
                            plain
                            class="btn ml-3" 
                            icon="pi pi-caret-right"
                            label="Start Modded"
                            @click="startModded"
                        />

                        <Button 
                            plain outlined
                            class="btn ml-3" 
                            icon="pi pi-caret-right"
                            label="Start Vanilla"
                            @click="startVanilla"
                        />

                        <Button 
                            plain outlined
                            class="btn ml-3 mt-4" 
                            icon="pi pi-file-edit"
                            label="Edit Config"
                            @click="setVisible(true)"
                        />
                    </div>
                </div>
            </div>
        </template>
    </Card>

    <DataView data-key="profile-selection">

    </DataView>

    <Dialog 
        modal class="no-drag"
        :style="dialogStyle"
        :block-scroll="true"
        :dismissable-mask="true"
        :show-header="false"
        v-model:visible="visible"
        v-model:draggable="draggable" 
        v-model:closable="closable"
    >
        <h2>Configuration File Editor</h2>
    </Dialog>
    
    <!-- <Splitter style="height: 350px; background: none; border: none;" class="mb-9 no-drag mx-auto">
        <SplitterPanel class="flex" :minSize="33">

        </SplitterPanel>

        <SplitterPanel class="flex align-items-center"> 
            <DataView data-key="profile-selection">

            </DataView>
        </SplitterPanel>
    </Splitter> -->
</div>
</template>

<style scoped>
/*.breadcrumb-container {
    width: 140%;
    height: 140%;
}*/

/*.selected-game {
    
}*/

.breadcrumb {
    margin-top: 20px;
    padding: 5px;
    justify-content: center;
    height: auto;
    background: none;
}

:deep(.breadcrumb *) {
    font-size: 19px;
    user-select: none;
    --wails-draggable: none;
}

.material-symbols-sharp {
    font-size: 26.5px;
    font-variation-settings:
    'FILL' 0,
    'wght' 270,
    'GRAD' 100,
    'opsz' 40
}

.current-game-card {
    background: none;
    width: fit-content;
}

:deep(.current-game-card .p-card-body) {
    padding: 15px 0px 0px 40px;
}

.current-game-thumbnail {
    user-select: none;
    width: 162.5px;
    border-radius: 4px;
    border: 2px outset var(--primary-color);
}

.btn {
    position: relative;
    border-radius: 5px;
    text-align: left;
}
</style>