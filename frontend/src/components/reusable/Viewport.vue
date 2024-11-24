<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@stores'

const appStore = useAppStore()
const { 
    sidebarOffsetPx,
    topbarHeight
} = storeToRefs(appStore)
</script>

<template>
    <div class="viewport">
        <slot></slot>
    </div>
</template>

<style scoped>
.viewport {
    max-height: calc(100vh - v-bind(topbarHeight));
    margin-left: calc(v-bind(sidebarOffsetPx));
    /* margin-right: 30px; */
}
</style>