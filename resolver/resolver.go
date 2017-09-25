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

	// Resolver is the root type resolver
	Resolver struct{}

	/*
		System type related resolvers
	*/
	systemResolver             struct{ system client.Client }
	systemVersionResolver      struct{ version types.Version }
	systemInfoResolver         struct{ info types.Info }
	systemInfoCommitResolver   types.Commit
	systemInfoPluginResolver   struct{ plugins types.PluginsInfo }
	systemInfoRegistryResolver struct{ registryConfig registry.ServiceConfig }
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

/*
	Common assertion funcs
*/

func ptrInt32(input interface{}) *int32 {
	switch input := input.(type) {
	case int:
		ptr := int32(input)
		return &ptr
	case int8:
		ptr := int32(input)
		return &ptr
	case int16:
		ptr := int32(input)
		return &ptr
	case int32:
		return &input
	case int64:
		ptr := int32(input)
		return &ptr
	case uint:
		ptr := int32(input)
		return &ptr
	case uint8:
		ptr := int32(input)
		return &ptr
	case uint16:
		ptr := int32(input)
		return &ptr
	case uint32:
		ptr := int32(input)
		return &ptr
	case uint64:
		ptr := int32(input)
		return &ptr
	default:
		// Wrong type
		return nil
	}
}

func ptrCommit(c types.Commit) *systemInfoCommitResolver {
	commit := systemInfoCommitResolver(c)
	return &commit
}

/*
	Specific assetion funcs
*/

// ptrIPNetString asserts []*registry.NetIPNet to *[]string
func ptrIPNetString(n []*registry.NetIPNet) *[]string {
	// Starts an slice with len 0 and then appends elements to it
	// So it will not output empty string
	// Such as:
	// "registryConfig": {
	//      "allowNondistributableArtifactsCidrs": [
	//        "", # This occurs when the slice is started with len 1
	//      ]
	ipNets := make([]string, 0)
	for _, ipNet := range n {
		ipNets = append(ipNets, ipNet.String())
	}
	return &ipNets
}
