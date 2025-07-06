package nexus

import (
	"context"
	"modm8/backend/app"
	"modm8/backend/utils"

	v1 "github.com/the-egg-corp/gonexus/v1"
)

type API struct {
	Ctx    context.Context
	Cache  map[string]any
	Client v1.NexusClient
}

func NewCache() map[string]any {
	return make(map[string]any, 0)
}

func NewAPI(ctx context.Context) (*API, error) {
	//fmt.Print("Gathering nexus key...\n")

	// TODO: A single file for a key is a bit odd. Store it in a better way. Win/Linux registry maybe?

	// Grab nexus key from `<configDir>/modm8/nex.key` file.
	key, err := utils.FindFirstValidLine(app.KeyPath())
	if err != nil {
		return nil, err
	}

	nexKey := ""
	if key != nil {
		nexKey = *key
	}

	client, err := v1.NewNexusClient(nexKey)
	if err != nil {
		return nil, err
	}

	return &API{
		Ctx:    ctx,
		Cache:  NewCache(),
		Client: *client,
	}, nil
}
