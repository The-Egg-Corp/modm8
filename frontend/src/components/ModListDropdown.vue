<script lang="ts" setup>
import { computed, ComputedRef } from 'vue'

export const filters: ComputedRef<{ label: string, value: string }[]> = computed(() => [{ 
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
    <Dropdown 
        class="no-drag"
        :options="filters"
        @change=""
    >
        <template #option>
            
        </template>

        <template #value>

        </template>
    </Dropdown>
</template>

<style scoped>

</style>