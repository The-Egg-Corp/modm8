<script lang="ts" setup>
import { computed } from 'vue'

import type { 
    ChangeEvent, 
    OptionItem, ValueItem,
    Theme, ThemeGroup
} from '@types'

import { usePreset } from '@primevue/themes'

//const settingStore = useSettingsStore()
//const { general } = settingStore

//const currentTheme = () => general.theme as Theme || groupedThemes.value[0]

const selectedTheme = computed<Theme>(() => ({
    label: 'Violet',
    value: 'aura-violet'
}))

const groupedThemes = computed<ThemeGroup[]>(() => [{
    label: 'Aura',
    themes: [{
        label: 'Violet',
        value: 'aura-violet'
    }, {
        label: 'Amber',
        value: 'aura-amber'
    }, {
        label: 'Emerald',
        value: 'aura-emerald'
    }, {
        label: 'Indigo',
        value: 'aura-indigo'
    }]
}, {
    label: 'Lara',
    themes: [{
        label: 'Violet',
        value: 'lara-violet'
    }, {
        label: 'Amber',
        value: 'lara-amber'
    }]
}])

//const PrimeVue = usePrimeVue()

const change = (e: ChangeEvent<Theme>) => {
    const newTheme = e.value
    
    // TODO: Use presets instead - https://primevue.org/theming/styled/#utils
    //PrimeVue.changeTheme(currentTheme().value, newTheme.value, 'theme-link')
    usePreset(newTheme.value)
    
    //setTheme(newTheme)
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