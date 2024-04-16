import { createI18n } from "vue-i18n"

//#region Import locales
import english from "./locales/en.json"
//#endregion

const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  legacy: false,
  messages: {
    en: english
  }
})

export default i18n