export type Alignment = "center" | "left" | "right"
export interface ChangeEvent<V> {
    originalEvent: Event
    value: V
}