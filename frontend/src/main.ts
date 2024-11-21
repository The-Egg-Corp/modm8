import { createApp } from 'vue'
import App from './App.vue'

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
import KeyFilter from 'primevue/keyfilter'
//#endregion

//#region Import styles
import "primeflex/primeflex.css"
import "primeicons/primeicons.css"

import './assets/styles/global.css'
import "./assets/styles/flags.css"

import { 
    defineAuraPreset, 
    presetPalette
} from './util'
//#endregion

const pinia = createPinia()
const app = createApp(App)

const DefaultPreset = defineAuraPreset({
    semantic: {
        primary: presetPalette('purple'),
        colorScheme: {
            light: { surface: presetPalette('zinc') },
            dark:  { surface: presetPalette('zinc') }
        }
    }
})

app.use(pinia)
app.use(i18n)
app.use(router)
app.use(PrimeVue, { 
    ripple: true,
    theme: {
        preset: DefaultPreset
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
app.directive('keyfilter', KeyFilter);

app.mount('#app')