import type { thunderstore } from "@backend/models.js"
import type { ObjValues, Prettify } from "./index.js"

// export type GameRunnerType = ObjValues<typeof GameRunner>
// export const GameRunner = {
//     STEAM:   "STEAM",
//     EPIC:    "EPIC",
//     XBOX_GAME_PASS: "XBOX_GP",
//     DIRECT:  "DIRECT",
//     OTHER:   "OTHER"
// } as const

export type BaseGame = {
    uuid: string
    /** The game's full name. Nexus refers to this as `name`. */
    title: string
    description?: string
    imageURL?: string
    path?: string
    favourited?: boolean
    installed?: boolean
    bepinexSetup?: boolean
    aliases?: string[]
    steamID?: number
    modLoader: string
}

export function newThunderstoreGame(game: ThunderstoreGame): GameContainer {
  return { platform: 'THUNDERSTORE', value: game }
}

export function newNexusGame(game: NexusGame): GameContainer {
  return { platform: 'NEXUS_MODS', value: game }
}

export type ThunderstoreGame = Prettify<BaseGame & {
    /** A unique short name for the game. Neither an ID or title. */
    identifier: string
    modCache?: thunderstore.StrippedPackage[]
}>

export type NexusGame = Prettify<BaseGame & {
    ID: number
    /** Equivalent to `identifier` in ThunderstoreGame. */
    domainName: string, 
    modCache?: any[] // TODO: Implement nexus package.
}>

export type ModPlatform = ObjValues<typeof ModPlatforms>
export const ModPlatforms = {
    TS: 'THUNDERSTORE',
    NEXUS: 'NEXUS_MODS'
} as const

export type GameContainer =
    { platform: typeof ModPlatforms.TS, value: ThunderstoreGame } | 
    { platform: typeof ModPlatforms.NEXUS, value: NexusGame }

// NOTE: May want to move this in future.
export type ModListTab = ObjValues<typeof ModListTabs>
export const ModListTabs = {
    ...ModPlatforms,
    PROFILE: 'PROFILE'
} as const