<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@stores'

const appStore = useAppStore()
const { 
    sidebarWidth 
} = storeToRefs(appStore)
</script>

<template>
    <div class="viewport">
        <slot></slot>
    </div>
</template>

<style scoped>
.viewport {
    margin-left: calc(v-bind(sidebarWidth) + 20px);
    margin-right: 20px;
}
</style>