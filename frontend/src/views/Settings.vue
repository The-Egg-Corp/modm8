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
}])

interface ChangeEvent<V> {
    originalEvent: Event
    value: V
}

const dropdownChange = (e: ChangeEvent<Country>) => {
    locale.value = e.value?.code
}
</script>

<template>
    <div class="container">
        <p class="header">{{ $t('keywords.settings') }}</p>

        <Dropdown @change="dropdownChange" class="no-drag-all-children w-full md:w-14rem" v-model="selectedCountry" :options="countries" optionLabel="name" placeholder="Select language">
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
                <div class="flex align-items-center">
                    <img 
                        :alt="slotProps.option.label" 
                        :class="`mr-2 flag-style flag flag-${slotProps.option.code.toLowerCase()}`" 
                        :src=flagPlaceholderURL
                    />

                    <div>{{ slotProps.option.name }}</div>
                </div>
            </template>
        </Dropdown>
    </div>
</template>

<style scoped>
.flag-style {
    width: 18px;
}

.container {
    position: absolute;
    top: 30px;
    left: 50%;
    transform: translate(-50%);
}

.container .header {
    user-select: none;
    font-size: 45px;
}
</style>