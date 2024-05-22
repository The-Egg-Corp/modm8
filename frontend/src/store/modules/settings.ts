import { core } from "@backend/models.js"
import { Context } from "@types"

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

const mutations = {
    setLocale(state: SettingsState, code: string) {
        state.locale = code
    },
    setTheme(state: SettingsState, theme: string) {
        state.theme = theme
    },
    setThreads(state: SettingsState, count: number) {
        state.performance.thread_count = count
    },
    setAcceleration(state: SettingsState, value: boolean) {
        state.performance.gpu_acceleration = value
    },
    setAnimationsEnabled(state: SettingsState, value: boolean) {
        state.misc.animations_enabled = value
    },
    setUpdateBehaviour(state: SettingsState, behaviour: core.UpdateBehaviour) {
        state.misc.update_behaviour = behaviour
    }
}

const actions = {
    setLocale(ctx: Context, code: string) {
        ctx.commit('setLocale', code)
    },
    setTheme(ctx: Context, theme: string) {
        ctx.commit('setTheme', theme)
    },
    setThreads(ctx: Context, count: number) {
        ctx.commit('setThreads', count)
    },
    setAcceleration(ctx: Context, value: boolean) {
        ctx.commit('setAcceleration', value)
    },
    setAnimationsEnabled(ctx: Context, value: boolean) {
        ctx.commit('setAnimationsEnabled', value)
    },
    setUpdateBehaviour(ctx: Context, behaviour: core.UpdateBehaviour) {
        ctx.commit('setUpdateBehaviour', behaviour)
    }
}

export type SettingsState = typeof state

export const SettingsModule = {
    namespaced: true,
    state,
    actions,
    mutations
}

export default SettingsModule