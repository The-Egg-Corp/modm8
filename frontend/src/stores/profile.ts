import { defineStore } from "pinia"

export interface ProfileState {
    
}

const state: ProfileState = {
    
}

export const useProfileStore = defineStore("ProfileStore", {
    state: () => state
})