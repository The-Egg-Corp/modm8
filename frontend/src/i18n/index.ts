import { Country } from "@types"

import { useSettingsStore } from "@stores"
import { storeToRefs } from "pinia"

import { ComputedRef, computed } from "vue"
import { createI18n } from "vue-i18n"

import en from "./locales/en.json"
import fr from "./locales/fr.json"
import de from "./locales/de.json"
import it from "./locales/it.json"
import es from "./locales/es.json"

const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  availableLocales: ["en", "de", "fr", "it", "es"],
  messages: { en, de, fr, it, es },
  legacy: false
})

export default i18n

export const { t, n, locale } = i18n.global

export const countries: ComputedRef<Country[]> = computed(() => [{ 
  name: t('languages.en'),
  code: 'en'
}, {
  name: t('languages.de'),
  code: 'de'
}, {
  name: t('languages.fr'),
  code: 'fr'
},{
  name: t('languages.it'),
  code: 'it'
},{
  name: t('languages.es'),
  code: 'es'
}])

const getSettingsStore = () => {
  const store = useSettingsStore()
  return {
    ...storeToRefs(store),
    ...store
  }
}

export const countryFromLocale = () => {
  const {
    general
  } = getSettingsStore()

  const lang = general.locale
  if (!lang) return countries.value[0]
  
  return countries.value.find(c => c.code === general.locale) || countries.value[0]
}

export const getCountry = computed(countryFromLocale) 

type CountryCode = typeof locale.value

/**
 * Updates the app-global language to another.
 * @param code The locale/language code to switch to, such as 'en' or 'fr'.
 */
export const changeLocale = async (code: string) => {
  const {
    setLocale
  } = getSettingsStore()

  locale.value = code as CountryCode
  setLocale(code)
}