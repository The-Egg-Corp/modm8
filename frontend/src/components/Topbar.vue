<script lang="ts" setup>
import router from "../router"
import ControlButtons from "./ControlButtons.vue"

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
            <ControlButtons/>
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
</style>