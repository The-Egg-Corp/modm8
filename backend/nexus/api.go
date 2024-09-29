package nexus

import "context"

type API struct {
	Ctx   context.Context
	Cache map[string]any
}

func NewAPI(ctx context.Context) *API {
	return &API{
		Ctx: ctx,
	}
}
