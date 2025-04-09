package nexus

import (
	"context"
	"fmt"
	"modm8/backend/app"
	"os"

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
	// TODO: Is this really the best way of doing this?
	// Grab nexus key from `<configDir>/modm8/nex.key` file
	fmt.Print("Reading nex.key file..\n")

	contents, err := os.ReadFile(app.KeyPath())
	if err != nil || contents == nil || len(contents) < 1 {
		return nil, err
	}

	key := string(contents)
	fmt.Printf("Nexus Key: %s\n\n", key)

	client, err := v1.NewNexusClient(key)
	if err != nil {
		return nil, err
	}

	return &API{
		Ctx:    ctx,
		Cache:  NewCache(),
		Client: *client,
	}, nil
}
