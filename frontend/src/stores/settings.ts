import { core } from "@backend/models.js"
import { defineStore } from 'pinia'

const state = {
    locale: 'en',
    theme: 'default-purple-dark',
    performance: {
        thread_count: 2,
        gpu_acceleration: true
    },
    misc: {
        animations_enabled: true,
        update_behaviour: 2
    }
}

const actions = {
    setLocale(code: string) {
        state.locale = code
    },
    setTheme(theme: string) {
        state.theme = theme
    },
    setThreads(count: number) {
        state.performance.thread_count = count
    },
    setAcceleration(value: boolean) {
        state.performance.gpu_acceleration = value
    },
    setAnimationsEnabled(value: boolean) {
        state.misc.animations_enabled = value
    },
    setUpdateBehaviour(behaviour: core.UpdateBehaviour) {
        state.misc.update_behaviour = behaviour
    }
}

export type SettingsState = typeof state

export const useSettingsStore = defineStore({
    id: 'SettingsStore',
    state: () => state,
    actions
})