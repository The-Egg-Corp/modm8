// eslint-disable-next-line @typescript-eslint/consistent-type-imports
import { thunderstore, v1 } from "@backend/models.js"

export type Package = thunderstore.StrippedPackage & {
    latestVersion: v1.PackageVersion
}

export interface TSPackageFile {
    fileName: string
    data: number[]
}

export interface TSPackageManifest extends TSPackageFile {
    author?: string
}