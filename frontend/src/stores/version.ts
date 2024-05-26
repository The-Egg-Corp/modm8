import { defineStore } from "pinia"
import { compare } from 'compare-versions'
import { version as pkgVer } from '../../package.json'

const state = {
    currentVersion: pkgVer,
    remoteVersion: ''
}

const actions = {
    async updateRemoteVersion() {
        // Call backend to fetch latest release version

        // Temporary
        const latest = await fetch("https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest").then(res => res.json())
        const ver = latest ? latest.tag_name : state.currentVersion

        state.remoteVersion = ver
    }
}

const getters = {
    newVersionAvailable: (s: VersionState) => s.remoteVersion == '' ? false : compare(s.remoteVersion, s.currentVersion, ">")
}

export type VersionState = typeof state

export const useVersionStore = defineStore("VersionStore", {
    state: () => state,
    actions,
    getters
})