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

        <Button class="plain-btn margin-lr no-drag" icon="pi pi-spin pi-home" @click="ToDashboard"/>
        <Button class="plain-btn margin-lr no-drag" icon="pi pi-spin pi-list" @click="ToGameSelection"/>
    </div>

    <div class="spacer"></div>

    <div class="bottom">
        <Button class="plain-btn margin-lr no-drag" icon="pi pi-spin pi-cog" @click="ToSettings"/>
    </div>
</div>
</template>

<style scoped>
.sidebar {
    z-index: 999;
    height: 100vh;
    width: fit-content;
    border-right: rgba(231, 231, 231, 0.885) 1px outset;
    background-color: rgba(25, 25, 26, 0.845);
}

.sidebar img {
    margin: 15px 15px 15px 15px;
}

.margin-lr {
    margin: 0px 15px 0px 15px;
}

.icon {
    width: 40px;
    height: 41.5px;
}

.top .p-button {
    margin-bottom: 10px;
}

.plain-btn {
    position: relative;
    color: rgba(218, 218, 218, 0.921);
    background-color: #18181bcd;
    border: solid rgba(139, 139, 139, 0.712) 1px;
    border-radius: 3.5px;
}

.spacer {
    flex: 1;
}

.bottom {
    padding-bottom: 20px;
}
</style>