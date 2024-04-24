import { RouteRecordRaw } from 'vue-router'

const Dashboard = () => import('../views/Dashboard.vue')
const GameSelection = () => import('../views/GameSelection.vue')
const Settings = () => import('../views/Settings.vue')

const ROUTES: RouteRecordRaw[] = [{
    path: '/', 
    name: 'Dashboard', 
    component: Dashboard
}, {
    path: '/settings',
    name: 'Settings',
    component: Settings
}, {
    path: '/game-selection',
    name: 'Game Selection',
    component: GameSelection
}]

export default ROUTES