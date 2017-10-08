package resolver

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
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
