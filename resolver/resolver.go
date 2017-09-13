package resolver

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var (
	ctx = context.Background()
)

type (
	Resolver       struct{}
	systemResolver struct {
		system client.Client
	}
	systemVersionResolver struct {
		version types.Version
	}
)

func (r *Resolver) System() *systemResolver {
	dkr, _ := client.NewEnvClient()
	return &systemResolver{system: *dkr}
}

func (r *systemResolver) Version() *systemVersionResolver {
	v, _ := r.system.ServerVersion(ctx)
	return &systemVersionResolver{version: v}
}
