<script lang="ts" setup>
import { 
    Save,
} from "@backend/core/AppSettings"

import { core } from "@backend/models"

import Dialog from "primevue/dialog"
import Divider from "primevue/divider"
import SelectButton from "primevue/selectbutton"

import LanguageDropdown from "./LanguageDropdown.vue"
import ThemeDropdown from "./ThemeDropdown.vue"

import { Ref, computed, ref } from "vue"

import { useDialog } from "@composables"
import { useAppStore, useSettingsStore } from "@stores"
import { Alignment, ChangeEvent, ValueItemLabeled } from "@types"
import { t } from "@i18n"

const { visible, draggable, closable } = useDialog()

const appStore = useAppStore()
const settingsStore = useSettingsStore()

const dialogStyle = {
    "margin-left": "70px",
    "width": 'auto',
    "min-width": '45rem'
}

const cardStyle = {
    "box-shadow": 'none',
    "overflow": 'hidden',
    "margin": '-15px',
    "background": 'none'
}

const dividerAlignment: Alignment = "center"

const accelChecked = ref(true)
const setAccel = (value: boolean) => {
    const store = useSettingsStore()
    store.setAcceleration(value)

    // Make Wails aware of new setting value by restarting the app (or prompt user?).
}

const threadCount = ref(2)
const setThreads = (count: number) => settingsStore.setThreads(count)

const animationsEnabled = ref(true)
const setAnimsEnabled = (value: boolean) => settingsStore.setAnimationsEnabled(value)

type Behaviour = ValueItemLabeled<core.UpdateBehaviour>

const updateBehaviour: Ref<Behaviour> = ref({ 
    label: t('settings.update-behaviour.option-1'),
    value: core.UpdateBehaviour.AUTO
})

const behaviours: Ref<Behaviour[]> = computed(() => [{
    label: t('settings.update-behaviour.option-1'),
    value: core.UpdateBehaviour.OFF
}, {
    label: t('settings.update-behaviour.option-2'),
    value: core.UpdateBehaviour.NOTIFY
}, {
    label: t('settings.update-behaviour.option-3'),
    value: core.UpdateBehaviour.AUTO
}])

const setUpdateBehaviour = (e: ChangeEvent<Behaviour>) => settingsStore.setUpdateBehaviour(e.value.value)

const applySettings = async() => {
    const t0 = performance.now()
    await Save()
    
    console.log(`Settings saved. Took ${performance.now() - t0}ms`)
}
</script>

<template>
    <div class="settings-container">
        <Dialog 
            modal class="settings-dialog no-drag"
            :style="dialogStyle"
            :block-scroll="true"
            :dismissable-mask="true"
            :show-header="false"
            v-model:visible="visible"
            v-model:draggable="draggable" 
            v-model:closable="closable"
        >
            <Card :style=cardStyle>
                <template #content>
                    <div class="card-content">
                        <div class="flex flex-column">
                            <h1 class="header">{{ $t('keywords.settings') }}</h1>
                            <p style="font-weight: 305; margin-bottom: 10px; margin-top: 10px; padding-left: 2px;">
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
                                    <InputSwitch v-model="animationsEnabled" @update:model-value="setAnimsEnabled"/>
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
                                    <InputSwitch v-model="accelChecked" @update:model-value="setAccel"/>
                                </div>
                            </div>

                            <div class="setting">
                                <div class="flex-item">
                                    <h3>{{ $t('settings.threads') }}</h3>
                                </div>
                                <div class="flex flex-row justify-content-end align-items-center gap-5">
                                    <InputText class="w-2" v-model.number="threadCount"/>
                                    <Slider class="w-12rem" v-model="threadCount" :min="2" :max="appStore.maxThreads" @change="setThreads"/>
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
                                <div class="flex-item">
                                    <SelectButton 
                                        aria-labelledby="basic" :options="behaviours.map(b => b.label)"
                                        v-model="updateBehaviour" @change="setUpdateBehaviour"
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
                    </div>
                </template>
            </Card>

            <div class="flex justify-content-end gap-3 mt-3">
                <Button class="w-full" type="button" :label="$t('keywords.close')" severity="secondary" @click="visible = false"></Button>
                <Button class="w-full" type="button" :label="$t('phrases.reset-all-default')" severity="secondary"></Button>
                <Button class="w-full" type="button" :label="$t('keywords.apply')" @click="applySettings"></Button>
            </div>
        </Dialog>
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

.settings-dialog {
    width: 42rem;
}

.card-content {
    justify-content: center;
}

.card-content .header {
    user-select: none;
    font-size: 45px;
    font-weight: 490;
    margin-top: 15px;
    margin-bottom: 5px;
    -webkit-font-smoothing: antialiased;
}

.setting {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding-left: 5px;
}

.flex-item {
    font-size: 15px;
}

.flex-item h3 {
    font-weight: 300;
    text-wrap: balance;
}

.category-divider {
    font-size: 22.5px;
    font-weight: 500;
    margin-top: -2px;
    margin-bottom: -2px;
    user-select: none;
}
</style>