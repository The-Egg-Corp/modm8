import { defineStore } from 'pinia'
import { ref } from 'vue'

import {
    SetAnimationsEnabled, 
    SetGPUAccel, 
    SetLocale,
    SetSteamInstallPath,
    SetNexusPersonalKey,
    SetThreads,
    SetUpdateBehaviour,
    SetGameSelectionLayout
} from '@backend/app/AppSettings.js'

import type { app } from "@backend/models.js"

export const useSettingsStore = defineStore("SettingsStore", () => {
    //#region State
    const general = ref({
        locale: 'en',
        animations_enabled: true,
        update_behaviour: 2,
    })
    
    const performance = ref({
        thread_count: 2,
        gpu_acceleration: true
    })
    
    const misc = ref({
        steam_install_path: '',
        nexus_personal_key: '',
        game_selection_layout: 'grid'
    })
    //#endregion

    //#region Actions
    
    // The 'update' parameter means updating the Viper config in the backend.
    // Saving of the `settings.toml` file is done upon clicking the 'Apply' button.
    function setLocale(code: string, update = true) {
        general.value.locale = code
        if (update) SetLocale(code)
    }

    // function setTheme(theme: Theme, save = true) {
    //     Object.assign(general.value.theme, theme)
    //     if (save) SetTheme(theme.value)
    // }

    function setAnimationsEnabled(value: boolean, update = true) {
        general.value.animations_enabled = value
        if (update) SetAnimationsEnabled(value)
    }

    function setUpdateBehaviour(behaviour: app.UpdateBehaviour, update = true) {
        general.value.update_behaviour = behaviour
        if (update) SetUpdateBehaviour(behaviour)
    }

    function setThreads(count: number, update = true) {
        performance.value.thread_count = count
        if (update) SetThreads(count)
    }

    function setAcceleration(value: boolean, update = true) {
        performance.value.gpu_acceleration = value
        if (update) SetGPUAccel(value)
    }

    function setSteamInstallPath(path: string, update = true) {
        misc.value.steam_install_path = path
        if (update) SetSteamInstallPath(path)
    }

    function setNexusPersonalKey(key: string, update = true) {
        misc.value.nexus_personal_key = key
        if (update) SetNexusPersonalKey(key)
    }

    function setGameSelectionLayout(layout: app.GameSelectionLayout, update = true) {
        misc.value.game_selection_layout = layout
        if (update) SetGameSelectionLayout(layout)
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
        setSteamInstallPath,
        setNexusPersonalKey,
        setUpdateBehaviour,
        setGameSelectionLayout
    }
})