import { defineStore } from 'pinia'
import { ref } from 'vue'

import {
    SetAnimationsEnabled, 
    SetGPUAccel, 
    SetLocale,
    SetNexusPersonalKey,
    SetTheme, 
    SetThreads,
    SetUpdateBehaviour,
    SetGameSelectionLayout
} from '@backend/app/AppSettings.js'

import type { app } from "@backend/models.js"

//import type { Theme } from "@types"

// export interface SettingsState {

// }

export const useSettingsStore = defineStore("SettingsStore", () => {
    //#region State
    const general = ref({
        locale: 'en',
        animations_enabled: true
    })
    
    const performance = ref({
        thread_count: 2,
        gpu_acceleration: true
    })
    
    const misc = ref({
        nexus_personal_key: '',
        update_behaviour: 2,
        game_selection_layout: 'grid'
    })
    //#endregion

    //#region Actions

    // 'Save' here means updating the viper config in the backend. Refer to the 'settings.go' file.
    function setLocale(code: string, save = true) {
        general.value.locale = code
        if (save) SetLocale(code)
    }

    // function setTheme(theme: Theme, save = true) {
    //     Object.assign(general.value.theme, theme)
    //     if (save) SetTheme(theme.value)
    // }

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

    function setUpdateBehaviour(behaviour: app.UpdateBehaviour, save = true) {
        misc.value.update_behaviour = behaviour
        if (save) SetUpdateBehaviour(behaviour)
    }

    function setGameSelectionLayout(layout: app.GameSelectionLayout, save = true) {
        misc.value.game_selection_layout = layout
        if (save) SetGameSelectionLayout(layout)
    }
    //#endregion

    return {
        general,
        performance,
        misc,
        setLocale,
        //setTheme,
        setAnimationsEnabled,
        setThreads,
        setAcceleration,
        setNexusPersonalKey,
        setUpdateBehaviour,
        setGameSelectionLayout
    }
})