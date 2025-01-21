<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@stores'

const appStore = useAppStore()
const { 
    sidebarWidthPx,
    sidebarMarginPx,
    sidebarOffsetPx,
    topbarHeightPx,
    topbarMarginPx
} = storeToRefs(appStore)
</script>

<template>
    <div class="viewport w-full">
        <slot></slot>
    </div>
</template>

<style scoped>
.viewport {
    position: fixed;
    top: v-bind(topbarHeightPx);
    padding-top: v-bind(topbarMarginPx);
    left: v-bind(sidebarWidthPx);
    padding-left: v-bind(sidebarMarginPx);
    padding-right: v-bind(sidebarOffsetPx);
    border-radius: 15px;
}
</style>