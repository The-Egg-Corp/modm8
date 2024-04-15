import { createApp } from 'vue'
import App from './App.vue'

//#region Import Vuetify
import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
//#endregion

//#region Import styles
import 'vuetify/styles'
import "@mdi/font/css/materialdesignicons.css";

import './assets/global.css';
//#endregion

const app = createApp(App)
app.use(createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: {
            mdi
        }
    },
    theme: {
        defaultTheme: "dark"
    }
}))

app.mount('#app')