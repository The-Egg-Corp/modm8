<script lang="ts" setup>
import { CardOverlay } from "@components"
import { Dialog } from "@composables"

defineProps<{
    dialog: Dialog
    installing: boolean
    lastInstalledMod: string
}>()
</script>

<template>
<CardOverlay
    class="installing-mod-card no-drag"
    v-model:visible="dialog.visible"
    v-model:closable="dialog.closable"
    v-model:draggable="dialog.draggable"
>
    <template #cardContent>
        <div v-if="installing" class="flex column justify-content-center align-items-baseline">
            <h1 class="mb-1 mt-2">Installing...</h1>
            <p class="mt-1" style="font-size: 18px">{{ lastInstalledMod }}</p>
        </div>
        <div v-else>
            <h1 class="mb-1 mt-2">Installed</h1>
            <p class="mt-1" style="font-size: 18px">{{ lastInstalledMod }}</p>
        </div>
    </template>

    <template #dialogContent>
        <div style="position: sticky; bottom: 0;" class="flex justify-content-center w-full">
            <div v-if="!installing" class="flex row gap-1 flex-grow-1">
                <Button class="w-full"
                    type="button" severity="secondary"
                    :label="$t('keywords.close')" @click="dialog.setVisible(false)"
                />
                <Button class="w-6"
                    type="button" severity="danger" icon="pi pi-trash"
                    label="Uninstall"
                />
            </div>
            <!-- <Button v-else class="flex flex-grow-1" 
                type="button" severity="danger" icon="pi pi-ban"
                :label="$t('keywords.cancel')" 
                @click=""
            /> -->
        </div>
    </template>
</CardOverlay>
</template>

<style scoped>
.installing-mod-card {
    padding: 0 1rem 1rem 1rem;
}
</style>