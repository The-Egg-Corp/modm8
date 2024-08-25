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
    <div class="flex column">
        <div class="flex flex-column">
            <h2 class="header">Editing: {{ fileName }}</h2>
            
        </div>

        <div v-for="(entries, section) in groupedEntries">
            <Divider v-if="section != '__root'" align="center" type="solid" class="mb-0 mt-0">
                <h2 class="category-divider">{{ section }}</h2>
            </Divider>

            <!-- Loop through each entry in the section -->
            <div v-for="(entry, key) in entries" :key="key">
                <div class="flex row justify-content-between pt-2">
                    <div class="flex column">
                        <p class="mt-0 mb-0" style="font-size: 20px">{{ key }}</p>
                        <p class="mt-1 mb-2" style="font-size: 16px">{{ getEntryDescription(entry.comments) }}</p>
                    </div>

                    <InputText :value="entry.value">
                
                    </InputText>
                </div>

            </div>
        </div>
    </div>
</template>

<style scoped>

</style>