export type Game = {
    identifier: string
    title?: string
    rating?: number
    description?: string,
    image?: string
}

export interface GameProps {
    items: Game[]
}