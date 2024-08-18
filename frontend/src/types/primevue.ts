export type Layout = 'grid' | 'list'

export type Alignment = "center" | "left" | "right"

export interface ChangeEvent<V> {
    originalEvent: Event
    value: V
}

export interface OptionItem<T> {
    index: number
    option: T
}

export interface ValueItem<V> {
    placeholder: string
    value: V
}

export interface ValueItemLabeled<V> {
    label: string
    value: V
}

export interface ItemProps<V> {
    items: V[]
}

export interface BreadcrumbPage {
    route: string
    home?: boolean
    label?: string
    icon?: string
    class?: string
    style?: string
}