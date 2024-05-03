<script lang="ts" setup>
import { 
    WindowMinimise,
    WindowUnmaximise, WindowMaximise,
    WindowIsMaximised, Quit
} from '../../wailsjs/runtime/runtime'

import router from "../router"

const CustomMinimise = async() => {
    if (await WindowIsMaximised()) WindowUnmaximise()
    else WindowMinimise()
}

const ToDashboard = () => router.push('/')
const ToGameSelection = () => router.push('/game-selection')
</script>

<template>
<div class="topbar drag">
    <Toolbar>
        <template #start>
            <div class="flex align-items-center">
                <img class="icon no-select" src="../assets/images/appicon.png">

                <Button class="no-drag button" :label="$t('topbar.dashboard')" plain text @click="ToDashboard"/>
                <Button class="no-drag button" :label="$t('topbar.select-game')" plain text @click="ToGameSelection"/>
            </div>
        </template>

        <template #end>
            <div class="control-buttons no-drag">
                <ButtonGroup>
                    <Button plain class="button" icon="pi pi-minus" @click="CustomMinimise"/>
                    <Button plain class="button" icon="pi pi-expand" @click="WindowMaximise"/>
                    <Button plain class="button" icon="pi pi-times" @click="Quit"/>
                </ButtonGroup>
            </div>
        </template>
    </Toolbar>
</div>
</template>

<style scoped>
.icon {
    width: 27px;
    height: 27px;
    margin: 0px 5px 0px 5px;
}

.p-toolbar {
    border-top: transparent;
    border-left: transparent;
    border-right: transparent;
    border-bottom: solid rgba(167, 167, 167, 0.574) 1px;
    border-radius: 0;
    padding: 0.4rem;
    background-color: #18181bcd;
}

.p-toolbar .button {
    color: rgba(201, 201, 201, 0.905);
    padding: 8px 10px 8px 10px;
    font-size: 17px;
}

.control-buttons {
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
}

.control-buttons .button {
    cursor: pointer;
    background: transparent;
    border: none;
    color: rgba(201, 201, 201, 0.905);
    margin-right: 0;
}
</style>