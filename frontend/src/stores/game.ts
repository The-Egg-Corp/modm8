import { Game } from '@types'
import { defineStore } from 'pinia'

import { SetFavouriteGames } from '@backend/core/Persistence'
import { reactive } from 'vue'

// Stores state for anything game-related.
const state = {
    selectedGame: {
        identifier: ''
    } as Game,
    games: reactive(new Map()) as Map<string, Game>
}

const actions = {
    // TODO: Possibly need to take in ID get from `state.games` instead ?
    setSelectedGame: (game: Game) => state.selectedGame = game,

    gamesAsArray: () => [...state.games.values()] as Game[],
    gameByID: (id: string) => state.games.get(id),

    isGameInstalled: (id: string) => actions.gameByID(id)?.installed || false,
    isFavouriteGame: (id: string) => actions.gameByID(id)?.favourited || false,

    toggleFavouriteGame: async (id: string) => {
        const game = actions.gameByID(id)
        if (!game) return // TODO: Implement proper error

        game.favourited = !game.favourited

        const arr = actions.favouriteGameIds()
        console.log(`Favourite games updated: ${arr}`)
        
        return SetFavouriteGames(arr)
    },

    favouriteGameIds: () => actions.gamesAsArray().reduce((acc: string[], game) => {
        if (!game.favourited) acc.push(game.identifier)
        return acc
    }, []),
}

export type GameState = typeof state

export const useGameStore = defineStore({
    id: 'GameStore',
    state: () => state,
    actions
})