<script lang="ts" setup>
import { storeToRefs } from "pinia"

import { v1 } from "@backend/models"
import { CardOverlay } from "@components"
import { Dialog } from "@composables"

import type { Nullable } from "@types"

import { useModListStoreTS } from '@stores'

const modListStoreTS = useModListStoreTS()
const { lastInstalledMod } = storeToRefs(modListStoreTS)

const props = defineProps<{
    dialog: Dialog
    installing: boolean
    lastInstalledMod: Nullable<v1.PackageVersion>
}>()

const installStatus = () => props.installing ? "Installing..." : "Installed"

function onClose(e: Event) {
    lastInstalledMod.value = null
    props.dialog.setVisible(false)
}
</script>

<template>
<CardOverlay
    class="installing-mod-card no-drag"
    v-model:visible="dialog.visible.value"
    v-model:closable="dialog.closable.value"
    v-model:draggable="dialog.draggable.value"
>
    <template #cardContent>
        <div class="flex column justify-content-center align-items-baseline">
            <h1 class="mb-0 mt-2">{{ installStatus() }} </h1>
            
            <p v-if="lastInstalledMod" class="mt-1 mb-1" style="font-size: 18px">
               {{ lastInstalledMod.full_name }}
            </p>
        </div>
    </template>

    <template #dialogContent>
        <div style="position: sticky; bottom: 0;" class="flex justify-content-center w-full">
            <div class="flex row gap-1 flex-grow-1">
                <Button class="w-full"
                    type="button" severity="secondary"
                    :label="$t('keywords.close')" 
                    :disabled="installing"
                    @click="onClose"
                />
                <!-- <Button class="w-6"
                    type="button" severity="danger" icon="pi pi-trash"
                    label="Uninstall" :disabled="installing"
                /> -->
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