import { createApp } from 'vue'
import App from './App.vue'

import { createPinia } from 'pinia'
import i18n from "./i18n"
import router from "./router"

//#region Import Framework & Components
import PrimeVue from 'primevue/config'

import Button from 'primevue/button'
import ButtonGroup from 'primevue/buttongroup'

import Dialog from 'primevue/dialog'
import Card from 'primevue/card'

import FloatLabel from 'primevue/floatlabel'
import Dropdown from 'primevue/dropdown'

import InputSwitch from 'primevue/inputswitch'
import InputText from 'primevue/inputtext'
import InputIcon from 'primevue/inputicon'
import IconField from 'primevue/iconfield'

import Tooltip from 'primevue/tooltip'
//#endregion

//#region Import styles
import "primeflex/primeflex.css"
import "primeicons/primeicons.css"
//import "@mdi/font/css/materialdesignicons.css";

import './assets/styles/global.css'
import "./assets/styles/flags.css"
//#endregion

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(i18n)
app.use(router)
app.use(PrimeVue, { ripple: true })

app.component('Button', Button)
app.component('ButtonGroup', ButtonGroup)

app.component('Dialog', Dialog)
app.component('Card', Card)

app.component('FloatLabel', FloatLabel)
app.component('Dropdown', Dropdown)

app.component('InputSwitch', InputSwitch)
app.component('InputText', InputText)
app.component('InputIcon', InputIcon)
app.component('IconField', IconField)

app.directive('Tooltip', Tooltip)

app.mount('#app')