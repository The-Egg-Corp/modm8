import { RouteRecordRaw } from 'vue-router'

import Dashboard from '../views/Dashboard.vue'
import GameSelection from '../views/GameSelection.vue'
import Settings from '../views/Settings.vue'

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