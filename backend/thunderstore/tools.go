package thunderstore

import (
	exp "github.com/the-egg-corp/thundergo/experimental"
)

// Collection of util methods relating to Thunderstore mod development.
type ThunderstoreTools struct {
}

func NewThunderstoreTools() *ThunderstoreTools {
	return &ThunderstoreTools{}
}

func (tools *ThunderstoreTools) ValidateManifest(author string, data []byte) (bool, []error, error) {
	return exp.ValidateManifest(author, data)
}

func (tools *ThunderstoreTools) ValidateReadme(data []byte) (bool, []error, error) {
	return exp.ValidateReadme(data)
}

func (tools *ThunderstoreTools) ValidateIcon(params []byte) (bool, error) {
	return exp.ValidateIcon(params)
}
