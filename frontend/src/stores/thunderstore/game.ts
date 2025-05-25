import { defineStore, storeToRefs } from 'pinia'
import { ref, computed } from 'vue'

import { Save, SetFavouriteGames } from '@backend/app/Persistence'
import { GetPersistence } from '@backend/app/Application.js'
import { ExistsAtPath } from '@backend/app/Utils.js'
import { BepinexInstalled } from '@backend/game/GameManager.js'
import type { thunderstore } from '@backend/models.js'

import type { Nullable, ThunderstoreGame } from '@types'
import { useGameStore } from '@stores'

export const useGameStoreTS = defineStore('GameStoreTS', () => {
    //#region State
    const games = ref<Map<string, ThunderstoreGame>>(new Map())

    const gameStore = useGameStore()
    const { selectedGame } = storeToRefs(gameStore)
    //#endregion

    //#region Getters
    const gameByID = (id: string): Nullable<ThunderstoreGame> => games.value.get(id)

    // TODO: Consider renaming this?
    const gamesAsArray = computed(() => [...games.value.values()] satisfies ThunderstoreGame[])

    // TODO: Evaluate if these are even useful.
    // const isBepinexSetup = computed(() => (id: string) => _gameByID(id)?.bepinexSetup || false)
    // const isGameInstalled = computed(() => (id: string) => _gameByID(id)?.installed || false)
    // const isFavouriteGame = computed(() => (id: string) => _gameByID(id)?.favourited || false)

    const favouriteGameIds = computed(() => gamesAsArray.value.reduce((acc: string[], game) => {
        if (game.favourited) acc.push(game.identifier)
        return acc
    }, []))
    //#endregion

    //#region Actions
    // TODO: Possibly need to take in ID, then get from `games` instead ?
    function setSelectedGame(game: ThunderstoreGame) {
        selectedGame.value.value = gameByID(game.identifier) ?? game
    }

    /** Fills the `modCache` for the currently selected game with the specified mods. */
    function updateModCache(mods: thunderstore.StrippedPackage[]) {
        if (selectedGame.value.type != 'THUNDERSTORE') {
            throw new Error('Could not update mod cache. Selected game is not of type `THUNDERSTORE`.')
        }

        const game = gameByID(selectedGame.value.value.identifier)
        if (!game) {
            throw new Error('Could not update mod cache. Selected game was not found in games.')
        }

        game.modCache = mods
    }

    async function toggleFavouriteGame(id: string) {
        const game = gameByID(id)
        if (!game) return // TODO: Implement proper error

        game.favourited = !game.favourited
        
        await SetFavouriteGames(favouriteGameIds.value)
        await Save()
    }

    async function initGames(gameList: ThunderstoreGame[]) {
        const persistence = await GetPersistence()
        games.value = new Map<string, ThunderstoreGame>()

        // Init game props.
        for (const game of gameList) {
            // TODO: Check game executable exists. For now, assume installed if game path is specified and exists.
            game.installed = !game.path ? false : await ExistsAtPath(game.path, true)
            game.favourited = await persistence.favourite_games.includes(game.identifier)

            if (game.path) {
                game.bepinexSetup = await BepinexInstalled(game.path)
            }

            // Ensure modCache is kept between mounts.
            const cachedGame = gameByID(game.identifier)
            if (cachedGame?.modCache) {
                game.modCache = cachedGame.modCache
            }

            games.value.set(game.identifier, game)
        }

        return games.value.size
    }
    //#endregion

    return {
        games,
        gamesAsArray,
        gameByID,
        // isBepinexSetup,
        // isGameInstalled,
        // isFavouriteGame,
        favouriteGameIds,
        setSelectedGame,
        toggleFavouriteGame,
        updateModCache,
        initGames
    }
})