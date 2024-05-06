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
        <Button class="button" icon="pi pi-minus" plain text @click="WindowMinimise"/>
        <Button class="button" icon="pi pi-expand" plain text @click="CustomMinimise"/>
        <Button class="button" icon="pi pi-times" plain text @click="Quit"/>
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
    background-color: transparent;
    border: rgba(211, 211, 211, 0.823) 1px outset;
    border-bottom-left-radius: 6px;
}

.control-buttons .button {
    cursor: pointer;
    background: transparent;
    color: rgba(201, 201, 201, 0.905);
    width: 46px;
    height: 38px;
    background-color: rgba(25, 25, 26, 0.845);
}
</style>