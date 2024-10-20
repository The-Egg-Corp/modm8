// eslint-disable-next-line @typescript-eslint/consistent-type-imports
import { thunderstore } from "@backend/models.js"

export interface Game {
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

export interface ThunderstoreGame extends Game {
    modCache?: thunderstore.StrippedPackage[]
}