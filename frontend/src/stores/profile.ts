import { defineStore } from "pinia"
import { ref } from "vue"

import { Nullable, Profile } from "@types"
import { useGameStore } from "@stores"

import { GetProfiles } from "@backend/profile/ProfileManager.js"

export interface ProfileState {
    profiles: Map<string, Profile>
}

// const mockProfiles = [
//     { name: 'Default', mods: { thunderstore: ['Owen3H-CSync'] } },
//     { name: 'testing', mods: {} },
//     { name: 'MEGALOPHOBIA', mods: {} },
//     { name: 'big futa cock mods', mods: {} }
// ]

export const useProfileStore = defineStore("ProfileStore", () => {
    const gameStore = useGameStore()

    // Key is the identifier of the game the profiles exist on.
    const profiles = ref<Profile[]>()
    const selectedProfile = ref<Nullable<Profile>>()

    function setSelectedProfile(prof: Profile) {
        selectedProfile.value = profiles.value?.find(p => p.name == prof.name) ?? prof
    }

    async function initProfiles() {
        try {
            let profs = await GetProfiles(gameStore.selectedGame.title)
            profiles.value = Object.entries(profs).map(([key, value]) => ({ ...value, name: key }))
        } catch (e: any) {
            console.error(e)
        }
    }

    return {
        profiles,
        selectedProfile,
        setSelectedProfile,
        initProfiles
    }
})