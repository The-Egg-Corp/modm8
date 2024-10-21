<script lang="ts" setup>
import { ref } from 'vue'

interface HoverProps {
    invert: boolean
}

const props = withDefaults(defineProps<HoverProps>(), { invert: false })
const outlined = ref(props.invert)
</script>

<template>
    <Button :outlined="outlined" 
        @mouseover="outlined = props.invert ? false : true" 
        @mouseleave="outlined = props.invert ? true : false"
    />
</template>