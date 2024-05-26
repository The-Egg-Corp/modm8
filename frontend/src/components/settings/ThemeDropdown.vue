<script lang="ts" setup>
import { ComputedRef, Ref, computed, ref } from 'vue'

import { useSettingsStore } from '@stores'
import { ChangeEvent, OptionItem, ValueItem } from '@types'

interface Theme {
    label: string
    value: string
}

interface ThemeGroup {
    label: string
    themes: Theme[]
}

const selectedTheme: Ref<Theme> = ref({
    label: 'Dark',
    value: 'aura-dark-purple'
})

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
        value: 'aura-dark-yellow'
    }, {
        label: 'Light',
        value: 'aura-light-yellow'
    }]
}])

const change = (e: ChangeEvent<Theme>) => {
    const settingsStore = useSettingsStore()
    
    settingsStore.setTheme(e.value.value)
    selectedTheme.value = e.value
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