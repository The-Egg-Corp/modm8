<script lang="ts" setup>
import { Ref, computed } from 'vue'

interface Theme {
    label: string
    value: string
}

interface ThemeGroup {
    label: string
    items: Theme[]
}

const selectedTheme = computed(() => 'aura-dark-purple')
const groupedThemes: Ref<ThemeGroup[]> = computed(() => [{
    label: 'Aura Purple',
    items: [{
        label: 'Dark',
        value: 'aura-dark-purple'
    }, {
        label: 'Light',
        value: 'aura-light-purple'
    }]
}, {
    label: 'Aura Yellow',
    items: [{
        label: 'Dark',
        value: 'aura-dark-yellow'
    }, {
        label: 'Light',
        value: 'aura-light-yellow'
    }]
}])

const labelFromValue = (value: string) => {
    let item: any
    const group = groupedThemes.value.find(g => {
        item = g.items.find(i => i.value == value)
        return item != null
    }) as ThemeGroup

    return `${group.label} (${item.label})`
}

const change = () => {
    
}
</script>

<template>
    <Dropdown 
        class="no-drag w-full md:w-14rem" 
        optionLabel="label" optionGroupLabel="label" optionGroupChildren="items"
        :placeholder="labelFromValue(selectedTheme)"
        :options="groupedThemes"
        v-model="selectedTheme"
        @change="change"
    >
        <template #optiongroup="slotProps">
            <div class="flex align-items-center">
                <div>{{ slotProps.option.label }}</div>
            </div>
        </template>
    </Dropdown>
</template>

<style scoped>

</style>