<script lang="ts" setup>
import Splitter from 'primevue/splitter'
import SplitterPanel from 'primevue/splitterpanel'
//import OutlineHoverButton from "./OutlineHoverButton.vue"

const emit = defineEmits(['selectThunderstore', 'selectNexus'])
</script>

<template>
<Splitter 
    class="platform-splitter"
    gutterSize.number="2"
    stateStorage="local"
    stateKey="platform-splitter-state"
>
    <SplitterPanel :min-size="25" class="platform-splitter-panel panel-ts">
        <div class="blur w-full h-full">
            <div class="flex column panel-text">
                Thunderstore
                <Button class="mt-3" label="Select" @click="emit('selectThunderstore')"/>
            </div>
        </div>
    </SplitterPanel>

    <SplitterPanel :min-size="35" class="platform-splitter-panel panel-nexus">
        <div class="blur w-full h-full">
            <div class="flex column panel-text">
                Nexus Mods
                <Button class="mt-3" label="Select" @click="emit('selectNexus')"/>
            </div>
        </div>
    </SplitterPanel>
</Splitter>
</template>

<style scoped>
.platform-splitter {
    height: 100vh;
    width: 100vw;
    background: transparent !important;
    border: none;
}

.platform-splitter-panel {
    display: flex;
    align-items: center;
    justify-content: center;
}

.panel-text {
    font-size: 55px;
    font-weight: 510;
    user-select: none;
    color: #ffffff;
    text-shadow: 0 0 5px rgb(25, 25, 25);
    text-wrap: nowrap;
}

.panel-ts {
    background-image: url('../../assets/images/ts.webp');
    background-position: center;
    background-position-y: 0px;
    background-repeat: no-repeat;
    background-size: cover;
}

.panel-nexus {
    background-image: url('../../assets/images/nexus.png');
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
}

.blur {
    display: flex;
    justify-content: center;
    align-items: center;
    backdrop-filter: brightness(25%) blur(12px); /* TODO: Also use blur here, but on GPU not CPU */
}

:deep(.p-splitter-gutter) { 
    background: #202125;
    border: 1px solid rgba(255, 255, 255, 0.259);
}

:deep(.p-splitter-gutter-handle) { 
    background: rgba(255, 255, 255, 0.617) !important;
}
</style>