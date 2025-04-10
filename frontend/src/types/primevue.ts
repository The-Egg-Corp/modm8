export type Layout = 'grid' | 'list'

export type Alignment = "center" | "left" | "right"

export type OptionItem<T> = {
    index: number
    option: T
}

export type ValueItem<V> = {
    placeholder: string
    value: V
}

export type ValueItemLabeled<V, L = string> = {
    label: L
    value: V
}

export type ItemProps<V> = {
    items: V[]
}

export type ChangeEvent<V> = {
    originalEvent: Event
    value: V
}

export type BreadcrumbPage = {
    route: string
    home?: boolean
    label?: string
    icon?: string
    class?: string
    style?: string
}

export type SemanticColorScheme = {
    colorScheme?: {
        light?: any
        dark?: any
    }
}

export type PresetOptions<PrimitiveTokens, SemanticTokens> = {
    primitive?: Partial<PrimitiveTokens>
    semantic?: Partial<SemanticTokens> & SemanticColorScheme
    [key: string]: any // For 'components' etc.
}