package resolver

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var (
	ctx = context.Background()
)

type (

	// Resolver is the root type resolver
	Resolver struct{}

	/* System type related resolvers */
	systemResolver             struct{ system client.Client }
	systemVersionResolver      struct{ version types.Version }
	systemInfoResolver         struct{ info types.Info }
	systemInfoCommitResolver   types.Commit
	systemInfoPluginResolver   struct{ plugins types.PluginsInfo }
	systemInfoRegistryResolver struct{ registryConfig registry.ServiceConfig }

	/*	Swarm type related resolvers */
	swarmResolver               struct{ Swarm swarm.Swarm }
	swarmSpecResolver           struct{ SwarmSpec swarm.Spec }
	orchestrationConfigResolver swarm.OrchestrationConfig
	raftConfigResolver          struct{ RaftConfig swarm.RaftConfig }
	dispatcherConfigResolver    struct{ DispatcherConfig swarm.DispatcherConfig }
	caConfigResolver            struct{ CaConfig swarm.CAConfig }
	externalCaResolver          struct{ ExternalCa *swarm.ExternalCA }
	encryptionConfigResolver    swarm.EncryptionConfig
	taskDefaultsResolver        struct{ TaskDefault swarm.TaskDefaults }
	swarmDriverResolver         struct{ LogDriver *swarm.Driver }
	jointTokensResolver         struct{ swarm.JoinTokens }
	unlockKeyResolver           types.SwarmUnlockKeyResponse // struct{ UnlockKey types.SwarmUnlockKeyResponse }
	swarmInfoResolver           struct{ Swarm swarm.Info }
	swarmClusterResolver        struct{ SwarmCluster *swarm.ClusterInfo }
	swarmVersionResolver        swarm.Version
	tlsInfoResolver             struct{ TLSInfo swarm.TLSInfo }
	remoteManagerResolver       struct{ RemoteManager swarm.Peer }
)

/* System type related objects */

// System retrieves a docker client
func (r *Resolver) System() *systemResolver {
	sys := clientInit()
	return &systemResolver{system: *sys}
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

/* Swarm type related objects */

func (r *Resolver) Swarm() *swarmResolver {
	cli := clientInit()
	sw, _ := cli.SwarmInspect(ctx)
	return &swarmResolver{Swarm: sw}
}

/* Docker Client */

func clientInit() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

/*	Specific assertion funcs */

// ptrIPNetString does type assertion []*registry.NetIPNet to *[]string
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

/*	Common assertion funcs */

// ptrString gets does type assertion,
// gets reference from []byte and string
func ptrString(input interface{}) *string {
	switch input := input.(type) {
	case []byte:
		ptr := string(input)
		return &ptr
	case string:
		return &input
	default:
		return nil
	}
}

// ptrCommit does type assertion of types.Commit
func ptrCommit(c types.Commit) *systemInfoCommitResolver {
	commit := systemInfoCommitResolver(c)
	return &commit
}

func ptrSwVer(c swarm.Version) *swarmVersionResolver {
	ver := swarmVersionResolver(c)
	return &ver
}

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
	case *int:
		value := *input
		ptr := int32(value)
		return &ptr
	case *int8:
		value := *input
		ptr := int32(value)
		return &ptr
	case *int16:
		value := *input
		ptr := int32(value)
		return &ptr
	case *int32:
		return input
	case *int64:
		value := *input
		ptr := int32(value)
		return &ptr
	case uintptr:
		value := input
		ptr := int32(value)
		return &ptr
	case *uint8:
		value := *input
		ptr := int32(value)
		return &ptr
	case *uint16:
		value := *input
		ptr := int32(value)
		return &ptr
	case *uint32:
		value := *input
		ptr := int32(value)
		return &ptr
	case *uint64:
		value := *input
		ptr := int32(value)
		return &ptr
	default:
		// Wrong type
		return nil
	}
}
