import { defineStore } from 'pinia'

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

import { core } from "@backend/models.js"
import { Theme } from "@types"
import { ref } from 'vue'

export interface SettingsState {

}

export const useSettingsStore = defineStore("SettingsStore", () => {
    //#region State
    const general = ref({
        locale: 'en',
        theme: {
            label: 'Dark',
            value: 'aura-dark-purple'
        } as Theme,
        animations_enabled: true,
    })
    
    const performance = ref({
        thread_count: 2,
        gpu_acceleration: true
    })
    
    const misc = ref({
        nexus_personal_key: '',
        update_behaviour: 2,
        game_selection_layout: 'grid',
    })
    //#endregion

    //#region Actions

    // 'Save' here means updating the viper config in the backend.
    // See: backend/core/settings.go
    function setLocale(code: string, save = true) {
        general.value.locale = code
        if (save) SetLocale(code)
    }

    function setTheme(theme: Theme, save = true) {
        Object.assign(general.value.theme, theme)
        if (save) SetTheme(theme.value)
    }

    function setAnimationsEnabled(value: boolean, save = true) {
        general.value.animations_enabled = value
        if (save) SetAnimationsEnabled(value)
    }

    function setThreads(count: number, save = true) {
        performance.value.thread_count = count
        if (save) SetThreads(count)
    }

    function setAcceleration(value: boolean, save = true) {
        performance.value.gpu_acceleration = value
        if (save) SetGPUAccel(value)
    }

    function setNexusPersonalKey(key: string, save = true) {
        misc.value.nexus_personal_key = key
        if (save) SetNexusPersonalKey(key)
    }

    function setUpdateBehaviour(behaviour: core.UpdateBehaviour, save = true) {
        misc.value.update_behaviour = behaviour
        if (save) SetUpdateBehaviour(behaviour)
    }

    function setGameSelectionLayout(layout: core.GameSelectionLayout, save = true) {
        misc.value.game_selection_layout = layout
        if (save) SetGameSelectionLayout(layout)
    }
    //#endregion

    return {
        general,
        performance,
        misc,
        setLocale,
        setTheme,
        setAnimationsEnabled,
        setThreads,
        setAcceleration,
        setNexusPersonalKey,
        setUpdateBehaviour,
        setGameSelectionLayout
    }
})