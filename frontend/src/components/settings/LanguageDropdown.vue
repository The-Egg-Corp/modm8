<script lang="ts" setup>
import { computed } from 'vue'
import { ChangeEvent, Country } from '@types'
import { 
    t,
    countries,
    changeLocale,
    getCountry
} from '@i18n'

import { Save, SetLocale } from '@backend/core/AppSettings'
    
const change = async (e: ChangeEvent<Country>) => {
    await changeLocale(e.value.code)

    // Save locale change to settings file.
    SetLocale(e.value.code)
    Save()
}

const flagPlaceholderURL = "https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
const placeholder = computed(() => t('settings.select-language'))
    
const alphabetSort = <T extends { name: string }>(arr: T[]) => arr.sort((a: T, b: T) => {
    if (a.name < b.name) return -1
    if (a.name > b.name) return 1

    return 0
})
</script>

<template>
    <Dropdown class="no-drag w-full md:w-14rem" optionLabel="name"
        :placeholder="placeholder"
        v-model="getCountry"
        :options="alphabetSort(countries)"
        @change="change"
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
