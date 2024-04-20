<template>
<div class="topbar">
    <Toolbar class="no-drag">
    <template #start>
        <div class="flex align-items-center gap-2">
            <Button plain label="Dashboard" text/>
            <Button plain label="Profiles" text/>
        </div>
    </template>
    </Toolbar>

    <div class="control-buttons no-drag">
        <ButtonGroup>
            <Button plain class="button" icon="pi pi-minus" @click="CustomMinimise"/>
            <Button plain class="button" icon="pi pi-expand" @click="WindowMaximise"/>
            <Button plain class="button" icon="pi pi-times" @click="Quit"/>
        </ButtonGroup>
    </div>
</div>

</template>

<script lang="ts" setup>
import { 
    WindowMinimise,
    WindowUnmaximise, WindowMaximise,
    WindowIsMaximised, Quit
} from '../../wailsjs/runtime/runtime'

const CustomMinimise = async() => {
    if (await WindowIsMaximised()) WindowUnmaximise()
    else WindowMinimise()
}
</script>

<style scoped>
.control-buttons {
    position: absolute;
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
}

.control-buttons .button {
    cursor: pointer;
    background: transparent;
    border: none;
    color: rgba(177, 177, 177, 0.889);
    padding-right: 0.5rem;
    padding-top: 1rem;
}

.topbar {
    margin-bottom: 0.5rem;
}

.p-toolbar {
    border: none;
    border-radius: 0;
    padding: 0.5rem;
}
</style>