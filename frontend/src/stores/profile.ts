import { defineStore } from "pinia"
import { ref } from "vue"

import type { Nullable, Profile } from "@types"
import { useGameStoreTS } from "@stores"

import { GetProfiles } from "@backend/profile/ProfileManager"

export interface ProfileState {
    profiles: Map<string, Profile>
}

export const useProfileStore = defineStore("ProfileStore", () => {
    const gameStoreTS = useGameStoreTS()

    // Key is the identifier of the game the profiles exist on.
    const profiles = ref<Profile[]>([])
    const selectedProfile = ref<Nullable<Profile>>(null)

    function setSelectedProfile(prof: Profile) {
        selectedProfile.value = profiles.value?.find(p => p.name == prof.name) ?? prof
    }

    async function initProfiles() {
        try {
            const profs = await GetProfiles(gameStoreTS.selectedGame.title)
            profiles.value = Object.entries(profs).map(([key, value]) => ({ ...value, name: key }))
        } catch (e: any) {
            console.error(e)
            // TODO: Add some sort of toast or error msg in the app itself.
        }
    }

    return {
        profiles,
        selectedProfile,
        setSelectedProfile,
        initProfiles
    }
})