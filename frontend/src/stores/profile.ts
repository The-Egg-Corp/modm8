import { defineStore, storeToRefs } from "pinia"
import { ref } from "vue"

import type { Nullable, Profile } from "@types"
import { useGameStore } from "@stores"

import { GetProfiles } from "@backend/profile/ProfileManager"

// export interface ProfileState {
//     // Key is the identifier of the game the profiles exist on.
//     profiles: Map<string, Profile>
// }

export const useProfileStore = defineStore("ProfileStore", () => {
    //#region Stores
    const gameStore = useGameStore()
    const { selectedGame } = storeToRefs(gameStore)
    //#endregion

    //#region State
    const profiles = ref<Profile[]>([]) // TODO: Maybe just use the exact same layout as Go (Map<string, Profile>) instead.
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
            // Convert go map to array of profiles.
            const profs = await GetProfiles(selectedGame.value.value.title)
            profiles.value = Object.entries(profs).map(([name, manifest]) => ({ name, ...manifest }))

            //console.log(`Profiles for ${gameStoreTS.selectedGame.title}:`, Object.keys(profs))
        } catch (e: any) {
            console.error(e)
            // TODO: Add some sort of toast or error msg in the app itself.
        }
    }
    //#endregion

    return {
        profiles,
        selectedProfile,
        profileByName,
        setSelectedProfile,
        initProfiles
    }
})