import { defineStore, storeToRefs } from 'pinia'
import { ref, computed } from 'vue'

import { GetPersistence } from '@backend/app/Application'
import { Save, SetFavouriteGames } from '@backend/appcore/Persistence'
import { ExistsAtPath } from '@backend/appcore/Utils'
import { BepinexInstalled } from '@backend/game/GameManager'
import type { thunderstore } from '@backend/models'

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
    // function setSelectedGame(game: ThunderstoreGame) {
    //     selectedGame.value.value = gameByID(game.identifier) ?? game
    // }

    /** Fills the `modCache` for the currently selected game with the specified mods. */
    function updateModCache(mods: thunderstore.StrippedPackage[]) {
        if (selectedGame.value.platform != 'THUNDERSTORE') {
            throw new Error('Could not update mod cache. Selected game platform is not `THUNDERSTORE`.')
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
        const { favourite_games } = await GetPersistence()
        games.value = new Map<string, ThunderstoreGame>()

        // Init game props.
        const len = gameList.length
        for (let i = 0; i < len; i++) {
            const game = gameList[i]
            
            // TODO: Check game executable exists. For now, assume installed if game path is specified and exists.
            game.installed = !game.path ? false : await ExistsAtPath(game.path, true)
            game.favourited = favourite_games.includes(game.identifier)

            if (game.path) {
                game.bepinexSetup = await BepinexInstalled(game.path)
            }

            // Ensure modCache is kept between mounts by using existing one if set.
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
        toggleFavouriteGame,
        updateModCache,
        initGames
    }
})