<script lang="ts" setup>
import { Ref, computed, ref, watch } from "vue"
import { storeToRefs } from "pinia"

import { app } from "@backend/models"
import { GetSettings } from "@backend/app/Application"
import { Load, SaveAndApply } from "@backend/app/AppSettings"

import { 
    CardOverlay, 
    LanguageDropdown, ThemeDropdown 
} from "@components"

import { useDialog } from "@composables"
import { useAppStore, useSettingsStore } from "@stores"
import { Alignment, ValueItemLabeled } from "@types"
import { t } from "@i18n"

import Divider from "primevue/divider"
import SelectButton from "primevue/selectbutton"
import Slider from "primevue/slider"

const { 
    setVisible,
    visible, closable, draggable 
} = useDialog('settings')

const settingsStore = useSettingsStore()
const {
    setThreads,
    setAcceleration,
    setAnimationsEnabled,
    setUpdateBehaviour
} = settingsStore

const appStore = useAppStore()
const { maxThreads } = storeToRefs(appStore)

const threadCount = ref(2)
const animationsEnabled = ref(true)
const accelChecked = ref(true)

const setAccel = (value: boolean) => {
    setAcceleration(value)

    // TODO: Make Wails aware of new setting value by restarting the app (or prompt user?).
}

//#region Update Behaviour
type Behaviour = ValueItemLabeled<app.UpdateBehaviour>

const updateBehaviour: Ref<Behaviour> = ref({ 
    label: t('settings.update-behaviour.option-1'),
    value: app.UpdateBehaviour.AUTO
})

const behaviours: Ref<Behaviour[]> = computed(() => [{
    label: t('settings.update-behaviour.option-1'),
    value: app.UpdateBehaviour.OFF
}, {
    label: t('settings.update-behaviour.option-2'),
    value: app.UpdateBehaviour.NOTIFY
}, {
    label: t('settings.update-behaviour.option-3'),
    value: app.UpdateBehaviour.AUTO
}])

// Every time the overlay is opened.
watch(visible, async (newVal: boolean) => {
    if (!newVal) return

    await Load() // Load the `settings.toml` file.
    await updateValues() // Refresh the refs with their respective loaded value.
})

const updateValues = async () => {
    const { general, performance } = await GetSettings()

    animationsEnabled.value = general.animations_enabled
    accelChecked.value = performance.gpu_acceleration
    threadCount.value = performance.thread_count
}
//#endregion

const applySettings = async() => {
    const t0 = performance.now()
    await SaveAndApply()

    console.info(`Settings saved and applied. Took ${performance.now() - t0}ms`)
    //console.info(`GOMAXPROCS: ${await GetMaxProcs()}`)
}

const dividerAlignment: Alignment = "center"
</script>

<template>
<div :class="['settings-container', { 'no-drag': visible }]">
    <CardOverlay 
        class="settings-dialog no-drag"
        v-model:visible="visible"
        v-model:closable="closable"
        v-model:draggable="draggable"
    >
        <template #cardContent>
            <div class="flex flex-column">
                <h1 class="header">{{ $t('keywords.settings') }}</h1>
                <p class="subheading">
                    Values set here are saved to the <b>settings.toml</b> file upon applying.
                </p>
            </div>

            <Divider :align="dividerAlignment" type="solid">
                <h2 class="category-divider">{{ $t('keywords.general') }}</h2>
            </Divider>

            <!-- #region General Section -->
            <div>
                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('keywords.language') }}</h3>
                    </div>
                    <div class="flex-item">
                        <LanguageDropdown/>
                    </div>
                </div>

                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('keywords.theme') }}</h3>
                    </div>
                    <div class="flex-item">
                        <ThemeDropdown/>
                    </div>
                </div>

                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('settings.animations-enabled') }}</h3>
                    </div>
                    <div class="flex-item">
                        <InputSwitch 
                            v-model="animationsEnabled" 
                            @update:model-value="setAnimationsEnabled(animationsEnabled)"
                        />
                    </div>
                </div>
            </div>
            <!-- #endregion -->

            <Divider :align="dividerAlignment" type="solid">
                <h2 class="category-divider">{{ $t('keywords.performance') }}</h2>
            </Divider>

            <!-- #region Performance Section -->
            <div>
                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('settings.gpu-acceleration') }}</h3>
                    </div>
                    <div class="flex-item">
                        <InputSwitch 
                            v-model="accelChecked" 
                            @update:model-value="setAccel(accelChecked)"
                        />
                    </div>
                </div>

                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('settings.threads') }}</h3>
                    </div>
                    <div class="w-5 flex flex-row align-items-center gap-4 justify-content-end">
                        <div style="font-size: 18px">{{ threadCount }}</div>
                        <Slider 
                            class="w-12rem" :min="2" :max="maxThreads"
                            v-model="threadCount" @change="setThreads(threadCount)"
                        />

                        <!-- @vue-ignore -->
                        <!-- <InputText type="number" style="max-width: 65px;" 
                            v-model.number="threadCount"
                            :min="2" :max="maxThreads" 
                        /> -->
                    </div>
                </div>
            </div>
            <!-- #endregion -->

            <Divider :align="dividerAlignment" type="solid">
                <h2 class="category-divider">{{ $t('keywords.misc') }}</h2>
            </Divider>
            
            <!-- #region Misc Section -->
            <div>
                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('settings.update-behaviour.label') }}</h3>
                    </div>
                    <div class="flex-item ml-3">
                        <SelectButton 
                            aria-labelledby="basic" :options="behaviours.map(b => b.label)"
                            v-model="updateBehaviour" @change="setUpdateBehaviour(updateBehaviour.value)"
                        />
                    </div>
                </div>

                <div class="setting">
                    <div class="flex-item">
                        <h3>{{ $t('settings.nexus-personal-key') }}</h3>
                    </div>
                    <div class="flex flex-row justify-content-end align-items-center">
                        <InputText type="password" style="width: 310px"/>
                    </div>
                </div>
            </div>
            <!-- #endregion -->
        </template>

        <template #dialogContent>
            <div class="flex gap-2 mt-3 justify-content-center">
                <Button class="w-7" type="button" :label="$t('keywords.close')" severity="secondary" @click="setVisible(false)"></Button>
                <Button class="w-full" type="button" icon="pi pi-refresh" :label="$t('phrases.reset-all-default')" severity="secondary"></Button>
                <Button class="w-full" type="button" :label="$t('keywords.apply')" @click="applySettings"></Button>
            </div>
        </template>
    </CardOverlay>
</div>
</template>

<style scoped>
:global(.p-dialog) {
    border: 1px solid #4d4d4d;
    border-radius: 9px;
}

:global(.p-dialog-content) {
    background: none;
    border-radius: 9px;
    padding: 1rem;
}

:global(.p-card-body) {
    padding-top: 0rem;
}

:global(.p-divider-content) {
    background: none;
}

:deep(.p-dialog-title) {
    user-select: none;
}

.settings-container {
    display: flex;
}

.setting {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding-left: 5px;
    padding-right: 5px;
}

.setting:hover {
    background: rgba(255, 255, 255, 0.05);
    border-radius: 5px;
}

.flex-item {
    font-size: 15px;
}

.flex-item h3 {
    font-weight: 300;
    text-wrap: balance;
    margin-top: 15px;
    margin-bottom: 15px;
}

.category-divider {
    font-size: 22.5px;
    font-weight: 500;
    margin-top: -2px;
    margin-bottom: -2px;
    user-select: none;
}

.subheading {
    font-weight: 305;
    margin-bottom: 0px;
    margin-top: 0px;
    padding-left: 2px;
}
</style>