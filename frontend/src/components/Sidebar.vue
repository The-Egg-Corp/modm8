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
    z-index: 999;
}

.settings .button {
    color: rgba(201, 201, 201, 0.905);
    background-color: #18181bcd;
    border: solid rgba(167, 167, 167, 0.574) 1px;
    border-radius: 3.5px;
}
</style>