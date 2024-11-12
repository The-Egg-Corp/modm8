import { OpenExternal } from "@backend/app/Application.js"

import type { PresetOptions } from "@types"
import { type TooltipOptions } from "primevue/tooltip"

import Aura from '@primevue/themes/aura'
import { definePreset } from "@primevue/themes"

type AuraOptions = PresetOptions<Aura.PrimitiveDesignTokens, Aura.SemanticDesignTokens>

// PrimeVue v4 does not provide typings for actions currently.
export const defineAuraPreset = (opts: AuraOptions) => definePreset(Aura, opts)

export const createPalette = (name: string) => ({
    50: `{${name}.50}`,
    100: `{${name}.100}`,
    200: `{${name}.200}`,
    300: `{${name}.300}`,
    400: `{${name}.400}`,
    500: `{${name}.500}`,
    600: `{${name}.600}`,
    700: `{${name}.700}`,
    800: `{${name}.800}`,
    900: `{${name}.900}`,
    950: `{${name}.950}`
})

const tooltipTextStyle = {
    fontSize: '14.5px'
}

const tooltipArrowStyle = {
    borderLeftColor: 'var(--primary-color)'
}

export const tooltipOpts = (text: string, textStyle = tooltipTextStyle, arrowStyle = tooltipArrowStyle): TooltipOptions => ({ 
    value: text,
    showDelay: 60,
    hideDelay: 5,
    pt: {
        text: {
            style: textStyle
        },
        arrow: {
            style: arrowStyle
        }
    }
})

type ArgsFunc = (...args: any[]) => void

export function debounce<T extends ArgsFunc>(func: T, delay = 100, now?: number) {
    let timeout: NodeJS.Timeout | null = null
    return function(this: any, ...args: Parameters<T>) {
        const ctx = this

        if (timeout) clearTimeout(timeout)
        else if (now) func.apply(ctx, args)

        const delayed = () => {
            if (!now) func.apply(ctx, args)
            timeout = null
        }

        timeout = setTimeout(delayed, delay)
    }
}

export const openLink = (path: string) => OpenExternal(path).catch(e => console.log(`Error opening folder: ${e}`))