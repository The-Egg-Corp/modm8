<template>
<div class="frameless-buttons">
    <ButtonGroup>
        <Button class="button" icon="pi pi-minus" @click="CustomMinimise"/>
        <Button class="button" icon="pi pi-expand" @click="WindowMaximise"/>
        <Button class="button" icon="pi pi-times" @click="Quit"/>
    </ButtonGroup>
</div>
</template>

<script lang="ts" setup>
import { WindowIsMaximised, WindowMinimise, WindowUnmaximise, WindowMaximise, Quit } from '../../wailsjs/runtime/runtime'

const CustomMinimise = async() => {
    const maximized = await WindowIsMaximised()
    if (maximized) {
        WindowUnmaximise()
        return
    }

    WindowMinimise()
}
</script>

<style scoped>
.frameless-buttons {
    position: absolute;
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
    padding-right: 10px;
}

.frameless-buttons .button {
    margin-right: 3px;
    cursor: pointer;
}
</style>