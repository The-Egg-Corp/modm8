export interface Country {
    name: string
    code: string
}

export interface Theme {
    label: string
    value: string
}

export interface ThemeGroup {
    label: string
    themes: Theme[]
}