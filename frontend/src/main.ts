import { createApp } from 'vue'
import App from './App.vue'

import './style.css';

import Buefy from 'buefy'
//import 'buefy/dist/buefy.css'

const app = createApp(App)

//@ts-ignore
app.use(Buefy)
app.mount('#app')