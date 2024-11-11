<script lang="ts" setup>
import router from "../../router"

import { tooltipOpts } from "../../util"

import { t } from '@i18n'
import { useDialog } from '@composables'
import { useAppStore } from "@stores"
import { storeToRefs } from "pinia"

const settings = useDialog('settings')
const appInfo = useDialog('app-info')

const appStore = useAppStore()
const { 
    sidebarExpanded, 
    sidebarWidth 
} = storeToRefs(appStore)

const Dashboard = () => router.push('/')
const GameSelection = () => router.push('/game-selection')
const ModDevTools = () => router.push('/mod-dev-tools')
</script>

<template>
<div v-if="sidebarExpanded" class="sidebar-expanded">
    <img class="icon no-select" src="../../assets/images/appicon.png">

    <div class="flex column h-full">
        <div class="top flex column">
            <Button 
                class="btn margin-lr no-drag" icon="pi pi-home" 
                :label="t('tooltips.sidebar.dashboard')"
                @click="Dashboard"
            />

            <Button 
                plain severity="secondary"
                class="btn margin-lr no-drag" 
                v-tooltip="tooltipOpts(t('tooltips.sidebar.game-selection'))"
                label="Games"
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
                severity="secondary"
                class="btn margin-lr no-drag" icon="pi pi-wrench"
                v-tooltip="tooltipOpts(t('tooltips.sidebar.mod-dev-tools'))"
                @click="ModDevTools"
            />

            <!-- <Button plain outlined class="btn margin-lr no-drag" icon="pi pi-upload"/> -->
        </div>

        <div class="spacer"></div>
 
        <div class="bottom flex column">
            <Button severity="secondary"
                class="btn margin-lr no-drag"
                icon="pi pi-info-circle"
                @click="appInfo.setVisible()"
            />

            <!-- <Button text class="btn margin-lr no-drag" :icon="themeMode === 'dark' ? 'pi pi-sun' : 'pi pi-moon'" @click="ToggleThemeMode"/> -->
            <Button
                severity="secondary"
                class="btn margin-lr spin-hover no-drag" icon="pi pi-cog" 
                :label="t('keywords.settings')"
                @click="settings.setVisible()"
            />
        </div>
    </div>
</div>

<div v-else class="sidebar">
    <img class="icon-alt no-select" src="../../assets/images/appicon.png">

    <div class="flex column h-full">
        <div class="top flex column">
            <Button 
                plain severity="secondary"
                class="btn-alt margin-lr no-drag" icon="pi pi-home" 
                v-tooltip="tooltipOpts(t('tooltips.sidebar.dashboard'))"
                @click="Dashboard"
            />

            <Button 
                plain severity="secondary"
                class="btn-alt margin-lr no-drag" 
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
                severity="secondary"
                class="btn-alt margin-lr no-drag" icon="pi pi-wrench"
                v-tooltip="tooltipOpts(t('tooltips.sidebar.mod-dev-tools'))"
                @click="ModDevTools"
            />

            <!-- <Button plain outlined class="btn margin-lr no-drag" icon="pi pi-upload"/> -->
        </div>

        <div class="spacer"></div>
 
        <div class="bottom flex column">
            <Button severity="secondary"
                class="btn-alt margin-lr no-drag"
                icon="pi pi-info-circle"
                @click="appInfo.setVisible()"
            />

            <!-- <Button text class="btn margin-lr no-drag" :icon="themeMode === 'dark' ? 'pi pi-sun' : 'pi pi-moon'" @click="ToggleThemeMode"/> -->
            <Button
                severity="secondary"
                class="btn-alt margin-lr spin-hover no-drag" icon="pi pi-cog" 
                v-tooltip="tooltipOpts(t('keywords.settings'))"
                @click="settings.setVisible()"
            />
        </div>
    </div>
</div>
</template>

<style scoped>
.sidebar {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: fixed;
    z-index: 999;
    height: 100vh;
    width: v-bind(sidebarWidth);
    border-right: 1px outset rgba(255, 255, 255, 0.45);
    background-color: var(--p-surface-800); /* #2c2d32; */
}

.sidebar-expanded {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: fixed;
    z-index: 999;
    height: 100vh;
    width: v-bind(sidebarWidth);
    padding-left: 20px;
}

.icon {
    width: 70px;
    height: 77px;
    margin: 20px 0px 20px 0px;
}

.icon-alt {
    width: 56px;
    height: 63px;
    margin: 15px 0px 20px 0px;
}

.btn {
    position: relative;
    border-radius: 3px;
    background: transparent !important;
    color: rgba(255, 255, 255, 0.85);
    border: none !important;
}

.btn-alt {
    position: relative;
    border-radius: 3px;
    border: 1px solid rgba(255, 255, 255, 0.3) !important;
}

.btn > * {
    font-size: 20px;
}

.top .p-button {
    margin-bottom: 10px;
}

.bottom .p-button {
    margin-top: 10px;
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
    font-size: 25px;
    font-variation-settings: 'FILL' 0, 'wght' 270, 'GRAD' 100, 'opsz' 40
}
</style>