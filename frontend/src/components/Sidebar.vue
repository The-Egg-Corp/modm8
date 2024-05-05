<script lang="ts" setup>
import router from "../router"

const ToSettings = () => router.push('/settings')
</script>

<template>
<div class="sidebar">
    <Button class="settings-btn button no-drag" icon="pi pi-spin pi-cog" plain @click="ToSettings"/>
</div>
</template>

<style scoped>
.sidebar {
    z-index: 999;
    display: flex;
    flex-direction: column;
    height: 100vh;
    width: fit-content;
    border: white 1px outset;
}

.settings-btn {
    position: relative;
    margin: 0px 20px 0px 20px;
    color: rgba(201, 201, 201, 0.905);
    background-color: #18181bcd;
    border: solid rgba(167, 167, 167, 0.574) 1px;
    border-radius: 3.5px;
}

</style>