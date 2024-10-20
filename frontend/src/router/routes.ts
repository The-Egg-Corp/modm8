import type { RouteRecordRaw } from 'vue-router'

const Dashboard = () => import('../views/Dashboard.vue')
const GameSelection = () => import('../views/GameSelection.vue')
const SelectedGame = () => import('../views/SelectedGame.vue')
const ModDevTools = () => import('../views/ModDevTools.vue')

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
    name: 'Mod Dev Tools',
    component: ModDevTools
}]

export default ROUTES