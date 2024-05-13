import { createStore, Commit } from 'vuex'

const state = {
    locale: localStorage.getItem('locale') || 'gb'
}

const mutations = {
    setLocale(state: State, code: string) {
        state.locale = code
        localStorage.setItem('locale', code)
    }
}

const actions = {
    setLocale({ commit }: { commit: Commit }, code: string) {
        commit('setLocale', code)
    }
}

export type State = typeof state

export default createStore<State>({ state, mutations, actions })