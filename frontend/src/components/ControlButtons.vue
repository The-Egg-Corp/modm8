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
        <Button plain text icon="pi pi-minus" @click="WindowMinimise"/>
        <Button plain text icon="pi pi-expand" @click="CustomMinimise"/>
        <Button plain text icon="pi pi-times" @click="Quit"/>
    </ButtonGroup>
</div>
</template>

<style scoped>
.control-buttons {
    position: fixed;
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
    border-left: rgba(211, 211, 211, 0.823) 1px solid;
    border-bottom: rgba(211, 211, 211, 0.823) 1px solid;
    border-bottom-left-radius: 5px;
    z-index: 1000;
    background-color: rgb(23, 23, 26);
}

.control-buttons .p-button {
    cursor: pointer;
    color: rgba(201, 201, 201, 0.905);
    width: 42px;
    height: 36px;
}

.control-buttons :hover {
    color: rgba(166, 105, 246, 0.863);
}
</style>