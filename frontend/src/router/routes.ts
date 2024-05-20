import { RouteRecordRaw } from 'vue-router'

const Dashboard = () => import('../views/Dashboard.vue')
const GameSelection = () => import('../views/GameSelection.vue')

const ROUTES: RouteRecordRaw[] = [{
    path: '/', 
    name: 'Dashboard', 
    component: Dashboard
}, {
    path: '/game-selection',
    name: 'Game Selection',
    component: GameSelection
}]

export default ROUTES