import { defineStore } from "pinia"

export interface ProfileState {
    // TODO: Create 'Profile' type and replace `any`
    profiles: Map<string, any>
}

export const useProfileStore = defineStore("ProfileStore", () => {
    return {}
})