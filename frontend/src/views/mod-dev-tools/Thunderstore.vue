<script lang="ts" setup>
import { ref } from 'vue'

import { TSPackageManifest, TSPackageFile } from '@types'
import { experimental } from '@frontend/wailsjs/go/models'

// TODO: Finish implementing pack/unpack functionality
import { 
    //UnpackZip, 
    ValidateIcon, 
    ValidateManifest, 
    ValidateReadme
} from '@backend/thunderstore/Tools'

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
<div class="mod-dev-tools column">
    <DataView dataKey="dev-package-list">

    </DataView>
</div>

<div class="no-drag">
    <FileUpload
        name="demo[]"
        accept=".zip"
        :multiple="false"
        :maxFileSize="1000000"
        :showUploadButton="false"
        @select=""
    >
        <template #content>

        </template>

        <template #empty>
            <div class="flex row pi pi-upload gap-2">
                <div>Drag and drop your mod zip file here.</div>
            </div>
        </template>
    </FileUpload>
</div>
</template>

<style scoped>

</style>