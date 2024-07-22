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