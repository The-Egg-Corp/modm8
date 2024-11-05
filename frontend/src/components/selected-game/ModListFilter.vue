<script lang="ts" setup>
import { computed } from 'vue'
import { ValueItemLabeled } from '@types'

const filters = computed<ValueItemLabeled<string>[]>(() => [{ 
    label: "Most Rated",
    value: "most-rated"
}, {
    label: "Most Downloaded",
    value: "most-downloaded"
}, {
    label: "Newest",
    value: "newest"
}, {
    label: "Last Updated",
    value: "last-updated"
}])
</script>

<template>
    <Dropdown  class="no-drag" :options="filters" @change="">
        <template #option>
            
        </template>

        <template #value>
            
        </template>
    </Dropdown>
</template>

<style scoped>

</style>