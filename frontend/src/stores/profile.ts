import { defineStore, storeToRefs } from "pinia"
import { ref } from "vue"

import type { Nullable, Profile } from "@types"
import { useGameStore } from "@stores"

import { GetProfiles } from "@backend/profile/ProfileManager"
import { platform } from "@backend/models"

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
            profiles.value = Object.entries(profs).map(([name, manifest]) => ({ 
                name, 
                mods: {
                    thunderstore: manifest.Mods[platform.ModPlatform.THUNDERSTORE],
                    nexus: manifest.Mods[platform.ModPlatform.NEXUS_MODS]
                }
            }))

            if (selectedProfile.value?.name) {
                setSelectedProfile(selectedProfile.value)
            }

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