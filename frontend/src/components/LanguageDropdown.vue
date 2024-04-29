<script lang="ts" setup>
import { ComputedRef, Ref, computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t, locale } = useI18n()
const flagPlaceholderURL = "https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"

interface Country {
    name: string
    code: string
}

const placeholder = computed(() => t('settings.select-language'))

const selectedCountry: Ref<Country | null> = ref(null)
const countries: ComputedRef<Country[]> = computed(() => [{ 
    name: t('languages.en'),
    code: 'gb'
}, {
    name: t('languages.de'),
    code: 'de'
}, {
    name: t('languages.fr'),
    code: 'fr'
}])

interface ChangeEvent<V> {
    originalEvent: Event
    value: V
}

const ChangeLocale = (e: ChangeEvent<Country>) => {
    locale.value = e.value?.code

    //#region Rebind the selected item with the updated value
    const updatedCountry = countries.value.find(c => c.code === e.value.code)
    if (!updatedCountry) return

    selectedCountry.value = updatedCountry
    //#endregion
}

const AlphabetSort = <T extends { name: string }>(arr: T[]) => arr.sort((a: T, b: T) => {
    if (a.name < b.name) return -1
    if (a.name > b.name) return 1

    return 0
})
</script>

<template>
    <Dropdown class="no-drag w-full md:w-14rem" optionLabel="name"
        :placeholder="placeholder"
        v-model="selectedCountry"
        :options="AlphabetSort(countries)"
        @change="ChangeLocale"
    >
        <template #option="selectedItem">
            <div class="flex no-drag align-items-center">
                <img
                    :alt="selectedItem.option.label" 
                    :class="`mr-2 flag-style flag flag-${selectedItem.option.code.toLowerCase()}`" 
                    :src=flagPlaceholderURL
                />

                <div class="no-select">{{ selectedItem.option.name }}</div>
            </div>
        </template>

        <template #value="selectedItem">
            <div v-if="selectedItem.value" class="flex align-items-center">
                <img
                    :alt="selectedItem.value.label" 
                    :class="`mr-2 flag-style flag flag-${selectedItem.value.code.toLowerCase()}`"
                    :src=flagPlaceholderURL
                />

                <div>{{ selectedItem.value.name }}</div>
            </div>

            <span v-else>{{ selectedItem.placeholder }}</span>
        </template>
    </Dropdown>
</template>

<style scoped>
.flag-style {
    width: 24px;
}
</style>