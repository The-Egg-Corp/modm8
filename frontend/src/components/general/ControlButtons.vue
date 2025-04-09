<script lang="ts" setup>
import { 
    Quit,
    WindowMinimise,
    WindowMaximise,
    WindowUnmaximise, 
    WindowIsMaximised,
} from '@runtime'

const CustomMinimise = async () => {
    const maximised = await WindowIsMaximised() 
    
    if (maximised) WindowUnmaximise()
    else WindowMaximise()
}
</script>

<template>
<div class="control-buttons no-drag gap-1">
    <Button text icon="pi pi-minus" @click="WindowMinimise"/>
    <Button text icon="pi pi-expand" @click="CustomMinimise"/>
    <Button text icon="pi pi-times" class="exit-btn" @click="Quit"/>
</div>
</template>

<style scoped>
.control-buttons {
    display: flex;
    z-index: 9999;
    margin: 5px;
    /* border-left: rgba(211, 211, 211, 0.823) 1px solid; */
    /* border-bottom: rgba(211, 211, 211, 0.823) 1px solid; */
    /* border-bottom-left-radius: 5px; */
    /* background-color: rgb(23, 23, 26); */
}

.control-buttons * {
    background-color: transparent;
}

.control-buttons .p-button {
    color: rgba(201, 201, 201, 0.8);
    width: 32px;
    padding: 0.25rem;
}

.control-buttons :nth-child(1) :hover {
    color: rgba(166, 105, 246, 1);
}

.control-buttons :nth-child(2) :hover {
    color: rgba(166, 105, 246, 1);
}

.control-buttons :nth-child(3) :hover {
    color: rgba(255, 93, 93, 1);
}
</style>