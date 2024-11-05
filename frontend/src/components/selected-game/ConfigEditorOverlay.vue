<script lang="ts" setup>
import { ref, watch } from 'vue'

import { BepinexConfigFiles, ParseBepinexConfig } from '@backend/game/GameManager'
import { game } from '@backend/models'

import type { Dialog } from '@composables'
import { CardOverlay, ConfigEditLayout } from '@components'

import { t } from '@i18n'
import { openLink } from "../../util"
import type { ThunderstoreGame, Nullable } from '@types'

const selectedConfig = ref<Nullable<game.BepinexConfig>>(null)
const selectedConfigName = ref<Nullable<string>>(null)

const configFiles = ref<string[]>([])

const props = defineProps<{
    dialog: Dialog
    selectedGame: ThunderstoreGame
}>()

watch(props.dialog.visible, async (newVal: boolean) => {
    if (!newVal) return
    configFiles.value = await getConfigFiles()
})

const closeConfigEditor = () => {
    props.dialog.setVisible(false)
    resetSelectedConfig()
}

const resetSelectedConfig = () => {
    selectedConfig.value = null
    selectedConfigName.value = null
}

const editConfig = async (path: string) => {
    try {
        // Read file contents from backend
        selectedConfig.value = await ParseBepinexConfig(path)
        selectedConfigName.value = getRelativeConfigPath(path)
    } catch(err: any) {
        // Show error toast
    }
}

const getConfigFiles = async () => {
    const gamePath = props.selectedGame.path

    if (!props.selectedGame.bepinexSetup || !gamePath) return []
    return await BepinexConfigFiles([gamePath, "BepInEx", "config"])
}

const getRelativeConfigPath = (absPath: string) => {
    const configPath = 'BepInEx/config'

    const normalizedFile = absPath.replace(/\\/g, '/')
    const startIndex = normalizedFile.indexOf(configPath) + configPath.length + 1

    return normalizedFile.substring(startIndex)
}

const dialogStyle = () => {
    if (selectedConfig) return {
        "margin-left": "10px",
        "width": 'auto',
        "min-width": '45rem',
        "max-width": '50rem'
    }
}
</script>

<template>
<CardOverlay class="no-drag"
    :dialogStyle="dialogStyle()"
    v-model:visible="dialog.visible.value"
    v-model:closable="dialog.closable.value"
    v-model:draggable="dialog.draggable.value"
>
    <template #cardContent>
        <div v-if="!selectedConfig" class="flex flex-column" style="max-height: calc(100vh - 180px);">
            <!-- #region Heading & Subheading -->
            <h1 class="header">{{ t('selected-game.config-editor.title') }}</h1>
            <p style="font-weight: 340; margin-bottom: 15px; margin-top: 3px; padding-left: 5px; user-select: none;">
                Choose the configuration file you would like to edit values for.
            </p>

            <div v-if="configFiles.length < 1" class="flex justify-content-center align-items-center">
                <p class="mb-3 mt-2" style="font-size: 18.5px; font-weight: 450; user-select: none;">No config files found!</p>
            </div>

            <div style="overflow-y: auto;">
                <div 
                    v-if="configFiles.length > 0"
                    v-for="(path, index) in configFiles" :key="index" 
                    class="flex flex-row pl-2 pr-2 justify-content-between align-items-center"
                    style="height: 58.5px"
                >
                    <p style="font-size: 18.5px; font-weight: 285; user-select: none;">{{ getRelativeConfigPath(path) }}</p>

                    <div class="flex gap-2">
                        <Button outlined plain 
                            icon="pi pi-file-edit"
                            style="width: 3rem; height: 2.6rem;"
                            v-tooltip.top="'Open File'"
                            @click="openLink(`file://${path}`)"
                        />

                        <Button plain
                            class="justify-content-center"
                            style="font-size: 18px; width: 6rem; height: 2.6rem;"
                            :label="t('selected-game.config-editor.edit-button')"
                            @click="editConfig(path)"
                        />
                    </div>
                </div>
            </div>
            <!-- #endregion -->
        </div>

        <ConfigEditLayout v-else :config="selectedConfig" :fileName="selectedConfigName">
            
        </ConfigEditLayout>
    </template>

    <template #dialogContent>
        <div style="position: sticky; bottom: 0;" class="flex justify-content-end gap-2">
            <Button v-if="!selectedConfig" 
                class="w-full" type="button" severity="secondary" 
                :label="t('keywords.close')" @click="closeConfigEditor()"
            />

            <Button v-else 
                class="w-5" type="button" severity="secondary" 
                icon="pi pi-arrow-left" :label="$t('keywords.back')"
                @click="resetSelectedConfig"
            />

            <Button v-if="selectedConfig"
                class="w-full" type="button" 
                :label="t('keywords.apply')" @click=""
            />
        </div>
    </template>
</CardOverlay>
</template>

<style scoped>

</style>