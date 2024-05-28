import { defineStore } from "pinia"

const state = {
    
}

export type ProfileState = typeof state

export const useProfileStore = defineStore("ProfileStore", {
    state: () => state
})