import { createApp } from 'vue'
import App from './App.vue'
import Aura from '@primevue/themes/aura'

import { createPinia } from 'pinia'
import i18n from "./i18n"
import router from "./router"

//#region Import Framework & Components
import PrimeVue from 'primevue/config'

import Button from 'primevue/button'

import Dialog from 'primevue/dialog'
import Card from 'primevue/card'

import Dropdown from 'primevue/select'

import ToggleSwitch from 'primevue/toggleswitch'
import InputText from 'primevue/inputtext'
import InputIcon from 'primevue/inputicon'
import IconField from 'primevue/iconfield'

import Tooltip from 'primevue/tooltip'
//#endregion

//#region Import styles
import "primeflex/primeflex.css"
import "primeicons/primeicons.css"

import './assets/styles/global.css'
import "./assets/styles/flags.css"

import { definePreset } from '@primevue/themes'
import { PresetOptions } from '@types'
//#endregion

const pinia = createPinia()
const app = createApp(App)

type AuraOptions = PresetOptions<Aura.PrimitiveDesignTokens, Aura.SemanticDesignTokens>

const palettes = {
    zinc: {
        0: '#ffffff',
        50: '{zinc.50}',
        100: '{zinc.100}',
        200: '{zinc.200}',
        300: '{zinc.300}',
        400: '{zinc.400}',
        500: '{zinc.500}',
        600: '{zinc.600}',
        700: '{zinc.700}',
        800: '{zinc.800}',
        900: '{zinc.900}',
        950: '{zinc.950}'
    },
    violet: {
        50: '{violet.50}',
        100: '{violet.100}',
        200: '{violet.200}',
        300: '{violet.300}',
        400: '{violet.400}',
        500: '{violet.500}',
        600: '{violet.600}',
        700: '{violet.700}',
        800: '{violet.800}',
        900: '{violet.900}',
        950: '{violet.950}'
    }
}

// PrimeVue v4 does not provide typings for actions currently.
const defineAuraPreset = (opts: AuraOptions) => definePreset(Aura, opts)
const AuraViolet = defineAuraPreset({
    semantic: {
        primary: palettes.violet,
        colorScheme: {
            light: {
                surface: palettes.zinc
            },
            dark: {
                surface: palettes.zinc
            }
        }
    }
})

app.use(pinia)
app.use(i18n)
app.use(router)
app.use(PrimeVue, { 
    ripple: true,
    theme: {
        preset: AuraViolet
    }
})

app.component('Dialog', Dialog)
app.component('Card', Card)

app.component('Button', Button)
app.component('Dropdown', Dropdown)
app.component('ToggleSwitch', ToggleSwitch)

app.component('InputText', InputText)
app.component('InputIcon', InputIcon)
app.component('IconField', IconField)

app.directive('Tooltip', Tooltip)

app.mount('#app')