import { Game } from '@types'
import { defineStore } from 'pinia'

import { Save, SetFavouriteGames } from '@backend/app/Persistence'
import { Ref, computed, ref } from 'vue'

export interface GameState {
    selectedGame: Game,
    games: Map<string, Game>
}

export const useGameStore = defineStore('GameStore', () => {
    //#region State
    const selectedGame = ref({
        identifier: ''
    }) as Ref<Game>

    const games = ref(new Map()) as Ref<Map<string, Game>>
    //#endregion

    //#region Getters
    const _gameByID = (id: string) => games.value.get(id)
    const gameByID = computed(() => _gameByID)

    const gamesAsArray = computed(() => [...games.value.values()] as Game[])

    const isGameInstalled = computed(() => (id: string) => _gameByID(id)?.installed || false)
    const isFavouriteGame = computed(() => (id: string) => {
        const game = _gameByID(id)
        return game?.favourited || false
    })

    const favouriteGameIds = computed(() => gamesAsArray.value.reduce((acc: string[], game) => {
        if (game.favourited) acc.push(game.identifier)
        return acc
    }, []))
    //#endregion

    //#region Actions
    // TODO: Possibly need to take in ID, then get from `state.games` instead ?
    function setSelectedGame(game: Game) {
        selectedGame.value = game
    }

    async function toggleFavouriteGame(id: string) {
        const game = _gameByID(id)
        if (!game) return // TODO: Implement proper error

        game.favourited = !game.favourited
        
        await SetFavouriteGames(favouriteGameIds.value)
        await Save()
    }
    //#endregion

    return {
        selectedGame,
        games,
        gamesAsArray,
        gameByID,
        isGameInstalled,
        isFavouriteGame,
        favouriteGameIds,
        setSelectedGame,
        toggleFavouriteGame
    }
})