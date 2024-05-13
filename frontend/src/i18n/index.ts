import store from "@store"
import { Country } from "@types"

import { ComputedRef, computed } from "vue"
import { createI18n } from "vue-i18n"

import en from "./locales/en.json"
import fr from "./locales/fr.json"
import de from "./locales/de.json"

const i18n = createI18n({
  locale: currentLocale,
  fallbackLocale: "en",
  availableLocales: ["en", "de", "fr"],
  messages: { en, fr, de },
  legacy: false,
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
}])

export const countryFromLocale = () => {
  const lang = store.state.locale
  return lang ? countries.value.find(c => c.code === lang) ?? countries.value[0] : countries.value[0]
}

export const currentLocale = computed(countryFromLocale)

type CountryCode = typeof locale.value

export const changeLocale = (code: string) => {
  locale.value = code as CountryCode 
  store.dispatch('setLocale', code)
}