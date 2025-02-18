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
    <div class="viewport">
        <slot></slot>
    </div>
</template>

<style scoped>
.viewport {
    width: 100%;
    height: calc(100vh - (v-bind(topbarHeightPx) * 2)); /* Stinky hack to keep equal spacing at the bottom as the top. Probably bc of padding fuckery */
    position: fixed;
    top: v-bind(topbarHeightPx);
    left: v-bind(sidebarWidthPx);
    padding-top: v-bind(topbarMarginPx);
    padding-left: v-bind(sidebarMarginPx);
    padding-right: v-bind(sidebarOffsetPx);
    border-radius: 15px;
}
</style>