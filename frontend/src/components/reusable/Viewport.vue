<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@stores'

const appStore = useAppStore()
const { 
    sidebarWidth 
} = storeToRefs(appStore)

const topbarHeight = '30px'
</script>

<template>
    <div class="viewport">
        <slot></slot>
    </div>
</template>

<style scoped>
.viewport {
    max-height: calc(100vh - v-bind(topbarHeight));
    margin-left: calc(v-bind(sidebarWidth) + 20px);
    margin-right: 20px;
}
</style>