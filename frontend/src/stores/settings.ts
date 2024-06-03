import { defineStore } from 'pinia'

import { core } from "@backend/models.js"
import { Layout } from "@types"
import {
    SetAnimationsEnabled, 
    SetGPUAccel, 
    SetLocale,
    SetNexusPersonalKey,
    SetTheme, 
    SetThreads,
    SetUpdateBehaviour,
    SetGameSelectionLayout
} from '@backend/core/AppSettings.js'

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

// 'Save' here means updating the viper config in the backend.
// See: backend/core/settings.go
const actions = {
    setLocale(code: string, save = true) {
        state.general.locale = code
        if (save) SetLocale(code)
    },
    setTheme(theme: string, save = true) {
        state.general.theme = theme
        if (save) SetTheme(theme)
    },
    setAnimationsEnabled(value: boolean, save = true) {
        state.general.animations_enabled = value
        if (save) SetAnimationsEnabled(value)
    },
    setThreads(count: number, save = true) {
        state.performance.thread_count = count
        if (save) SetThreads(count)
    },
    setAcceleration(value: boolean, save = true) {
        state.performance.gpu_acceleration = value
        if (save) SetGPUAccel(value)
    },
    setNexusPersonalKey(key: string, save = true) {
        state.misc.nexus_personal_key = key
        if (save) SetNexusPersonalKey(key)
    },
    setUpdateBehaviour(behaviour: core.UpdateBehaviour, save = true) {
        state.misc.update_behaviour = behaviour
        if (save) SetUpdateBehaviour(behaviour)
    },
    setGameSelectionLayout(layout: Layout, save = true) {
        state.misc.game_selection_layout = layout
        if (save) SetGameSelectionLayout(layout)
    }
}

export type SettingsState = typeof state

export const useSettingsStore = defineStore("SettingsStore", {
    state: () => state,
    actions,
})