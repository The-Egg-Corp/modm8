import { defineStore } from 'pinia'

import { core } from "@backend/models.js"
import { Layout } from "@types"

const state = {
    general: {
        locale: 'en',
        theme: 'aura-dark-purple',
        animations_enabled: true,
    },
    performance: {
        thread_count: 2,
        gpu_acceleration: true
    },
    misc: {
        nexus_personal_key: '',
        update_behaviour: 2,
        game_selection_layout: 'grid',
    }
}

const actions = {
    setLocale(code: string) {
        state.general.locale = code
    },
    setTheme(theme: string) {
        state.general.theme = theme
    },
    setAnimationsEnabled(value: boolean) {
        state.general.animations_enabled = value
    },
    setThreads(count: number) {
        state.performance.thread_count = count
    },
    setAcceleration(value: boolean) {
        state.performance.gpu_acceleration = value
    },
    setNexusPersonalKey(key: string) {
        state.misc.nexus_personal_key = key
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