<script lang="ts" setup>
import router from "../router"

const ToSettings = () => router.push('/settings')
</script>

<template>
<div class="settings">
    <Button plain class="button no-drag" icon="pi pi-spin pi-cog" @click="ToSettings"/>
</div>
</template>

<style scoped>
.settings {
    position: absolute;
    left: 1.1rem;
    bottom: 1.1rem;
}

.settings .button {
    background-color: transparent;
    color: white;
    border: none;
    border-radius: 3.5px;
}
</style>