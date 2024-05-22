// TODO: Move from Vuex to Pinia
import { createStore } from 'vuex'

import SettingsModule from './modules/settings.js'
// import ProfileModule from './modules/profile.js'
// import VersionModule from './modules/version.js'

// Defaults. These will usually be different at runtime.
const state = {
    selectedGame: null
}

const modules = {
    settings: SettingsModule,
    // profile: ProfileModule,
    // version: VersionModule
}

// The root state of all modules.
export type AppState = typeof state
export type Modules = typeof modules

export { 
   modules 
}

export default createStore<AppState>({ 
    strict: false,
    state,
    modules
})