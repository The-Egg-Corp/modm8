<script lang="ts" setup>
import router from "../router"

const ToDashboard = () => router.push('/')
const ToGameSelection = () => router.push('/game-selection')
const ToSettings = () => router.push('/settings')
</script>

<template>
<div class="sidebar flex-column">
    <div class="top flex-column">
        <img class="icon no-select" src="../assets/images/appicon.png">
        <Button class="plain-btn no-drag" icon="pi pi-spin pi-home" @click="ToDashboard"/>
        <Button class="plain-btn no-drag" icon="pi pi-spin pi-list" @click="ToGameSelection"/>
    </div>

    <div class="spacer"></div>

    <ButtonGroup class="bottom">
        <Button class="plain-btn no-drag" icon="pi pi-spin pi-cog" @click="ToSettings"/>
    </ButtonGroup>
</div>
</template>

<style scoped>
.sidebar {
    z-index: 999;
    height: 100vh;
    width: fit-content;
    border-right: white 1px outset;
    background-color: rgba(25, 25, 26, 0.845);
}

.icon {
    margin: 0px 20px 10px 20px;
    width: 42px;
    height: 42px;
}

.plain-btn {
    position: relative;
    margin: 0px 20px 0px 20px;
    color: rgba(218, 218, 218, 0.921);
    background-color: #18181bcd;
    border: solid rgba(139, 139, 139, 0.712) 1px;
    border-radius: 3.5px;
}

.top {
    padding-top: 20px;
    gap: 10px;
}

.spacer {
    flex: 1;
}

.bottom {
    padding-bottom: 20px;
}
</style>