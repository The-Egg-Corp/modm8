<script lang="ts" setup>
import { 
    Quit,
    WindowMinimise,
    WindowMaximise,
    WindowUnmaximise, 
    WindowIsMaximised,
} from '@runtime'

import ButtonGroup from 'primevue/buttongroup'

let windowMaximised: boolean

const CustomMinimise = async () => {
    const maximised = await WindowIsMaximised() 
    if (maximised) WindowUnmaximise()
    else WindowMaximise()

    windowMaximised = maximised
}
</script>

<template>
<div class="control-buttons no-drag">
    <ButtonGroup>
        <Button text icon="pi pi-minus" @click="WindowMinimise"/>
        <Button text :icon="`pi ${windowMaximised ? 'pi-minimise' : 'pi-expand'}`" @click="CustomMinimise"/>
        <Button text icon="pi pi-times" @click="Quit"/>
    </ButtonGroup>
</div>
</template>

<style scoped>
.control-buttons {
    margin-top: 4px;
    margin-right: 4px;
    z-index: 9999;
    position: fixed;
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
    /* border-left: rgba(211, 211, 211, 0.823) 1px solid; */
    /* border-bottom: rgba(211, 211, 211, 0.823) 1px solid; */
    /* border-bottom-left-radius: 5px; */
    /* background-color: rgb(23, 23, 26); */
}

.control-buttons * {
    background-color: transparent;
}

.control-buttons .p-button {
    cursor: pointer;
    color: rgba(201, 201, 201, 0.905);
    width: 40px;
}

.control-buttons :hover {
    color: rgba(166, 105, 246, 0.863);
    background-color: transparent;
}
</style>