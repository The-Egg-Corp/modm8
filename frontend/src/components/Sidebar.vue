<script lang="ts" setup>
import router from "../router"
import { useDialog } from '@composables'

import { tooltipOpts } from "../../src/util"
import { t } from '@i18n'

const settings = useDialog('settings')
const appInfo = useDialog('app-info')

const Dashboard = () => router.push('/')
const GameSelection = () => router.push('/game-selection')
const ModDevTools = () => router.push('/mod-dev-tools')
</script>

<template>
<div class="sidebar flex column">
    <img class="icon no-select" src="../assets/images/appicon.png">

    <div class="flex column h-full">
        <div class="top flex column">
            <Button 
                plain outlined 
                class="btn margin-lr no-drag" icon="pi pi-spin pi-home" 
                v-tooltip="tooltipOpts(t('tooltips.sidebar.dashboard'))"
                @click="Dashboard"
            />
    
            <Button 
                plain outlined 
                class="btn margin-lr no-drag" icon="pi pi-spin pi-list" 
                v-tooltip="tooltipOpts(t('tooltips.sidebar.game-selection'))"
                @click="GameSelection"
            >
                <template #icon>
                    <span class="material-symbols-sharp">stadia_controller</span>
                </template>
            </Button>
        </div>
    
        <div class="spacer"></div>

        <div class="top flex column">
            <Button 
                plain outlined 
                class="btn margin-lr no-drag" icon="pi pi-spin pi-wrench"
                v-tooltip="tooltipOpts(t('tooltips.sidebar.mod-dev-tools'))"
                @click="ModDevTools"
            />

            <!-- <Button plain outlined class="btn margin-lr no-drag" icon="pi pi-spin pi-upload"/> -->
        </div>

        <div class="spacer"></div>
 
        <div class="bottom flex column">
            <Button plain outlined
                class="btn margin-lr no-drag"
                style="border: none;"
                icon="pi pi-info-circle"
                @click="appInfo.setVisible()"
            />

            <!-- <Button text class="btn margin-lr no-drag" :icon="themeMode === 'dark' ? 'pi pi-sun' : 'pi pi-moon'" @click="ToggleThemeMode"/> -->
            <Button 
                plain outlined 
                class="btn margin-lr no-drag" icon="pi pi-spin pi-cog" 
                v-tooltip="tooltipOpts(t('keywords.settings'))"
                @click="settings.setVisible()"
            />
        </div>
    </div>
</div>
</template>

<style scoped>
.sidebar {
    position: relative;
    z-index: 9998;
    height: 100vh;
    width: fit-content;
    border-right: rgba(231, 231, 231, 0.885) 1px outset;
    background-color: rgb(28, 29, 33);
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

/*.plain-btn {
    color: rgba(218, 218, 218, 0.921);
    background-color: #18181bcd;
    border: solid rgba(139, 139, 139, 0.712) 1px;
    border-radius: 3px;
}*/

.spacer {
    flex: 1;
}

.bottom {
    padding-bottom: 20px;
}

.material-symbols-sharp {
    font-variation-settings:
    'FILL' 0,
    'wght' 270,
    'GRAD' 100,
    'opsz' 40
}
</style>