import { defineStore } from 'pinia'

import { core } from "@backend/models.js"
import { Layout } from "@types"

const state = {
    general: {
        locale: 'en',
        theme: 'default-purple-dark',
    },
    performance: {
        thread_count: 2,
        gpu_acceleration: true
    },
    misc: {
        animations_enabled: true,
        update_behaviour: 2,
        game_selection_layout: 'grid'
    }
}

const actions = {
    setLocale(code: string) {
        state.general.locale = code
    },
    setTheme(theme: string) {
        state.general.theme = theme
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
    },
    setGameSelectionLayout(layout: Layout) {
        state.misc.game_selection_layout = layout
    }
}

export type SettingsState = typeof state

export const useSettingsStore = defineStore("SettingsStore", {
    state: () => state,
    actions,
})