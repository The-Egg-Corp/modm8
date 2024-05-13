export interface ChangeEvent<V> {
    originalEvent: Event
    value: V
}

export interface Country {
    name: string
    code: string
}