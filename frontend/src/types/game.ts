export interface Game {
    identifier: string
    title?: string
    description?: string,
    image?: string
}

export interface GameProps {
    items: Game[]
}