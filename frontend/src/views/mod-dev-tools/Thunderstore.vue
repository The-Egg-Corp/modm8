<script lang="ts" setup>
import { computed, ref } from 'vue'

import FileUpload, { FileUploadUploaderEvent } from 'primevue/fileupload'
import DataView from 'primevue/dataview'

import { Viewport } from '@components'
import { TSPackageManifest, TSPackageFile } from '@types'
import { experimental } from '@backend/models'

// TODO: Finish implementing pack/unpack functionality
import { 
    //UnpackZip, 
    ValidateIcon, 
    ValidateManifest, 
    ValidateReadme,
} from '@backend/thunderstore/Tools'

import { GetFilesInZip } from '@backend/app/Utils'

// Validates each required file for uploading a Thunderstore mod.
function ValidateFiles(icon: experimental.IconValidatorParams, readme: TSPackageFile, manifest: TSPackageManifest) {
    ValidateIcon(icon)
    ValidateReadme(readme.data)

    if (!manifest.author) return
    ValidateManifest(manifest.author, manifest.data)
}

const items = ref([{ 
    label: 'Your Packages', 
    icon: 'pi pi-box' 
}, {
    label: 'Package Validater',
    icon: 'pi pi-check-square' 
}])

type FileMap = {
    [key: string]: string;
}

const files = ref<FileMap>()
const zipName = ref<string>()

const packageIcon = computed(() => files.value ? files.value['icon.png'] : null)
const packageIconSrc = computed(() => `data:image/png;base64,${packageIcon.value}`)

const onUploadZip = async (event: FileUploadUploaderEvent) => {
    const file = Array.isArray(event.files) ? event.files[0] : event.files
    console.log(`Uploaded: ${file.name}`)

    const buf = await file.arrayBuffer()
    const data = [...new Uint8Array(buf)]

    // Unzip in backend returning map[string][]byte
    files.value = await GetFilesInZip(data) as unknown as FileMap
    zipName.value = file.name
}

const requiredFilesExist = computed(() => Object.values(getRequiredFiles()).every(x => !!x))
const getRequiredFiles = () => {
    if (!files.value) return {}

    return {
        icon: files.value['icon.png'],
        manifest: files.value['manifest.json'],
        readme: files.value['README.md'],
        changelog: files.value['CHANGELOG.md']
    }
}

// Refers to 10GB. Even for big mods, reaching this will be quite rare.
const MAX_ZIP_BYTES = 1e+10
const ICON_WIDTH = 128
</script>

<template>
<Viewport class="mod-dev-tools">
    <p class="zip-selection-header mt-0 mb-3">Package Validator</p>

    <FileUpload auto customUpload accept=".zip"
        :maxFileSize="MAX_ZIP_BYTES"
        :showUploadButton="false"
        @uploader="onUploadZip"
    >
        <template #content>
            <div v-if="files" class="flex column align-items-center">
                <img v-if="packageIcon" class="pt-5"
                    :width="ICON_WIDTH" :src="packageIconSrc">
                </img>
                
                <p class="zip-name">{{ zipName }}</p>
            </div>
        </template>

        <template #empty>
            <div class="flex column align-items-center justify-content-center">
                <div v-if="!files" class="flex row gap-2 align-items-center">
                    <div class="pi pi-upload"></div>
                    <p>Drag and drop your mod zip file here.</p>
                </div>
            </div>
        </template>
    </FileUpload>

    <div class="flex column">
        <p class="output-header mt-4 mb-1">Results</p>
        <div class="flex row align-items-center gap-2">
            <p class="mt-1" style="font-size: 18px;">Status:</p>

            <p v-if="requiredFilesExist" class="mt-1" style="color: green; font-size: 18px;">Valid</p>
            <p v-else class="mt-1" style="color: red; font-size: 18px;">Invalid!</p>
        </div>

        <div class="flex column">
            <p class="output-list-item-name mt-1 mb-1">CHANGELOG.md</p>
            <p class="output-list-item-name mt-1 mb-1">README.md</p>
            <p class="output-list-item-name mt-1 mb-1">manifest.json</p>
        </div>
    </div>

    <!-- <DataView dataKey="dev-package-list">
        
    </DataView> -->
</Viewport>
</template>

<style scoped>
.mod-dev-tools {
    display: flex;
    flex-direction: column;
}

.zip-selection-header {
    font-size: 30px;
    font-weight: 540;
    user-select: none;
}

.output-header {
    font-size: 25px;
    font-weight: 540;
    user-select: none;
}

.output-list-item-name {
    font-size: 18px;
    font-weight: 540;
}

.zip-name {
    font-size: 24px;
    font-weight: 540;
    user-select: none;
}

:deep(.p-fileupload-header) {
    justify-content: center;
    padding-bottom: 0;
    padding-top: 2rem;
}

:deep(.p-fileupload-content) {
    gap: 0;
    width: 25rem;

}
</style>