import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

import { Save, SetFavouriteGames } from '@backend/app/Persistence'
import { GetPersistence } from '@backend/app/Application.js'
import { ExistsAtPath } from '@backend/app/Utils.js'
import { BepinexInstalled } from '@backend/game/GameManager.js'
import type { thunderstore } from '@backend/models.js'

import type { ThunderstoreGame } from '@types'

// export interface GameState {
//     selectedGame: ThunderstoreGame,
//     games: Map<string, ThunderstoreGame>
// }

export const useGameStoreTS = defineStore('GameStoreTS', () => {
    //#region State
    const games = ref<Map<string, ThunderstoreGame>>(new Map())
    const selectedGame = ref<ThunderstoreGame>({
        title: 'Placeholder',
        identifier: 'tf2',
        steamID: 0
    })
    //#endregion

    //#region Getters
    const _gameByID = (id: string) => games.value.get(id)
    const gameByID = computed(() => _gameByID)

    // TODO: Consider renaming this?
    const gamesAsArray = computed(() => [...games.value.values()] as ThunderstoreGame[])

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
        selectedGame.value = _gameByID(game.identifier) ?? game
    }

    /** Fills the `modCache` for the currently selected game with the specified mods. */
    function updateModCache(mods: thunderstore.StrippedPackage[]) {
        const game = _gameByID(selectedGame.value.identifier)
        if (!game) return // TODO: Implement proper error

        game.modCache = mods
    }

    async function toggleFavouriteGame(id: string) {
        const game = _gameByID(id)
        if (!game) return // TODO: Implement proper error

        game.favourited = !game.favourited
        
        await SetFavouriteGames(favouriteGameIds.value)
        await Save()
    }

    async function initGames(gameList: ThunderstoreGame[]) {
        games.value = new Map<string, ThunderstoreGame>()

        const persistence = await GetPersistence()
        const len = gameList.length
        
        // Init game props.
        for (let i = 0; i < len; i++) {
            const game = gameList[i]
            
            // TODO: Check game executable exists. For now, assume installed if game path is specified and exists.
            game.installed = !game.path ? false : await ExistsAtPath(game.path, true)
            game.favourited = await persistence.favourite_games.includes(game.identifier)

            if (game.path) {
                game.bepinexSetup = await BepinexInstalled(game.path)
            }

            // Ensure modCache is kept between mounts.
            const cachedGame = _gameByID(game.identifier)
            if (cachedGame?.modCache) {
                game.modCache = cachedGame.modCache
            }

            games.value.set(game.identifier, game)
        }

        return games.value.size
    }
    //#endregion

    return {
        selectedGame,
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