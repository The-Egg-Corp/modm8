import { ValueItemLabeled } from "./primevue.js"

export interface Country {
    name: string
    code: string
}

export type Theme = ValueItemLabeled<string>

export interface ThemeGroup {
    label: string
    themes: Theme[]
}