package thunderstore

import (
	exp "github.com/the-egg-corp/thundergo/experimental"
)

// Collection of util methods relating to Thunderstore mod development.
type Tools struct {
}

func NewTools() *Tools {
	return &Tools{}
}

func (tools *Tools) ValidateManifest(author string, data []byte) (bool, []error, error) {
	return exp.ValidateManifest(author, data)
}

func (tools *Tools) ValidateReadme(data []byte) (bool, []error, error) {
	return exp.ValidateReadme(data)
}

func (tools *Tools) ValidateIcon(params []byte) (bool, error) {
	return exp.ValidateIcon(params)
}
