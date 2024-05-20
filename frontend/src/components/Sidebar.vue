<script lang="ts" setup>
import router from "../router"
import { useDialog } from '@composables'

const { setVisible } = useDialog()

const Dashboard = () => router.push('/')
const GameSelection = () => router.push('/game-selection')
//const Settings = () => router.push('/settings')

const tooltipOpts = (text: string) => ({ 
    value: text, 
    showDelay: 60, 
    hideDelay: 5
})
</script>

<template>
<div class="sidebar flex-column">
    <div class="top flex-column">
        <img class="icon no-select" src="../assets/images/appicon.png">

        <Button 
            plain outlined 
            class="btn margin-lr no-drag" icon="pi pi-spin pi-home" 
            v-tooltip="tooltipOpts('Dashboard')"
            @click="Dashboard"
        />

        <Button 
            plain outlined 
            class="btn margin-lr no-drag" icon="pi pi-spin pi-list" 
            v-tooltip="tooltipOpts('Game Selection')"
            @click="GameSelection"
        />
    </div>

    <!-- <div class="spacer"></div>

    <div class="top flex-column">
        <Button plain outlined class="btn margin-lr no-drag" icon="pi pi-spin pi-wrench"/>
        <Button plain outlined class="btn margin-lr no-drag" icon="pi pi-spin pi-upload"/>
    </div> -->

    <div class="spacer"></div>

    <div class="bottom flex-column">
        <!-- <Button text class="btn margin-lr no-drag" :icon="themeMode === 'dark' ? 'pi pi-sun' : 'pi pi-moon'" @click="ToggleThemeMode"/> -->
        <Button 
            plain outlined 
            class="btn margin-lr no-drag" icon="pi pi-spin pi-cog" 
            v-tooltip="tooltipOpts('Settings')"
            @click="setVisible(true)"/>
    </div>
</div>
</template>

<style scoped>
.sidebar {
    position: relative;
    z-index: 999;
    height: 100vh;
    width: fit-content;
    border-right: rgba(231, 231, 231, 0.885) 1px outset;
    background-color: rgb(23, 23, 26);
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

.bottom .p-button {
    margin-top: 10px;
}

.btn {
    position: relative;
    border-radius: 3px;
}

.plain-btn {
    color: rgba(218, 218, 218, 0.921);
    background-color: #18181bcd;
    border: solid rgba(139, 139, 139, 0.712) 1px;
    border-radius: 3px;
}

.spacer {
    flex: 1;
}

.bottom {
    padding-bottom: 20px;
}
</style>