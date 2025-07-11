<script lang="ts" setup>
import { onBeforeMount, onBeforeUnmount } from "vue"

import * as Steam from '@backend/steam/SteamLauncher'
import * as GameManager from "@backend/game/GameManager"
import { GetPathToProfile } from "@backend/profile/ProfileManager"
import { loaders } from "@frontend/wailsjs/go/models"

import { useDialog } from '@composables'
import {
    Viewport,
    ModList,
    ProfileManager,
    ModInstallationOverlay, 
    ConfigEditorOverlay
} from "@components"

import { 
    useAppStore,
    useProfileStore,
    useGameStore,
    useModListStoreTS
} from "@stores"

import { storeToRefs } from "pinia"

const appStore = useAppStore()
const { sidebarMarginPx } = storeToRefs(appStore)

const profileStore = useProfileStore()
const { selectedProfile } = storeToRefs(profileStore)

const gameStore = useGameStore()
const { selectedGame } = storeToRefs(gameStore)

const modListStoreTS = useModListStoreTS()
const { installing, lastInstalledMod } = storeToRefs(modListStoreTS)

const configEditorDialog = useDialog('selected-game-config-editor')
const installingModDialog = useDialog('selected-game-installing-mod')

const canStartModded = () => {
    const profileMods = selectedProfile.value?.mods
    if (!profileMods) return false

    const tsMods = (profileMods?.thunderstore?.length || 0) > 0
    const nexusMods = (profileMods?.nexus?.length || 0) > 0

    // Only start modded if we have at least one mod from either platform.
    return tsMods || nexusMods
}

// TODO: This is temporary and not reliable :)
const gameThumbnail = () => selectedGame.value.value.imageURL
    ? `https://raw.githubusercontent.com/ebkr/r2modmanPlus/develop/src/assets/images/game_selection/${selectedGame.value.value.imageURL}` 
    : ""

const launchSteamGame = async (modded: boolean, args?: string[]) => {
    // TODO: These returns should probably reject.
    if (!selectedGame.value) return console.error('Error launching Steam game! Invalid or no selected game.')

    const id = selectedGame.value.value.steamID
    if (!id || id < 10) {
        return console.error(`Error launching Steam game!\n\nCannot use invalid ID \`${id}\`.`)
    }

    // Steam games (apps) IDs are in multiples of 10, anything in-between is used for depots.
    if (id % 10 !== 0) {
        return console.error(`Error launching Steam game!\n\nCannot use depot ID \`${id}\` as an app ID. Must be a multiple of 10.`)
    }

    if (!selectedProfile.value) {
        return console.error(`Error launching Steam game!\n\nProfile is null. Make sure it still exists and is selected.`)
    }

    const profPath = await GetPathToProfile(selectedGame.value.value.title, selectedProfile.value?.name)
    const instructions = await GameManager.GetModLoaderInstructions(loaders.ModLoaderType.BEPINEX, profPath)

    console.log(`Attempting to launch game '${selectedGame.value.value.title}'. Steam ID: ${id}`)
    return await Steam.LaunchGame(id, modded ? instructions.ModdedParams : instructions.VanillaParams)
}

onBeforeUnmount(() => {
    // This is needed currently, otherwise a profile from one game could stay selected on another game.
    selectedProfile.value = null
})

onBeforeMount(async () => {
    // The API fetch here will only happen on first mount when mod cache is empty.
    // Any subsequent mounts will use the populated mod cache.
    await modListStoreTS.refreshMods(true)

    const selectedGameTitle = selectedGame.value?.value.title || "No game selected"
    const tsModsAmt = selectedProfile.value?.mods.thunderstore?.length || 0
    const nexusModsAmt = selectedProfile.value?.mods.nexus?.length || 0

    const profStr = `${selectedProfile.value?.name} [${tsModsAmt + nexusModsAmt} mods]`
    console.log(`[SelectedGame]\nMounted game: ${selectedGameTitle}\nSelected profile: ${profStr}`)

    // Ensure no overlays are still shown.
    configEditorDialog.setVisible(false)
    installingModDialog.setVisible(false)
})
</script>

<template>
<Viewport :class="['selected-game', { 'no-drag': configEditorDialog.visible || installingModDialog.visible }]">
    <div class="flex column collapsible-content">
        <Card class="selected-game-card no-drag">
            <template #title>
                <p class="selected-game-card-header mt-0 mb-2">
                    {{ $t('selected-game.currently-selected') }}
                </p>
            </template>
            
            <template #content>
                <div class="flex column no-drag align-items-center">
                    <div class="game-thumbnail-container">
                        <img class="selected-game-thumbnail-background" :src="gameThumbnail()"/>
                        <img class="selected-game-thumbnail-foreground" :src="gameThumbnail()"/>
                    </div>

                    <div class="mt-3">
                        <p class="selected-game-title mt-0 mb-0">{{ selectedGame.value.title }}</p> 
                        
                        <div class="flex column mt-2 flex-grow-1">
                            <div class="flex row gap-2">
                                <Button plain class="btn flex-1" 
                                    icon="pi pi-caret-right"
                                    :label="$t('selected-game.start-modded-button')"
                                    :disabled="!canStartModded()"
                                    @click="launchSteamGame(true)"
                                />
        
                                <Button plain class="btn" severity="secondary"
                                    icon="pi pi-caret-right"
                                    :label="$t('selected-game.start-vanilla-button')"
                                    @click="launchSteamGame(false)"
                                />
                            </div>
    
                            <Button class="btn mt-2"
                                icon="pi pi-file-edit"
                                :label="$t('selected-game.config-button')"
                                @click="configEditorDialog.setVisible(true)"
                            />
                        </div>
                    </div>
                </div>
            </template>
        </Card>

        <ProfileManager @profileSelected="modListStoreTS.refreshPage()"/>
    </div>

    <div class="flex mod-list-container">
        <ModList :installingModDialog="installingModDialog"/>
    </div>

    <!-- Opened when ModList (see above) calls to install a mod. -->
    <ModInstallationOverlay 
        :dialog="installingModDialog" 
        :installing="installing"
        :lastInstalledMod="lastInstalledMod!"
    />

    <!-- Opened when the "Edit Configs" button is pressed. -->
    <ConfigEditorOverlay 
        :dialog="configEditorDialog" 
        :selectedGame="selectedGame"
    />
</Viewport>
</template>

<style scoped>
.material-symbols-sharp {
    font-size: 26.5px;
    font-variation-settings: 'FILL' 0, 'wght' 270, 'GRAD' 100, 'opsz' 40
}

.selected-game {
    display: flex;
    flex-direction: row;
    gap: v-bind(sidebarMarginPx); /* Spacing between Card/ProfileManager and the DataView. */
}

.collapsible-content {
    /*height: calc(100vh - 80px);*/
    min-width: 500px; /* Ensure it never shrinks while out, the mod list can suck it. */
    background: var(--collapsible-content-color);
    border-radius: 12px;
    padding: 10px 20px 10px 20px;
}

.selected-game-card {
    flex-shrink: 0;
    background: none;
    box-shadow: none;
    justify-content: center;
    align-items: center;
}

.selected-game-card-header {
    font-size: 36px;
    font-weight: 540;
    user-select: none;
    justify-self: center;
}

.selected-game-card :deep(.p-card-body) {
    padding: 0px;
    gap: 0;
}

.game-thumbnail-container {
    position: relative;
}

.selected-game-thumbnail-foreground {
    position: relative;
    min-width: 155px;
    max-width: 40%;
    max-height: 200px;
    border-radius: 3px;
    user-select: none;
    border: 1px ridge rgba(255, 255, 255, 0.55);
}

.selected-game-thumbnail-background {
    position: absolute;
    z-index: -1;
    filter: blur(4px);
    width: auto;
    height: 100%;
}

.selected-game-title {
    font-size: 26px;
    font-weight: 330;
    text-shadow: 0px 0px 10px rgba(255, 255, 255, 0.45);
    justify-self: center;
}

.btn {
    border-radius: 4px;
    text-align: left;
}

.justify-left {
    justify-content: left !important;
}

.mod-list-container {
    max-width: 100%;
}

:deep(.p-tabmenu-tablist) {
    background: none !important;
}

:deep(.p-tabmenu-item-label) {
    text-wrap: nowrap;
}
</style>