<script lang="ts" setup>
import { game } from '@backend/models'
import { computed } from 'vue'

import InputText from 'primevue/inputtext'
import Divider from 'primevue/divider'
import { Nullable } from 'primevue/ts-helpers'

interface ConfigProps {
    config: game.BepinexConfig
    fileName: Nullable<string>
}

const props = defineProps<ConfigProps>()

const groupedEntries = computed(() => {
    const grouped: Record<string, Record<string, game.BepinexConfigEntry>> = {}
    const entries = Object.entries(props.config.entries)

    for (const entry of entries) {
        const [section, key] = entry[0].split('.')

        if (!grouped[section]) {
            grouped[section] = {}
        }

        grouped[section][key] = entry[1]
    }

    return grouped
})

const getEntryDescription = (comments: string[]) => {
    const arr = comments.map(c => c.replaceAll('#', "").trim()).filter(c => 
        c != "" && !c.startsWith('Setting type:') && !c.startsWith('Default value:')
    )

    return arr
}
</script>

<template>
    <div class="flex column" style="max-height: calc(100vh - 180px);">
        <div class="flex flex-column">
            <h1 class="header mb-2">Config Editor</h1>
            <h3 class="mt-0 mb-0">Currently editing: {{ fileName?.replace('.cfg', '') }}</h3>
        </div>

        <div style="overflow-y: auto; scrollbar-width: none;">
            <div v-for="(entries, section) in groupedEntries">
                <Divider v-if="section != '__root'" align="center" type="solid" class="mb-0 mt-0">
                    <h2 class="category-divider">{{ section }}</h2>
                </Divider>
    
                <!-- Loop through each entry in the section -->
                <div v-for="(entry, key) in entries" :key="key">
                    <div class="flex row pt-2 justify-content-between align-items-center">
                        <div class="flex column">
                            <p class="mt-0 mb-0" style="font-size: 20px; color: var(--primary-color)">{{ key }}</p>
                            <div v-for="comment in getEntryDescription(entry.comments)"> 
    
                            </div>
                        </div>
    
    
                        <InputText class="ml-3" style="font-size: 16.5px;" :value="entry.value">
                    
                        </InputText>
                    </div>
    
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>

</style>