import store, { modules } from "@store"
import { Country } from "@types"

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
  messages: { en, fr, de, it, es },
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

export const countryFromLocale = () => {
  const lang = modules.settings.state.locale
  if (!lang) return countries.value[0]

  return countries.value.find(c => c.code === lang) || countries.value[0]
}

export const getCountry = computed(countryFromLocale) 

type CountryCode = typeof locale.value

/**
 * Changes the current locale within the Vuex store.
 * @param code The locale/language code to switch to, such as 'en' or 'fr'.
 */
export const changeLocale = async (code: string) => {
  locale.value = code as CountryCode 
  await store.dispatch('settings/setLocale', code)
}