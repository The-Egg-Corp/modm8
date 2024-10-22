import type { RouteRecordRaw } from 'vue-router'

const Dashboard = () => import('../views/Dashboard.vue')
const GameSelection = () => import('../views/GameSelection.vue')
const SelectedGame = () => import('../views/SelectedGame.vue')

const ModDevTools = () => import('../views/mod-dev-tools/PlatformSelection.vue')
const ModDevToolsThunderstore = () => import('../views/mod-dev-tools/Thunderstore.vue')
const ModDevToolsNexus = () => import('../views/mod-dev-tools/Nexus.vue')

const ROUTES: RouteRecordRaw[] = [{
    path: '/',
    name: 'Dashboard',
    component: Dashboard
}, {
    path: '/game-selection',
    name: 'Game Selection',
    component: GameSelection
}, {
    path: '/selected-game',
    name: 'Selected Game',
    component: SelectedGame
}, {
    path: '/mod-dev-tools',
    name: 'Mod Dev Tools (Platform Selection)',
    component: ModDevTools
}, {
    path: '/mod-dev-tools/thunderstore',
    name: 'Mod Dev Tools (Thunderstore)',
    component: ModDevToolsThunderstore
}, {
    path: '/mod-dev-tools/nexus',
    name: 'Mod Dev Tools (Nexus)',
    component: ModDevToolsNexus
}]

export default ROUTES