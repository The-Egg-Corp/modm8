// eslint-disable-next-line @typescript-eslint/consistent-type-imports
import { thunderstore } from "@backend/models.js"

export const GameRunner = {
    STEAM:   "STEAM",
    EPIC:    "EPIC",
    XBOX_GAME_PASS: "XBOX_GP",
    DIRECT:  "DIRECT",
    OTHER:   "OTHER"
} as const

export type GameRunnerType = typeof GameRunner[keyof typeof GameRunner]

export interface BaseGame {
    platform: ModPlatform
    title: string // The game's full name. Nexus refers to this as `name`.
    description?: string
    image?: string
    path?: string
    favourited?: boolean
    installed?: boolean
    bepinexSetup?: boolean
    aliases?: string[]
    steamID?: number
}

export interface ThunderstoreGame extends BaseGame {
    identifier: string // A unique short name for the game. Neither an ID or title.
    modCache?: thunderstore.StrippedPackage[]
}

export interface NexusGame extends BaseGame {
    ID: number
    domainName: string, // Equivalent to `identifier` in ThunderstoreGame.
    modsCount: number,
    modCache?: any[] // TODO: Implement nexus package.
}

export type ModPlatform = typeof ModPlatforms[keyof typeof ModPlatforms]
export const ModPlatforms = {
    TS: 'THUNDERSTORE',
    NEXUS: 'NEXUS_MODS'
} as const

// TODO: Can we enforce `type` to be one of ModPlatform?
// export type GameContainer =
//     { type: 'THUNDERSTORE', value: ThunderstoreGame } | 
//     { type: 'NEXUS_MODS', value: NexusGame }

// NOTE: May want to move this in future.
export type ModListTab = typeof ModListTabs[keyof typeof ModListTabs];
export const ModListTabs = {
    ...ModPlatforms,
    PROFILE: 'PROFILE'
} as const