import { createApp } from 'vue'
import App from './App.vue'

import i18n from "./i18n/index"

//#region Import framework
import PrimeVue from 'primevue/config'
//#endregion

//#region Import styles
//import "@mdi/font/css/materialdesignicons.css";
import "primevue/resources/themes/aura-dark-amber/theme.css"

import './assets/global.css';
//#endregion

const app = createApp(App)

app.use(i18n)
app.use(PrimeVue)

app.mount('#app')