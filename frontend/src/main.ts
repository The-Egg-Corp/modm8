import { createApp } from 'vue'
import App from './App.vue'

import i18n from "./i18n"
import router from "./router"

//#region Import framework
import PrimeVue from 'primevue/config'

import Button from 'primevue/button'
import ButtonGroup from 'primevue/buttongroup'
import Card from 'primevue/card'
import Toolbar from 'primevue/toolbar'
import FloatLabel from 'primevue/floatlabel'
import InputText from 'primevue/inputtext'
import DataView from 'primevue/dataview'
//#endregion

//#region Import styles
import "primevue/resources/themes/aura-dark-amber/theme.css"
import "primeicons/primeicons.css";
//import "@mdi/font/css/materialdesignicons.css";

import './assets/global.css';
//#endregion

const app = createApp(App)

app.use(i18n)
app.use(router)
app.use(PrimeVue, { ripple: true })

app.component('Button', Button)
app.component('ButtonGroup', ButtonGroup)
app.component('Card', Card)
app.component('Toolbar', Toolbar)
app.component('FloatLabel', FloatLabel)
app.component('InputText', InputText)
app.component('DataView', DataView)

app.mount('#app')