package resolver

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var (
	ctx = context.Background()
)

type (

	// Resolver is the main type resolver
	Resolver struct{}

	systemResolver struct {
		system client.Client
	}
	systemVersionResolver struct {
		version types.Version
	}
	systemInfoResolver struct {
		info types.Info
	}
	systemInfoCommitResolver types.Commit
	systemInfoPluginResolver struct {
		plugins types.PluginsInfo
	}
	systemInfoRegistryResolver struct {
		registryConfig registry.ServiceConfig
	}
)

// System retrieves a docker client
func (r *Resolver) System() *systemResolver {
	dkr, _ := client.NewEnvClient()
	return &systemResolver{system: *dkr}
}

// Version retrieves server version related info
func (r *systemResolver) Version() *systemVersionResolver {
	v, _ := r.system.ServerVersion(ctx)
	return &systemVersionResolver{version: v}
}

// Info retrieves a docker host info
func (r *systemResolver) Info() *systemInfoResolver {
	i, _ := r.system.Info(ctx)
	return &systemInfoResolver{info: i}
}
