import { OpenExternal } from "@backend/app/Application.js"
import { TooltipOptions } from "primevue/tooltip"

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