export interface Game {
    identifier: string
    title?: string
    description?: string,
    image?: string,
    path?: string,
    favourited?: boolean
    installed?: boolean
    bepinexSetup?: boolean
}