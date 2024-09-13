import { thunderstore } from "@backend/models.js"

export interface BaseGame {
    steamID: number
    identifier: string
    title: string
    description?: string
    image?: string
    path?: string
    favourited?: boolean
    installed?: boolean
    bepinexSetup?: boolean
    aliases?: string[]
}

export interface ThunderstoreGame extends BaseGame {
    modCache?: thunderstore.StrippedPackage[]
}