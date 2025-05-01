import { defineStore } from "pinia"
import { ref } from "vue"

import type { Nullable, Profile } from "@types"
import { useGameStoreTS } from "@stores"

import { GetProfiles } from "@backend/profile/ProfileManager"

// export interface ProfileState {
//     // Key is the identifier of the game the profiles exist on.
//     profiles: Map<string, Profile>
// }

export const useProfileStore = defineStore("ProfileStore", () => {
    //#region Stores
    const gameStoreTS = useGameStoreTS()
    //#endregion

    //#region State
    const profiles = ref<Profile[]>([])
    const selectedProfile = ref<Nullable<Profile>>(null)
    //#endregion

    //#region Getters
    const profileByName = (name: string) => profiles.value?.find(p => p.name == name) satisfies Nullable<Profile>
    //#endregion

    //#region Actions
    function setSelectedProfile(prof: Profile) {
        selectedProfile.value = profileByName(prof.name) ?? prof
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
    //#endregion

    return {
        profiles,
        selectedProfile,
        setSelectedProfile,
        initProfiles
    }
})