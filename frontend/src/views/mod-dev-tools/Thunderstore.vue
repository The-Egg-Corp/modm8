<script lang="ts" setup>
import { ref } from 'vue'

import FileUpload from 'primevue/fileupload'
import DataView from 'primevue/dataview'

import { TSPackageManifest, TSPackageFile } from '@types'
import { experimental } from '@frontend/wailsjs/go/models'

// TODO: Finish implementing pack/unpack functionality
import { 
    //UnpackZip, 
    ValidateIcon, 
    ValidateManifest, 
    ValidateReadme
} from '@backend/thunderstore/Tools'

import { useAppStore } from "@stores"
import { storeToRefs } from "pinia"

const appStore = useAppStore()
const { sidebarWidth } = storeToRefs(appStore)

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
</script>

<template>
<div class="mod-dev-tools column mt-5 pl-5">
    <div class="mb-4">
        <FileUpload
            name="demo[]"
            accept=".zip"
            :multiple="false"
            :maxFileSize="1000000"
            :showUploadButton="false"
            @select=""
        >
            <template #content>
                <div class="pi pi-upload">
                    
                </div>
            </template>
    
            <template #empty>
                <p>Drag and drop your mod zip file here.</p>
            </template>
    
        </FileUpload>
    </div>

    <DataView dataKey="dev-package-list">
        
    </DataView>
</div>
</template>

<style scoped>
.mod-dev-tools {
    margin-left: v-bind(sidebarWidth); /* Account for sidebar */
    top: 350px;
}
</style>