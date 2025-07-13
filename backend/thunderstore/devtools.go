package thunderstore

import (
	exp "github.com/the-egg-corp/thundergo/experimental"
)

// Collection of util methods relating to Thunderstore mod development.
type ThunderstoreDevTools struct {
	PackageValidator *PackageValidator
}

func NewThunderstoreDevTools() *ThunderstoreDevTools {
	return &ThunderstoreDevTools{PackageValidator: NewPackageValidator()}
}

type PackageValidator struct {
}

func NewPackageValidator() *PackageValidator {
	return &PackageValidator{}
}

func (v *PackageValidator) ValidateManifest(author string, data []byte) (bool, []error, error) {
	return exp.ValidateManifest(author, data)
}

func (v *PackageValidator) ValidateReadme(data []byte) (bool, []error, error) {
	return exp.ValidateReadme(data)
}

func (v *PackageValidator) ValidateIcon(params []byte) (bool, error) {
	return exp.ValidateIcon(params)
}
