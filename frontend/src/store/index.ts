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

type ActionParams = { 
    commit: Commit 
}

const actions = {
    setLocale(params: ActionParams, code: string) {
        params.commit('setLocale', code)
    },
    setTheme(params: ActionParams, theme: string) {
        params.commit('setTheme', theme)
    },
    setThreads(params: ActionParams, count: number) {
        params.commit('setThreads', count)
    },
    setAcceleration(params: ActionParams, value: boolean) {
        params.commit('setAcceleration', value)
    },
}

export type State = typeof state

export default createStore<State>({ state, mutations, actions })