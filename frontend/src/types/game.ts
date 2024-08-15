import { thunderstore } from "@backend/models.js"

export interface Game {
    id: number
    identifier: string
    title?: string
    description?: string
    image?: string
    path?: string
    favourited?: boolean
    installed?: boolean
    bepinexSetup?: boolean
    aliases?: string[]
    modCache?: thunderstore.StrippedPackage[]
}