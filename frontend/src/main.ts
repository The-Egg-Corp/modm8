import { createApp } from 'vue'
import App from './App.vue'

import { createPinia } from 'pinia'
import i18n from "./i18n"
import router from "./router"

//#region Import Framework & Components
import PrimeVue from 'primevue/config'

import Dialog from 'primevue/dialog'
import Card from 'primevue/card'
import TabMenu from 'primevue/tabmenu'

import Button from 'primevue/button'
import Dropdown from 'primevue/select'
import ToggleSwitch from 'primevue/toggleswitch'

import FloatLabel from 'primevue/floatlabel'

import InputGroup from "primevue/inputgroup"
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
import Skeleton from 'primevue/skeleton'
//#endregion

const pinia = createPinia()
const app = createApp(App)

const DefaultPreset = defineAuraPreset({
    primitive: {
        // zinc: {
        //     50: "white",
        //     100: "#91919B",
        //     200: "#6C6C77",
        //     300: "#52525B",
        //     400: "#46464D",
        //     500: "#36363C",
        //     600: "#212125",
        //     700: "#18181B",
        //     800: "#121214",
        //     900: "#0E0E10",
        //     950: "#0C0C0E"
        // },
        zinc: {
            50: "white",
            100: "#d7d7db",
            200: "#bdbdc1",
            300: "#a6a6ab",
            400: "#7a7a82",
            500: "#64646d",
            600: "#47474f",
            700: "#3b3b41",
            800: "#2b2b30",
            900: "#19191c",
            950: "#0a0a0b"
        }
    },
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
app.component('TabMenu', TabMenu)

app.component('Button', Button)
app.component('Dropdown', Dropdown)
app.component('ToggleSwitch', ToggleSwitch)

app.component('FloatLabel', FloatLabel)

app.component('InputGroup', InputGroup)
app.component('InputText', InputText)
app.component('InputIcon', InputIcon)
app.component('IconField', IconField)

app.component('Skeleton', Skeleton)

app.directive('Tooltip', Tooltip)
app.directive('keyfilter', KeyFilter)

app.mount('#app')