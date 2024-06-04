<script lang="ts" setup>
import { ComputedRef, computed } from 'vue'
import { usePrimeVue } from 'primevue/config'

import { useSettingsStore } from '@stores'
import { 
    ChangeEvent, 
    OptionItem, ValueItem,
    Theme, ThemeGroup
} from '@types'

const currentTheme = () => {
    const { general } = useSettingsStore()
    return general.theme as Theme || groupedThemes.value[0]
}

const selectedTheme: ComputedRef<Theme> = computed(currentTheme)
const groupedThemes: ComputedRef<ThemeGroup[]> = computed(() => [{
    label: 'Aura Purple',
    themes: [{
        label: 'Dark',
        value: 'aura-dark-purple',
    }, {
        label: 'Light',
        value: 'aura-light-purple'
    }]
}, {
    label: 'Aura Yellow',
    themes: [{
        label: 'Dark',
        value: 'aura-dark-amber'
    }, {
        label: 'Light',
        value: 'aura-light-amber'
    }]
}])

const PrimeVue = usePrimeVue()

const change = (e: ChangeEvent<Theme>) => {
    const newTheme = e.value
    
    PrimeVue.changeTheme(currentTheme().value, newTheme.value, 'theme-link')

    const { setTheme } = useSettingsStore()
    setTheme(newTheme)
}

const getGroup = (theme: Theme) => groupedThemes.value.find(g => g.themes.find(t => t.value == theme.value))
const fullLabel = (theme: Theme) => {
    const group = getGroup(theme)
    if (!group) return theme.value

    return `${group.label} (${theme.label})`
}
</script>

<template>
    <Dropdown
        class="no-drag w-full md:w-14rem" 
        optionLabel="label" optionGroupLabel="label" optionGroupChildren="themes"
        :placeholder="fullLabel(selectedTheme)"
        :options="groupedThemes"
        v-model="selectedTheme"
        @change="change"
    >
        <template #optiongroup="slotProps: OptionItem<ThemeGroup>">
            <div class="flex align-items-center">
                <div>{{ slotProps.option.label }}</div>
            </div>
        </template>

        <template #value="slotProps: ValueItem<Theme>">
            <div class="flex align-items-center">
                <div>{{ fullLabel(slotProps.value) }}</div>
            </div>
        </template>
    </Dropdown>
</template>

<style scoped>

</style>