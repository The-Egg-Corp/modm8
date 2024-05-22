<script lang="ts" setup>
import { Ref, ref } from 'vue'

const themes: Ref<string[]> = ref(['Default Purple - Dark', 'Default Purple - Light'])
const currentTheme = ref('Default')

const placeholder = ref('Default Purple - Dark')
</script>

<template>
    <!-- <Dropdown class="no-drag w-full md:w-14rem" optionLabel="name"
        :placeholder="placeholder"
        v-model="currentTheme"
        :options="themes"
        @change=""
    >
        <template #option="selectedItem">
            <div class="flex no-drag align-items-center">
                <div class="no-select">{{ selectedItem.option.name }}</div>
            </div>
        </template>

        <template #value="selectedItem">
            <div v-if="selectedItem.value" class="flex align-items-center">
                <div>{{ selectedItem.value.name }}</div>
            </div>

            <span v-else>{{ selectedItem.placeholder }}</span>
        </template>
    </Dropdown> -->
</template>

<style scoped>

</style>