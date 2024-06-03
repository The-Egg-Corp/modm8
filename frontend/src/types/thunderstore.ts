export interface TSPackageFile {
    fileName: string
    data: number[]
}

export interface TSPackageManifest extends TSPackageFile {
    author?: string
}