import { createStore, Commit } from 'vuex'
import { SetLocale, Save } from '@backend/core/AppSettings'

const state = {
    locale: 'en'
}

const mutations = {
    setLocale(state: State, code: string) {
        state.locale = code
        SetLocale(code)
        Save()
    }
}

const actions = {
    setLocale({ commit }: { commit: Commit }, code: string) {
        commit('setLocale', code)
    }
}

export type State = typeof state

export default createStore<State>({ state, mutations, actions })