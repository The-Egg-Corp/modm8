<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@stores'

const appStore = useAppStore()
const { 
    sidebarOffsetPx,
    sidebarMarginPx,
    topbarHeightPx
} = storeToRefs(appStore)
</script>

<template>
    <div class="viewport">
        <slot></slot>
    </div>
</template>

<style scoped>
.viewport {
    margin-top: v-bind(topbarHeightPx);
    margin-left: v-bind(sidebarOffsetPx);
    margin-right: v-bind(sidebarMarginPx);
}
</style>