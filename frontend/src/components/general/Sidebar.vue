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
    sidebarWidthPx
} = storeToRefs(appStore)

const Dashboard = () => router.push('/')
const GameSelection = () => router.push('/game-selection')
const ModDevTools = () => router.push('/mod-dev-tools')
</script>

<template>
<div v-if="sidebarExpanded" class="sidebar expanded">
    <img class="icon no-select" src="../../assets/images/appicon.png">

    <div class="btn-container gap-5">
        <div class="top flex column">
            <Button outlined size="large"
                class="btn no-drag"
                icon="pi pi-info-circle" iconPos="top"
                :label="t('keywords.about')" 
                @click="appInfo.setVisible()"
            />
        </div>

        <div class="flex column gap-1">
            <Button outlined size="large"
                class="btn no-drag" 
                icon="pi pi-th-large"
                :label="t('tooltips.sidebar.dashboard')"
                @click="Dashboard"
            />

            <Button outlined size="large"
                class="btn no-drag"
                :label="t('keywords.games')"
                @click="GameSelection"
            >
                <template #icon>
                    <span class="material-symbols-sharp">stadia_controller</span>
                </template>
            </Button>

            <!-- <Button outlined
                class="btn no-drag"
                icon="pi pi-wrench" iconPos="top"
                label="Mod Developer Tools" 
                @click="ModDevTools"
            /> -->

            <!-- <Button plain outlined class="btn no-drag" icon="pi pi-upload"/> -->
        </div>

        <div class="spacer"/>
 
        <div class="bottom flex column">
            <!-- <Button text iconPos="top" label=" " class="btn no-drag" :icon="themeMode === 'dark' ? 'pi pi-sun' : 'pi pi-moon'"/> -->
            <Button outlined
                class="btn no-drag" 
                icon="pi pi-cog"
                :label="$t('keywords.settings')"
                @click="settings.setVisible()"
            />
        </div>
    </div>
</div>

<div v-else class="sidebar collapsed">
    <img class="icon no-select" src="../../assets/images/appicon.png">

    <div class="flex column h-full mb-4">
        <div class="top flex column mb-4">
            <Button outlined raised
                class="btn no-drag" icon="pi pi-home" 
                v-tooltip="tooltipOpts(t('tooltips.sidebar.dashboard'))"
                @click="Dashboard"
            />
        </div>

        <div class="top flex column">
            <Button outlined raised
                class="btn no-drag" 
                v-tooltip="tooltipOpts(t('tooltips.sidebar.game-selection'))"
                @click="GameSelection"
            >
                <template #icon>
                    <span class="material-symbols-sharp">stadia_controller</span>
                </template>
            </Button>

            <Button outlined raised
                class="btn no-drag" icon="pi pi-wrench"
                v-tooltip="tooltipOpts(t('tooltips.sidebar.mod-dev-tools'))"
                @click="ModDevTools"
            />

            <!-- <Button plain outlined class="btn no-drag" icon="pi pi-upload"/> -->
        </div>

        <div class="spacer"></div>
 
        <div class="bottom flex column">
            <Button outlined raised
                class="btn no-drag"
                icon="pi pi-info-circle"
                @click="appInfo.setVisible()"
            />

            <!-- <Button text class="btn no-drag" :icon="themeMode === 'dark' ? 'pi pi-sun' : 'pi pi-moon'" @click="ToggleThemeMode"/> -->
            <Button outlined raised
                class="btn spin-btn-icon no-drag" icon="pi pi-cog" 
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
    width: v-bind(sidebarWidthPx);
}

.collapsed {
    height: 100vh;

    background-color: var(--sidebar-bg-color);
}

.expanded {
    height: 100vh;

    background-color: var(--sidebar-bg-color);
}

.collapsed .icon {
    width: 48px;
    margin: 15px 0px 15px 0px;
}

.expanded .icon {
    width: 75px;
    margin: 30px 0px 0px 0px;
}

.btn-container {
    display: flex;
    flex-direction: column;
    height: 100%;
    margin: 20px 0px 20px 0px;
}

.expanded .btn {
    position: relative;
    border-radius: 3px;
    background: transparent !important;
    color: rgba(209, 209, 209, 0.971);
    border: none !important;
}

.collapsed .btn {
    position: relative;
    border-radius: 3px;
    border: none !important;
    color: var(--p-surface-50);
    background-color: var(--sidebar-btn-color);
    height: 39px;
    width: 43px;
}

.top .p-button {
    margin-bottom: 10px;
}

.bottom .p-button {
    margin-top: 10px;
}

.collapsed .btn :first-child {
    font-size: 18px;
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

.material-symbols-sharp {
    font-size: 24px !important;
    font-variation-settings: 'FILL' 0, 'wght' 450, 'GRAD' 0, 'opsz' 48
}
</style>