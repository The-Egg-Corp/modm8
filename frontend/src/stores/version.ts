import { defineStore } from "pinia"
import { compare } from 'compare-versions'
import { version as pkgVer } from '../../package.json'
import { computed, ref } from "vue"

export type VersionState = {
    currentVer: string,
    remoteVer: string
}

type GithubRelease = {
    tag_name: string
}

export const useVersionStore = defineStore("VersionStore", () => {
    const currentVer = ref(pkgVer)
    const remoteVer = ref('')

    const newVersionAvailable = computed(() => remoteVer.value == '' ? false : compare(remoteVer.value, currentVer.value, ">"))

    async function updateRemoteVersion() {
        // Call backend to fetch latest release version

        // Temporary (requires auth while repo is private)
        const latest = await fetch("https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest").then(res => res.json()) as GithubRelease
        remoteVer.value = !latest ? currentVer.value : latest.tag_name
    }

    return {
        currentVer,
        remoteVer,
        newVersionAvailable,
        updateRemoteVersion
    }
})