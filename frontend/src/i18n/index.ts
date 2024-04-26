import { createI18n } from "vue-i18n"

//#region Import locales
import en from "./locales/en.json"
import fr from "./locales/fr.json"
//#endregion

const i18n = createI18n({
  locale: "fr",
  fallbackLocale: "en",
  availableLocales: ["en", "fr"],
  legacy: false,
  messages: { en, fr }
})

export default i18n