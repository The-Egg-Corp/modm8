<script lang="ts" setup>
import { 
    Quit,
    WindowMinimise,
    WindowMaximise,
    WindowUnmaximise, 
    WindowIsMaximised,
} from '../../wailsjs/runtime/runtime'

const CustomMinimise = async() => await WindowIsMaximised() 
    ? WindowUnmaximise()
    : WindowMaximise()
</script>

<template>
<div class="control-buttons no-drag">
    <ButtonGroup>
        <Button plain class="button" icon="pi pi-minus" @click="WindowMinimise"/>
        <Button plain class="button" icon="pi pi-expand" @click="CustomMinimise"/>
        <Button plain class="button" icon="pi pi-times" @click="Quit"/>
    </ButtonGroup>
</div>
</template>

<style scoped>
.control-buttons {
    position: absolute;
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
    background-color: #1e1e21;
    border-radius: 5px;
}

.control-buttons .button {
    cursor: pointer;
    background: transparent;
    border: none;
    color: rgba(201, 201, 201, 0.905);
    margin-right: 0;
}
</style>