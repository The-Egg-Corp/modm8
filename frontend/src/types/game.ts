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

export interface Game {
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

export interface ThunderstoreGame extends Game {
    identifier: string // A unique short name for the game. Neither an ID or title.
    modCache?: thunderstore.StrippedPackage[]
}

export interface NexusGame extends Game {
    ID: number
    domainName: string, // Equivalent to `identifier` in ThunderstoreGame.
    modsCount: number,
    modCache?: any[] // TODO: Implement nexus package.
}

export type AnyGame = ThunderstoreGame | NexusGame

// NOTE: May want to move this in future.
export type ModListTabType = typeof ModListTabs[keyof typeof ModListTabs];
export const ModListTabs = {
    PROFILE: 'PROFILE',
    TS: 'THUNDERSTORE',
    NEXUS: 'NEXUS_MODS'
} as const