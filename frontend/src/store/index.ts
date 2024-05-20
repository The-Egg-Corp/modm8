import { createStore, Commit } from 'vuex'

// Defaults. These will usually be different at runtime.
const state = {
    settings: {
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
}

const mutations = {
    setLocale(state: State, code: string) {
        state.settings.locale = code
    },
    setTheme(state: State, theme: string) {
        state.settings.theme = theme
    },
    setThreads(state: State, count: number) {
        state.settings.performance.thread_count = count
    },
    setAcceleration(state: State, value: boolean) {
        state.settings.performance.gpu_acceleration = value
    },
}

type Context = { 
    commit: Commit 
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
}

export type State = typeof state

export default createStore<State>({ state, mutations, actions })