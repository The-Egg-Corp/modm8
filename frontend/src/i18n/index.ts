import { createI18n } from "vue-i18n"

//#region Import locales
import en from "./locales/en.json"
import fr from "./locales/fr.json"
import de from "./locales/de.json"
//#endregion

const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  availableLocales: ["en", "fr", "de"],
  messages: { en, fr, de },
  legacy: false,
})

export default i18n