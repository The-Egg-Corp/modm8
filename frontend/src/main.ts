import { createApp } from 'vue'
import App from './App.vue'

import i18n, { changeLocale, currentLocale } from "./i18n"
import router from "./router"
import store from "./store"

//#region Import Framework & Components
import PrimeVue from 'primevue/config'

import Button from 'primevue/button'
import ButtonGroup from 'primevue/buttongroup'

import Sidebar from 'primevue/sidebar'
import Toolbar from 'primevue/toolbar'

import Card from 'primevue/card'
import FloatLabel from 'primevue/floatlabel'
import Dropdown from 'primevue/dropdown'

import InputText from 'primevue/inputtext'
import InputIcon from 'primevue/inputicon'
import IconField from 'primevue/iconfield'
//#endregion

//#region Import styles
import "primeflex/primeflex.css"
import "primeicons/primeicons.css"
//import "@mdi/font/css/materialdesignicons.css";

import './assets/styles/global.css'
import "./assets/styles/flags.css"
//#endregion

const app = createApp(App)

app.use(i18n)
app.use(router)
app.use(PrimeVue, { ripple: true })
app.use(store)

app.component('Button', Button)
app.component('ButtonGroup', ButtonGroup)

app.component('Sidebar', Sidebar)
app.component('Toolbar', Toolbar) // AKA 'Topbar'

app.component('Card', Card)
app.component('FloatLabel', FloatLabel)
app.component('Dropdown', Dropdown)

app.component('InputText', InputText)
app.component('InputIcon', InputIcon)
app.component('IconField', IconField)

app.mount('#app')