<script lang="ts" setup>
import { computed } from 'vue'

import Divider from 'primevue/divider'

import { game } from '@backend/models'
import type { Nullable } from '@types'

interface ConfigProps {
    config: game.BepinexConfig
    fileName?: Nullable<string>
}

const props = defineProps<ConfigProps>()

type BepinexConfigEntry = game.BepinexConfigEntry & { checked: boolean }

const groupedEntries = computed(() => {
    const grouped: Record<string, Record<string, BepinexConfigEntry>> = {}
    const entries = Object.entries(props.config.entries)

    for (const entry of entries) {
        const [section, key] = entry[0].split('.')

        if (!grouped[section]) {
            grouped[section] = {}
        }

        const newEntry = entry[1] as BepinexConfigEntry
        newEntry.checked = asBool(newEntry.default_value)

        grouped[section][key] = newEntry
    }

    return grouped
})

const getEntryDescription = (comments: string[]) => {
    const arr = comments.map(c => c.replaceAll('#', "").trim()).filter(c => 
        c != "" && !c.startsWith('Setting type:') && !c.startsWith('Default value:')
    )

    arr.shift()
    return arr
}

const isBool = (str: string) => {
    str = str.trim()
    return str == "true" || str == "false"
}

const asBool = (str: string) => {
    str = str.trim()
    return str.toLowerCase() == "true"
}
</script>

<template>
<div class="config-edit-layout">
    <div class="flex column">
        <div class="flex row align-items-center justify-content-between">
            <h1 class="header mb-1">Config Editor</h1>
            <div class="flex row gap-2 align-items-baseline">
                <p class="file-name mt-0 mb-0">{{ fileName }}</p>
            </div>
        </div>
    </div>

    <div class="scroll-hidden pb-4">
        <div v-for="(entries, section) in groupedEntries">
            <Divider v-if="section != '__root'" align="center" type="solid" class="mb-0 mt-0">
                <p class="category-divider mb-1 mt-1">{{ section }}</p>
            </Divider>

            <!-- Loop through each entry in the section -->
            <div v-for="(entry, key) in entries" :key="key">
                <div class="flex row pt-2 justify-content-between align-items-center">
                    <!-- Container of the key and comments -->
                    <div class="mb-1" style="width: 60%;">
                        <p class="entry-key mt-0 mb-1">{{ key }}</p>
                        
                        <div v-if="entry.comments">
                            <div class="entry-description">{{ entry.comments[0].replaceAll('#', '') }}</div>
                            <div class="entry-description" v-for="comment in getEntryDescription(entry.comments)">{{ comment }}</div>
                        </div>
                    </div>

                    <div v-if="isBool(entry.default_value)" class="flex align-items-center">
                        <ToggleSwitch v-model="entry.checked"/>
                    </div>
                    <div v-else class="flex flex-grow-1 align-items-center">
                        <InputText class="ml-3 flex-grow-1" style="font-size: 16px;" :value="entry.value"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<style scoped>
.config-edit-layout {
    display: flex;
    flex-direction: column;
    max-height: calc(100vh - 180px);
}

.scroll-hidden {
    overflow-y: auto;
    scrollbar-width: none;
}

.entry-key {
    font-weight: 420;
    font-size: 18px;
    color: var(--p-primary-color);
}

.entry-description {
    font-weight: 250;
    font-size: 16px;
}

.category-divider {
    position: sticky;
    font-size: 22px;
    font-weight: 600;
}

.file-name {
    font-size: 24px;
    color: var(--p-primary-color);
    text-shadow: 0px 0px 18px rgba(255, 255, 255, 0.3);
}
</style>