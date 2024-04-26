<script lang="ts" setup>
import { Ref, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t, locale } = useI18n()
const flagPlaceholderURL = "https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"

interface Country {
    name: string
    code: string
}

const selectedCountry: Ref<any> = ref()
const countries: Ref<Country[]> = ref([{ 
    name: t('languages.en'),
    code: 'gb'
}, { 
    name: t('languages.fr'),
    code: 'fr'
}, {
    name: t('languages.de'),
    code: 'de'
}])

interface ChangeEvent<V> {
    originalEvent: Event
    value: V
}

const ChangeLocale = (e: ChangeEvent<Country>) => {
    locale.value = e.value?.code
}
</script>

<template>
    <Dropdown 
        @change="ChangeLocale" class="no-drag w-full md:w-14rem" 
        optionLabel="name"
        :placeholder="$t('settings.select-language')"
        v-model="selectedCountry" 
        :options="countries" 
    >
        <template #value="slotProps">
            <div v-if="slotProps.value" class="flex align-items-center">
                <img 
                    :alt="slotProps.value.label" 
                    :class="`mr-2 flag-style flag flag-${slotProps.value.code.toLowerCase()}`"
                    :src=flagPlaceholderURL
                />

                <div>{{ slotProps.value.name }}</div>
            </div>

            <span v-else>{{ slotProps.placeholder }}</span>
        </template>

        <template #option="slotProps">
            <div class="flex no-drag align-items-center">
                <img 
                    :alt="slotProps.option.label" 
                    :class="`mr-2 flag-style flag flag-${slotProps.option.code.toLowerCase()}`" 
                    :src=flagPlaceholderURL
                />

                <div class="no-select">{{ slotProps.option.name }}</div>
            </div>
        </template>
    </Dropdown>
</template>

<style scoped>
.flag-style {
    width: 24px;
}
</style>